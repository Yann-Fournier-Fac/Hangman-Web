package hangman

import (
	"strings"
)

func Compa(lettre string, jeu *Hangman) { // comparer le mot rentrer par le joueur

	if len(lettre) >= 2 {
		
		// comparer le stop (si le joueur veux arreter)
		if len(lettre) == len(jeu.Mot) { // Comparer les mots de meme longeur

			var cpt2 int // compteur de lettre correspondentes
			for i := 0; i < len(jeu.Mot); i++ {
				let := strings.ToLower(string(lettre[i]))
				if let == string(jeu.Mot[i]) {
					cpt2++
				}
			}
			if cpt2 == len(jeu.Mot) { // verification du nbr de lettre correspondentes
				jeu.Lettremanque = 0
			} else {
				jeu.Cpt += 2
			}

			jeu.Wrong = ""

		} else if len(lettre) != len(jeu.Mot) {
			jeu.Cpt += 2
			jeu.Wrong = "Ce n'est pas le bon mot"
			//fmt.Println(Purple + "Cette lettre à déjà été rentrée" + Reset)
		}

	} else { // Sinon comparer la lettre rentrer par le joueur

		var cpt4 int = 0 // compteur de lettre correspondentes

		lettre = strings.ToLower(lettre) // mettre la lettre en minuscule (si besoin)

		// On cherche si la lettre a déjà été rentrée
		for i := 0; i < len(jeu.Lettre); i++ {
			if lettre == jeu.Lettre[i] {
				cpt4++
				break
			}
		}

		if cpt4 == 0 { // Si elle est nouvelle

			jeu.Lettre = append(jeu.Lettre, lettre) // ajout à lettre déjà entrer à notre liste

			let := 0
			byt := []byte(jeu.MotATrouve)
			byt2 := []byte(lettre)
			// On ajoute la lettre au Mot a trouver
			for i := 0; i < len(jeu.Mot); i++ {
				if lettre == string(jeu.Mot[i]) { // si la lettre corespond

					byt[i] = byte(byt2[0]) // Ajout de la lettre au mot a trouver

					jeu.Lettremanque-- // decremante lettre manquante
					let++              // Si une lettre correspond : ++
				}
			}

			jeu.MotATrouve = ""
			// byte to string
			for i := 0; i < len(byt); i++ {
				jeu.MotATrouve += string(byt[i])
			}

			// Pour savoir si aucune lettre ne correspond et qu'il faut augmenter le cpt d'erreur de un
			if let == 0 {
				jeu.Cpt++
			}

			// On regarde si la lettre était déjà afficher
			cpt5 := 0 // compter cbm de fois la lettre est afficher
			for i := 0; i < len(jeu.Affich); i++ {
				if jeu.Affich[i] == lettre {
					cpt5++
				}
			}
			jeu.Lettremanque += cpt5 // Puis on ajoute a lettremanque car les lettres afficher
			//ne sont pas consider comme des lettres manquantes
			
			jeu.Wrong = ""

		} else {
			jeu.Wrong = "Cette lettre à déjà été rentrée"
			// 	fmt.Println(Purple + "Cette lettre à déjà été rentrée" + Reset)
		}
	}
}
