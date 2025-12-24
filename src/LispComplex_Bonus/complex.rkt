#lang racket

;; Module fonctionnel pour les nombres complexes

(provide make-complex
         real
         imag
         add
         sub
         mul
         divc
         conjc
         equal?
         format-complex)


;; Constructeur : un complexe est une paire (a . b)
(define (make-complex a b)
  (cons a b))

;; Sélecteurs
(define (real z) (car z))
(define (imag z) (cdr z))


;; -------------------------
;; Opérations fonctionnelles
;; -------------------------

;; Addition
(define (add z1 z2)
  (make-complex
   (+ (real z1) (real z2))
   (+ (imag z1) (imag z2))))

;; Soustraction
(define (sub z1 z2)
  (make-complex
   (- (real z1) (real z2))
   (- (imag z1) (imag z2))))

;; Multiplication
(define (mul z1 z2)
  (make-complex
   (- (* (real z1) (real z2))
      (* (imag z1) (imag z2)))
   (+ (* (real z1) (imag z2))
      (* (imag z1) (real z2)))))

;; Division
(define (divc z1 z2)
  (let* ([a (real z1)]
         [b (imag z1)]
         [c (real z2)]
         [d (imag z2)]
         [den (+ (* c c) (* d d))])
    (if (= den 0)
        (error "Erreur : division par un nombre complexe nul")
        (make-complex
         (/ (+ (* a c) (* b d)) den)
         (/ (- (* b c) (* a d)) den)))))

;; Conjugué
(define (conjc z)
  (make-complex (real z) (- (imag z))))

;; égalité exacte
(define (equal? z1 z2)
  (and (= (real z1) (real z2))
       (= (imag z1) (imag z2))))



;; Formate un flottant avec 2 décimales en forçant 
(define (fmt x)
  (let ([v (if (< (abs x) 1e-12) 0.0 x)])
    (real->decimal-string v 2)))

;; format complex
(define (format-complex z)
  (string-append (fmt (real z)) " + " (fmt (imag z)) "i"))
