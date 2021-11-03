package server

import (
	"github.com/bwmarrin/discordgo"
  "github.com/jaredfolkins/badactor"
)

type CozyBot struct {
	Disc *discordgo.Session
  BadActor *badactor.Studio
}
