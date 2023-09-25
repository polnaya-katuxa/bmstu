run:
	go run cmd/lab_05/main.go

bench:
	mkdir -p "data"
	go run -gcflags '-N -l' cmd/test/main.go -test.benchtime=100x

plot:
	mkdir -p "images"
	gnuplot -persist scripts/memory.gpi
	gnuplot -persist scripts/time_best.gpi
	gnuplot -persist scripts/time_worst.gpi
	gnuplot -persist scripts/time_best_winograd.gpi
	gnuplot -persist scripts/time_worst_winograd.gpi

test:
	go test ./...  -v