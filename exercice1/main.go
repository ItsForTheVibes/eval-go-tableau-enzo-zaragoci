package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	fmt.Printf("Vie moyenne : %.2f\n", vieMoyenne)
	fmt.Println()

	fmt.Println("Soldats avec moins de 800 PV :", faibles)

	// Exercice 3
	fmt.Println()
	degats := demanderNombre("Degats de l'ennemi : ")

	attaquerEquipe(&equipe, degats)

	// Exercice 4
	afficherEtat(equipe)

	// Exercice 5
	fmt.Println()
	fmt.Println("=== BATAILLE ===")
	fmt.Println()

	nombreAttaques := demanderNombre("Nombre d'attaques ennemies : ")

	for i := 1; i <= nombreAttaques; i++ {
		fmt.Println()
		degats = demanderNombre("Attaque " + strconv.Itoa(i) + " : ")

		attaquerEquipe(&equipe, degats)

		fmt.Println()
		fmt.Println("=== APRES L'ATTAQUE", i, "===")

		afficherEtat(equipe)
	}

	// Exercice 6
	vivants := compterVivants(equipe, 0)

	// Exercice 7
	fmt.Println()
	fmt.Println("=== FIN DE LA BATAILLE ===")
	fmt.Println()

	fmt.Println("Nombre de soldats vivants :", vivants)
	fmt.Println("Nombre de soldats KO :", 6-vivants)
	fmt.Println()

	if peutContinuer(equipe) {
		fmt.Println("L'equipe peut continuer le combat!, bonne chance!")
	} else {
		fmt.Println("Tous les soldats sont KO, RIP, try again.")
		fmt.Println("La bataille est terminee!")
	}
}

// Permet de demander un nombre valide
func demanderNombre(message string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(message)

		texte, _ := reader.ReadString('\n')
		texte = strings.TrimSpace(texte)

		nombre, erreur := strconv.Atoi(texte)

		if erreur != nil {
			fmt.Println("Erreur: Et bah non, faut rentrer un nombre entier!")
			continue
		}

		if nombre < 0 {
			fmt.Println("Erreur: tu peux arreter d'essayer de casser mon code merci")
			continue
		}

		return nombre
	}
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
	fmt.Println("=== ETAT DE L'EQUIPE ===")
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
	if index >= 6 {
		return 0
	}

	if equipe[index].vie > 0 {
		return 1 + compterVivants(equipe, index+1)
	}

	return compterVivants(equipe, index+1)
}

func peutContinuer(equipe [6]Soldat) bool {
	for i := 0; i < 6; i++ {
		if equipe[i].vie > 0 {
			return true
		}
	}

	return false
}
