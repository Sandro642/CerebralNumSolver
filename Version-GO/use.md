# Exemple d'utilisation de CerebralNumSolver

Le CerebralNumSolver est un outil puissant pour résoudre des problèmes numériques complexes de manière intelligente. Voici comment vous pouvez l'utiliser dans votre propre projet :

```go
// Importez le package contenant la classe CerebralNumSolver
package main

import (
	"fmt"
)

func main() {
	// Créez un nouvel objet CerebralNumSolver
	solver := NewCerebralNumSolver()

	// Ajoutez des propositions de nombres à résoudre
	solver.AddGuess(10)
	solver.AddGuess(25)
	solver.AddGuess(14)
	solver.AddGuess(8)
	solver.AddGuess(19)

	// Calculez le nombre cible en utilisant la méthode CalculateTarget
	target := solver.CalculateTarget()

	// Affichez le résultat
	fmt.Printf("Le nombre cible est : %d\n", target)
}
```

Dans cet exemple, nous importons le package contenant la classe CerebralNumSolver. Ensuite, nous créons un nouvel objet CerebralNumSolver à l'aide de la fonction NewCerebralNumSolver(). Ensuite, nous ajoutons des propositions de nombres à l'objet solver à l'aide de la méthode AddGuess(). Enfin, nous utilisons la méthode CalculateTarget() pour calculer le nombre cible basé sur les propositions fournies, et nous affichons le résultat.

N'hésitez pas à essayer cet exemple dans votre environnement Go pour découvrir le fonctionnement de CerebralNumSolver et voir comment il peut vous aider à résoudre des problèmes numériques de manière intelligente !
