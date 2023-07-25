# Exemple d'importation de CerebralNumSolver dans un projet JavaScript 📥

Créez un nouveau fichier votre_script_principal.js dans le même répertoire que index.html et ajoutez le code suivant :

```
// Importez la classe CerebralNumSolver depuis le fichier CerebralNumSolver.js
import CerebralNumSolver from './CerebralNumSolver.js';

// Exemple d'utilisation
const solver = new CerebralNumSolver();
solver.addGuess(3);
solver.addGuess(7);
solver.addGuess(1);
const target = solver.calculateTarget();
console.log(`Le CerebralNumSolver calcule le nombre cible : ${target}`);
```