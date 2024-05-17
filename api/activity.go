package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddActivity(c *gin.Context) {
	auth, msg, user := Authorize(c.GetHeader("Authorization"))
	if !auth || !user.IsAdmin {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
	}

	activityName := c.PostForm("name")
	actStartTime, _ := strconv.ParseInt(c.PostForm("start_time"), 10, 64)
	actEndTime, _ := strconv.ParseInt(c.PostForm("end_time"), 10, 64)
	actDiscount, _ := strconv.ParseFloat(c.PostForm("discount"), 64)
	activityDescription := c.PostForm("description")

	activity := Activity{
		Name:        activityName,
		Description: activityDescription,
		StartTime:   actStartTime,
		EndTime:     actEndTime,
		Discount:    actDiscount,
	}
	DB.Create(&activity)
	c.JSON(200, gin.H{
		"message": "add activity",
		"status":  "ok",
		"data":    activity,
	})
}

func UpdateActivity(c *gin.Context) {
	auth, msg, user := Authorize(c.GetHeader("Authorization"))
	if !auth || !user.IsAdmin {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
	}

	actID, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	actName := c.PostForm("name")
	actStartTime, _ := strconv.ParseInt(c.PostForm("start_time"), 10, 64)
	actEndTime, _ := strconv.ParseInt(c.PostForm("end_time"), 10, 64)
	actDiscount, _ := strconv.ParseFloat(c.PostForm("discount"), 64)
	actDescription := c.PostForm("description")

	activity := Activity{ID: actID}
	DB.Model(&activity).Updates(Activity{
		Name:        actName,
		Description: actDescription,
		StartTime:   actStartTime,
		EndTime:     actEndTime,
		Discount:    actDiscount,
	})

	c.JSON(200, gin.H{
		"message": "update activity",
		"status":  "ok",
		"data":    activity,
	})
}

func DeleteActivity(c *gin.Context) {
	auth, msg, user := Authorize(c.GetHeader("Authorization"))
	if !auth || !user.IsAdmin {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
	}

	actID, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	var activity Activity
	DB.First(&activity, actID)
	DB.Delete(&activity)

	c.JSON(200, gin.H{
		"message": "delete activity",
		"status":  "ok",
	})
}
