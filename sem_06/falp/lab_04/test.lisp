(defclass ring ()
	(
		(value :accessor value :initarg :value)
		(module :accessor module :initarg :module)
	)
)

(defmethod add ((x ring) y)
		(make-instance 'ring :value (mod (+ (value x) (value y)) (module x)) :module (module y))
)

(defmethod mul ((x ring) y)
	(make-instance 'ring :value (mod (* (value x) (value y)) (module x)) :module (module y))
		
)

(defmethod neg ((x ring))
		(make-instance 'ring :value (- (module x) (value x)) :module (module x))
)

(defun nod (x y)
	(
		cond ((= x y) x)
		((< x y) (nod x (- y x)))
		(t (nod (- x y) y))
	)
)

 
; int gcd (int a, int b, int & x, int & y) {
; 	if (a == 0) {
; 		x = 0; y = 1;
; 		return b;
; 	}
; 	int x1, y1;
; 	int d = gcd (b%a, a, x1, y1);
; 	x = y1 - (b / a) * x1;
; 	y = x1;
; 	return d;
; }

(defun getx (a b x y)
	(cond
		((= a 0) (list a b 0 1))
		(t (let ((res (getx (mod b a) a 0 0))) (list a (cadr res) (- (cadddr res) (* (/ (- b (mod b a)) a) (caddr res))) (caddr res))))
	)
)

(defmethod inv ((x ring))
	(let ((res (getx (value x) (module x) 0 0))) 
		(cond
			((not (= (cadr res) 1)) nil)
			(t ( make-instance 'ring :value (mod (+ (mod (caddr res) (module x)) (module x)) (module x)) :module (module x)))
		)
	)
)

(defmethod div ((x ring) y)
		(mul x (inv y))
)

(defclass monomial ()
   (
      (coefficient :accessor coefficient :initarg :coefficient)
      (power :accessor power :initarg :power)
   )
)

; (defmethod to_string_sign ((object monomial))
;    (format NIL " + ~dx^~d" (if (< (coefficient object) 0) '- '+) (abs (coefficient object)) (power object))
; )

(defmethod to_string ((object monomial))
   (format NIL " + ~dx^~d" (value (coefficient object)) (power object))
)

(defclass polynomial ()
   (
      (monomials :accessor monomials :initarg :monomials)
   )
)

(defun polynomial_from_list (lst module)
	(make-instance 'polynomial :monomials (mapcar
		#'(lambda (el) (make-instance 'monomial :coefficient (make-instance 'ring :value (car el) :module module) :power (cdr el)))
		lst
	))
)

(defmethod to_string ((object polynomial))
   (cond
	   	((null (monomials object)) "0")
	   	(t 
	   		(reduce 
			   	#'(lambda (res el) (concatenate 'string res (to_string el)))
			   	(cdr (monomials object))
			   	:initial-value (to_string (car (monomials object)))
			)
	   	)
   	)
)

; (defvar m1 (make-instance 'monomial :coefficient 4 :power 3))

; (print (to_string m1))

(defun add_monomial (m p)
   (cond 
      ((zerop (value (coefficient m))) p)
      (t (make-instance 'polynomial :monomials (cons m (monomials p))))
   )
)

(defmethod get_first ((object polynomial))
   (car (monomials object))
)

(defmethod drop_first ((object polynomial))
   (make-instance 'polynomial :monomials (cdr (monomials object)))
)

(defmethod empty ((object polynomial))
   (null (monomials object))
)

(defun polynomial_bin_sum (p1 p2 res)
   (cond
      ((and (empty p1) (empty p2)) 
      	(make-instance 'polynomial :monomials (reverse (monomials res)))
      ) 
      ((empty p1)
      	(polynomial_bin_sum p1 (drop_first p2) (add_monomial (get_first p2) res))
      ) 
      ((empty p2) 
      	(polynomial_bin_sum (drop_first p1) p2 (add_monomial (get_first p1) res))
      )
      ((> (power (get_first p1)) (power (get_first p2))) 
      	(polynomial_bin_sum (drop_first p1) p2 (add_monomial (get_first p1) res))
      )
      ((< (power (get_first p1)) (power (get_first p2))) 
      	(polynomial_bin_sum p1 (drop_first p2) (add_monomial (get_first p2) res))
      )
      (t
       (polynomial_bin_sum (drop_first p1) (drop_first p2) (add_monomial (make-instance 'monomial :coefficient (add (coefficient (get_first p1)) (coefficient (get_first p2))) :power (power (get_first p1))) res))
      ) 
   )
)

(defun polynomial_bin+ (p1 p2)
   (polynomial_bin_sum p1 p2 (make-instance 'polynomial :monomials ()))
)

(defun polynomial+ (p &rest prest)
   (reduce
      #'polynomial_bin+
      prest
      :initial-value p
   )
)

(defun polynomial_neg (p)
   (make-instance 'polynomial :monomials (mapcar
      #'(lambda (el) (make-instance 'monomial :coefficient (neg (coefficient el)) :power (power el)))
      (monomials p)
   ))
)

(defun polynomial_bin- (p1 p2)
   (polynomial_bin+ p1 (polynomial_neg p2))
)

(defun polynomial- (p &rest prest)
   (reduce
      #'polynomial_bin-
      prest
      :initial-value p
   )
)

(defmethod polynomial_monomial* ((object polynomial) m)
   (make-instance 'polynomial :monomials (mapcar
      #'(lambda (el) (make-instance 'monomial :coefficient (mul (coefficient el) (coefficient m)) :power (+ (power el) (power m))))
      (monomials object)
   ))
)

(defun polynomial_bin* (p1 p2)
   (reduce
      #'polynomial_bin+
      (mapcar
         #'(lambda (el) (polynomial_monomial* p2 el))
         (monomials p1)
      )
      :initial-value (make-instance 'polynomial :monomials ())
   )
)

(defun polynomial* (p &rest prest)
   (reduce
      #'polynomial_bin*
      prest
      :initial-value p
   )
)

(defmethod divide_monomial ((p1 polynomial) p2)
	(make-instance 'monomial :coefficient (div (coefficient (get_first p1)) (coefficient (get_first p2))) :power (- (power (get_first p1)) (power (get_first p2))))
)

(defun polynomial_bin_div (p1 p2 res)
   (cond
   	  ((empty p2)
   	  	(error "division by zero")
   	  )
      ((empty p1) 
      	(cons (make-instance 'polynomial :monomials (reverse (monomials res))) (make-instance 'polynomial :monomials ()))
      )
      ((< (power (get_first p1)) (power (get_first p2)))
      	(cons (make-instance 'polynomial :monomials (reverse (monomials res))) p1)
      )
      (t
      	(polynomial_bin_div (polynomial_bin- p1 (polynomial_monomial* p2 (divide_monomial p1 p2))) p2 (add_monomial (divide_monomial p1 p2) res))
      )
   )
)

(defun polynomial_bin/ (p1 p2)
   (polynomial_bin_div p1 p2 (make-instance 'polynomial :monomials ()))
)

(defun polynomial/ (p &rest prest)
   (reduce
      #'(lambda (res el) (car (polynomial_bin/ res el)))
      prest
      :initial-value p
   )
)

(defclass sum ()
	((p1 :accessor p1 :initarg :p1)
	 (p2 :accessor p2 :initarg :p2)
	)
)

(defclass sub ()
	((p1 :accessor p1 :initarg :p1)
	 (p2 :accessor p2 :initarg :p2)
	)
)

(defclass mul ()
	((p1 :accessor p1 :initarg :p1)
	 (p2 :accessor p2 :initarg :p2)
	)
)

(defclass div ()
	((p1 :accessor p1 :initarg :p1)
	 (p2 :accessor p2 :initarg :p2)
	)
)

(defgeneric evaluate (expr)
  (:method ((expr sum))
    (polynomial_bin+ (evaluate (p1 expr))
       (evaluate (p2 expr))))
  (:method ((expr sub))
    (polynomial_bin- (evaluate (p1 expr))
       (evaluate (p2 expr))))
  (:method ((expr mul))
    (polynomial_bin* (evaluate (p1 expr))
       (evaluate (p2 expr))))
  (:method ((expr div))
    (car (polynomial_bin/ (evaluate (p1 expr))
       (evaluate (p2 expr)))))
  (:method ((expr polynomial))
    expr))

; (defvar x (make-instance 'ring :value 0 :module 4))
; (defvar y (make-instance 'ring :value 1  :module 4))
; (print (value (mul x y)))

(defvar p1 (polynomial_from_list '((1 . 3) (1 . 2) (1 . 0)) 3))
; (defvar p1 (polynomial_from_list '((2 . 1) (1 . 0)) 3))

(print (to_string p1))

(defvar p2 (polynomial_from_list '((1 . 1) (2 . 0)) 3))

(print (to_string p2))
(print (to_string (polynomial+ p1 p2)))

(print (to_string (polynomial* p1 p2)))

; (defvar res (make-instance 'polynomial :monomials ()))

; (print (to_string (add_monomial (divide_monomial p1 p2) res)))

; (print (to_string (polynomial_bin- p1 (polynomial_monomial* p2 (divide_monomial p1 p2)))))

; (print (to_string (car (polynomial_bin/ p1 p2))))
(print (to_string (polynomial/ p1 p2)))

; (print (to_string (car (polynomial_bin/ p1 p2))))

; (print (to_string (cdr (polynomial_bin/ p1 p2))))

; (defvar p3 (make-instance 'sum :p1 p1 :p2 p2))

; (print (to_string (evaluate p3)))

; (defvar p4 (make-instance 'sub :p1 p1 :p2 p2))

; (print (to_string (evaluate p4)))

; (defvar p5 (make-instance 'mul :p1 p1 :p2 p2))

; (print (to_string (evaluate p5)))

; (defvar p6 (make-instance 'div :p1 p1 :p2 p2))

; (print (to_string (evaluate p6)))

; (load "~/quicklisp/setup.lisp")
; (ql:quickload "fiveam")

; (fiveam:test polynomial+
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '((1 . 3))) (polynomial_from_list '((1 . 2))))) "1x^3 + 1x^2"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '((1 . 3) (5 . 2))) (polynomial_from_list '((-7 . 2))))) "1x^3 - 2x^2"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '((7 . 2))) (polynomial_from_list '((-7 . 2))))) "0"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '()) (polynomial_from_list '((-7 . 2))))) "-7x^2"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '()) (polynomial_from_list '()))) "0"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '((2 . 2) (2 . 1) (3 . 0))) (polynomial_from_list '((8 . 2) (8 . 1))))) "10x^2 + 10x^1 + 3x^0"))
;     (fiveam:is (equalp (to_string (polynomial+ (polynomial_from_list '((1 . 3) (5 . 2))) (polynomial_from_list '((-7 . 2))) (polynomial_from_list '((-7 . 3))))) "-6x^3 - 2x^2"))
; )

; (fiveam:test polynomial-
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '((1 . 3))) (polynomial_from_list '((1 . 2))))) "1x^3 - 1x^2"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '((1 . 3) (7 . 2))) (polynomial_from_list '((7 . 2))))) "1x^3"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '((7 . 2))) (polynomial_from_list '((7 . 2))))) "0"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '()) (polynomial_from_list '((-7 . 2))))) "7x^2"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '()) (polynomial_from_list '()))) "0"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '((2 . 2) (2 . 1) (3 . 0))) (polynomial_from_list '((8 . 2) (8 . 1))))) "-6x^2 - 6x^1 + 3x^0"))
;     (fiveam:is (equalp (to_string (polynomial- (polynomial_from_list '((1 . 3) (5 . 2))) (polynomial_from_list '((-7 . 2))) (polynomial_from_list '((-7 . 3))))) "8x^3 + 12x^2"))
; )

; 	(fiveam:test polynomial*
;     (fiveam:is (equalp (to_string (polynomial* (polynomial_from_list '((1 . 1) (-1 . 0))) (polynomial_from_list '((1 . 1) (-1 . 0))))) "1x^2 - 2x^1 + 1x^0"))
;     (fiveam:is (equalp (to_string (polynomial* (polynomial_from_list '((1 . 1) (-1 . 0))) (polynomial_from_list '()))) "0"))
;     (fiveam:is (equalp (to_string (polynomial* (polynomial_from_list '()) (polynomial_from_list '()))) "0"))
;     (fiveam:is (equalp (to_string (polynomial* (polynomial_from_list '((5 . 2) (-3 . 1) (4 . 0))) (polynomial_from_list '((1 . 1))))) "5x^3 - 3x^2 + 4x^1"))
;     (fiveam:is (equalp (to_string (polynomial* (polynomial_from_list '((4 . 3) (1 . 2) (-15 . 0))) (polynomial_from_list '((13 . 4) (6 . 3) (-2 . 2))))) "52x^7 + 37x^6 - 2x^5 - 197x^4 - 90x^3 + 30x^2"))
; )

; (fiveam:test polynomial/
;     (fiveam:is (equalp (to_string (polynomial/ (polynomial_from_list '()) (polynomial_from_list '((1 . 1) (-1 . 0))))) "0"))
;     (fiveam:is (equalp (to_string (polynomial/ (polynomial_from_list '((1 . 3) (-12 . 2) (-42 . 0))) (polynomial_from_list '((1 . 1) (-3 . 0))))) "1x^2 - 9x^1 - 27x^0"))
; )

; (fiveam:test polynomial_bin/
;     (fiveam:is (equalp (to_string (car (polynomial_bin/ (polynomial_from_list '()) (polynomial_from_list '((1 . 1) (-1 . 0)))))) "0"))
;     (fiveam:is (equalp (to_string (car (polynomial_bin/ (polynomial_from_list '((1 . 3) (-12 . 2) (-42 . 0))) (polynomial_from_list '((1 . 1) (-3 . 0)))))) "1x^2 - 9x^1 - 27x^0"))
;     (fiveam:is (equalp (to_string (cdr (polynomial_bin/ (polynomial_from_list '((1 . 3) (-12 . 2) (-42 . 0))) (polynomial_from_list '((1 . 1) (-3 . 0)))))) "-123x^0"))
; )

; (fiveam:run!)

