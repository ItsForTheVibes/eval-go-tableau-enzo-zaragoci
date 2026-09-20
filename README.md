Évaluation Go

Zaragoci

Enzo

Exercices

Exercice1 :

Dans notre fonction afficherEquipe cela se fait en deux parties

Afficher la banniere de lequipe pour separer les exercices
Et chaque personnage qui saffiche avec leur nom vie et attaque.

Exercice2:

Dans le deuxieme exercice on veut creer une fonction Analyse ou on rassemblera ces fonctions suivantes:

trouverPlusDeVie: Cette fonction nous permet de voir quel soldat possede le plus de pv
trouverPlusDAttaque: Cette fonction permet de retrouver quel soldat possede le plus dattaque
calculerVieMoyenne: cette fonction permet de retrouver la moyenne de pv de lequipe
compterFaibles: Cette fonction permet de compter le nombre de soldats faibles dans lequipe (pv <800)

Exercice3:

Dans le troisieme exercice on ajoute une attaque ennemie sur toute lequipe.

On demande a lutilisateur combien de degats lennemi va faire
attaquerEquipe: Cette fonction retire les degats a chaque soldat encore vivant
Si la vie dun soldat descend en dessous de 0 on la remet a 0
Si le soldat arrive a 0 pv il est considere comme KO et ne perdra plus de vie

On utilise aussi un pointeur sur equipe pour pouvoir modifier directement les pv des soldats.

Exercice4:

Dans le quatrieme exercice on ajoute la fonction afficherEtat.

afficherEtat: Cette fonction permet dafficher la vie actuelle de chaque soldat
Si le soldat a plus de 0 pv on affiche son nombre de pv
Si le soldat a 0 pv on affiche KO a la place

Cela nous permet de voir letat de lequipe apres une attaque.

Exercice5:

Dans le cinquieme exercice on cree le deroulement de la bataille.

On demande dabord combien dattaques lennemi va effectuer
Une boucle for permet de repeter les attaques le nombre de fois demande
Pour chaque attaque on demande les degats
On utilise attaquerEquipe pour enlever les pv
Ensuite on utilise afficherEtat pour voir letat de lequipe apres chaque attaque

Exercice6:

Dans le sixieme exercice on ajoute la fonction compterVivants.

compterVivants: Cette fonction permet de compter combien de soldats sont encore vivants
Un soldat est vivant si ses pv sont superieurs a 0
Cette fonction utilise la recursivite au lieu dune boucle for
La fonction se rappelle elle meme en augmentant lindex pour verifier le soldat suivant

Quand lindex arrive a 6 cela veut dire que tous les soldats ont ete verifies et la fonction sarrete.

Exercice7:

Dans le dernier exercice on ajoute la fonction peutContinuer.

peutContinuer: Cette fonction verifie si au moins un soldat est encore vivant
Si elle trouve un soldat avec plus de 0 pv elle retourne true
Si tous les soldats sont KO elle retourne false

A la fin on affiche le nombre de soldats vivants et le nombre de soldats KO.

Si au moins un soldat est vivant lequipe peut continuer le combat, sinon la bataille est terminee.