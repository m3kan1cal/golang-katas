## Kata introduction

In mathematics, the Fibonacci numbers, commonly denoted Fn form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1.

The goal in this kata is to write a function that takes one parameter called **steps**, and returns a number from the Fibonacci sequence, based on the parameter **steps**, which determines the position in Fibonacci number.

## Getting started: Docker environment

Launch the Docker container in to interactive mode.

```bash

```

Run the tests at the command line to make sure everything is green.

```bash
cd fibonacci
go test ./test -v
```

Create the module by changing the build architecture to match your system. In `Makefile`, change the `GOOS=darwin` value to match to your environment (`linux` if you're on Linux, and so on.)

```bash
cd fibonacci
make build
```

Run the program to make sure the correct Fibonacci sequence numbers are returned.

```bash
./bin/kata
```

## Getting started: local environment

Run the tests to make sure everything is green.

```bash
cd fibonacci
go test ./test -v
```

Create the module by changing the build architecture to match your system. In `Makefile`, change the `GOOS=darwin` value to match to your environment (`linux` if you're on Linux, and so on.)

```bash
cd fibonacci
make build
```

Run the program to make sure the correct Fibonacci sequence numbers are returned.

```bash
./bin/kata
```
