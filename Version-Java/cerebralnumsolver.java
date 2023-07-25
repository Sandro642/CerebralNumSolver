import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

// Classe CerebralNumSolver
public class CerebralNumSolver {
    private List<Integer> guesses;

    // Constructeur de la classe CerebralNumSolver
    public CerebralNumSolver() {
        guesses = new ArrayList<>();
    }

    // Méthode pour ajouter un nombre à la liste des propositions
    public void addGuess(int guess) {
        guesses.add(guess);
    }

    // Méthode pour calculer le nombre cible en utilisant une fonction mathématique complexe
    public int calculateTarget() {
        // Exemple d'une fonction mathématique complexe (ici : médiane des propositions)
        return median(guesses);
    }

    // Fonction "median" pour calculer la médiane d'une liste d'entiers
    private int median(List<Integer> numbers) {
        // Ici, vous pouvez implémenter votre propre algorithme de calcul de médiane
        // en triant la liste et en trouvant l'élément du milieu.
        // Pour simplifier, nous utilisons la méthode de tri de Collections pour cet exemple.
        Collections.sort(numbers);

        // Maintenant, nous retournons simplement l'élément du milieu
        return numbers.get(numbers.size() / 2);
    }