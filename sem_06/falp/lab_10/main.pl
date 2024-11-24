:- use_module(library(dif)).	% Sound inequality
:- use_module(library(clpfd)).	% Finite domain constraints
:- use_module(library(clpb)).	% Boolean constraints
:- use_module(library(chr)).	% Constraint Handling Rules
:- use_module(library(when)).	% Coroutining
:- dynamic fact/2.

fact1(0, Acc, Acc) :- !.
fact1(N, Acc, Res) :- N #> 0, Acc #=< Res, Res #> 0, Acc1 #= N * Acc, N1 #= N - 1, fact1(N1, Acc1, Res).

fact(N, Res) :- fact1(N, 1, Res).

fib1(0, AccA, _, AccA) :- !.
fib1(1, _, AccB, AccB) :- !.
fib1(N, AccA, AccB, Res) :- N #> 1, AccA #< Res, Res #> 0, AccB1 #= AccA + AccB, N1 #= N - 1, fib1(N1, AccB, AccB1, Res).

% fib(N, Res) :- N #< 0, N1 #= -N, Res1 #= -Res, fib1(N1, 0, 1, Res1), !.
fib(N, Res) :- fib1(N, 0, 1, Res).