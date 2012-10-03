(define (fibb n)
	(if (< num 2)
		(if (>= num 0)
			num
			(-
				(fibb (+ n 2))
				(fibb (+ n 1))
			)
		)
		(+
			(fibb (- n 2))
			(fibb (- n 1))
		)
	)
)
