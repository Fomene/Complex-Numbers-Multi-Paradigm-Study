// Définition d'un nombre complexe sous forme de structure.
// re = partie réelle
// im = partie imaginaire
#[derive(Debug, Clone, Copy)]
pub struct Complex {
    pub re: f64,
    pub im: f64,
}

// Méthodes associées (style orienté objet de Rust).
impl Complex {
    // Constructeur
    pub fn new(re: f64, im: f64) -> Complex {
        Complex { re, im }
    }

    // Conjugué : a + bi -> a - bi
    pub fn conj(&self) -> Complex {
        Complex {
            re: self.re,
            im: -self.im,
        }
    }

    // Retourne une chaîne formatée : "a.b + c.di"
    pub fn format(&self) -> String {
        let re = if self.re.abs() < 1e-12 { 0.0 } else { self.re };
        let im = if self.im.abs() < 1e-12 { 0.0 } else { self.im };
        format!("{:.2} + {:.2}i", re, im)
    }
}


// Surcharge d'opérateurs


use std::ops::{Add, Sub, Mul, Div};

// Addition : (a+c, b+d)
impl Add for Complex {
    type Output = Complex;

    fn add(self, other: Complex) -> Complex {
        Complex {
            re: self.re + other.re,
            im: self.im + other.im,
        }
    }
}

// Soustraction
impl Sub for Complex {
    type Output = Complex;

    fn sub(self, other: Complex) -> Complex {
        Complex {
            re: self.re - other.re,
            im: self.im - other.im,
        }
    }
}

// Multiplication
impl Mul for Complex {
    type Output = Complex;

    fn mul(self, other: Complex) -> Complex {
        Complex {
            re: self.re * other.re - self.im * other.im,
            im: self.re * other.im + self.im * other.re,
        }
    }
}

// Division (on vérifie que le dénominateur n'est pas nul)
impl Div for Complex {
    type Output = Complex;

    fn div(self, other: Complex) -> Complex {
        let denom = other.re * other.re + other.im * other.im;

        if denom == 0.0 {
            panic!("Erreur : division par un nombre complexe nul");
        }

        Complex {
            re: (self.re * other.re + self.im *other.im) / denom,
            im: (self.im * other.re - self.re * other.im) / denom,
        }
    }
}

// Égalité : (a == c) et (b == d)
impl PartialEq for Complex {
    fn eq(&self, other: &Self) -> bool {
        self.re == other.re && self.im == other.im
    }
}
