package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"transaction/entity"
)


func handleerror(err error,msg string)  {
	if err != nil {
		log.Fatalf("%s:%s",msg,err)
	}
}

func Rabbit(uID string,gID string,sID string,number string) {
	app := entity.App{
		UID:uID,
		GID:gID,
		SID:sID,
		Number:number,
	}

	//连接rabbit
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleerror(err,"failed to connection to rabbit")
	defer conn.Close()

	//打开通道
	ch,err := conn.Channel()
	handleerror(err,"failed to open a channel")
	defer ch.Close()

	//声明消息队列
	queue,err := ch.QueueDeclare("goods",true,false,false,false,nil)
	handleerror(err,"failed to declare a queue")

	body,err := json.Marshal(app)
	handleerror(err,"failed to encoding json")

	err = ch.Publish("",queue.Name,false,false,amqp.Publishing{
		DeliveryMode:amqp.Persistent,
		ContentType:"text/plain",
		Body:body,
	})
	handleerror(err,"failed to publish an information")

	log.Printf("publish an information:%s",string(body))
}


