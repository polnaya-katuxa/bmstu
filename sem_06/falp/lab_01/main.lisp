; 2
; 2 element
(car (cdr `(1 2 3 4 5 6)))

; 3 element
(car (cdr (cdr `(1 2 3 4 5 6))))

; 4 element
(car (cdr (cdr(cdr `(1 2 3 4 5 6)))))

; 3
(CAADR ' ((blue cube) (red pyramid))) ; RED
(CDAR '((abc) (def) (ghi))) ; NIL
(CADR ' ((abc) (def) (ghi))) ; (DEF)
(CADDR ' ((abc) (def) (ghi))) ; (GHI)

; 4
(list 'Fred 'and 'Wilma) (list 'Fred '(and Wilma)) (cons Nil Nil) ; (FRED AND WILMA) (FRED (AND WILMA)) (NIL)
(cons T Nil) ; (T)
(cons Nil T) ; (NIL . T)
(list Nil) ; (NIL)
(cons ' (T) Nil) ; ((T))
(list ' (one two) ' (free temp)) ; ((ONE TWO) (FREE TEMP))
(cons 'Fred '(and Wilma)) (cons 'Fred '(Wilma)) (list Nil Nil) ; (FRED AND WILMA) (FRED WILMA) (NIL NIL)
(list T Nil) ; (T NIL)
(list Nil T) ; (NIL T)
(cons T (list Nil)) ; (T NIL) 
(list '(T) Nil) ; ((T) NIL)
(cons '(one two) '(free temp)) ; ((ONE TWO) FREE TEMP)

; 5
; (f arl ar2 ar3 ar4) => ((arl ar2) (ar3 ar4))
((lambda (ar1 ar2 ar3 ar4) (list (list ar1 ar2) (list ar3 ar4))) 1 2 3 4)

(defun f1 (ar1 ar2 ar3 ar4) (list (list ar1 ar2) (list ar3 ar4)))

; (f arl ar2) => ((arl) (ar2))
((lambda (ar1 ar2) (list (list ar1) (list ar2))) 1 2)

(defun f2 (ar1 ar2) (list (list ar1) (list ar2)))

; (f arl) => (((arl)))
((lambda (ar1) (list (list (list ar1)))) 1)

(defun f3 (ar1) (list (list (list ar1))))
