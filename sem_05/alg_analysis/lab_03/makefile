run:
	go run cmd/lab_02/main.go

bench:
	go run cmd/test/main.go -test.benchtime=1000x

plot:
	gnuplot -persist scripts/time_best.gpi
	gnuplot -persist scripts/time.gpi
	gnuplot -persist scripts/memory.gpi
	gnuplot -persist scripts/time_worst.gpi

test:
	go test ./...  -v