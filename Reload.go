package hangman

func Reload(je *Hangman) {
	je.Cpt = 0
	je.Lettre = []string{}
	je.Pos = Hang(je.Cpt, je.Pendu)
	je.Affich = []string{}
	je.Wrong = ""
}