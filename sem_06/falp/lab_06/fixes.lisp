(defun get_list (l)
	(
		cond ((null l) nil)
			 ((atom (car l)) (get_list (cdr l)))
			 (t (caar l))
	)
)

; с вложенностью

(defun get_inside (l)
	(
		cond ((null l) nil)
		((atom (car l)) (car l))
		(t (get_inside (car l)))
	)
)

(defun get_first_atom (l)
	(
		cond ((null l) nil)
		((atom (car l)) (get_first_atom (cdr l)))
		(t (get_inside (car l))
	)
))


;(3 4 (9 7) 8) -> 9
;(3 4 (((5) 6) 7) 8) -> 5
;(3 4 ((() 6) 7) 8) -> 

;(3 4 (9 7) 8) -> (8 (7 9) 4 3)
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defun mul_numbers1 (l num acc)
	(
		cond ((null l) acc)
		(t  (mul_numbers1 (cdr l) num (cons (* (car l) num) acc)))
	)
)

(defun my_mul1 (l num)
	(mul_numbers1 l num nil)
)

(defun mul_numbers2 (l num acc)
	(
		cond ((null l) acc)
		((numberp (car l)) (mul_numbers2 (cdr l) num (cons (* (car l) num) acc)))
		((atom (car l)) (mul_numbers2 (cdr l) num (cons (car l) acc)))
		(t (mul_numbers2 (cdr l) num (cons (mul_numbers2 (car l) num nil) acc)))
	)
)

(defun my_mul2 (l num)
	(mul_numbers2 l num nil)
)

;;;;;;;;;;;;;;;;;;;;;;;;;;;

; (defun select_between (lst l r)
; 	(
; 		cond ((null lst) nil)
; 			 ((or (< l (car lst) r) (< r (car lst) l)) (cons (car lst) (select_between (cdr lst) l r)))
; 			 (t (select_between (cdr lst) l r))
; 	)
; )

(defun select_between1 (lst l r acc)
	(
		cond ((null lst) acc)
			 ((or (< l (car lst) r) (< r (car lst) l)) (select_between1 (cdr lst) l r (cons (car lst) acc)))
			 (t (select_between1 (cdr lst) l r acc))
	)
)

(defun select_between (lst l r)
	(
		select_between1 lst l r nil
	)
)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defun rec_add (lst)
	(
		cond ((null lst) 0)
			 ((numberp (car lst)) (+ (car lst) (rec_add (cdr lst))))
			 ((atom (car lst)) (rec_add (cdr lst)))
			 (t (+ (rec_add (car lst)) (rec_add (cdr lst))))
	)
)

(defun add (lst acc)
(	cond ((null lst) acc)
	((numberp (car lst)) (add (cdr lst) (+ (car lst) acc)))
	(t (add (cdr lst) acc))
))

(defun rec_add (lst) (add lst 0))





