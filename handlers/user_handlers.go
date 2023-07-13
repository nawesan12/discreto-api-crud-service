package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"discreto-api-crud-service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret") // Reemplaza esto con tu propia clave secreta JWT

func SetupUserRoutes(r *gin.Engine, DB *gorm.DB) {
	r.GET("/users", func(c *gin.Context) {
		GetUsers(c, DB)
	})
	r.POST("/users", func(c *gin.Context) {
		CreateUser(c, DB)
	})
	r.POST("/login", func(c *gin.Context) {
		Login(c, DB)
	})
	r.GET("/me", Authenticate(), func(c *gin.Context) {
		Me(c, DB)
	})
	r.PUT("/me", Authenticate(), func(c *gin.Context) {
		UpdateProfile(c, DB)
	})
}

func GetUsers(c *gin.Context, DB *gorm.DB) {
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context, DB *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context, DB *gorm.DB) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding request to struct: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var user models.User
	if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		log.Printf("Error finding user with email: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("Error comparing password hashes: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing JWT token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Me(c *gin.Context, DB *gorm.DB) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(c.MustGet("token").(string), &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	var user models.User
	if err := DB.First(&user, claims["id"]).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context, db *gorm.DB) {
	var user models.User
	id := getUserId(c)

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var reqBody struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		// Agrega más campos aquí según sea necesario
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// user.FirstName = reqBody.FirstName
	// user.LastName = reqBody.LastName
	// Asegúrate de actualizar cualquier otro campo que hayas agregado

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, ok := claims["user_id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.Abort()
				return
			}
			c.Set("user_id", userId)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getUserId(c *gin.Context) string {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Fatal("cannot find user_id in context")
	}
	return userId.(string)
}
