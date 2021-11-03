package bot

import "github.com/bwmarrin/discordgo"

type Command struct {
  Id         string   // id for db (e.g. when Repeatable==true)
	Name       string
	Ids        []string
  Repeatable bool     // only allows one executions, hints to other answer. (4 jokes + better performance with lower db accesses)
	Handler    func(m *Message) *discordgo.Message
}

type Message struct {
	Session     *discordgo.Session
	Message     *discordgo.MessageCreate
	Command     *Command
	CommandData *[]string
	Reply       *discordgo.Message
	Error       error
	Closed      bool
}
