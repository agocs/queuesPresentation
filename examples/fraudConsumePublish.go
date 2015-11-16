package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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
	rx_queue, err2 := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	if err2 != nil {
		panic(err2.Error())
	}
	//END3 OMIT

	//EXCHANGESTART OMIT
	ex_name := "postFraudExchange"
	kind := "topic"
	channel.ExchangeDeclare(ex_name, kind, false, false, false, false, nil)
	//EXCHANGEEND OMIT

	consumer := ""
	autoAck := false
	cons_exclusive := false
	noLocal := false
	cons_noWait := false
	msgs, err3 := channel.Consume(rx_queue.Name, consumer, autoAck, cons_exclusive, noLocal, cons_noWait, nil)
	if err3 != nil {
		panic(err3.Error())
	}

	for d := range msgs {
		log.Printf("Got a message: %s", d.Body)

		//START4 OMIT
		key := ""
		//Advanced Fraud Detection
		if coinflip() {
			key = "Transactions.Success"
		} else {
			key = "Transactions.Failure"
		}
		exchange := ex_name
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
		log.Printf("Published one message with the key: %s\n", key)
		//END4 OMIT

		d.Ack(false)
	}

}

func coinflip() bool {
	return rand.Intn(2) == 0

}
