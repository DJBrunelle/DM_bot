package bot

import (
	"DM_bot/api/answer"
	"DM_bot/config"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session
var yesOrNoVal *yesNoStruct

const nolan = "135923889795497984"
const daniel = "196128092874342401"
const cel = "92086825551687680"
const derek = "90648400776671232"
const nino = "91360391333945344"
const han = "137690439909113856"
const sarah = "108729905713303552"
const twysper = "90967921043451904"
const david = "217053054699175936"
const benji = "90968733580820480"

type yesNoStruct struct {
	Answer string `json:"answer"`
	Image  string `json:"image"`
}

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

	if strings.HasSuffix(m.Content, "?") {
		answer, image := answer.YesNoMaybe()
		_, _ = s.ChannelMessageSend(m.ChannelID, answer)
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "yes") || strings.Contains(strings.ToLower(m.Content), "ya") {
		image := answer.Yes()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "no") || strings.Contains(strings.ToLower(m.Content), "nah") {
		image := answer.No()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "maybe") {
		image := answer.Maybe()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

}

func help() string {
	return "!ping - command to check if bot is active\nroll n - rolls an n sided dice for user\nyes/no/maybe - sends a gif based on command\n ? - sends random yes no maybe answer"

}

//roll an n sided dice
func roll(n int) int {
	return rand.Intn(n) + 1
}
