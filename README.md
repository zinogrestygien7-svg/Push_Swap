# push-swap

Projet en Go : trier une pile d'entiers avec **deux piles** (`a` et `b`) et un jeu restreint d'instructions, en produisant le moins d'instructions possible.

Le projet fournit deux programmes :

- **`push-swap`** : reçoit des entiers en argument et affiche, une instruction par ligne, une suite d'instructions qui trie la pile `a`.
- **`checker`** : reçoit des entiers en argument, lit des instructions sur l'entrée standard, les exécute, puis affiche `OK` si `a` est triée et `b` vide, sinon `KO`.

Seule la bibliothèque standard de Go est utilisée.

## Compilation

```bash
go build -o push-swap ./pushswap
go build -o checker ./checker
```

Sous Windows, ajouter `.exe` aux noms de sortie (`push-swap.exe`, `checker.exe`).

## Utilisation

```bash
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

Les nombres peuvent être donnés dans un seul argument (`"2 1 3"`) ou dans plusieurs (`2 1 3`). Le premier nombre se trouve au sommet de la pile `a`.

### Comportement

| Situation | `push-swap` | `checker` |
|---|---|---|
| Aucun argument | n'affiche rien | n'affiche rien |
| Argument non entier, doublon | `Error` sur la sortie d'erreur | `Error` sur la sortie d'erreur |
| Instruction inconnue ou mal écrite | non applicable | `Error` sur la sortie d'erreur |
| Pile déjà triée | n'affiche rien | `OK` |
| Pile triée après exécution, `b` vide | non applicable | `OK` |
| Autre cas | non applicable | `KO` |

## Instructions

| Instruction | Effet |
|---|---|
| `pa` | prend le sommet de `b` et le place au sommet de `a` |
| `pb` | prend le sommet de `a` et le place au sommet de `b` |
| `sa` / `sb` / `ss` | échange les deux premiers éléments de `a` / `b` / des deux |
| `ra` / `rb` / `rr` | le premier élément devient le dernier |
| `rra` / `rrb` / `rrr` | le dernier élément devient le premier |

## Organisation du code

```
.
├── go.mod
├── stack/       Pile, instructions de base, Machine (a + b), Apply, IsSorted
├── input/       Lecture et validation des arguments (entiers, doublons)
├── sorter/      Algorithmes de tri et nettoyage des instructions
├── checker/     Programme checker
└── pushswap/    Programme push-swap
```

- **`stack`** : le type `Stack` (l'index 0 est le sommet) avec ses quatre mouvements de base (échange, rotation, rotation inverse, transfert), et `Machine` qui contient les piles `a` et `b` et traduit chaque instruction du sujet. `Apply` exécute une instruction donnée sous forme de texte et renvoie une erreur si elle n'existe pas.
- **`input`** : `ParseArgs` transforme les arguments en pile et refuse les valeurs non entières et les doublons.
- **`sorter`** : choisit un algorithme selon la taille de la pile (voir ci-dessous).

## Algorithmes

Le choix dépend du nombre d'éléments :

1. **Pile déjà triée** : aucune instruction.
2. **Petites piles (BFS)** : recherche en largeur sur tous les états possibles des deux piles, par nombre d'instructions croissant. Le premier état trié rencontré correspond au plus court programme possible : le résultat est **optimal**. Réservé aux très petites piles, car le nombre d'états explose avec la taille.
3. **Piles moyennes** : on envoie les plus petits éléments dans `b` (en les remontant par le chemin le plus court), on trie les trois restants, puis on ramène tout avec `pa`.
4. **Grandes piles (chunks)** : chaque valeur est remplacée par son rang (ce qui gère les nombres négatifs et les grands écarts). On envoie les éléments dans `b` par tranches de rangs successives, puis on ramène à chaque fois le plus grand élément de `b` vers `a` par le chemin le plus court (`rb` ou `rrb`).

Les seuils (`bfsLimit`, `smallLimit`) et la taille des tranches (`chunkSize`) sont réglés par mesure, avec des tests dédiés (`TestTuneSmallLimit`, `TestTuneChunkSize`).

### Nettoyage des instructions

`Optimize` repasse sur la liste produite :

- deux instructions qui s'annulent sont supprimées (`ra` `rra`, `pb` `pa`, `sa` `sa`...) ;
- deux instructions indépendantes sont fusionnées (`ra` `rb` devient `rr`, `rra` `rrb` devient `rrr`, `sa` `sb` devient `ss`).

## Résultats mesurés

| Pile | Instructions |
|---|---|
| exemple du sujet (`4 67 3 87 23`) | 6 |
| 100 nombres aléatoires | environ 550 |
| 500 nombres aléatoires | environ 4 900 |

## Tests

```bash
go test ./...
```

Les tests couvrent les instructions de base, la lecture des arguments, et le tri : toutes les permutations de 2 à 5 éléments, des piles aléatoires jusqu'à 500 éléments (avec nombres négatifs), et le nettoyage des instructions. Chaque test rejoue les instructions produites sur une copie de la pile, comme le fait `checker`, pour vérifier que le résultat est bien trié.

Pour voir les mesures de réglage :

```bash
go test -v -run TestTune ./sorter
```
