// Implémentation en Go selon un paradigme procédural.
// Les opérations sont définies comme des méthodes explicites
// car Go ne permet pas la surcharge d'opérateurs.

package main

import (
	"fmt"
	"math"
)

// Structure représentant un nombre complexe.
// re = partie réelle
// im = partie imaginaire
type Complex struct {
	Re, Im float64
}

// Fonction de création d'un nombre complexe.
// Utilisée comme "constructeur".
func NewComplex(re, im float64) Complex {
	return Complex{Re: re, Im: im}
}

// Addition de deux nombres complexes.
// (a + bi) + (c + di) = (a+c) + (b+d)i
func (z Complex) Add(other Complex) Complex {
	return Complex{
		Re: z.Re + other.Re,
		Im: z.Im + other.Im,
	}
}

// Soustraction :
// (a + bi) - (c + di) = (a-c) + (b-d)i
func (z Complex) Sub(other Complex) Complex {
	return Complex{
		Re: z.Re - other.Re,
		Im: z.Im - other.Im,
	}
}

// Multiplication :
// (a + bi)(c + di) = (ac - bd) + (ad + bc)i
func (z Complex) Mul(other Complex) Complex {
	return Complex{
		Re: z.Re*other.Re - z.Im*other.Im,
		Im: z.Re*other.Im + z.Im*other.Re,
	}
}

// Division :
// (a + bi) / (c + di) = [(ac + bd) + (bc - ad)i] / (c² + d²)
// On vérifie seulement que le dénominateur n'est pas nul.
func (z Complex) Div(other Complex) Complex {
	denom := other.Re*other.Re + other.Im*other.Im
	if denom == 0 {
		// J'utilise panic ici car Go n'a pas d'exceptions classiques.
		// Pour le projet, ça suffit largement.
		panic("Erreur : division par un nombre complexe nul")
	}

	return Complex{
		Re: (z.Re*other.Re + z.Im*other.Im) / denom,
		Im: (z.Im*other.Re - z.Re*other.Im) / denom,
	}
}

// Conjugué :
// conj(a + bi) = a - bi
func (z Complex) Conj() Complex {
	return Complex{Re: z.Re, Im: -z.Im}
}

// Vérifie si deux complexes sont égaux.
// On compare simplement les deux parties.
func (z Complex) Equals(other Complex) bool {
	return z.Re == other.Re && z.Im == other.Im
}

// Méthode pour afficher un complexe sous une forme lisible.
// Format : "a + bi"
func (z Complex) String() string {
	re := z.Re
	im := z.Im
	if math.Abs(re) < 1e-12 {
		re = 0
	}
	if math.Abs(im) < 1e-12 {
		im = 0
	}

	return fmt.Sprintf("%.2f + %.2fi", re, im)
}
