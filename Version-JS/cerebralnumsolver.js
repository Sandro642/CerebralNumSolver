// Classe CerebralNumSolver
class CerebralNumSolver {
    constructor() {
      this.guesses = [];
    }
  
    // Méthode pour ajouter un nombre à la liste des propositions
    addGuess(guess) {
      this.guesses.push(guess);
    }
  
    // Méthode pour calculer le nombre cible en utilisant une fonction mathématique complexe
    calculateTarget() {
      // Exemple d'une fonction mathématique complexe (ici : médiane des propositions)
      return this.median(this.guesses);
    }
  
    // Fonction "median" pour calculer la médiane d'une liste de nombres
    median(numbers) {
      // Ici, vous pouvez implémenter votre propre algorithme de calcul de médiane
      // en triant la liste et en trouvant l'élément du milieu.
      // Pour simplifier, nous utilisons une fonction de tri rapide (QuickSort) pour cet exemple.
      this.quickSort(numbers, 0, numbers.length - 1);
  
      // Maintenant, nous retournons simplement l'élément du milieu
      return numbers[Math.floor(numbers.length / 2)];
    }
  
    // Fonction de tri rapide (QuickSort)
    quickSort(arr, low, high) {
      if (low < high) {
        const pivot = this.partition(arr, low, high);
        this.quickSort(arr, low, pivot - 1);
        this.quickSort(arr, pivot + 1, high);
      }
    }
  
    partition(arr, low, high) {
      const pivot = arr[high];
      let i = low - 1;
      for (let j = low; j < high; j++) {
        if (arr[j] < pivot) {
          i++;
          [arr[i], arr[j]] = [arr[j], arr[i]];
        }
      }
      [arr[i + 1], arr[high]] = [arr[high], arr[i + 1]];
      return i + 1;
    }
  }