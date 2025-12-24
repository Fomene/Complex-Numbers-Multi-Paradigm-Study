# Projet : Nombres complexes en Go, Rust et Lisp (Racket)

Implantation des opérations de base sur les nombres complexes (addition, soustraction, multiplication, division, conjugue, égalité) dans trois langages/paradigmes. Les entrées et sorties ont le même format dans les trois versions. Les tests sont regroupés dans `src/tests/` (`test1.in` ... `test6.in` et leurs `testX.out`).

## Arborescence
```
Complex-Numbers-Multi-Paradigm-Study/
├── README.md
├── rapport/
│   └── main.tex
└── src/
    ├── GoComplex/
    ├── RustComplex/
    ├── LispComplex_Bonus/
    └── tests/
        ├── test1.in / test1.out
        ├── test2.in / test2.out
        └── ... (jusqu'a test6)
```

## Go (procédural)
Téléchargement : https://go.dev/dl/  
VS Code : extension "Go" (Go Team)

Compilation :
```
cd src/GoComplex
go build -o complex-go.exe .
```
Éxécution avec un test :
```
go run . < ../tests/testX.in
```
Comparer la sortie avec `../tests/testX.out`.

## Rust (POO)
Téléchargement (rustup) : https://www.rust-lang.org/tools/install  
VS Code : extension "rust-analyzer"

Compilation (release) :
```
cd src/RustComplex
cargo build --release
```
Éxécution avec un test :
```
cargo run --quiet < ../tests/testX.in
```
ou, après build :
```
./target/release/rustcomplex < ../tests/testX.in
```
Comparer la sortie avec `../tests/testX.out`.

## Scheme / Racket (fonctionnel)
Téléchargement : https://racket-lang.org/download/  
VS Code : "Magic Racket"

Éxécution :
```
cd src/LispComplex_Bonus
racket main.rkt < ../tests/testX.in
```
Comparer la sortie avec `../tests/testX.out`.

## Notes
- Division par zéro : message d'erreur puis poursuite du programme.
- Affichage : arrondi à deux décimales ; les valeurs très petites sont forcées à `0.00` pour éviter `-0.00`.
