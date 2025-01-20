# Advent of Code 2024

### How to run

Run from the project root and use flags to select day, part and dataset.

`go run . -d 12 -p 2 -s 1`

`-d` or `-day` to select the day, from 1 to 25

`-p` or `-puzzle` to select the first or second puzzle, expects 1 or 2 (defaults to 1)

`-s` or `-sample` to select sample data, expects 1+ (most days only have 1 sample dataset, some have more) - omit this flag to run on the full dataset

---

### Disclaimers

I have been learning Go in my spare time for about a year, coming from being a PHP/WordPress/web developer. I mostly tried to solve these problems on my own, therefore they don't necessarily represent the best possible ways to solve, or idiomatic Go code, as I am still learning. Where I have looked for help for others on the Reddit I have made comments in the code to give credit and thanks to those individuals for sharing their solutions.

At time of writing the following days are not completely solved:

##### Day 9

Runs correctly on all the samples I could find but seems to produce an answer that is off by a minute fraction on the full dataset.

##### Day 17

My nemesis! I think just running my part 2 solution for many, many hours would eventually find the answer - but as it stands I am exploring other options, ie starting to learn some CUDA.

##### Day 25

Can't complete until I finish day 17...
