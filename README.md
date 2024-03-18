This repo contains two Go programs:

- `duotrigordle-all` suggests a first guess for Duotrigordle which maximises the expected number of _solvable_ boards
- `duotrigordle-specific` reports _solvable_ boards for a given sequence of one or more guesses

A board is _solvable_ if there is only one remaining possible guess.

Guess responses are colour-coded with ANSI terminal codes.

## Instructions

To build:

```
go build -o . ./...
```

To run:

```
./duotrigordle-all
```

```
./duotrigordle-specific guess [guess...]
```

## Results

### Best first guess

```
$ ./duotrigordle-all
Dictionaries: 14855 valid guesses, 2653 solutions
Best guess: BURET
14 of 2653 boards (0.5%) are solvable on next guess
15.6% chance of solvable board on grid of 32
899.2 average possible solutions per board
```

### Specific guess sequences

```
$ ./duotrigordle-specific buret
Dictionaries: 14855 valid guesses, 2653 solutions
Guesses: [BURET]
BURET BERET
BURET BERTH
BURET BITER
BURET BLUER
BURET BRUTE
BURET BUYER
BURET DEBUT
BURET REBUT
BURET RUBLE
BURET STRUT
BURET THROB
BURET TRIBE
BURET TUBER
BURET TURBO
14 of 2653 boards (0.5%) are solvable on next guess
15.6% chance of solvable board on grid of 32
899.2 average possible solutions per board
```

```
$ ./duotrigordle-specific soare until
Dictionaries: 14855 valid guesses, 2653 solutions
Guesses: [SOARE UNTIL]
SOARE UNTIL AFOUL
[...snip...]
SOARE UNTIL WHORL
203 of 2653 boards (7.7%) are solvable on next guess
92.2% chance of solvable board on grid of 32
26.3 average possible solutions per board
```

```
$ ./duotrigordle-specific files prang dutch
Dictionaries: 14855 valid guesses, 2653 solutions
Guesses: [FILES PRANG DUTCH]
FILES PRANG DUTCH ACTED
[...snip...]
FILES PRANG DUTCH WRONG
922 of 2653 boards (34.8%) are solvable on next guess
100.0% chance of solvable board on grid of 32
4.7 average possible solutions per board
```
