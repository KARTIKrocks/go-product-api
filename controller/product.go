package controller

import (
	"github.com/gin-gonic/gin"
	"goProductAPI/models"
	"net/http"
	"strconv"
)

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

// CRUD operation for controller
func (repo *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	err := models.CreateProduct(repo.Products, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, repo.Products)
}
func (repo *ProductRepo) ReadProduct(c *gin.Context) {
	c.JSON(http.StatusOK, repo.Products)
}
func (repo *ProductRepo) ReadProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	product := models.ReadProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}
func (repo *ProductRepo) UpdateProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	var product models.Product
	c.BindJSON(&product)
	if product.Id != int(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updateProduct := models.UpdateProductById(repo.Products, &product)
	c.JSON(http.StatusOK, &updateProduct)
}
func (repo *ProductRepo) DeleteProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	product := models.DeleteProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}
