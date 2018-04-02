package answer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type answerStruct struct {
	Answer string `json:"answer"`
	Image  string `json:"image"`
}

//YesNoMaybe returns an answer as a string (yes, no, or maybe)
//and returns a gif associated with answer
func YesNoMaybe() (string, string) {
	answer := getAnswer("")
	return answer.Answer, answer.Image
}

//Yes returns a gif signalling 'yes'
func Yes() string {
	answer := getAnswer("yes")
	return answer.Image
}

//No returns a gif signalling 'no'
func No() string {
	answer := getAnswer("no")
	return answer.Image
}

//Maybe returns a gif signalling 'Maybe'
func Maybe() string {
	answer := getAnswer("maybe")
	return answer.Image
}

// Requests URL from Yes/No API and returns answer with gif
func getAnswer(input string) answerStruct {
	var url string
	//Set url based on input
	if input == "" {
		url = "https://yesno.wtf/api"
	} else if input == "yes" {
		url = "https://yesno.wtf/api?force=yes"
	} else if input == "no" {
		url = "https://yesno.wtf/api?force=no"
	} else if input == "maybe" {
		url = "https://yesno.wtf/api?force=maybe"
	}
	response, err := http.Get(url)
	var answer *answerStruct

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return *answer
	}

	//data holds json
	data, _ := ioutil.ReadAll(response.Body)

	//update answer with json
	err = json.Unmarshal(data, &answer)

	return *answer
}
