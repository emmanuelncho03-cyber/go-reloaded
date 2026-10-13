# Go Reloaded

Ce programme lit un fichier texte, applique les règles du projet Go Reloaded, puis écrit le texte corrigé dans un autre fichier. Il utilise uniquement les outils intégrés à Go.

## Ouvrir le projet

Télécharger le dossier du projet, puis l’ouvrir dans VS Code. Ouvrir le terminal dans ce dossier.

## Compiler

```sh
go build .
```

## Lancer le programme

```sh
go run . examples/input.txt resultat.txt
```

Le premier chemin désigne le fichier à corriger. Le second désigne le fichier où enregistrer le résultat. L’exemple fourni est dans `examples/input.txt` ; le résultat attendu est dans `examples/output.txt`.

Le programme prend en charge les conversions `(hex)` et `(bin)`, les commandes `(up)`, `(low)` et `(cap)` avec ou sans nombre, la ponctuation, les apostrophes et la règle `a` vers `an`.

## Vérifier avec les tests

```sh
go test ./...
```

Pour vérifier aussi la qualité du code :

```sh
go vet ./...
```
