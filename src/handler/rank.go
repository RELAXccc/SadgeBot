package handler


import (
  "SadgeBot/src/bot"
  "strconv"
  "log"
  "github.com/bwmarrin/discordgo"
  "github.com/replit/database-go"
)


func RankHandler(message *bot.Message) *discordgo.Message {
  	

  

    matches, _ := database.ListKeys("log_" + message.Message.Author.ID)

    exp := 0
    for _, match := range matches{

      svalue, err := database.Get(match)

      if err != nil{
                log.Println(match)

        log.Println(err)
        continue
      }
      
      i , err := strconv.Atoi(svalue) // todo: maybe handle error, should never happen tho
      if err != nil{
        database.Delete(match)
        continue
      }
      exp += i
    }
    log.Println("EXP???:")



    level := (float64(exp) / float64(2000)) + 1 

    log.Println(exp)

    content := message.Message.Message.Author.Mention() + " you have " + strconv.Itoa(exp) + "EXP so your level is: " + strconv.Itoa(int(level))
     answer, _ := message.Session.ChannelMessageSend(message.Message.Message.ChannelID, content) // todo: maybe handle error, should never happen tho

     return answer
}