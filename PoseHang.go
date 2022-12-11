package hangman

import (
	"bufio"
	"io"
	"os"
)

// Récupération des positions du Pendu
// on renvoi un tableau de string ou chaque string correspond a une ligne du fichier hangman.txt
func Hangmanpose() []string {

	var pose []string // Creation d'un tableau de string vide

	for i := 0; i < 8; i++ {
		pose = append(pose, "")
	}

	file, err := os.Open("hangman.txt") // On ouvre le fichier file

	if err != nil { // On gère l'erreur
		panic(err)
	}

	defer file.Close() // On ferme le fichier

	reader := bufio.NewReader(file) // On lit le *fichier

	for {
		line, _, err := reader.ReadLine() // Lire ligne par ligne

		if err == io.EOF { // On arrête la boucle si erreur (a la fin du ficher)
			break
		}

		pose = append(pose, string(line)) // On ajoute la ligne à notre tableau de string
		//On transforme "line" en string car c'est un tableau de byte

	}
	return pose // enfin on retourne le tableau de string
}

// Affichage du Pendu
// Donc on affiche seulement les lignes que l'on souhaite
func Hang(hp int, pose []string) []string {
	tab := []string{}
	// hp correspond a notre compteur d'erreur
	a := 0 + hp*8
	b := 7 + hp*8
	// On affiche les lignes que l'on souhaite (multiple de 8)
	for i := a; i <= b; i++ {
		tab = append(tab, pose[i])
	}
	return tab
}
