package main

import "fmt"

type Soldat struct {
	nom     string
	vie     int
	attaque int
}

func main() {
	equipe := [6]Soldat{
		{"Arthas", 1200, 250},
		{"Kael", 850, 320},
		{"Thrall", 1500, 180},
		{"Sylvanas", 700, 400},
		{"Garrosh", 1000, 280},
		{"Jaina", 500, 450},
	}
	//ex1
	afficherEquipe(equipe)
	//ex2
	analyse(equipe [6]Soldat)
}

func afficherEquipe(equipe [6]Soldat) {
	fmt.Println("=== EQUIPE ===")
	fmt.Println()

	for i := 0; i < 6; i++ {
		fmt.Println(equipe[i].nom)
		fmt.Println("Vie :", equipe[i].vie)
		fmt.Println("Attaque :", equipe[i].attaque)
		fmt.Println()
	}
}

func analyse(equipe [6]Soldat) {
	plusDeVie := trouverPlusDeVie(equipe)
}

func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	plusdevie := equipe[0]

	for i := 1; i < 6; i++ {
		if equipe[i].vie > plusdevie.vie {
			plusdevie = equipe[i]
		}
	}
	return plusdevie
}

/*func trouverPlusDAttaque(equipe [6]Soldat) Soldat

func calculerVieMoyenne(equipe [6]Soldat) float64

func compterFaibles(equipe [6]Soldat) int*/
