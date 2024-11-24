(defun get_reversed (l r) 
	(cond
		((null l) r)
		(t (get_reversed 
			(cdr l) (cons (car l) r))
		)
		  
	)
)

(defun my_reverse (l)
	(get_reversed l nil)
)

(defun my_reverse (l &optional (r nil)) 
	(cond
		((null l) r)
		(t (get_reversed 
			(cdr l) (cons (car l) r))
		)  
	)
)

(defun get_list (l)
	(
		cond ((null l) nil)
			 ((and (listp (car l)) (car l)) (car l))
			 (t (get_list (cdr l)))
	)
)

(defun mul_numbers (l num)
	(
		cond ((null l) nil)
			 (t  (cons (* (car l) num) (mul_numbers (cdr l) num)))
	)
)

(defun mul_numbers (l num)
	(
		cond ((null l) nil)
			 ((numberp (car l)) (cons (* (car l) num) (mul_numbers (cdr l) num)))
			 (t (cons (car l) (mul_numbers (cdr l) num)))
	)
)


(defun select_between (lst l r)
	(
		cond ((null lst) nil)
			 ((< l (car lst) r) (cons (car lst) (select_between (cdr lst) l r)))
			 (t (select_between (cdr lst) l r))
	)
)

(defun rec_add (lst)
	(
		cond ((null lst) 0)
			 ((numberp (car lst)) (+ (car lst) (rec_add (cdr lst))))
			 (t (rec_add (cdr lst)))
	)
)

(defun rec_add (lst)
	(
		cond ((null lst) 0)
			 ((numberp (car lst)) (+ (car lst) (rec_add (cdr lst))))
			 ((listp (car lst)) (+ (rec_add (car lst)) (rec_add (cdr lst))))
			 (t (rec_add (cdr lst)))
	)
)

(defun drop-last (x)
	(cond 
		((null x) nil)
		((null (cdr x)) nil)
		(t (cons (car x) (drop-last (cdr x))))
	)
)

