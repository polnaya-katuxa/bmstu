(defun minus_10_if_num (x)
	(if (numberp x) (- x 10) x)
)

(defun minus_10 (x) 
	(mapcar #'minus_10_if_num x) 
)


(defun square_num (x)
	(if (numberp x) (* x x) x)
)

(defun square_list (x) 
	(mapcar #'square_num x) 
)

(defun mul_list (x mul)
	(mapcar #'(lambda (num) (* num mul)) x)
)

(defun mul_list (x mul)
	(mapcar #'(lambda (num) (if (numberp num) (* num mul) num)) x)
)

3 2 1


(defun drop_last (x) (reverse (cdr (reverse x))))

(defun fun (x)
	(let ((r (reverse x)))
		(equalp (cdr r) (drop_last x))
	)
)

(defun is_palindrome (x)
	(maplist #'fun x)
)


(defun set_equal (x y)
	()
)




(defun is_in_set(elem src_set) 
	(not (null (member elem src_set)))
)

(defun is_subset(set1 set2)
    (reduce #'(lambda (x y) (and x y)) 
            (mapcar #'(lambda (x) (in_set x set2)) set1) :initial-value t))

(defun are_equal_sets(s1 s2)
    (and (is_subset s1 s2) (is_subset s2 s1)))



(defun in_set(elem src_set) 
    (reduce #'(lambda (x y) (or x y)) 
            (mapcar #'(lambda (x) (equal x elem)) src_set) 
            :initial-value Nil))

(defun is_subset(set1 set2)
    (reduce #'(lambda (x y) (and x y)) 
            (mapcar #'(lambda (x) (in_set x set2)) set1)))

(mapcar #'(lambda (x) (in_set x set2)) set1)

(defun f (l x)
	(not (equal (car (member x l)) elem))
)

(defun is_subset (s1 s2)
	(not ())
	(not (find nil (mapcar #'(lambda (x) (in_set x set2)) set1) :test #'equal))
)

(defun set_equal(set1 set2)
    (and (is_subset set1 set2) (is_subset set2 set1)))


:initial-value T


(defun select_between(lst left right)
        (sort (reduce #'(lambda (res_lst elem) 
            (if (< left elem right) (cons elem res_lst) res_lst)) 
                lst :initial-value ()) #'<))



(defun select_between (lst l r)
	()
)


(defun cartesian (l1 l2)
	(reduce #'append
        (mapcar
                (lambda (x) (mapcar
                        (lambda (y) (cons x y))
                        l2
                ))
                l1
        )
    :initial-value ())
)

(defun f (l)
 (reduce
  (lambda (res elem) (+ res (length elem)))
  l
  :initial-value 0
 )
)