package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Ecriture(structure *Hangman) {
	content, err := json.Marshal(structure)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("save.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Lecture() Hangman {
	// var hangman Hangman
	// content, err := os.ReadFile("save.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err2 := json.Unmarshal(content, &hangman)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// return hangman
	content, err := ioutil.ReadFile("save.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into payload
	var payload Hangman
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}
