package main

import (
	"fmt"

	"github.com/Thadchanon/learn-basic-go/movie"

	"github.com/Thadchanon/learn-basic-go/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
