package main

import (
	"SadgeBot/src/server"
	"flag"
	"fmt"
  "log"
	"os"
  "time"
	"github.com/bwmarrin/discordgo"
  "github.com/jaredfolkins/badactor"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.Parse()
}




var Bot server.CozyBot

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	Bot.Disc = dg


  // studio capacity
	var sc int32
	// director capacity
	var dc int32

	sc = 1024
	dc = 1024

  st := badactor.NewStudio(sc)

  blockRule := &badactor.Rule{
		Name:        "spamBlock",
		Message:     "slow cowboy",
		StrikeLimit: 2,
		ExpireBase:  time.Second * 30,
		Sentence:    time.Minute * 1, // todo: test out some settings xD
		Action:      &Bot,
	}
	st.AddRule(blockRule)

   kickRule := &badactor.Rule{
		Name:        "spamKick",
		Message:     "goodbye",
		StrikeLimit: 10, // after 10 blocks in 10 minutes
		ExpireBase:  time.Minute * 10,
		Sentence:    time.Hour * 1, // gets kicked on every new message for 1 Hour
		Action:      &Bot,
	}
	st.AddRule(kickRule)

	err = st.CreateDirectors(dc)
	if err != nil {
		log.Fatal(err)
	}



  Bot.BadActor = st

	Bot.Run()
}
