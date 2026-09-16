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
	afficherEquipe(equipe)
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
