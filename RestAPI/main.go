package main

import (
	"net/http"
	"time"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	_ "myproject/docs"
)

//=================================
// response
//=================================

type Good struct {
	ID          string  `json:"id" example:"1"`
	Name        string  `json:"name" example:"Стол"`
	Description string  `json:"description" example:"Обычный деревянный стол"`
	Price       float32 `json:"price" example:"10000.0"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"id not found"`
}

//=================================
// data
//=================================

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryID  int
	Price       int
}

//=================================
// API functions
//=================================

// @Summary Получить все товары
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Success 200 {array} Good
// @Router /products [get]
func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// @Summary Получить товар по ID
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Success 200 {array} Good
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [get]
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	var product Product

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Добавить новый товар
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Param good body Good true "Информация о товаре"
// @Success 200 {array} Good
// @Failure 404 {object} ErrorResponse
// @Router /products [post]
func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var category Category
	if err := db.First(&category, newProduct.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid category ID"})
		return
	}

	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)

}

// @Summary Обновить существующий товар
// @Description Обновляет данные товара по ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара"
// @Param good body Good true "Новые данные товара"
// @Success 200 {object} Good
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [put]
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := db.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// @Summary Удалить товар
// @Description Удаляет товар по ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Success 200 {object} Good
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [delete]
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})

}

//=================================
// JWT
//=================================

var jwtKey = []byte("very_secret_phrase")

var users = map[string]string{
	"admin": "admin",
	"user1": "qwerty",
}

var roles = map[string]string{
	"admin": "admin",
	"user":  "user",
}

// structs

type Credentials struct {
	Username string
	Password string
	Role     string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	Role string `json:"role"`
}

// funcs
func generateToken(username string, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Role:     role, // Включаем роль в токен
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// @Summary      Залогин юзера
// @Description  Аутентификация пользователя и возврат JWT токена при усехе
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      Credentials  true  "Данные пользователя для входа"
// @Success      200  {object} ErrorResponse
// @Failure      400  {object} ErrorResponse
// @Failure      401  {object} ErrorResponse
// @Failure      500  {object} ErrorResponse
// @Router       /login [post]
func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	storedPassword, ok := users[creds.Username]
	if !ok || storedPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	role, roleExists := roles[creds.Username]
	if !roleExists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "role not assigned"})
		return
	}

	token, err := generateToken(creds.Username, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
				c.Abort()
				return
			}

			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// @Summary      Регистрация юзера
// @Description  Регистрация нового пользователя через указание юзера и пароля
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      Credentials  true  "Данные пользователя для входа"
// @Success      201  {object} ErrorResponse
// @Failure      400  {object} ErrorResponse
// @Failure      409  {object} ErrorResponse
// @Router       /signup [post]
func signup(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if _, exists := users[creds.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
		return
	}

	role := "user" // дефолтная роль

	if creds.Role != "" {
		role = creds.Role
	}

	users[creds.Username] = creds.Password
	roles[creds.Username] = role

	c.JSON(http.StatusCreated, gin.H{"message": "user signed up successfully"})
}

func roleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		claims := &Claims{}

		log.Println(tokenString)

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "error while encoding token"})
			c.Abort()
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		} else {
			log.Println("invalid token")
		}

		if claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// @Summary      Обновление токена
// @Description  Обновляет устаревший JWT токен
// @Tags         auth
// @Produce      json
// @Param        Authorization  header    string  true  "Токен авторизации"
// @Success      200  {object} ErrorResponse
// @Failure      400  {object} ErrorResponse
// @Failure      401  {object} ErrorResponse
// @Failure      500  {object} ErrorResponse
// @Router       /refresh [post]
func refresh(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{"message": "token not expired enough"})
		return
	}

	newToken, err := generateToken(claims.Username, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

//=================================
// Database
//=================================

var db *gorm.DB

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func initDB() {
	dsn := "host=postgres user=postgres password=admin dbname=restapi port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Product{}, &Category{})
}

//=================================
// main
//=================================

func main() {
	initDB() // инициализация бд
	router := gin.Default()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// запросы авторизации
	router.POST("/login", login)
	router.POST("/signup", signup)
	router.POST("/refresh", refresh)

	// Проверка авторизации
	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		router.GET("/products", getProducts)
		router.GET("/products/:id", getProductByID)
		router.POST("/products", roleMiddleware("admin"), createProduct)
		router.PUT("/products/:id", roleMiddleware("admin"), updateProduct)
		router.DELETE("/products/:id", roleMiddleware("admin"), deleteProduct)
	}

	router.Run(":8080")
}
