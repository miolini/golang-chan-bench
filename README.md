golang-chan-bench
=================

Simple Golang Channel Benchmark

This is simple Golang channel benchmark.

## Instructions

```
# make test
```


## Results

| Date  | Machine   | OS | Go Version | GOMAXPROCS=1 ns | GOMAXPROCS=2 ns |
|------:|:----------|----|------------|:------------------:|:------------------:|
| 2014-02-11 | MacBook Pro 2010 Mid Core i5 @ 2.4GHz | OSX 10.9 Mavericks  | 1.2 | 87 | 131 |
| 2014-02-11 | CubieTruck AllWinner A20 Dual @ 1GHz | Fedora Kernel 3.4.75 | 1.2 | 1219 | 1076 |
| 2014-02-11 | Xeon E5-2680 x 2 @ 2.70GHz | Ubuntu 12.04 LTS | 1.2 | 80 | 125 |
| 2014-02-12 | Intel Atom N270 @ 1.6GHz | Ubuntu 12.04 LTS | 1.2 | 542 | 538 |
| 2014-02-12 | Intel Core i5 M 520 @ 2.4GHz | Ubuntu 12.04 LTS | 1.2 | 82 | 157|

## Contacts

Please, send results to mio@volmy.com. Thank you.
