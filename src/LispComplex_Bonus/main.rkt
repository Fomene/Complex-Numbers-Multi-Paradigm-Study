#lang racket
(require "complex.rkt")

;; Lecture d'une ligne, transformation en liste de nombres
(define (read-complex-line)
  (let* ([line (read-line)]
         [parts (map string->number (string-split line))])
    (make-complex (first parts) (second parts))))

;; Lire les deux complexes
(define z1 (read-complex-line))
(define z2 (read-complex-line))

;; Affichage comme Go/Rust
(displayln (string-append "z1 = " (format-complex z1)))
(displayln (string-append "z2 = " (format-complex z2)))

(displayln (string-append "Addition       : " (format-complex (add z1 z2))))
(displayln (string-append "Soustraction   : " (format-complex (sub z1 z2))))
(displayln (string-append "Multiplication : " (format-complex (mul z1 z2))))

(if (equal? z2 (make-complex 0 0))
    (displayln "Division       : erreur (division par un nombre complexe nul)")
    (displayln (string-append "Division       : " (format-complex (divc z1 z2)))))

(displayln (string-append "Conjugué z1    : " (format-complex (conjc z1))))
(displayln (string-append "Conjugué z2    : " (format-complex (conjc z2))))
(displayln (string-append "Égalité z1=z2  : " (if (equal? z1 z2) "true" "false")))
