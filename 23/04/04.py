import math
import re
import sys
from collections import defaultdict

acc = 0
cards = defaultdict(lambda: 1)

for row in sys.stdin:
    match = re.match("Card\s+(\d+):\s+([\d ]+) \| ([\d ]+)", row)

    card, winner, have = match.groups()
    card = int(card)
    cards[card] += 0

    winner = set(re.split(r"\s+", winner))
    have = set(re.split(r"\s+", have))

    hits = len(winner & have)
    for copy in range(card + 1, card + hits + 1):
        cards[copy] += cards[card]

    points = int(math.pow(2, hits - 1))
    acc += points

print(acc)
print(sum(cards.values()))
