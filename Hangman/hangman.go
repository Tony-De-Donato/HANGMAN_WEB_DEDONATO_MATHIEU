package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	//
	fmt.Println("Bienvenue dans notre jeu hangman.")
	menu()

}

//_______________________________________________________________________________________________________________________________________

func menu() {
	//affichage des différents choix, récupération de l'entrrée utilisateur et gestion de cette dernière
	fmt.Println("Que souhaitez vous faire ?")
	fmt.Println("1 : lancer une nouvelle partie")
	fmt.Println("2 : voir les règles")
	fmt.Println("3 ou autre : arrêter le terminal")
	var reponse int
	fmt.Scan(&reponse)
	switch reponse {
	case 1:
		choix_personnage()
	case 2:
		affichage_regle()
	}
}
func affichage_regle() {
	//affichage des règles du jeu, renvoie ensuite au menu
	fmt.Println("Le jeu du Hangman consiste à trouver un mot choisit aléatoirement par l'ordinateur.\nVous avez 10 tentatives pour trouver le mot.\nSi vous trouvez le mot avant d'avoir utilisé toutes vos tentatives, vous gagnez.\nSi vous n'avez plus de tentatives, vous perdez.\nBonne chance !!!")
	time.Sleep(6 * time.Second)
	menu()
}

// _______________________________________________________________________________________________________________________________________
func transforme_en_liste(fichier *os.File) []string {
	//transforme un fichier en liste, chaque ligne du fichier est un nouvel élément de la liste
	var liste []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		liste = append(liste, scanner.Text())
	}
	return liste
}

func liste_position(fichier *os.File) []string {
	// permet de créer une liste contenant toutes les positions présentes dans le fichier pris en paramètre
	var liste []string
	scanner := bufio.NewScanner(fichier)
	var stockage string
	var compteur int
	for scanner.Scan() {
		compteur++
		stockage += string(scanner.Text()) + "\n"
		if compteur == 8 {
			liste = append(liste, stockage)
			compteur = 0
			stockage = ""
		}
		// le code ci-dessus regroupe les lignes du fichier par 9 (la hauteur d'une position) et les ajoute à la liste
	}
	return liste
}
func choix_personnage() {
	file1, err1 := os.Open("words_1.txt")
	if err1 != nil {
		log.Fatal(err1)
	} //ouverture du premier fichier texte + gestion d'erreur
	file1_2, err1_2 := os.Open("words2_1.txt")
	if err1_2 != nil {
		log.Fatal(err1_2)
	} //ouverture du deuxième fichier texte + gestion d'erreur
	file1_3, err1_3 := os.Open("words3.txt")
	if err1_3 != nil {
		log.Fatal(err1_3)
	} //ouverture du troisième fichier texte + gestion d'erreur
	file1_liste := transforme_en_liste(file1)
	file2_liste := transforme_en_liste(file1_2)
	file3_liste := transforme_en_liste(file1_3)
	file1_liste = append(file1_liste, file2_liste...)
	file1_liste = append(file1_liste, file3_liste...)
	//regroupement des trois listes provenants des fichiers dans un seul
	file1.Close()
	file1_2.Close()
	file1_3.Close()
	//fermeture de l'accés trois fichier
	file2, err2 := os.Open("pos_hangman.txt")
	if err2 != nil {
		log.Fatal(err2)
	} // ouverture du fichier contenant les positions + gestion d'erreur
	var liste_des_positions []string
	liste_des_positions = liste_position(file2) // ajout de toutes les positions du hangman
	file2.Close()
	//
	fmt.Println("Tapez le nom de votre personnage : ")
	var nom string
	fmt.Scan(&nom) // récupération de l'entrée utilisateur pour nommer le hangman
	var personnage HangManData
	mot := nouveau_mot(file1_liste)                                          // récupération d'un mot aléatoire dans la liste des trois fichiers textes
	personnage.Init(nom, mot, word_with_blank(mot), 10, liste_des_positions) // initialisation de la stucture contenant les informations relatives au jeu
	lancement_jeu(personnage)
}

func lancement_jeu(h HangManData) {

	for h.Attempts > 0 { // la boucle tourne tant qu'il reste des essais au joueur
		if h.Word == h.ToFind {
			h.Victoire()
		} else {
			h.jouer_tour()
		}
	}
	if h.Attempts == 0 {
		h.perdu()
	}
}

func (h *HangManData) jouer_tour() {

	fmt.Println("\nIl vous reste", h.Attempts, "essais.")
	fmt.Println("Voici la pose actuelle de", h.nom, ":\n ")
	fmt.Println(h.ActualPosition)
	fmt.Println("Voici les lettres que vous avez déja essayées :", h.UsedLetter)
	fmt.Println("Voici ce que vous avez trouvé du mot :", h.Word)
	fmt.Println("\nProposez une lettre (coute un essai) ou un mot (coute deux essais) :")
	var lettre string
	fmt.Scan(&lettre)
	if len(lettre) == 1 { // si le joueur propose une lettre
		h.AjoutLettre(lettre)      // elle est ajoutée à la liste des lettres déja utilisées si elle n'y est pas déjà
		if h.verifletter(lettre) { // si la lettre est présente dans le mot
			fmt.Println(" \nCette lettre fait bien partie du mot, bravo !")
			h.remplace(lettre) // elle remplace les blancs dans le mot montré au joueur
		} else {
			fmt.Println("Vous vous êtes trompé.")
			h.Attempts -= 1 //le nombre d'essais baisse de 1
			if h.Attempts != 0 {
				h.ActualPosition = h.HangmanPositions[10-h.Attempts] // mise à jour de la position du pendu
			}
		}
	} else if len(lettre) > 1 { // si le joueur propose un mot
		if lettre == h.ToFind { // si c'est le bon mot
			h.Word = lettre // le mot que voit le joueur est remplacé par la proposition (permet de faire gagner le joueur à la fin du tour)
		} else { //si ce n'est pas le bon mot
			fmt.Println("Ce n'est pas le bon mot, vous perdez deux essais.")
			h.Attempts -= 2 // le nombre d'essais diminue de 2
			if h.Attempts < 0 {
				h.Attempts = 0 // permet de gérer le fait que le nombre d'essais descende en dessous de 0
			}
			if h.Attempts != 0 {
				h.ActualPosition = h.HangmanPositions[10-h.Attempts] //mise à jour de la position du pendu
			}
		}
	}
	time.Sleep(2 * time.Second)
	fmt.Println("______________________________________________________________")
}

