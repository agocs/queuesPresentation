package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//START1 OMIT
	url := "amqp://guest:guest@" + "localhost" + ":5672"
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	//END1 OMIT

	//START2 OMIT
	channel, err1 := connection.Channel()
	if err1 != nil {
		panic(err1.Error())
	}
	defer channel.Close()
	//END2 OMIT

	//START3 OMIT
	name := "TransactionFirehose"
	durable := true
	autoDelete := false
	exclusive := false
	noWait := false
	queue, err2 := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	if err2 != nil {
		panic(err2.Error())
	}
	//END3 OMIT

	//START4 OMIT
	exchange := ""
	key := queue.Name
	mandatory := true
	immediate := false
	msg := amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "UTF-8",
		Body:            []byte("blahblah"),
		DeliveryMode:    amqp.Transient,
		Priority:        0,
	}
	channel.Publish(exchange, key, mandatory, immediate, msg)
	log.Println("Published one message to the queue!")
	//END4 OMIT

}
