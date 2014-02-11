golang-chan-bench
=================

Simple Golang Channel Benchmark

This is simple Golang channel benchmark.

## Instructions

```
# make
# GOMAXPROCS=1 ./chan_bench 
# GOMAXPROCS=2 ./chan_bench
```


## Results

| Date  | Machine                             | OS                  | GOMAXPROCS=1 msg/s | GOMAXPROCS=2 msg/s |
|------:|:-------------------------------|---------------------|:------------------:|:------------------:|
| 2014-02-11 | MacBook Pro 2010 Mid Core i5 2.4GHz | OSX 10.9 Mavericks  |           11300000 |            7300000 |
| 2014-02-11 | CubieTruck                          | Fedora Kernel 3.4.75 |            910000 |            1020000 |
| 2014-02-11 | Dual Xeon E5-2680 2.70GHZ           | Ubuntu 12.04 LTS | 12400000 | 7800000 |


## Contacts

Please, send results to mio@volmy.com. Thank you.
