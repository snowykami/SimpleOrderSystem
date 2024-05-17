package server

import (
	"github.com/gin-gonic/gin"
	"main/api"
	"strconv"
	"time"
)

type GrabOrderEvent struct {
	Order    *api.Order
	RecvChan chan GrabItemResponse
}

type GrabItemResponse struct {
	H    gin.H
	Code int
}

// GrabItem 抢商品
func GrabItem(c *gin.Context) {
	// 获取用户id
	auth, msg, user := api.Authorize(c.GetHeader("Authorization"))
	if !auth {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
		return
	}
	// 获取商品id
	itemId, _ := strconv.ParseInt(c.PostForm("item_id"), 10, 64)

	actId, _ := strconv.ParseInt(c.PostForm("activity_id"), 10, 64)
	// 抢购商品
	recvChan := make(chan GrabItemResponse)
	go GrabItemService(recvChan, user.ID, itemId, actId)
	// 等待抢购结果
	resp := <-recvChan
	c.JSON(resp.Code, resp.H)
}

func GrabItemService(ch chan GrabItemResponse, userId int64, itemId int64, actId int64) {
	// 生成订单
	// 将订单放入消息队列
	OrderChan <- GrabOrderEvent{
		Order: &api.Order{
			UserID:     userId,
			ItemID:     itemId,
			ActivityID: actId,
			CreatedAt:  time.Now().Unix(),
		},
		RecvChan: ch,
	}
}
