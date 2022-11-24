package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Findwords(name string) string {
	Mot := ""
	words := Readfile(name)

	if len(words) != 0 {

		// trirage d'un indice < len(words)
		rand.Seed(time.Now().UnixNano())
		nbr := rand.Intn(len(words))

		Mot = words[nbr]
		return Mot // puis on retourne le mot

	} else { // Si words est vide
		return "Veuillez relancer le jeux"
	}
}

func Readfile(name string) []string { // name : nom d'un fichier.txt

	words := []string{} // tableau qui contiendra les mots

	contents, err := os.Open(name) // Ouverture fichier

	if err != nil { // Gestion erreur
		fmt.Println("File reading error", err)
		fmt.Println("Veuillez relencer le jeux")
		return words // on retourne les mots
	}

	defer contents.Close() // Fermeture du fichier

	scanner := bufio.NewScanner(contents) // Transformation de contents en tableau de byte
	scanner.Split(bufio.ScanWords)        // Separer le tableau de byte par ligne

	// Ajouter chaque ligne au tableau words (une ligne correspond à un mot)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil { // gestion d'erreur
		fmt.Println(err)
	}

	return words // On retourne les mots
}

