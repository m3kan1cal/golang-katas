## Kata introduction

Taken from : http://codingdojo.org/kata/FizzBuzz/. Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz.” For numbers which are multiples of both three and five print “FizzBuzz."

## Getting started

Create the module.

```bash
cd fizzbuzz
chmod +x ./gomod.sh
./gomod.sh
```

Change the build architecture to match your system. In `Makefile`, change the `GOOS=darwin` value to match to your environment (`linux` if you're on Linux, and so on.)

Build the binaries.

```bash
make build
```

Run the binaries.

```bash
./bin/kata
```
