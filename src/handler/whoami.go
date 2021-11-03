package handler


import (
  "SadgeBot/src/bot"
  "github.com/bwmarrin/discordgo"
)


func WhoamiHandler(message *bot.Message) *discordgo.Message {
  	content := "hello " + message.Message.Message.Author.ID +",\nyea, you're no more than a fucking number for me\ndeal with it ¯\\_(ツ)_/¯"
    answer, _ := message.Session.ChannelMessageSend(message.Message.Message.ChannelID, content)

    return answer
}