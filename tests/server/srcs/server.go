package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	zmq "github.com/pebbe/zmq4"
)

type Input struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	zctx, _ := zmq.NewContext()

	socket, _ := zctx.NewSocket(zmq.REQ)

	fmt.Println("Binding socket on port 5555")
	socket.Bind("tcp://*:5555")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON((fiber.Map{"Welcome": "Hello!"}))
	})

	app.Post("/get-length", func(ctx *fiber.Ctx) error {
		input := new(Input)
		if err := ctx.BodyParser(input); err != nil {
			panic(err)
		}
		fmt.Println("before snding...")
		socket.Send(input.Text, 0)
		fmt.Println("after snding...")

		if msg, err := socket.Recv(0); err != nil {
			panic(err)
		} else {
			fmt.Println("Received msg->", msg)
			return ctx.JSON((fiber.Map{"Length": msg}))
		}
	})

	fmt.Println("Starting listening on port 3000")
	app.Listen(":3000")
}
