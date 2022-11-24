package main

import (
	"fmt"
	"hangman"
)

func main() {

	var Hangman hangman.Hangman

	Hangman.Mot = hangman.Findwords()
	Hangman.MotATrouve = ""
	Hangman.Pendu = hangman.Hangmanpose()
	Hangman.Pendu = hangman.Space_transform(Hangman.Pendu)
	Hangman.Cpt = 0
	Hangman.Affich = []string{}
	for i := 0; i < len(Hangman.Mot); i++ {
		Hangman.MotATrouve += "_"
	}
	Hangman.Lettre = []string{}
	Hangman.Pos = []string{}
	hangman.NLetter(&Hangman)

	for i := 0; i < len(Hangman.Pendu); i++ {
		fmt.Print(Hangman.Pendu[i])
		fmt.Printf("\n")
	}

	
}
