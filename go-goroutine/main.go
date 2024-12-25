package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)
type Message struct{
	Data string `json:"data"`
}

type Pubsub struct{
	subs []chan Message
	mu   sync.Mutex
}

func (pb *Pubsub) Subscribe() chan Message{
	pb.mu.Lock()
	defer pb.mu.Unlock()
	ch := make(chan Message, 1)
	pb.subs = append(pb.subs, ch)
	return ch
}
func (pb *Pubsub) Publisher(msg *Message) {
	pb.mu.Lock()
	defer pb.mu.Unlock()
	for _, sub := range pb.subs{
		sub <- *msg
	}
}

func main(){
	// Goroutine schedule every 10 secs with Cronjob --> try to check out "crontab guru"
	c := cron.New(cron.WithSeconds())
	c.AddFunc("10 * * * * *", func(){
		fmt.Println("Goroutine with cron every 10 seconds")
	})
	c.Start()

	// Pubsub pattern to broadcast the publisher message to every subscriber channel
	app := fiber.New()
	pubsub := &Pubsub{}

	app.Post("/publisher", func(c *fiber.Ctx) error {
		message := new(Message)
		if err := c.BodyParser(message); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		pubsub.Publisher(message)
		return c.JSON(fiber.Map{
			"message": "add to subscriber",
		})
	})

	sub := pubsub.Subscribe()
	go func(){
		for msg := range sub{
			fmt.Println("Receive Message: ", msg)
		}
	}()

	app.Listen(":8080")
}