func (h HangManData) perdu() {
	// affichage ne cas de défaite, le mot est dévoilé au joueur puis ce dernier est renvoyé au menu de relance
	fmt.Println("\nVous avez perdu !!!")
	fmt.Println("Le mot était", h.ToFind)
	time.Sleep(3 * time.Second)
	relance()
}

//_________________________________________________________________________________________________________________________________________

type HangManData struct {
	nom              string   // nom du Hangman
	Word             string   // mot composé de '_', ex: H_ll_
	ToFind           string   // mot à trouver
	Attempts         int      // nombre de tentatives restantes
	HangmanPositions []string // liste contenant les position de "pos_hangman.txt"
	ActualPosition   string   // position actuelle du Hangman
	UsedLetter       []string // liste des lettre déjà proposées par l'utilisateur
}

func (h *HangManData) Init(nom string, a_trouver string, mot_actuel string, tentatives int, liste_pose []string) {
	// initialisation du hangman
	h.nom = nom
	h.ToFind = a_trouver
	h.Word = mot_actuel
	h.Attempts = tentatives
	h.HangmanPositions = liste_pose
	h.ActualPosition = liste_pose[0]
}
func word_with_blank(mot string) string {
	//prend en paramètre un mot et le renvoie après avoir remplacé un certain nombre de lettre par des "_"
	var liste []string
	for _, element := range mot {
		liste = append(liste, string(element)) //transforme le mot en liste de byte
	}
	n := len(mot)/2 - 1
	var nouveau_mot string
	i := 0
	for i < len(mot)-n {
		index := random(len(mot) - 1) // choisit une lettre du mot aléatoirement
		if liste[index] != "_" {      // si elle n'a pas déjà été remplacée, elle l'est
			liste[index] = "_"
			i++ // augmentation du compteur
		}
	}
	for _, element := range liste {
		nouveau_mot += string(element) // création du nouveau mot
	}
	return nouveau_mot
}
func nouveau_mot(fichier []string) string {
	// permet de choisir aléatoirement un mot dans une liste de mots
	random := random(len(fichier))
	mot := fichier[random]
	return mot
}
func random(i int) int {
	// renvoie un entier pseudo-aléatoire entre 0 et i
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(i)
	return random
}

//_________________________________________________________________________________________________________________________________________

func (h *HangManData) verifletter(letter string) bool {
	// vérifie que la lettre soit dans le nom
	for _, i := range h.ToFind {
		if letter == string(i) {
			return true
		}
	}
	return false
}

func (h *HangManData) DejaDansMot(letter string) bool {
	// vérifie si la lettre est déja dans le mot proposé, non utilisé
	retour := false
	for _, i := range h.Word {
		if string(i) == letter {
			retour = true
			break
		}
	}
	return retour
}

func (h *HangManData) remplace(lettre string) {
	// remplace les blancs dans le mot par la lettre entrée si elle correspond
	var nouveau_mot string
	for i := 0; i < len(h.ToFind); i++ { // parcourt les lettres du mot
		if string(h.ToFind[i]) == lettre { //compare la lettre entrée à celle présente dans le mot à trouver
			nouveau_mot += lettre // si elle correspond, elle est ajoutée au nouveau mot
		} else {
			nouveau_mot += string(h.Word[i]) // si elle ne correspond pas, c'est la lettre déjà présente qui est ajoutée
		}
	}
	h.Word = nouveau_mot // remplacement du mot
}

func (h *HangManData) LettreUtilise(letter string) bool {
	// vérifie si la lettre a déjà été entrée
	for _, i := range h.UsedLetter {
		if i == letter {
			return true
		}
		break
	}
	return false
}

func (h *HangManData) AjoutLettre(lettre string) {
	// ajoute la lettre à la liste des lettres déjà utilisées
	if h.LettreUtilise(lettre) == true {
		fmt.Println("Cette lettre à déjà été utilisée.")

	} else if h.LettreUtilise(lettre) == false {
		h.UsedLetter = append(h.UsedLetter, lettre)
	}
}
func relance() {
	// affichage du menu de relance + gestion de l'entrée utilisateur
	fmt.Println("\nQue souhaitez vous faire ?\n1 : relancer une partie\n2 : Quitter")
	var rep int
	fmt.Scan(&rep)
	switch rep {
	case 1:
		choix_personnage()
	case 2:
		Quit()
	}
}

func (h *HangManData) Victoire() {
	if h.ToFind == h.Word { // vérifie si le mot est trouvé
		fmt.Printf("\nVous avez trouvé le mot %s, félicitations.", h.ToFind) //affichage du mot
		relance()                                                            // renvoie au menu de relance
	}
}

func Quit() {
	// permet de quitter
	os.Exit(0)
}
