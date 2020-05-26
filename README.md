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

You can invoke the simulator by specifying the three required parameters on the
command line:

```
    bj <# of decks> <# of players> <# of hands> [quiet]
```

or you can specify a setup file:

```
    bj <setup file>
```

In the setup file you can specify the three required parameters as well as
names, starting chip values and strategies for your players. Each player can
have a different strategy, so you can test different strategies at the same
time.

A sample setup.json file is included.

When not in quiet mode, the output is quite noisy, so if you are going to
simulate millions of hands, use quiet.

## Strategy files

The ./strategies directory holds the strategy files. At this point, there is
only st-basic.json. Using st-basic.json as a template, you can create your own
strategies for testing. A player can specify which strategy they are using in
the setup.json file, by adding

```
    "strategy": "strategy-name"
```

to their configuration. The simulator looks for your strategy file in the
./strategies directory and prefixes the strategy name with "st-" and adds
".json" as the extension.

Strategy files use abbreviations for actions:

```
    d = double
    h = hit
    s = stand
    v = split
```

A strategy consists of a name and three different maps:

    - hard: The hard map is indexed by the hardValue of the player's hand.
    - pair: The pair map is indexed by the value of the first card in the player's hand.
    - soft: The soft map is indexed by the softValue of the player's hand minus 11

The simulator chooses which map to used based on the situation. It tries to use
the pair map first (hand must be only two cards and have match ShortNames), then
it chooses the soft map if there is an ace in the hand, otherwise it chooses the
hard map.

The second index of all strategy maps is the value of the dealer up card
starting with 2, and going through 11, for a total of ten actions per hand
index.

### Limitations

This was just for fun to get familiar with go, so it is not a full fledged
simulator.

-   When splitting aces, you are supposed to only get one card per hand. This
    sim allows you to do whatever you wish with your hands after splitting them.

-   func documentation is missing. I didn't read the Go recommendation for
    documenting functions, types, and interfaces, and all the code I saw in the
    docs was comment-less, so I left out documentation for now. Code comments
    exist though.

### Results

Using probabilities, you can calculate the chance of being dealt a blackjack as
about 4.83%. Running the simulator for fifty million hands, the observed chances
of getting a blackjack are about 5.03%. This went up from 4.79% after I
implemented splits. My guess is this will go done when I correctly account for
split aces.

I didn't really look at expected win, loss and push percentages, so not sure how
close those are to expectations either.
