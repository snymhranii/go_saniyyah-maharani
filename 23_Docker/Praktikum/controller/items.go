package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model"
	"project_structure/model/payload"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetItemsController(c echo.Context) error {
	items, e := usecase.GetListProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items":  items,
	})
}

func GetItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	item, err := usecase.GetProduct(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get product",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items":  item,
	})
}

// create new product
func CreateItemController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	payload := payload.ProductRequest{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	item, err := usecase.CreateItem(payload.Name, payload.Description, payload.Price, payload.Stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"items":   item,
	})
}

// delete product by id
func DeleteItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete item",
			"errorDescription": err,
			"errorMessage":     "Sorry, the product cannot be item",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

// update product by id
func UpdateItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	product := model.Product{}
	c.Bind(&product)
	product.ID = uint(id)

	updateStockRequest := payload.UpdateStockRequest{}
	if err := c.Bind(&updateStockRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update item data",
			"errorDescription": err,
			"errorMessage":     "Sorry, the product data cannot be changed",
		})
	}

	if err := usecase.UpdateProduct(&product, updateStockRequest.Stock); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update item",
			"errorDescription": err,
			"errorMessage":     "Sorry, the item cannot be changed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
	})
}
