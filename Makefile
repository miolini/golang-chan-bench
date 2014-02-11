test:
	GOMAXPROCS=1 time go test chan_bench_test.go --test.bench ".*" --test.benchtime=10s
	GOMAXPROCS=2 time go test chan_bench_test.go --test.bench ".*" --test.benchtime=10s