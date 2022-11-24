package hangman

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
