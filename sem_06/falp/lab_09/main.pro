domains
	name = symbol.
	gender = symbol.
	human = human(name, gender).
	
predicates
	nondeterm parent(human, human).
	nondeterm gender(name, gender).
	
	nondeterm grandparent_gender(name, name, gender).
	
	nondeterm grandmother(name, name).
	nondeterm grandfather(name, name).
	nondeterm grandparent(name, name).
	
	nondeterm grandparent_from_mother_gender(name, name, gender).
	nondeterm grandmother_from_mother(name, name).
	nondeterm grandparent_from_mother(name, name).
	
	nondeterm sibling(name, name).
	nondeterm partner(name, name).
	nondeterm aunt(name, name).
	nondeterm grandmaN(name, name, integer).
	nondeterm grandN(name, name, integer).
	
	nondeterm max2(integer, integer, integer).
	nondeterm max2_cut(integer, integer, integer).
	
	nondeterm max3(integer, integer, integer, integer).
	nondeterm max3_cut(integer, integer, integer, integer).
	
clauses
	parent(human(nigel, male), human(eliza, female)).
	parent(human(marianna, female), human(eliza, female)).
	%parent(human(nigel, male), human(dabby, female)).
	%parent(human(marianna, female), human(dabby, female)).
	parent(human(nigel, male), human(donnie, male)).
	parent(human(marianna, female), human(donnie, male)).
	
	parent(human(sofie, female), human(marianna, female)).
	parent(human(sir, male), human(nigel, male)).
	parent(human(sir, male), human(mimi, female)).
	parent(human(sirness, female), human(nigel, male)).
	parent(human(sirness, female), human(mimi, female)).
	
	parent(human(grandsir, male), human(sir, male)).
	parent(human(grandsirness, female), human(sirness, female)).
	parent(human(grandgrandsirness, female), human(grandsirness, female)).
	
	gender(Name, Gender) :- parent(human(Name, Gender), _); parent(_, human(Name, Gender)).
	
	grandparent_gender(HumanName, Name, Gender) :- parent(human(Name, Gender), Parent), parent(Parent, human(HumanName, _)).
	
	grandmother(HumanName, Name) :- grandparent_gender(HumanName, Name, female).
	grandfather(HumanName, Name) :- grandparent_gender(HumanName, Name, male).
	grandparent(HumanName, Name) :- grandparent_gender(HumanName, Name, _).
	
	grandparent_from_mother_gender(HumanName, Name, Gender) :- parent(human(Name, Gender), human(ParentName, female)), parent(human(ParentName, female), human(HumanName, _)).
	
	grandmother_from_mother(HumanName, Name) :- grandparent_from_mother_gender(HumanName, Name, female).
	grandparent_from_mother(HumanName, Name) :- grandparent_from_mother_gender(HumanName, Name, _).
	
	sibling(Name1, Name2) :- parent(human(NameF, female), human(Name1, _)), parent(human(NameF, female), human(Name2, _)), 
				 parent(human(NameM, male), human(Name1, _)), parent(human(NameM, male), human(Name2, _)).
				 
	partner(Name1, Name2) :- parent(human(Name1, _), Child), parent(human(Name2, _), Child), Name1 <> Name2, !.	
	
	aunt(Name1, Name2) :- sibling(Name1, Name), parent(human(Name, _), human(Name2, _)), Name1 <> Name.
	
	grandmaN(Grand, Name, 0) :- parent(human(Grand, _), human(Name, _)).
	grandmaN(Grand, Name, N) :- N1 = N - 1,
				    parent(human(Grand, _), human(PName, _)),
				    grandmaN(PName, Name, N1).
	grandN(Grand, Name, N) :- gender(Grand, female), grandmaN(Grand, Name, N).
	
	max2(A, B, A) :- A > B.
	max2(A, B, B) :- B >= A.
	
	max3(A, B, C, A) :- A >= B, A >= C.
	max3(A, B, C, B) :- B > A, B >= C.
	max3(A, B, C, C) :- C > A, C > B.
	
	max2_cut(A, B, A) :- A > B, !.
	max2_cut(_, B, B) :- !.
	
	max3_cut(A, B, C, A) :- A >= B, A >= C, !.
	max3_cut(_, B, C, B) :- B >= C, !.
	max3_cut(_, _, C, C) :- !.
goal
	%grandmother(nigel, Grandmother).
	% grandparent(donnie, GrandParent).
	% grandmother_from_mother(donnie, GrandParent).
	
	%sibling(eliza, Who).
	%partner(Who, nigel).
	%aunt(Who, eliza).
	grandN(Who, Name, 3).
	%gender(grandgrandsirness, F).
	
	%max2(2, 4, Res).
	%max3(3, 2, 3, Res).
	
	%max2_cut(2, 4, Res).
	%max3_cut(1, 2, 3, Res).