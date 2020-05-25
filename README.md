# Blackjack Simulator

I interviewed for a company recently, and one of the questions was about how I
would design a blackjack simulator. What classes, interfaces, functionality etc.
When I was in college, I spent a summer working on such a simulator. During that
time we would visit Las Vegas and practice our card counting skills. I was never
banned from a casino, so I guess my skills were not that great.

Anyway, the company that I interviewed for also used go as part of their stack.
Inspired by the memory of that summer, and the desire to pick up go, I created
this command line blackjack simulator as my first project in go.

## Installation

This is my first go project, so I'm not sure if I have my environment set just
right. Since I use a regular user in my docker containers, I had to create the
.cache dir with the correct permissions and link it with another volume entry
before running the container. Here is the command I used:

```
    docker run \
        -i \
        --rm \
        -t \
        -u `id -u`:`id -g` \
        -v "$(pwd)/cache":/.cache \
        -v `pwd`:"/bj" \
        -v /tmp/bashrc:"/bj/.bashrc" \
        -w /bj \
        golang:latest \
        bash --rcfile /bj/.bashrc -i
```

See my
[docker-scripts repo](https://github.com/marallyn/docker-scripts "marallyn/docker-scripts")
for my current scripts for running docker.

Clone the repo and run:

```
    go install
```

## Usage

```
    bj <# of decks> <# of players> <# of hands> [quiet]
```

When not in quiet mode, the output is quite noisy, so if you are going to
simulate millions of hands, use quiet.

### Limitations

This was just for fun to get familiar with go, so it is not a full fledged
simulator.

-   Stategy is hard coded. The dealer strategy and player strategies are
    constant and embedded in the code. The plan before shelving this project is
    to learn a bit about io and maps by reading strategies from files.

-   When splitting aces, you are supposed to only get one card per hand. This
    sim allows you to do whatever you wish with your hands after splitting them.

-   Comments are missing (func documentation). I didn't read the golang
    recommendation for commenting functions, types, and interfaces, and all the
    code I saw in the docs was comment-less. So I left out documentation for
    now. Code comments exist though.

### Results

Using probabilities, you can calculate the chance of being dealt a blackjack as
about 4.83%. Running the simulator for about a million hands, the observed
chances of getting a blackjack are about 4.79%. Not sure what to attribute the
difference to, but it is kind of close.

I didn't really look at expected win, loss and push percentages, so not sure how
close those are to expectations either.
