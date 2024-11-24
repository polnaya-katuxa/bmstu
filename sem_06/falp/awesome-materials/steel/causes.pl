list_find([H|T], Num) :- H \= Num, !, list_find(T, Num).
list_find([H|_], Num):- H = Num , !.
list_find([], _) :- false.

plate_by_group(Group, Name, Ap, Fn, Cost) :- plate_properties(PlateGroup, List, _),
	                                         list_find(List, Group),
	                                         plate(PlateGroup, Name, Ap, Fn, Cost).

vc(Metal, Plate, Cover, Res) :- metal(Metal),
                                metal_properties(Group, Metal, _, _),
                                plate_by_group(Group, Plate, _, fn(_, FnMin, FnMax), _).


q(Metal, Plate, Cover, Res) :- metal(Metal).

power(Vc, Ap, Fn, Kc, Res) :-
	Res #= Vc * Ap * Fn * Kc.

plate_by_metal(MetalName,MaxPower, MaxResis, MinCost,MaxCost,Plate_Name):-
	metal_name(Tag, MetalName),
	metal_properties(Group, Tag, properties(_, Vc, Kc, _, _)),
	plate_by_group(Group, Plate_Name, ap(Ap), fn(Fn),Resistence ,cost(Cost)),
    power(Vc, Ap, Fn, Kc,Power),
	Power #< MaxPower,
	Resistence #< MaxResis,
    Cost #> MinCost,
    Cost #< MaxCost.