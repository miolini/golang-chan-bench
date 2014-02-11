all:
	go test chan_bench_test.go --test.bench ".*" --test.benchtime=10s