## Kata introduction

In mathematics, the Fibonacci numbers, commonly denoted Fn form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1.

The goal in this kata is to write a function that takes one parameter called **steps**, and returns a number from the Fibonacci sequence, based on the parameter **steps**, which determines the position in Fibonacci number.

## Getting started: Docker environment

Launch the Docker container in to interactive mode.

```bash
cd fibonacci
workdir="/go/src/github.com/golang-katas/fibonacci"
docker run --rm -it \
    --name go-fibonacci \
    --volume $PWD:$workdir \
    --workdir $workdir \
    golang
```

Run the tests at the command line to make sure everything is green.

```bash
go test ./test -v
```

Create the module and build the binaries.

```bash
make build
```

Run the program to make sure the correct Fibonacci sequence numbers are returned.

```bash
./bin/kata
```

After running the above, you should get the following results (20 steps.)

```bash
0 is the Fibonnaci result for 0 steps
1 is the Fibonnaci result for 1 steps
1 is the Fibonnaci result for 2 steps
2 is the Fibonnaci result for 3 steps
3 is the Fibonnaci result for 4 steps
5 is the Fibonnaci result for 5 steps
8 is the Fibonnaci result for 6 steps
13 is the Fibonnaci result for 7 steps
21 is the Fibonnaci result for 8 steps
34 is the Fibonnaci result for 9 steps
55 is the Fibonnaci result for 10 steps
89 is the Fibonnaci result for 11 steps
144 is the Fibonnaci result for 12 steps
233 is the Fibonnaci result for 13 steps
377 is the Fibonnaci result for 14 steps
610 is the Fibonnaci result for 15 steps
987 is the Fibonnaci result for 16 steps
1597 is the Fibonnaci result for 17 steps
2584 is the Fibonnaci result for 18 steps
4181 is the Fibonnaci result for 19 steps
```

Experiment by editing the Go code in `cmd` and `pkg` packages, then rinse and repeat the steps above to rebuild and run.

## Getting started: local environment

Before you can run this locally, you'll need Go and build tooling installed.

Run the tests to make sure everything is green.

```bash
cd fibonacci
go test ./test -v
```

Create the module and build the binaries. Be sure to change the build architecture to match your system. In `Makefile`, change the `GOOS=linux` value to match to your environment (`darwin` if you're on macOS, and so on.)

```bash
make build
```

Run the program to make sure the correct Fibonacci sequence numbers are returned.

```bash
./bin/kata
```

After running the above, you should get the following results (20 steps.)

```bash
0 is the Fibonnaci result for 0 steps
1 is the Fibonnaci result for 1 steps
1 is the Fibonnaci result for 2 steps
2 is the Fibonnaci result for 3 steps
3 is the Fibonnaci result for 4 steps
5 is the Fibonnaci result for 5 steps
8 is the Fibonnaci result for 6 steps
13 is the Fibonnaci result for 7 steps
21 is the Fibonnaci result for 8 steps
34 is the Fibonnaci result for 9 steps
55 is the Fibonnaci result for 10 steps
89 is the Fibonnaci result for 11 steps
144 is the Fibonnaci result for 12 steps
233 is the Fibonnaci result for 13 steps
377 is the Fibonnaci result for 14 steps
610 is the Fibonnaci result for 15 steps
987 is the Fibonnaci result for 16 steps
1597 is the Fibonnaci result for 17 steps
2584 is the Fibonnaci result for 18 steps
4181 is the Fibonnaci result for 19 steps
```

Experiment by editing the Go code in `cmd` and `pkg` packages, then rinse and repeat the steps above to rebuild and run.
