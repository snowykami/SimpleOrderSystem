package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryItem(c *gin.Context) {
	var items []Item
	itemId := c.Query("id")
	if itemId != "" {
		DB.First(&items, itemId)
	} else {
		DB.Find(&items)
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"status":  "ok",
		"data":    items,
	})
}

// AddItem Add item
func AddItem(c *gin.Context) {
	auth, msg, user := Authorize(c.GetHeader("Authorization"))
	if !auth || !user.IsAdmin {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
		return
	}

	itemName := c.PostForm("name")
	itemPrice, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	itemStock, _ := strconv.ParseInt(c.PostForm("stock"), 10, 64)

	item := Item{Name: itemName, Price: itemPrice, Stock: itemStock}
	DB.Create(&item)
	c.JSON(200, gin.H{
		"message": "ok",
		"status":  "ok",
		"data":    item,
	})
}

// UpdateItem Update item
func UpdateItem(c *gin.Context) {
	auth, msg, user := Authorize(c.GetHeader("Authorization"))
	if !auth || !user.IsAdmin {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
	}

	itemID, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	itemName := c.PostForm("name")
	itemPrice, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	itemStock, _ := strconv.ParseInt(c.PostForm("stock"), 10, 64)

	item := Item{ID: itemID}
	DB.Model(&item).Updates(Item{Name: itemName, Price: itemPrice, Stock: itemStock})
	c.JSON(200, gin.H{
		"message": "ok",
		"status":  "ok",
		"data":    item,
	})
}
