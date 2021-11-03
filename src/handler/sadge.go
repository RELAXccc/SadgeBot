package handler


import (
  "SadgeBot/src/bot"
  "github.com/bwmarrin/discordgo"
  "math/rand"
  "fmt"
)


func SadgeHandler(message *bot.Message) *discordgo.Message {


answers := make([]string, 0)
answers = append(answers,
    "really sadge",
    "not sadge at all",
    "verry sadge",
    "maybe a little bit... happy",
    "maybe a little bit sadge",
    "verry happy",
    "playing lol",
    "headshoting noobs",
    "just sadge",
    "sadge'nt")

content := fmt.Sprint("right now i'm ", answers[rand.Intn(len(answers))])
    answer, _ := message.Session.ChannelMessageSend(message.Message.Message.ChannelID, content)

    return answer
}