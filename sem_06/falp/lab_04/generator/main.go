package main

import (
	"fmt"
	"math/rand"
	"os"
)

const script = `
(defun add_monomial (m p)
   (cond 
      ((zerop (car m)) p)
      (t (cons m p))
   )
)

; По возрастанию степеней
(defun sum_pol1 (p1 p2 res)
   (cond
      ((and (null p1) (null p2)) res) ; Если оба пустые, возвращаем результат
      ((null p1) (sum_pol1 p1 (cdr p2) (cons (car p2) res))) ; Если первый пустой, добавляем элемент второго
      ((null p2) (sum_pol1 (cdr p1) p2 (cons (car p1) res))) ; Если второй пустой, добавляем элемент первого
      ((> (cdar p1) (cdar p2)) (sum_pol1 (cdr p1) p2 (cons (car p1) res)))
      ((< (cdar p1) (cdar p2)) (sum_pol1 p1 (cdr p2) (cons (car p2) res)))
      (t (sum_pol1 (cdr p1) (cdr p2) (add_monomial (cons (+ (caar p1) (caar p2)) (cdar p1)) res)))
   )
)

(defun sum_pol_bin (p1 p2)
   (reverse (sum_pol1 p1 p2 NIL))
)

(defun pol_neg (p)
   (mapcar
      #'(lambda (el) (cons (- (car el)) (cdr el)))
      p
   )
)

(defun sub_pol_bin (p1 p2)
   (sum_pol_bin p1 (pol_neg p2))
)

(defun sum_pol (p &rest prest)
   (reduce
      #'sum_pol_bin
      prest
      :initial-value p
   )
)

(defun sub_pol (p &rest prest)
   (reduce
      #'sub_pol_bin
      prest
      :initial-value p
   )
)

(defun mul_pol_by_mon (m p)
   (mapcar
      #'(lambda (el) (cons (* (car el) (car m)) (+ (cdr el) (cdr m))))
      p
   )
)

(defun mul_pol_bin (p1 p2)
   (reduce
      #'sum_pol_bin
      (mapcar
         #'(lambda (el) (mul_pol_by_mon el p2))
         p1
      )
      :initial-value ()
   )
)


(defun mul_pol (p &rest prest)
   (reduce
      #'mul_pol_bin
      prest
      :initial-value p
   )
)

; p1, p2 - reversed
(defun div_coef (p1 p2)
   (cons (/ (caar p1) (caar p2)) (- (cdar p1) (cdar p2)))
)

(defun div_pol_bin1 (p1 p2 res)
   (cond
      ((or (null p1) (null p2)) NIL)
      ((< (cdar p1) (cdar p2)) (list (reverse res) p1))
      (t (div_pol_bin1 (sub_pol_bin p1 (mul_pol_by_mon (div_coef p1 p2) p2)) p2 (cons (div_coef p1 p2) res)))
   )
)

(defun div_pol_bin (p1 p2)
   (div_pol_bin1 p1 p2 NIL)
)

(defun div_pol (p &rest prest)
   (reduce
      #'(lambda (res el) (car (div_pol_bin res el)))
      prest
      :initial-value p
   )
)

(print (div_pol %s))
`

func generatePolynomial(n int) string {
	res := "'("

	for i := n; i > 0; i-- {
		res += fmt.Sprintf("(%d . %d)", rand.Intn(1000), i)
		if i != n-1 {
			res += " "
		}
	}

	res += ")"

	return res
}

func generatePolynomials(n, power int) string {
	res := ""

	for i := 0; i < n; i++ {
		res += generatePolynomial(power)
		if i != n-1 {
			res += " "
		}
	}

	return res
}

func main() {
	var power, n int

	fmt.Print("Enter power: ")
	fmt.Scan(&power)

	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	file, _ := os.Create("test.lisp")
	fmt.Fprintf(file, script, generatePolynomials(n, power))
	file.Close()
}
