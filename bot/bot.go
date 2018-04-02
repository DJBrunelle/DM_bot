package bot

import (
	"DM_bot/api/answer"
	"DM_bot/api/doggo"
	"DM_bot/api/insult"
	"DM_bot/config"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	active     bool
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
	active = true
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

	if m.Content == "insult norlando" {
		nolan, _ := s.User(users["nolan"])
		_, _ = s.ChannelMessageSend(m.ChannelID, nolan.Mention()+" is a nolan")
		return
	}

	if m.Content == "bot status" {
		if active {
			_, _ = s.ChannelMessageSend(m.ChannelID, "On")
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Off")
		}

		return
	}

	if m.Content == "bot off" && active {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Shutting off...")
		active = false
		return
	}
	if m.Content == "bot on" && !active {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Turning on...")
		active = true
		return
	}

	if active == false {
		return
	}

	if strings.ToLower(m.Content) == "doggo" {
		dog := doggo.Doggo("")
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}
	if strings.ToLower(m.Content) == "shibe" {
		dog := doggo.Doggo("shiba")
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}
	if strings.ToLower(m.Content) == "poofer" {
		dog := doggo.Doggo("pomeranian")
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}
	if strings.ToLower(m.Content) == "boofer" {
		r := rand.Intn(3)
		var dog string
		if r == 0 {
			dog = doggo.Doggo("stbernard")
		} else if r == 1 {
			dog = doggo.Doggo("husky")
		} else if r == 2 {
			dog = doggo.Doggo("eskimo")
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}
	if strings.ToLower(m.Content) == "borker" {
		r := rand.Intn(3)
		var dog string
		if r == 0 {
			dog = doggo.Doggo("rottweiler")
		} else if r == 1 {
			dog = doggo.Doggo("germanshepherd")
		} else if r == 2 {
			dog = doggo.Doggo("doberman")
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}

	if strings.ToLower(m.Content) == "pupper" {
		r := rand.Intn(4)
		var dog string
		if r == 0 {
			dog = doggo.Doggo("papillon")
		} else if r == 1 {
			dog = doggo.Doggo("pomeranian")
		} else if r == 2 {
			dog = doggo.Doggo("pekinese")
		} else if r == 3 {
			dog = doggo.Doggo("poodle/miniature")
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, dog)
		return
	}

	//Static bot responses
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		prefixHandler(s, m)
		return
	}

	//Rolls an n sided dice for user
	if strings.HasPrefix(strings.ToLower(m.Content), "roll") {
		rollHandler(s, m)
		return
	}

	//insult person
	if strings.HasPrefix(strings.ToLower(m.Content), "insult") {
		userID := ""
		msg := strings.Split(strings.ToLower(m.Content), " ")
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
		return
	}

	if strings.HasSuffix(m.Content, "?") {
		answer, image := answer.YesNoMaybe()
		_, _ = s.ChannelMessageSend(m.ChannelID, answer)
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
		return
	}

	if strings.HasPrefix(strings.ToLower(m.Content), "yes") || strings.HasPrefix(strings.ToLower(m.Content), "ya") {
		image := answer.Yes()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
		return
	}

	if strings.HasPrefix(strings.ToLower(m.Content), "no") || strings.HasPrefix(strings.ToLower(m.Content), "nah") {
		image := answer.No()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "maybe") {
		image := answer.Maybe()
		_, _ = s.ChannelMessageSend(m.ChannelID, image)
		return
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
	return "bot status - indicates whether bot is on or off\nbot on/off - turns bot on or off\ninsult <name> - insults person\nroll n - rolls an n sided dice for user\nyes/no/maybe - sends a gif based on command\n ? - sends random yes no maybe answer"
}

//roll an n sided dice
func roll(n int) int {
	return rand.Intn(n) + 1
}
