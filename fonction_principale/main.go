package main

// Demander un nom d'utilisateur sur le Menu
// (faire un score board)
// faire une save
// faire un truc qui recupere le buttom du menu

import (
	"fmt"
	"hangman"
	"net/http"
	"text/template"
)

// type Page struct {
// 	title string
// 	body  []byte
// }

func Objet() hangman.Hangman {

	var Hangman hangman.Hangman

	Hangman.Mot = ""
	Hangman.MotATrouve = ""
	Hangman.Pendu = hangman.Hangmanpose()
	Hangman.Cpt = 0
	Hangman.Affich = []string{}
	Hangman.Lettre = []string{}
	Hangman.Pos = []string{}

	return Hangman
}

var Je = Objet()
var Token int = 0

func main() {

	fmt.Printf("\n")
	fmt.Println("http://localhost:8080/menu/")
	fmt.Printf("\n")

	http.HandleFunc("/menu/", menuHandler)
	http.HandleFunc("/me/", meHandler)
	http.HandleFunc("/niveau/", niveauHandler)
	http.HandleFunc("/niv/", nivHandler)
	http.HandleFunc("/jeu/", jeuHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/win/", winHandler)
	http.HandleFunc("/lose/", loseHandler)

	//CSS en static
	//http.Handle("/", http.FileServer(http.Dir("/Web/style/")))
	//http.Handle("/Web/style/", http.StripPrefix("/Web/style/", nil))
	http.ListenAndServe(":8080", nil)

}

var tpmlMenu = template.Must(template.ParseFiles("./Web/menu.html"))
var tpmlNiveau = template.Must(template.ParseFiles("./Web/niveau.html"))
var tpmlJeu = template.Must(template.ParseFiles("./Web/Jeu.html"))
var tpmlWin = template.Must(template.ParseFiles("./Web/win.html"))
var tpmlLose = template.Must(template.ParseFiles("./Web/lose.html"))

// Affichage de la page
func menuHandler(w http.ResponseWriter, r *http.Request) {
	err := tpmlMenu.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//button := r.FormValue("boutton")
	// if button == "New Game" {
	// 	Token = 1
	// }
	// fmt.Print(Token)

	//http.Redirect(w, r, "/menu/", http.StatusFound)
	//http.Redirect(w, r, "/pas_acces/", http.StatusFound)

}

func meHandler(w http.ResponseWriter, r *http.Request) {
	button := r.FormValue("bouton")
	if button == "New Game" {
		http.Redirect(w, r, "/niveau/", http.StatusFound)
	} else if button == "Continuer une Sauvegarde" {
		fmt.Println("partie Restaurée")
		// Reload les données
		//http.Redirect(w, r, "/jeu/", http.StatusFound)
		http.Redirect(w, r, "/niveau/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/menu/", http.StatusFound)
	}
}

func niveauHandler(w http.ResponseWriter, r *http.Request) {
	//if Token == 1 {
	err := tpmlNiveau.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// } else {
	// 	http.Redirect(w, r, "/menu/", http.StatusFound)
	// }
}

func nivHandler(w http.ResponseWriter, r *http.Request) {
	niveau := r.FormValue("niveau")

	if niveau[len(niveau)-1] == '1' {
		Je.Mot = hangman.Findwords("words.txt")
		Token = 2
		// fmt.Println(Je.Mot)

	} else if niveau[len(niveau)-1] == '2' {
		Je.Mot = hangman.Findwords("words2.txt")
		Token = 2
		//fmt.Println(Je.Mot)

	} else if niveau[len(niveau)-1] == '3' {
		Je.Mot = hangman.Findwords("words3.txt")
		Token = 2
		//fmt.Println(Je.Mot)
	}

	n := len(Je.Mot)/2 - 1
	Je.Lettremanque = len(Je.Mot) - n

	Je.MotATrouve = ""

	for i := 0; i < len(Je.Mot); i++ {
		Je.MotATrouve += "_"
	}

	hangman.NLetter(&Je)

	//fmt.Print(Token)
	//fmt.Println(niveau)
	http.Redirect(w, r, "/jeu/", http.StatusFound)
}

func jeuHandler(w http.ResponseWriter, r *http.Request) {
	//if Token == 2 {
	Je.Pos = hangman.Hang(Je.Cpt, Je.Pendu)
	err := tpmlJeu.Execute(w, Je)
	if err != nil {
		fmt.Println("il y a une erreur")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// } else {
	// 	http.Redirect(w, r, "/menu/", http.StatusFound)
	// }
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	lettre := r.FormValue("letter")
	if lettre == "Sauvegarder" {
		fmt.Println("Partie Sauvegarder")
		//hangman.Save(Je)
		http.Redirect(w, r, "/menu/", http.StatusFound)
	} else {
		hangman.Compa(lettre, &Je)
		//fmt.Println(lettre)
		if (Je.Cpt < 10) && (Je.Lettremanque != 0) {
			http.Redirect(w, r, "/jeu/", http.StatusFound)
		} else if Je.Cpt > 9 {
			hangman.Reload(&Je)
			http.Redirect(w, r, "/lose/", http.StatusFound)
		} else if Je.Lettremanque == 0 {
			hangman.Reload(&Je)
			http.Redirect(w, r, "/win/", http.StatusFound)
		}
	}
}

func winHandler(w http.ResponseWriter, r *http.Request) {
	err := tpmlWin.Execute(w, Je)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loseHandler(w http.ResponseWriter, r *http.Request) {
	Je.Pos = hangman.Hang(Je.Cpt, Je.Pendu)
	err := tpmlLose.Execute(w, Je)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
