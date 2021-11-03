package server

import (
	"fmt"
  "github.com/gofiber/fiber/v2"


	"github.com/bwmarrin/discordgo"
)

func (cb *CozyBot) Run() {

	// Register the messageCreate func as a callback for MessageCreate events.
	cb.Disc.AddHandler(cb.messageCreate)

	// In this example, we only care about receiving message events.
	cb.Disc.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := cb.Disc.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
  defer cb.Disc.Close()
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋! Greetings, SadgeBot 🙂")
    })

    app.Listen(":8000")

	// Cleanly close down the Discord session.
	
}
