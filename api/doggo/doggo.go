package doggo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type doggoStruct struct {
	URL string `json:"message"`
}

//Doggo returns a URL linking to a pic of a doggo
func Doggo() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}

	var doggo *doggoStruct
	//data holds json
	data, _ := ioutil.ReadAll(response.Body)

	//update answer with json
	err = json.Unmarshal(data, &doggo)

	fmt.Println(doggo.URL)

	return doggo.URL
}
