package bot

import (
	"DM_bot/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session
var yesOrNoVal *yesNoStruct

const nolan = "GutsmansAss"
const daniel = "iel"
const cel = "Firefly"
const derek = "Derek"
const nino = "Fajitasforme"
const han = "beardo"
const sarah = "Sarah"
const twysper = "Twysper"
const david = "Alucard1557"
const benji = "( ͡ ° ͜ ʖ ͡ ° )"

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
		answer, image := yesOrNo()
		_, _ = s.ChannelMessageSend(m.ChannelID, answer)
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "yes") || strings.Contains(strings.ToLower(m.Content), "ya") {
		image := yes()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "no") || strings.Contains(strings.ToLower(m.Content), "nah") {
		image := no()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
	}

	if strings.Contains(strings.ToLower(m.Content), "maybe") {
		image := maybe()
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

func yesOrNo() (string, string) {
	response, err := http.Get("https://yesno.wtf/api")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", ""
	}
	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &yesOrNoVal)

	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}
	return yesOrNoVal.Answer, yesOrNoVal.Image
}

func yes() string {
	response, err := http.Get("https://yesno.wtf/api?force=yes")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}
	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &yesOrNoVal)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return yesOrNoVal.Image
}

func no() string {
	response, err := http.Get("https://yesno.wtf/api?force=no")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}
	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &yesOrNoVal)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return yesOrNoVal.Image
}

func maybe() string {
	response, err := http.Get("https://yesno.wtf/api?force=maybe")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}
	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &yesOrNoVal)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return yesOrNoVal.Image
}
