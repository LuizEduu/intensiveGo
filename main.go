package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/luizeduu/intensiveGo/internal/infra/database"
	"github.com/luizeduu/intensiveGo/internal/useCase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	calculateFinalPriceUseCase := useCase.NewCalculateFinalPriceUseCase(orderRepository)

	input := useCase.OrderInput{
		ID:    "any_id",
		Price: 10,
		Tax:   5,
	}

	output, err := calculateFinalPriceUseCase.Execute(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)

}
