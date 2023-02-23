package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ISSCrew is used to Struct Json data in it
type ISSCrew struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	People  []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//if res.Status == "200" {
	fmt.Printf("HTTP: %s\n", res.Status)
	//}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := ISSCrew{}
	jsonErr := json.Unmarshal(body, &people1)
	// Print struct with its definitions
	fmt.Printf("%+v", people1)
	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
		log.Fatal(jsonErr)
	}

	people1.People = append(people1.People,
		struct {
			Craft string `json:"craft"`
			Name  string `json:"name"`
		}{"Central Line", "Rowan Atkinson"},
		struct {
			Craft string `json:"craft"`
			Name  string `json:"name"`
		}{"Picadilly Line", "Ozhan Karaman"},
	)
	people1.Number = len(people1.People)

	//fmt.Println(people1.People[3].Name, people1.Number)
	fmt.Println(people1.Number)
	fmt.Println(people1.People[1].Name)
	fmt.Println("Last Crew: ", people1.People[len(people1.People)-1].Name)
	for _, v := range people1.People {
		fmt.Println(v.Craft, v.Name)
	}

	jsonBytes, jsonErr := json.Marshal(people1)
	if jsonErr != nil {
		log.Fatalf("unable to unmarshall error: %s", jsonErr.Error())
		log.Fatal(jsonErr)
	}
	// if you not convert to string it will be jsonBytes string
	fmt.Printf("%+v", string(jsonBytes))

	err = ioutil.WriteFile("./iss-updated.json", jsonBytes, 0644)
	if err != nil {
		panic(err)
	}

}
