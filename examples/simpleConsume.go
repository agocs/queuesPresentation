package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//START OMIT1
	url := "amqp://guest:guest@" + "localhost" + ":5672"
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	//END OMIT1

	//START OMIT2
	channel, err1 := connection.Channel()
	if err1 != nil {
		panic(err1.Error())
	}
	defer channel.Close()
	//END OMIT2

	//START OMIT3
	name := "TransactionFirehose"
	durable := true
	autoDelete := false
	exclusive := false
	noWait := false
	queue, err2 := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	if err2 != nil {
		panic(err2.Error())
	}
	//END OMIT3

	//START4 OMIT
	consumer := ""
	autoAck := true
	cons_exclusive := false
	noLocal := false
	cons_noWait := false
	msgs, err3 := channel.Consume(queue.Name, consumer, autoAck, cons_exclusive, noLocal, cons_noWait, nil)
	if err3 != nil {
		panic(err3.Error())
	}

	for d := range msgs {
		log.Printf("Got a message: %s", d.Body)
	}
	//END4 OMIT

}
