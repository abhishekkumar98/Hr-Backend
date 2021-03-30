package controller

import (
	"log"
	"net/http"
	"product/model"
	"product/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []model.Register
	err := service.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(c *gin.Context) {
	var product model.Register

	c.BindJSON(&product)
	log.Printf(" Prod to be Created %v", product)
	err := service.CreateProduct(&product)
	if err != nil {
		log.Println("Error", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("Product Created %v", product)
		c.JSON(http.StatusOK, product)

	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Register
	err := service.GetProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
func GetUserByName(c *gin.Context) {
	name := c.Params.ByName("firstname")
	var user model.Register
	err := service.GetUserByName(&user, name)
	if err != nil {
		c.AbortWithStatus(http.StatusFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func UpdateProduct(c *gin.Context) {
	var product model.Register
	id := c.Params.ByName("id")
	err := service.GetProduct(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = service.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product model.Register
	id := c.Params.ByName("id")
	err := service.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
