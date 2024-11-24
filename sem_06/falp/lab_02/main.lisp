(defun getHypot (k1 k2) (sqrt (+ (* k1 k1) (* k2 k2))))
(getHypot 4 3)

(defun f-to-c (temp) (* 5/9 (- temp 320)))
;  5/9*(f-320) 9/5*c+32.0
(defun c-to-f (temp) (+ (* 9/5 temp) 32.0))
