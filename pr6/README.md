![image](https://github.com/user-attachments/assets/3c46f99a-84b8-4834-abf0-d1afdcf35436)## Практическая работа №6
### Лапин Д.С. ПИМО-01-24

Код 
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Category    string
}

var products = []Product{

	{ID: "1", Name: "Солнцезащитные очки Ray-Ban Aviator", Description: "Классические авиаторы с металлической оправой и темными линзами.", Category: "Солнцезащитные очки"},
	{ID: "2", Name: "Контактные линзы Acuvue Oasys", Description: "Дышащие двухнедельные контактные линзы для комфортного ношения.", Category: "Контактные линзы"},
	{ID: "3", Name: "Оправа для очков Gucci GG0010O", Description: "Стильная оправа из ацетата с фирменным дизайном Gucci.", Category: "Оправы для очков"},
	{ID: "4", Name: "Очки для чтения Foster Grant", Description: "Легкие и удобные очки для чтения с различными диоптриями.", Category: "Очки для чтения"},
	{ID: "5", Name: "Солнцезащитные очки Oakley Holbrook", Description: "Спортивные очки с высокой ударопрочностью и защитой от УФ-лучей.", Category: "Солнцезащитные очки"},
	{ID: "6", Name: "Контактные линзы Biofinity", Description: "Месячные силикон-гидрогелевые линзы с высоким уровнем увлажнения.", Category: "Контактные линзы"},
	{ID: "7", Name: "Оправа для очков Ray-Ban Wayfarer", Description: "Иконическая пластиковая оправа в стиле ретро.", Category: "Оправы для очков"},
	{ID: "8", Name: "Компьютерные очки Gunnar Optiks", Description: "Очки с фильтром синего света для защиты глаз при работе за компьютером.", Category: "Защитные очки"},
	{ID: "9", Name: "Контактные линзы Dailies Total 1", Description: "Ежедневные одноразовые линзы с уникальной технологией увлажнения.", Category: "Контактные линзы"},
	{ID: "10", Name: "Солнцезащитные очки Polaroid PLD 1013/S", Description: "Очки с поляризованными линзами для четкого и контрастного зрения.", Category: "Солнцезащитные очки"},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

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

1. GET products - получение всех товаров

![image](https://github.com/user-attachments/assets/d1ed399c-4d84-4987-88e9-79552a9b67fd)


2. GET products/{id} - получение информации о товаре по id

![image](https://github.com/user-attachments/assets/63e42d53-6894-448d-a873-05e68b92b2d7)


3. POST products - создание нового товаров

![image](https://github.com/user-attachments/assets/82af3bd5-95da-4247-8fe9-405236586dba)

![image](https://github.com/user-attachments/assets/8309e8c8-0b84-4330-b21c-983f6cb50e5f)


4. PUT products/{id} - обновление информации о существубщем товаре

![image](https://github.com/user-attachments/assets/b27c58d7-6e5b-49fe-b611-33a47282dc91)

![image](https://github.com/user-attachments/assets/d4c7adde-cd34-4af7-9815-a060a4b8fa62)


5. DELETE products/{id} - удаление товара

![image](https://github.com/user-attachments/assets/cdb8b25d-ecb8-436a-802a-3d692bf9370a)

![image](https://github.com/user-attachments/assets/864a7f30-b961-4702-9460-ef184c8ef225)

