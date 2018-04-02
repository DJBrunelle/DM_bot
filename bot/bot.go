package bot

import (
	"DM_bot/api/answer"
	"DM_bot/api/insult"
	"DM_bot/config"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	botID      string
	goBot      *discordgo.Session
	yesOrNoVal *yesNoStruct
	users      = map[string]string{
		"nolan":   "135923889795497984",
		"daniel":  "196128092874342401",
		"cel":     "92086825551687680",
		"derek":   "90648400776671232",
		"nino":    "91360391333945344",
		"han":     "137690439909113856",
		"sarah":   "108729905713303552",
		"twysper": "90967921043451904",
		"david":   "217053054699175936",
		"benji":   "90968733580820480",
		"shawn":   "122520927689900033",
	}
)

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
		prefixHandler(s, m)
	}

	//Rolls an n sided dice for user
	if strings.HasPrefix(strings.ToLower(m.Content), "roll") {
		rollHandler(s, m)
	}

	//insult person
	if strings.HasPrefix(strings.ToLower(m.Content), "insult") {
		userID := ""
		msg := strings.Split(m.Content, " ")
		if len(msg) == 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Possible people to insult: ")
			var ppl []string
			for k := range users {
				ppl = append(ppl, k)
				ppl = append(ppl, "\n")
			}
			pplMsg := strings.Join(ppl[:], "")
			_, _ = s.ChannelMessageSend(m.ChannelID, pplMsg)
			return
		}
		for k, v := range users {
			if msg[1] == k {
				userID = v
			}
		}
		if userID != "" {
			insult := insult.Insult()
			user, _ := s.User(userID)
			_, _ = s.ChannelMessageSend(m.ChannelID, user.Mention()+" "+insult)
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

func prefixHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, help())
	}

	if m.Content == "!meme" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "https://memegen.link/buzz/memes/memes_everywhere.jpg")
	}
}

func rollHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	strRoll := strings.Split(m.Content, " ")
	var i int
	var n int
	var err error

	if len(strRoll) == 1 {
		n = 20
		i = roll(n)
	} else if n, err = strconv.Atoi(strRoll[1]); err == nil {
		i = roll(n)
	} else {
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Username+" rolled a "+strconv.Itoa(i)+" on a "+strconv.Itoa(n)+" sided dice")
	if i == 1 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "lol "+m.Author.Mention()+" is a loser")
	} else if i == n {
		_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" flippin nailed it")
	}
}

func help() string {
	return "roll n - rolls an n sided dice for user\nyes/no/maybe - sends a gif based on command\n ? - sends random yes no maybe answer"
}

//roll an n sided dice
func roll(n int) int {
	return rand.Intn(n) + 1
}
