import sys

ROCK = 1
PAPER = 2
SCISSOR = 3
LOSS = 0
DRAW = 3
WIN = 6

mapping1 = {
    "A": {  # ROCK
        "X": ROCK + DRAW,
        "Y": PAPER + WIN,
        "Z": SCISSOR + LOSS,
    },
    "B": {  # PAPER
        "X": ROCK + LOSS,
        "Y": PAPER + DRAW,
        "Z": SCISSOR + WIN,
    },
    "C": {  # SCISSOR
        "X": ROCK + WIN,
        "Y": PAPER + LOSS,
        "Z": SCISSOR + DRAW,
    },
}

mapping2 = {
    "A": {  # ROCK
        "X": LOSS + SCISSOR,
        "Y": DRAW + ROCK,
        "Z": WIN + PAPER,
    },
    "B": {  # PAPER
        "X": LOSS + ROCK,
        "Y": DRAW + PAPER,
        "Z": WIN + SCISSOR,
    },
    "C": {  # SCISSOR
        "X": LOSS + PAPER,
        "Y": DRAW + SCISSOR,
        "Z": WIN + ROCK,
    },
}
with open(sys.argv[1], "r") as f:
    acc1 = 0
    acc2 = 0
    for row in f.readlines():
        their, expected = row.strip().split(" ")
        current1 = mapping1[their][expected]
        current2 = mapping2[their][expected]

        acc1 += current1
        acc2 += current2


print(f"Part 1: {acc1}")
print(f"Part 2: {acc2}")
