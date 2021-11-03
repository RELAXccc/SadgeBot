package server

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/replit/database-go"
)

func (cb *CozyBot) LogXp(session *discordgo.Session, message *discordgo.MessageCreate) error {
	uid, err := strconv.ParseInt(message.Author.ID, 10, 64)
	if err != nil {
		return err
	}

	//todo: think about an better exp algorythm


	err = database.Set("log_"+message.Author.ID+"_"+strconv.FormatInt(time.Now().Unix(), 10), strconv.Itoa(len(strings.Trim(message.Content, " ,.-!?"))+100))

	if err != nil {
		return err
	}
	log.Println("ID:")
	log.Println(uid)
	log.Println("Chars:")
	//ParseMessage(message.Content)

	log.Println(len(strings.Trim(message.Content, " ,.-!?")))
	log.Println("ID: " + message.Author.ID)
	log.Println("ID: " + message.Author.ID)

	return nil
}
