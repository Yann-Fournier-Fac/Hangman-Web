package hangman

import (
	"math/rand"
	"time"
)

// Affichages de N lettres aléatoires pour deviner le Mot au début
func NLetter(jeu *Hangman) {

	n := len(jeu.Mot)/2 - 1 // Le nombre de lettre à afficher
	tab := []int{}          // creation d'un tableau qui contiendra les indices à afficher
	boolean := false

	// Conversion string to byte
	Mot_byte := []byte(jeu.MotATrouve)

	// On affiche aucune lettre car le mot est trop petit
	if n == 0 {

	} else {

		// On recupere n numéros < len(Mot)
		// Ce sont les indices des lettres à afficher
		for !boolean { // La boucle s'arrete quand tous les éléments de tab sont différents

			// recuperation au hazard des n numeros
			for i := 0; i < n; i++ {

				rand.Seed(time.Now().UnixNano())     // "planter une graine" aléatoire en fonction de l'heure
				radomInt2 := rand.Intn(len(jeu.Mot)) // random des numeros
				tab = append(tab, radomInt2)         // puis ajout a notre tableau tab
			}
			boolean = Alldiff(tab) // tous les indices doivent évidemment etre different (pour bien avoir n numeros)
		}

		// Puis on met les lettres dans le bon tableau et on ajoute les lettres au tableau affiche
		for i := 0; i < len(tab); i++ {

			Mot_byte[tab[i]] = jeu.Mot[tab[i]]

			jeu.Affich = append(jeu.Affich, string(jeu.Mot[tab[i]]))
		}

		// Conversion byte to str
		str := ""
		for i := 0; i < len(Mot_byte); i++ {
			str += string(Mot_byte[i])
		}
		jeu.MotATrouve = str
	}
}

// retourne true si tous les élément du tableau tab sont différents sinon retourne false
// On compare tous les éléments entre eux
func Alldiff(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] == tab[j] {
				return false
			}
		}
	}
	return true
}
