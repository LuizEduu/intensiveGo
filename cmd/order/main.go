package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/luizeduu/intensiveGo/internal/infra/database"
	"github.com/luizeduu/intensiveGo/internal/useCase"
	"github.com/luizeduu/intensiveGo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	defer db.Close()

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	calculateFinalPriceUseCase := useCase.NewCalculateFinalPriceUseCase(orderRepository)

	channel, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	rabbitMqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(channel, rabbitMqChannel) //abre uma nova trhead para não travar a principal
	go rabbitmq.Consume(channel, rabbitMqChannel) //abre uma nova trhead para não travar a principal
	rabbitMqWorker(rabbitMqChannel, calculateFinalPriceUseCase)

}

func rabbitMqWorker(messageChannel chan amqp.Delivery, uc *useCase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range messageChannel {
		var input useCase.OrderInput

		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Mensagem processada e salvada no banco de dados", output)

	}

}
