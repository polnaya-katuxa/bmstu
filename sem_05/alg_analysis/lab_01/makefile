run:
	go run cmd/lab_01/main.go

bench:
	go run cmd/test/main.go

plot:
	gnuplot -persist scripts/time.gpi
	gnuplot -persist scripts/memory.gpi
	gnuplot -persist scripts/time2.gpi
	gnuplot -persist scripts/memory2.gpi

test:
	go test ./...  -v