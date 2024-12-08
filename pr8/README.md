# Практическая работа №7
## Лапин Д.С. ПИМО-01-24


### 1. GET products - получение всех товаров

![image](https://github.com/user-attachments/assets/bd3659dc-c6a1-4a47-a9a4-1e4875ddb748)




### sql запрос
```sql
DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INT NOT NULL,
	price INT,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO categories (id, name) VALUES
(1, 'Солнцезащитные очки'),
(2, 'Контактные линзы'),
(3, 'Оправы для очков'),
(4, 'Очки для чтения'),
(5, 'Защитные очки');

INSERT INTO products (id, name, description, category_id, price)
VALUES
(1, 'Солнцезащитные очки Ray-Ban Aviator', 'Классические авиаторы с металлической оправой и темными линзами.', 1, 100000),
(2, 'Контактные линзы Acuvue Oasys', 'Дышащие двухнедельные контактные линзы для комфортного ношения.', 2, 5000),
(3, 'Оправа для очков Gucci GG0010O', 'Стильная оправа из ацетата с фирменным дизайном Gucci.', 3, 12000),
(4, 'Очки для чтения Foster Grant', 'Легкие и удобные очки для чтения с различными диоптриями.', 4, 15000),
(5, 'Солнцезащитные очки Oakley Holbrook', 'Спортивные очки с высокой ударопрочностью и защитой от УФ-лучей.', 1, 6000),
(6, 'Контактные линзы Biofinity', 'Месячные силикон-гидрогелевые линзы с высоким уровнем увлажнения.', 2, 10000),
(7, 'Оправа для очков Ray-Ban Wayfarer', 'Иконическая пластиковая оправа в стиле ретро.', 3, 10000),
(8, 'Компьютерные очки Gunnar Optiks', 'Очки с фильтром синего света для защиты глаз при работе за компьютером.', 5, 7432),
(9, 'Контактные линзы Dailies Total 1', 'Ежедневные одноразовые линзы с уникальной технологией увлажнения.', 2, 13243),
(10, 'Солнцезащитные очки Polaroid PLD 1013/S', 'Очки с поляризованными линзами для четкого и контрастного зрения.', 1, 9000);

```


### Код 
```go
package main

import (
	"net/http"
	"time"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	var product Product

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
	}
	c.JSON(http.StatusOK, product)
}

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

// key, users, roles

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

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

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
	//Products []Product `gorm:"foreignKey:CategoryID"` // Связь с продуктами
}

func initDB() {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
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

```
