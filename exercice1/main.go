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

	// Exercice 1
	afficherEquipe(equipe)

	// Exercice 2
	plusDeVie := trouverPlusDeVie(equipe)
	plusDAttaque := trouverPlusDAttaque(equipe)
	vieMoyenne := calculerVieMoyenne(equipe)
	faibles := compterFaibles(equipe)

	fmt.Println("=== ANALYSE ===")
	fmt.Println()

	fmt.Println("Soldat avec le plus de vie :", plusDeVie.nom)
	fmt.Println("Vie :", plusDeVie.vie)
	fmt.Println()

	fmt.Println("Soldat avec la plus grande attaque :", plusDAttaque.nom)
	fmt.Println("Attaque :", plusDAttaque.attaque)
	fmt.Println()

	fmt.Printf("Vie moyenne :", vieMoyenne)
	fmt.Println()

	fmt.Println("Soldats avec moins de 800 PV :", faibles)

	// Exercice 3
	var degats int

	fmt.Println()
	fmt.Print("Dégâts de l'ennemi : ")
	fmt.Scanln(&degats)

	attaquerEquipe(&equipe, degats)

	// Exercice 4
	afficherEtat(equipe)

	// Exercice 5
	var nombreAttaques int

	fmt.Println()
	fmt.Println("=== BATAILLE ===")
	fmt.Println()

	fmt.Print("Nombre d'attaques ennemies : ")
	fmt.Scanln(&nombreAttaques)

	for i := 1; i <= nombreAttaques; i++ {
		fmt.Print("Attaque ", i, " : ")
		fmt.Scanln(&degats)

		attaquerEquipe(&equipe, degats)

		fmt.Println()
		fmt.Println("=== APRÈS L'ATTAQUE", i, "===")
		fmt.Println()

		afficherEtat(equipe)
	}
	// Exercice 6
	vivants := compterVivants(equipe, 0)

	fmt.Println()
	fmt.Println("Nombre de soldats vivants :", vivants)
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

func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	plusDeVie := equipe[0]

	for i := 1; i < 6; i++ {
		if equipe[i].vie > plusDeVie.vie {
			plusDeVie = equipe[i]
		}
	}

	return plusDeVie
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	plusDAttaque := equipe[0]

	for i := 1; i < 6; i++ {
		if equipe[i].attaque > plusDAttaque.attaque {
			plusDAttaque = equipe[i]
		}
	}

	return plusDAttaque
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	total := 0

	for i := 0; i < 6; i++ {
		total = total + equipe[i].vie
	}

	return float64(total) / 6
}

func compterFaibles(equipe [6]Soldat) int {
	faibles := 0

	for i := 0; i < 6; i++ {
		if equipe[i].vie < 800 {
			faibles++
		}
	}

	return faibles
}

func attaquerEquipe(equipe *[6]Soldat, degats int) {
	for i := 0; i < 6; i++ {
		if equipe[i].vie > 0 {
			equipe[i].vie = equipe[i].vie - degats

			if equipe[i].vie <= 0 {
				equipe[i].vie = 0
				fmt.Println(equipe[i].nom, "est KO !")
			}
		}
	}
}

func afficherEtat(equipe [6]Soldat) {
	fmt.Println()
	fmt.Println("=== ÉTAT DE L'ÉQUIPE ===")
	fmt.Println()

	for i := 0; i < 6; i++ {
		if equipe[i].vie > 0 {
			fmt.Println(equipe[i].nom, ":", equipe[i].vie, "PV")
		} else {
			fmt.Println(equipe[i].nom, ": KO")
		}
	}
}

func compterVivants(equipe [6]Soldat, index int) int {
	if index == 6 {
		return 0
	}

	if equipe[index].vie > 0 {
		return 1 + compterVivants(equipe, index+1)
	}

	return compterVivants(equipe, index+1)
}
