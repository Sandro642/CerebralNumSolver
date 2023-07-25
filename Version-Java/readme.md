# Exemple d'importation de CerebralNumSolver en Java 📥

<a href="https://github.com/Sandro642/CerebralNumSolver/releases/download/1.0.0/CNSAlgorithm.jar">Installer le jar.</a>

```
public class Main {
    public static void main(String[] args) {
        // Créez un nouvel objet CerebralNumSolver
        CerebralNumSolver solver = new CerebralNumSolver();

        // Ajoutez des propositions de nombres à résoudre
        solver.addGuess(3);
        solver.addGuess(7);
        solver.addGuess(1);

        // Calculez le nombre cible en utilisant la méthode calculateTarget()
        int target = solver.calculateTarget();

        // Affichez le résultat
        System.out.println("Le CerebralNumSolver calcule le nombre cible : " + target);
    }
}
```
