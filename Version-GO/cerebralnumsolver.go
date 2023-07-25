package main

// Classe CerebralNumSolver
type CerebralNumSolver struct {
	guesses []int // Liste pour stocker les nombres proposés
}

// Constructeur de la classe CerebralNumSolver
func NewCerebralNumSolver() *CerebralNumSolver {
	return &CerebralNumSolver{}
}

// Méthode pour ajouter un nombre à la liste des propositions
func (c *CerebralNumSolver) AddGuess(guess int) {
	c.guesses = append(c.guesses, guess)
}

// Méthode pour calculer le nombre cible en utilisant une fonction mathématique complexe
func (c *CerebralNumSolver) CalculateTarget() int {
	// Exemple d'une fonction mathématique complexe (ici : médiane des propositions)
	return c.median(c.guesses)
}

// Fonction "median" pour calculer la médiane d'une liste d'entiers
func (c *CerebralNumSolver) median(numbers []int) int {
	// Ici, vous pouvez implémenter votre propre algorithme de calcul de médiane
	// en triant la liste et en trouvant l'élément du milieu.
	// Pour simplifier, nous utilisons une fonction de tri rapide (QuickSort) de Go
	// pour cet exemple.
	quickSort(numbers, 0, len(numbers)-1)

	// Maintenant, nous retournons simplement l'élément du milieu
	return numbers[len(numbers)/2]
}

// Fonction de tri rapide (QuickSort)
func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
