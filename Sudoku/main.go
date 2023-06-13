package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	solutions := true
	if !(len(os.Args) == 10) { //Vérifier la commande envoyé soit bien un sudoku !(len(os.Args)==10) car le sudoku a 9 colonnes et os.Args donne le chemin du fichier 1er argument donc 10 argument.
		AfficheError()
	} else {
		tableau := os.Args[1:]
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if len(tableau) != 9 {
					solutions = false
					break
				}
				if len(tableau[i]) != 9 {
					solutions = false
					break
				} else if !((tableau[i][j] >= 48 && tableau[i][j] <= 57) || tableau[i][j] == 46) {
					solutions = false
				}
			}
		}
		if solutions {
			grille := GrilleInt(os.Args[1:]) //créer la grille en appelant la fn GrilleInt
			grille2 := grille
			if ChercheValeurs(&grille) { // Appeler la fn ChercheValeurs avec la variable grille avec un pointer (&)
				if ChercheValeurs2(&grille2) {
					for i := 0; i < 9; i++ {
						for j := 0; j < 9; j++ {
							if !(grille[i][j] == grille2[i][j]) {
								solutions = false
								break
							}
						}
					}
				}
			}
			if solutions {
				AfficheGrille(grille)
			}
		} else {
			AfficheError()
		}
	}
}
func GrilleInt(argument []string) [9][9]int { //Entré : os.Args[1:] qui donne un tableau de string; sortie : un tableau d'un tableau d'entier
	tableau := [9][9]int{}
	for j := 0; j < 9; j++ { // Parcours éléments de argument
		for i := 0; i < 9; i++ { //Parcours éléments de argument
			if argument[j][i] == 46 { //Cherche dans la var argument(donné dans l'entré de la fn) les points
				tableau[j][i] = 0 // Ajoute un zéro à la position [j][i] dans le tableau
			} else {
				tableau[j][i] = int(argument[j][i] - 48) // Ajoute le chiffre de la position argument[j][i] à la position [j][i] dans le tableau
			}
		}
	}
	return tableau // Renvoie un tableau d'un tableau d'entier
}
func AfficheGrille(grille [9][9]int) { // Entré : variable grille
	for ligne := 0; ligne < 9; ligne++ { // Parcours des lignes de la grille
		for col := 0; col < 9; col++ { // Parcours des colonnes de la grille
			z01.PrintRune(rune((grille[ligne][col]) + 48))
			z01.PrintRune(rune(32))
		}
		z01.PrintRune('\n') // Retour à la ligne à chaque nouvelle ligne
	}
}
func AfficheError() {
	str := "Error"
	for i := 0; i < len(str); i++ { // Parcours des valeurs de la string str
		z01.PrintRune(rune(str[i]))
	}
}
func ChercheValeurs(grille *[9][9]int) bool { //Entré: variable grille avec pointer (*) afin d'utiliser la variable grille crée dans la fn main; Sortie : valeur booléenne
	if !CaseVide(grille) { // Vérifier si il y a encore des 0
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grille[i][j] == 0 { // Chercher les 0
				for valeur := 9; valeur >= 1; valeur-- { // Parcours les 9 possibilitées
					grille[i][j] = valeur    //Attribuer une possibilité au 0 trouvé au-dessus
					if VerifGrille(grille) { //Tester si la grille est bonne (pas de double)
						if ChercheValeurs(grille) { //Relancer la fn (qui réverifiera si il y a encore des 0)
							return true
						}
					}
					grille[i][j] = 0 //Si la fn n'a pas trouvé un bonne valeur elle remet un zéro à la place de la valeur possibilité et recommence
				}
				return false
			}
		}
	}
	return false
}
func ChercheValeurs2(grille2 *[9][9]int) bool { //Entré: variable grille avec pointer (*) afin d'utiliser la variable grille crée dans la fn main; Sortie : valeur booléenne
	if !CaseVide(grille2) { // Vérifier si il y a encore des 0
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grille2[i][j] == 0 { // Chercher les 0
				for valeur := 1; valeur <= 9; valeur++ { // Parcours les 9 possibilitées
					grille2[i][j] = valeur    //Attribuer une possibilité au 0 trouvé au-dessus
					if VerifGrille(grille2) { //Tester si la grille est bonne (pas de double)
						if ChercheValeurs2(grille2) { //Relancer la fn (qui réverifiera si il y a encore des 0)
							return true
						}
					}
					grille2[i][j] = 0 //Si la fn n'a pas trouvé un bonne valeur elle remet un zéro à la place de la valeur possibilité et recommence
				}
				return false
			}
		}
	}
	return false
}

func CaseVide(grille *[9][9]int) bool { //Entré: variable grille avec un pointer (*) afin d'utiliser la variable grille crée dans la fn main; Sortie : valeur booléenne
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grille[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func VerifGrille(grille *[9][9]int) bool { //Entré : variable grille avec un pointer (*) afin d'utiliser la variable grille crée dans la fn main; Retour : valeur booléenne
	//Vérifier les valeurs ligne par ligne
	for ligne := 0; ligne < 9; ligne++ {
		cpt := [10]int{} //Créer d'un tableau de 10 valeurs (10 car il y a 10possibilités : 0,1,2,3,4,5,6,7,8,9)
		for col := 0; col < 9; col++ {
			cpt[grille[ligne][col]]++ // Ajouter 1 à la valeur de l'index du chiffre (exemple : grille[0][1]=3 alors index=3, alors cpt=[0,0,0,1,0,0,0,0,0,0])
		}
		if Doubles(cpt) { // Regarder si la table cpt possède des doubles
			return false
		}
	}
	//Vérifier les valeurs colonne par colonne
	for col := 0; col < 9; col++ {
		cpt := [10]int{} //Créer d'un tableau de 10 valeurs (10 car il y a 10possibilités : 0,1,2,3,4,5,6,7,8,9)
		for ligne := 0; ligne < 9; ligne++ {
			cpt[grille[ligne][col]]++ // Ajouter 1 à la valeur de l'index du chiffre (exemple : grille[0][1]=3 alors index=3, alors cpt=[0,0,0,1,0,0,0,0,0,0])
		}
		if Doubles(cpt) { // Regarder la table cpt possède des doubles
			return false
		}
	}
	//Vérifier les valeurs par carré de 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			cpt := [10]int{} //Créer d'un tableau de 10 valeurs (10 car il y a 10possibilités : 0,1,2,3,4,5,6,7,8,9)
			for ligne := i; ligne < i+3; ligne++ {
				for col := j; col < j+3; col++ {
					cpt[grille[ligne][col]]++ // Ajouter 1 à la valeur de l'index du chiffre (exemple : grille[0][1]=3 alors index=3, alors cpt=[0,0,0,1,0,0,0,0,0,0])
				}
				if Doubles(cpt) { // Regarder la table cpt possède des doubles
					return false
				}
			}
		}
	}
	return true
}

func Doubles(cpt [10]int) bool { //Entré : la table cpt; Sortie : valeur booléene
	for i, nombre := range cpt { //Parcours la table cpt, avec i l'index, nombre la valeur à l'index i
		if i == 0 { // Si l'index est à 0 passé car c'est normal qu'il y ait encore des zéros dans la grille
			continue
		}
		if nombre > 1 { //Si la valeur est supérieure à 1 cela signifie qu'il y a plusieurs fois ce chiffre donc la grille est fausse
			return true
		}
	}
	return false
}
