#!/bin/bash

for i in $( seq $1 $2); do
echo -n $i" " >> time.txt
./cg-practice -d $i -i $3 -m >> time.txt || (echo "Program failed."; exit 1)
done