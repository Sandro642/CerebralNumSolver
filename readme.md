# CerebralNumSolver 👨‍💻

CerebralNumSolver est une classe Go (Golang) développée par Sandro642, un développeur passionné et inventif. Il s'agit d'un solveur numérique intelligent qui offre une approche cérébrale pour résoudre des problèmes complexes impliquant des nombres.

## Fonctionnalités 🎯

- **Classe CerebralNumSolver** : Cette classe permet de stocker et de gérer les propositions de nombres soumises par l'utilisateur ou générées par le programme.

- **Constructeur NewCerebralNumSolver()** : Un constructeur est fourni pour initialiser un nouvel objet CerebralNumSolver et ainsi commencer une nouvelle résolution de problème.

- **Méthode AddGuess(guess int)** : Cette méthode permet d'ajouter un nombre à la liste des propositions enregistrées dans l'objet CerebralNumSolver.

- **Méthode CalculateTarget() int** : La méthode CalculateTarget() calcule le nombre cible en utilisant une fonction mathématique complexe basée sur les propositions enregistrées. Pour le moment, le calcul est effectué en trouvant la médiane des nombres proposés.

- **Fonction "median"** : La fonction median(numbers []int) int est utilisée par CalculateTarget() pour calculer la médiane d'une liste d'entiers. Elle utilise l'algorithme de tri rapide (QuickSort) pour simplifier le calcul.

## Comment ça marche ? 🤔

Le CerebralNumSolver est conçu pour résoudre des problèmes numériques en utilisant une approche cérébrale, ce qui signifie qu'il essaie d'adopter une stratégie intelligente basée sur les données fournies. L'utilisateur peut ajouter des propositions de nombres à l'objet CerebralNumSolver à l'aide de la méthode AddGuess(). Ensuite, la méthode CalculateTarget() utilise une fonction mathématique complexe pour déterminer le nombre cible à atteindre. Pour l'exemple actuel, la médiane des propositions est utilisée comme fonction mathématique.

## Comment importer les libs ? 🤔

Consulter les fichiers readme pour comprendre.

<a href="https://github.com/Sandro642/CerebralNumSolver/Version-GO/readme.md/">Version Go</a>
<a href="https://github.com/Sandro642/CerebralNumSolver/Version-Java/readme.md/">Version Java</a>
<a href="https://github.com/Sandro642/CerebralNumSolver/Version-JS/readme.md/">Version JS</a>

## Auteur et Contribution 💡

Ce projet a été entièrement développé par Sandro642, un développeur passionné et créatif. Il a conçu CerebralNumSolver comme une solution pour résoudre des problèmes numériques d'une manière intelligente et réfléchie.

## Prochainement

🔜 Une version Java de CerebralNumSolver est en cours de développement, afin de permettre aux développeurs Java de bénéficier de cette approche cérébrale pour résoudre leurs propres problèmes numériques.

🔜 Une version JavaScript de CerebralNumSolver est également prévue pour faciliter l'intégration du solveur dans des projets basés sur Node.js ou dans des applications Web.

🚀 N'hésitez pas à utiliser CerebralNumSolver dans vos projets en Go, et restez à l'affût des prochaines versions en Java et JavaScript pour résoudre des problèmes numériques de manière innovante !
