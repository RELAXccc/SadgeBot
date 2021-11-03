package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
  "github.com/replit/database-go"
	"log"
  "strings"
	"time"
)

var deleteCmdDuration = 10 + time.Second

//delete message (optional with duration delay)
func (m *Message) Delete(d ...*time.Duration) {
	if len(d) > 0 && d[0] != nil {
		time.Sleep(*d[0])
	}
	err := m.Session.ChannelMessageDelete(m.Message.ChannelID, m.Message.ID)
	if err != nil {
		log.Println(err)
	}
}

//add command reply
func (m *Message) CommandReply(d ...*time.Duration) {
	if len(d) > 0 && d[0] != nil {
		time.Sleep(*d[0])
	}
	err := m.Session.ChannelMessageDelete(m.Message.ChannelID, m.Message.ID)
	if err != nil {
		log.Println(err)
	}
}

// update command reply
func (m *Message) UpdateCommandReply() {
	content := m.Message.Author.Mention() + " wants "

	if m.Command != nil {
		log.Println(m.Command)
		content += "to " + m.Command.Name
	} else {
		content += "something"
	}

	if m.Error != nil {
		content += "\n" + m.Error.Error()
	}

	if m.Reply == nil {
		reply, err := m.Session.ChannelMessageSend(m.Message.ChannelID, content)
		if err != nil {
			m.Error = err
			return
		}
		m.Reply = reply
	} else {
		edit := discordgo.MessageEdit{
			Content: &content,
			ID:      m.Reply.ID,
			Channel: m.Reply.ChannelID,
		}
		reply, err := m.Session.ChannelMessageEditComplex(&edit)
		if err != nil {
			m.Error = err
			return
		}
		m.Reply = reply
	}

	// delete info reply
	if m.Closed {
		time.Sleep(deleteCmdDuration)
		err := m.Session.ChannelMessageDelete(m.Reply.ChannelID, m.Reply.ID)
		if err != nil {
			log.Println(err)
		}
	}
}


// handle command
func (m *Message) HandleCommand() {
	if m.Command.Handler == nil{
		m.Error = errors.New("5️⃣0️⃣1️⃣ command Not Implemented")
	}else{
  go func(){
    //non-repeatables exception
    if !m.Command.Repeatable{
      usage, err := database.Get(m.Command.Id + "_" + m.Message.Message.Author.ID + "_" + m.Message.ChannelID)
      if err == nil{
        log.Println("command already used")
        log.Println(usage)
        answerIds := strings.Split(usage,"_")
        if len(answerIds) != 2{
          return
        }

        m.Session.ChannelMessageSendReply(m.Message.ChannelID, "Are you stupid? You already asked my this!", &discordgo.MessageReference{
          MessageID : answerIds[0],
          ChannelID:  answerIds[1],
        } )

        return
      }
      answer := m.Command.Handler(m)

      
      database.Set(m.Command.Id + "_" + m.Message.Message.Author.ID + "_" + answer.ChannelID , answer.ID + "_" + answer.ChannelID)
    }else{
      m.Command.Handler(m)
    }
  }()

    
	}
	m.Closed = true
	m.UpdateCommandReply()
}
