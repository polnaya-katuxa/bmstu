domains 
 list = integer*.

predicates
 nondeterm length1(list, integer, integer).
 nondeterm length(list, integer).
 
 nondeterm sum1(list, integer, integer).
 nondeterm sum(list, integer).
 
 nondeterm sum_odd_pos1(list, integer, integer, integer).
 nondeterm sum_odd_pos(list, integer).
 nondeterm sum_odd_pos2(list, integer, integer).
 nondeterm sum_odd_pos_new(list, integer).
 
 nondeterm list_of_bigger(list, integer, list).
 
 nondeterm del_all(list, integer, list).
 nondeterm del_single(list, integer, list).
 
 nondeterm union(list, list, list).
 nondeterm union_new(list, list, list).
 
clauses
 length1([_ | T], Acc, Res) :- !, Acc1 = Acc + 1, length1(T, Acc1, Res).
 length1([], Acc, Acc) :- !.
 length(L, Res) :- !, length1(L, 0, Res).
 
 sum1([], Acc, Acc) :- !.
 sum1([H | T], Acc, Res) :- !, Acc1 = Acc + H, sum1(T, Acc1, Res).
 sum(L, Res) :- !, sum1(L, 0, Res).
 
 sum_odd_pos1([], _, Acc, Acc) :- !.
 sum_odd_pos1([_ | T], Pos, Acc, Res) :- Pos mod 2 = 0, !, Pos1 = Pos + 1, sum_odd_pos1(T, Pos1, Acc, Res).
 sum_odd_pos1([H | T], Pos, Acc, Res) :- !, Pos1 = Pos + 1, Acc1 = Acc + H, sum_odd_pos1(T, Pos1, Acc1, Res).
 sum_odd_pos(L, Res) :- !, sum_odd_pos1(L, 0, 0, Res).
 
 sum_odd_pos2([], Acc, Res) :- Res = Acc, !.
 sum_odd_pos2([_ | [H | T]], Acc, Res) :- !, Acc1 = Acc + H, sum_odd_pos2(T, Acc1, Res).
 sum_odd_pos_new(L, Res) :- !, sum_odd_pos2(L, 0, Res).
 
 list_of_bigger([], _, []) :- !.
 list_of_bigger([H | T], N, [H | ResT]) :- H > N, !, list_of_bigger(T, N, ResT).
 list_of_bigger([_ | T], N, Res) :- !, list_of_bigger(T, N, Res).
 
 del_all([], _, []) :- !.
 del_all([H | T], N, Res) :- H = N, !, del_all(T, N, Res).
 del_all([H | T], N, [H | ResT]) :- !, del_all(T, N, ResT).
 
 del_single([], _, []) :- !.
 del_single([H | T], N, T) :- H = N, !.
 del_single([H | T], N, [H | ResT]) :- !, del_single(T, N, ResT).
 
 union([], [], []) :- !.
 union([H1 | T1], [], [H1 | ResT]) :- !, union(T1, [], ResT).
 union([], [H2 | T2], [H2 | ResT]) :- !, union([], T2, ResT).
 union([H1 | T1], [H2 | T2], [H1, H2 | ResT]) :- !, union(T1, T2, ResT).
 
 union_new([], L, L) :- !.
 union_new([H | T], L, [H | ResT]) :- !, union(T, L, ResT).

goal
 %length([1, 2, 3], Is).
 %sum([1, 2, 3], Is).
 %sum_odd_pos([1, 0, 2, 0, 5, 3], Is).
 %sum_odd_pos_new([1, 0, 2, 0, 5, 3], Is).
 %list_of_bigger([1, 0, 2, 0, 5, 3], 2, Is).
 %del_single([1, 0, 2, 0, 5, 2, 6], 2, Is).
 %del_all([1, 0, 2, 0, 5, 2, 6], 2, Is).
 %union([1, 2, 3], [1, 0, 2, 0, 5, 2, 6], Is).
 union_new([1, 2, 3], [1, 0, 2, 0, 5, 2, 6], Is).