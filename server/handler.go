package server

import (
	"github.com/gin-gonic/gin"
	"main/api"
)

func initOrderHandler(orderEventChan chan GrabOrderEvent) {
	for {
		orderEvent := <-orderEventChan
		// 存储订单
		api.Logger.Info("receive orderEvent: ", orderEvent.Order)
		order := orderEvent.Order

		// 检查活动是否存在或者是否过期
		activity := &api.Activity{}
		api.DB.First(activity, order.ActivityID)
		if activity.ID == 0 || activity.EndTime < order.CreatedAt || activity.StartTime > order.CreatedAt {
			orderEvent.RecvChan <- GrabItemResponse{
				// 活动不存在或者已经过期，或者还没开始
				H: gin.H{
					"message": "activity not exists or expired or not start yet",
				},
				Code: 400,
			}
			api.Logger.Info("activity not exists or expired")
			continue
		}

		// 检查库存
		item := &api.Item{}
		api.DB.First(item, order.ItemID)
		if item.Stock <= 0 {
			orderEvent.RecvChan <- GrabItemResponse{
				H: gin.H{
					"message": "stock not enough",
				},
				Code: 400,
			}
			api.Logger.Info("stock not enough")
			continue
		}

		// 库存充足，减库存
		item.Stock--
		api.DB.Save(item)
		// 保存订单
		api.DB.Save(orderEvent.Order)
		orderEvent.RecvChan <- GrabItemResponse{
			H: gin.H{
				"message": "grab item success",
			},
			Code: 200,
		}
		api.Logger.Info("save orderEvent success")
	}
}
