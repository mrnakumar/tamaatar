#!/bin/bash

# Useful while developing.
# Builds ui and backend and launches the main
frontEnd='frontEnd'
[ ! -d "$frontEnd" ] && mkdir $frontEnd
cd ui || exit 1
ng build || exit 2
mv dist/ui/* ../$frontEnd/ || exit 3
cd ../backend || exit 4
go build main.go || exit 5
mv main ../ && cd ..
./main || exit 6