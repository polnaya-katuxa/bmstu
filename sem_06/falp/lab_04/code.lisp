(defun random_dice ()
	(+ (random 6) 1)
)

(defun random_dice_pair ()
	(cons (random_dice) (random_dice))
)

(defun play_dice (player) 
	(let ((result (random_dice_pair))) 
		(print (list player 'dice  'is result))
		(if (or (equal result (cons 6 6)) (equal result (cons 1 1)))
			(play_dice player)
			(+ (car result) (cdr result))
		)
	)
)

(defun absolute_win (dice_s)
	(or (= dice_s 7) (= dice_s 11))
)

(defun play_game ()
	(let ((result1 (play_dice 1)))
		(if (absolute_win result1)
			(print "player 1 absolutely won")
			(let ((result2 (play_dice 2)))
				(if (absolute_win result2)
					(print "player 2 absolutely won")
					(cond 
						((< result1 result2) (print "player 2 won"))
						((> result1 result2) (print "player 1 won"))
						(t (print "draw"))
					)
				)
			)
		)
		nil
	)
)





(defun drop_last (x) (reverse (cdr (reverse x))))

(defun is_palindrome (x) 
	(let ((r (reverse x)))
		 (or (<= (length x) 1) (and (equalp (car x) (car r)) (is_palindrome (drop_last (cdr x)))))
	)
)


(defun get_cap_by_country (country table)
	(cdr (assoc country table :test #'equalp))
)

(defun get_country_by_cap (cap table)
	(car (rassoc cap table :test #'equalp))
)

Напишите функцию, которая умножает на заданное число-аргумент первый числовой элемент списка из заданного 3-х элементного списка- аргумента, 
когда
(a) все элементы списка --- числа,
(б) элементы списка -- любые объекты.

(defun mul_number1 (x l) 
	(cond
		((not (numberp x)) nil)
		((not (listp l)) nil)
		((/= (length l) 3) nil)
		(t (setf (car l) (* x (car l))))
	)
)

(defun get_first_num (l)
	(cond
		((null l) nil)
		((numberp (car l)) l)
		(t (get_first_num (cdr l)))
	)
)

(defun mul_number2 (x l) 
	(cond
		((not (numberp x)) nil)
		((not (listp l)) nil)
		((/= (length l) 3) nil)
		(t (let
			((mul (get_first_num l)))
			(setf (car mul) (* x (car mul)))
			)
		)
	)
)



