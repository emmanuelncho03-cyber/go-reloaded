# Go Reloaded

Ce programme lit un fichier texte, applique les règles du projet, puis écrit le résultat dans un autre fichier. Il utilise uniquement la bibliothèque standard de Go.

## Ouvrir le projet

Ouvre le dossier du projet dans VS Code, puis ouvre un terminal à la racine du dossier.

## Exemple simple

```sh
go run . examples/input.txt resultat.txt
```

`examples/input.txt` contient le texte à transformer. Le résultat attendu est dans `examples/output.txt`. Le programme crée `resultat.txt` au lancement ; ce fichier généré n'a pas besoin d'être ajouté au dépôt.

## Exemple complet

`examples/input-demo-complet.txt` montre les principales transformations. Son résultat attendu est fourni dans `resultat-demo-complet.txt`.

```sh
go run . examples/input-demo-complet.txt resultat.txt
```

Le programme prend en charge `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, les nombres de mots, la ponctuation, les apostrophes et la règle `a` vers `an`.

## Compiler et vérifier

```sh
go build .
go test ./...
go vet ./...
```
