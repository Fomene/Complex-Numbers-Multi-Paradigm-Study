mod complex;
use complex::Complex;
use std::io;

fn main() {
    // Lecture depuis l'entrée standard :
    // a b
    // c d
    let mut input = String::new();

    // Lecture du premier complexe
    io::stdin().read_line(&mut input).unwrap();
    let parts1: Vec<f64> = input
        .trim()
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();

    input.clear();

    // Lecture du second complexe
    io::stdin().read_line(&mut input).unwrap();
    let parts2: Vec<f64> = input
        .trim()
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();

    let z1 = Complex::new(parts1[0], parts1[1]);
    let z2 = Complex::new(parts2[0], parts2[1]);

    // Affichage identique à Go
    println!("z1 = {}", z1.format());
    println!("z2 = {}", z2.format());

    println!("Addition       : {}", (z1 + z2).format());
    println!("Soustraction   : {}", (z1 - z2).format());
    println!("Multiplication : {}", (z1 * z2).format());
    if z2.re == 0.0 && z2.im == 0.0 {
        println!("Division       : erreur (division par un nombre complexe nul)");
    } else {
        println!("Division       : {}", (z1 / z2).format());
    }
    println!("Conjugué z1    : {}", z1.conj().format());
    println!("Conjugué z2    : {}", z2.conj().format());
    println!("Égalité z1=z2  : {}", z1 == z2);
}
