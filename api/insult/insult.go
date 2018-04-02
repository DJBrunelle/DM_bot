package insult

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Insult returns a random insult from insult.mattbas.org
func Insult() string {
	response, err := http.Get("https://insult.mattbas.org/api/insult")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}

	data, _ := ioutil.ReadAll(response.Body)
	return string(data[:len(data)])
}
