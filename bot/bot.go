package bot

import (
	"DM_bot/config"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

const nolan = "GutsmansAss"
const daniel = "iel"

//Start runs the gobot
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

//MessageHandler listens for messages sent in discord chat
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botID {
		return
	}

	//Static bot responses
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Content == "!help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, help())
		}
		//ping to see if bot is running
		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong "+m.Author.Username)
		}

		if m.Content == "!meme" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "https://memegen.link/buzz/memes/memes_everywhere.jpg")
		}
	}

	//Rolls an n sided dice for user
	if strings.HasPrefix(m.Content, "roll") {
		strRoll := strings.Split(m.Content, " ")

		if len(strRoll) <= 1 {
			return
		}
		if n, err := strconv.Atoi(strRoll[1]); err == nil {
			i := roll(n)
			_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Username+" rolled a "+strconv.Itoa(i)+" on a "+strconv.Itoa(n)+" sided dice")
			if i == 1 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "lol "+m.Author.Mention()+" is a loser")
			} else if i == n {
				_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" flippin nailed it")
			}
		}
	}

}

func help() string {
	return "!ping - command to check if bot is active\nroll n - rolls an n sided dice for user"

}

//roll an n sided dice
func roll(n int) int {
	return rand.Intn(n) + 1
}
