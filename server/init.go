package server

var OrderChan = make(chan GrabOrderEvent, 100)

func StartServer() {
	go initRouter()
	go initOrderHandler(OrderChan) // 初始化消息队列
	select {}
}
