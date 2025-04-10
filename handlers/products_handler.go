package handlers

import (
	"inventory/models"
	"inventory/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{service: services.NewProductService()}
}

func (h *ProductHandler) SaveProduct(c *gin.Context) {

	product := models.Product{}

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong body for product",
		})
		return
	}

	savedProduct, err := h.service.SaveProduct(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, savedProduct)

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	product := models.Product{}

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong body for product",
		})
		return
	}

	updatedProduct, err := h.service.UpdateProduct(id, &product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)

}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	product, err := h.service.GetProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, products)

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	err = h.service.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted the product",
	})

}
