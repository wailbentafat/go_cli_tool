package project

import (
	"fmt"
	
)

// Template prints the filename (not used currently)


// InMain generates the main.go content
func InMain() string {
	content := `package main

import (
	"fmt"
	"your_project_path/utils"
)

func main() {
	fmt.Println("Hello, World!")
}
`
	return content
}

// InitDB generates the database connection file content
func InitDB() string {
	content := `package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}
`
	return content
}

// InitRedis generates the Redis initialization file content
func InitRedis() string {
	content := `package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func InitializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func SetCache(key string, value string, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

func GetCache(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}
`
	return content
}

// InitAuthRoutes generates the authentication routes file content
func InitAuthRoutes(projectName string) string {
	content := fmt.Sprintf(`package routes

import (
	"github.com/gin-gonic/gin"
	"%s/auth/controllers" 
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
`, projectName)
	return content
}

// InitAuthModels generates the user model file content
func InitAuthModels() string {
	content := `package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   ` + "`gorm:\"primaryKey\"`" + `
	Email     string ` + "`json:\"email\" gorm:\"unique\"`" + `
	Password  string ` + "`json:\"password\"`" + `
}
`
	return content
}

// InitAuthControllers generates the authentication controllers file content
func InitAuthControllers(projectName string) string {
	content := fmt.Sprintf(`package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"%s/utils/db"
	"%s/utils/jwt"
	"%s/utils/redis"
	"%s/auth/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username string ` + "`json:\"username\"`" + `
	jwt.RegisteredClaims
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	redis.SetCache(tokenString, "logged_in", 30*time.Minute)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}
`, projectName, projectName, projectName, projectName)
	return content
}

// InitMiddleware generates the middleware for JWT authentication
func InitMiddleware(projectName string) string {
	content := fmt.Sprintf(`package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"%s/utils/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
`, projectName)
	return content
}

// InitJWT generates the JWT utilities file content
func InitJWT() string {
	content := `package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string ` + "`json:\"email\"`" + `
	jwt.RegisteredClaims
}

var JwtSecretKey = []byte(os.Getenv("JWT_KEY"))

func CreateToken(email string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorUnverifiable)
		}
		return JwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}

	return claims, nil
}
`
	return content
}
