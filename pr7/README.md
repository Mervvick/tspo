# Практическая работа №7
## Лапин Д.С. ПИМО-01-24



### 1. POST signup - регистрация пользователя с ролью user

Дефолтный юзер уже существует

![image](https://github.com/user-attachments/assets/1417e697-1883-43a9-bac0-d9f4fe0a2edd)

Новый юзер

![image](https://github.com/user-attachments/assets/e539f6fd-d082-46e9-8981-f2fc69303ce8)


### 2. POST login - получение токена для юзера

Верные данные

![image](https://github.com/user-attachments/assets/6bed7ee0-adf3-48be-a452-89f650b651b3)

Неверный пароль

![image](https://github.com/user-attachments/assets/c68e1777-5d4e-48db-9b86-998f201df59c)


### 3. GET products - получение всех товаров (user)

Доступно для роли user

![image](https://github.com/user-attachments/assets/ef676eee-f234-45c0-a206-4bc352a3207f)

### 4. DELETE products/{id} - удалении товара (user)

У роли user нет доступа к удалению

![image](https://github.com/user-attachments/assets/89957459-63d1-40c0-893e-49c026c1c666)

### 5. POST signup - регистрация пользователя с ролью admin

![image](https://github.com/user-attachments/assets/c556106d-0a1d-4f37-843b-0957debe6ba9)

### 6. POST login - получение токена для админа

![image](https://github.com/user-attachments/assets/2f1aea4f-a25a-49c0-968b-6bd0e21f9210)

### 7. DELETE products/{id} - удалении товара (admin)

У админа есть доступ к удалению товаров

![image](https://github.com/user-attachments/assets/c9e1d379-6423-482d-b5a8-86c7800db7dd)

Проверка

![image](https://github.com/user-attachments/assets/5b2b9b2d-0828-48f5-b5f6-63591b5cd5cf)



### Код 
```go
package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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

var jwtKey = []byte("very_secret_phrase")

var users = map[string]string{
	"admin": "admin",
	"user1": "qwerty",
}

var roles = map[string]string{
	"admin": "admin",
	"user":  "user",
}

type Product struct {
	ID          string
	Name        string
	Description string
	Category    string
	Price       int
}

var products = []Product{

	{ID: "1", Name: "Солнцезащитные очки Ray-Ban Aviator", Description: "Классические авиаторы с металлической оправой и темными линзами.", Category: "Солнцезащитные очки", Price: 100000},
	{ID: "2", Name: "Контактные линзы Acuvue Oasys", Description: "Дышащие двухнедельные контактные линзы для комфортного ношения.", Category: "Контактные линзы", Price: 5000},
	{ID: "3", Name: "Оправа для очков Gucci GG0010O", Description: "Стильная оправа из ацетата с фирменным дизайном Gucci.", Category: "Оправы для очков", Price: 12000},
	{ID: "4", Name: "Очки для чтения Foster Grant", Description: "Легкие и удобные очки для чтения с различными диоптриями.", Category: "Очки для чтения", Price: 15000},
	{ID: "5", Name: "Солнцезащитные очки Oakley Holbrook", Description: "Спортивные очки с высокой ударопрочностью и защитой от УФ-лучей.", Category: "Солнцезащитные очки", Price: 6000},
	{ID: "6", Name: "Контактные линзы Biofinity", Description: "Месячные силикон-гидрогелевые линзы с высоким уровнем увлажнения.", Category: "Контактные линзы", Price: 10000},
	{ID: "7", Name: "Оправа для очков Ray-Ban Wayfarer", Description: "Иконическая пластиковая оправа в стиле ретро.", Category: "Оправы для очков", Price: 10000},
	{ID: "8", Name: "Компьютерные очки Gunnar Optiks", Description: "Очки с фильтром синего света для защиты глаз при работе за компьютером.", Category: "Защитные очки", Price: 7432},
	{ID: "9", Name: "Контактные линзы Dailies Total 1", Description: "Ежедневные одноразовые линзы с уникальной технологией увлажнения.", Category: "Контактные линзы", Price: 13243},
	{ID: "10", Name: "Солнцезащитные очки Polaroid PLD 1013/S", Description: "Очки с поляризованными линзами для четкого и контрастного зрения.", Category: "Солнцезащитные очки", Price: 9000},
}

func main() {
	router := gin.Default()

	router.POST("/login", login)
	router.POST("/signup", signup)
	router.POST("/refresh", refresh)

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

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)

}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

```
