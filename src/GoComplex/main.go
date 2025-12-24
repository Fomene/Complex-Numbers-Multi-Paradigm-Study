package main

import (
	"fmt"
)

func main() {

	// Lecture des données depuis l'entrée standard (stdin)
	// Format attendu :
	// a b
	// c d
	// (a + bi) et (c + di) sont les deux complexes

	var a, b, c, d float64

	// On lit les quatre nombres
	_, err1 := fmt.Scan(&a, &b)
	_, err2 := fmt.Scan(&c, &d)

	if err1 != nil || err2 != nil {
		fmt.Println("Erreur : impossible de lire les nombres depuis l'entrée.")
		return
	}

	z1 := NewComplex(a, b)
	z2 := NewComplex(c, d)

	fmt.Println("z1 =", z1)
	fmt.Println("z2 =", z2)

	fmt.Println("Addition       :", z1.Add(z2))
	fmt.Println("Soustraction   :", z1.Sub(z2))
	fmt.Println("Multiplication :", z1.Mul(z2))
	if z2.Re == 0 && z2.Im == 0 {
		fmt.Println("Division       : erreur (division par un nombre complexe nul)")
	} else {
		fmt.Println("Division       :", z1.Div(z2))
	}
	fmt.Println("Conjugué z1    :", z1.Conj())
	fmt.Println("Conjugué z2    :", z2.Conj())
	fmt.Println("Égalité z1=z2  :", z1.Equals(z2))
}
