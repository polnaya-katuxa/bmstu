#!/bin/bash

alphas="0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1.0"
#betas="0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1.0"
ks="0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1.0"
times="1 2 3"
n=4

cd ..

for alpha in $alphas; do
    #for beta in $betas; do
    	for k in $ks; do
	    	for time in $times; do
	    		go run ./cmd/lab_07 -f=$n -m=true -a=$alpha -k=$k -t=$time >> "data/paramQ$n.txt" #-b=$beta
	    	done
	    done
    #done
done