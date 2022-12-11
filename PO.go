package hangman

//import "net/http"

type Hangman struct {
	Mot          string
	MotATrouve   string
	Cpt          int
	Lettremanque int
	Wrong        string
	Pendu        []string
	Pos          []string
	Affich       []string
	Lettre       []string
}

// func (hang Hangman) ServeHTTP(r http.ResponseWriter, w *http.Request) {

// }
