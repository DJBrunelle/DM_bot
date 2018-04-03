package kitty

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type kittyStruct struct {
	URL string `xml:"data>images>image>url"`
}

//Kitty returns a URL linking to a pic of a doggo
func Kitty() string {
	response, err := http.Get("http://thecatapi.com/api/images/get?format=xml&results_per_page=1")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}

	var kitty *kittyStruct
	//data holds json
	data, _ := ioutil.ReadAll(response.Body)

	//update answer with json
	err = xml.Unmarshal(data, &kitty)

	return kitty.URL
}
