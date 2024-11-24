predicates
	nondeterm fact1(integer, integer, integer).
	nondeterm fact(integer, integer).

	nondeterm fib1(integer, integer, integer, integer).
	nondeterm fib(integer, integer).
	
clauses
	fact1(0, Acc, Acc) :- !.
	fact1(N, Acc, Res) :- N > 0, Acc1 = N * Acc, N1 = N - 1, fact1(N1, Acc1, Res).

	fact(N, Res) :- fact1(N, 1, Res).

	fib1(0, AccA, _, AccA) :- !.
	fib1(1, _, AccB, AccB) :- !.
	fib1(-1, _, AccB, AccB) :- !.
	fib1(N, AccA, AccB, Res) :- N > 1, AccB1 = AccA + AccB, N1 = N - 1, fib1(N1, AccB, AccB1, Res).
	fib1(N, AccA, AccB, Res) :- N < -1, AccB1 = AccA + AccB, N1 = N + 1, fib1(N1, AccB, AccB1, Res).

	fib(N, Res) :- N < 0, fib1(N, 0, -1, Res), !.
	fib(N, Res) :- fib1(N, 0, 1, Res).
	
goal
	fact(5, N).
	
	