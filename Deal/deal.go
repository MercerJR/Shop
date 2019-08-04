package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
	"transaction/Dao"
	"transaction/entity"
)

func dealerror(err error,msg string)  {
	if err != nil {
		log.Fatalf("%s:%s",msg,err)
	}
}

func main()  {
	//连接rabbit
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	dealerror(err,"failed to connect to rabbit")
	defer conn.Close()

	//打开通道
	ch,err := conn.Channel()
	dealerror(err,"failed to open a channel")
	defer ch.Close()

	//声明一个消息队列
	queue,err := ch.QueueDeclare("goods",true,false,false,false,nil)
	dealerror(err,"failed to declare a queue")

	err = ch.Qos(1,0,false)
	dealerror(err,"failed to configure Qos")

	msgch,err := ch.Consume(queue.Name,"",false,false,false,false,nil)
	dealerror(err,"failed to register a consumer")

	stopchan := make(chan bool)

	go func() {
		log.Printf("Consumer ready,PID:%d",os.Getpid())
		for d := range msgch{
			log.Printf("receve an message:%s",string(d.Body))

			app := &entity.App{}
			err := json.Unmarshal(d.Body,app)
			dealerror(err,"failed to decoding json")

			order := &entity.Order{
				CId:app.UID,
				GID:app.GID,
				SID:app.SID,
				Number:app.Number,
				Time:time.Now(),
			}

			shop := &entity.Shop{
				GID:app.GID,
				SID:app.SID,
				GNum:app.Number,
			}

			Dao.Insert(order)
			Dao.Update(shop)

			log.Printf("information:%s",string(d.Body))

			if err := d.Ack(false);err != nil {
				log.Printf("error acknowledging message:%s",err)
			}else {
				log.Printf("acknowledged message")
			}
		}
	}()
	<-stopchan
}
