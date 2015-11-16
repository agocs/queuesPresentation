package main

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
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
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go consume(connection, i, wg)
		wg.Add(1)
	}
	wg.Wait()

}

func consume(conn *amqp.Connection, workerNum int, wg sync.WaitGroup) {
	//START OMIT2
	channel, err1 := conn.Channel()
	if err1 != nil {
		panic(err1.Error())
	}
	defer channel.Close()
	defer wg.Done()
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
		log.Printf("Worker %v got a message: %s", workerNum, d.Body)
	}
	//END4 OMIT
}
