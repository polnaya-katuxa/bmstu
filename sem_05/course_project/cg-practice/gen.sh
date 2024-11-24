#!/bin/bash

for i in $( seq $1 $2); do
echo Generating $i...
./cg-practice -d 3 -i $3 -r $i -o $(printf "%02d%02d" $4 $i).png || (echo "Program failed."; exit 1)
done