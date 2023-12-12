import sys
from collections import defaultdict

kinds = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
}


def kind_value(kind):
    return int(kind) if kind.isdigit() else kinds[kind]


def hand_value(hand):
    cards = list(hand)

    count = defaultdict(int)
    for kind in cards:
        count[kind] += 1

    counts = sorted(count.values(), reverse=True)

    value = 0
    match counts:
        case [5]:  # Five of a kind
            value = 6
        case [4, 1]:  # Four of a kind
            value = 5
        case [3, 2]:  # Full house
            value = 4
        case [3, 1, 1]:  # Three of a kind
            value = 3
        case [2, 2, 1]:  # Two pair
            value = 2
        case [2, 1, 1, 1]:  # One pair
            value = 1
        case [1, 1, 1, 1, 1]:  # High card
            value = 0

    tie_breaker = tuple(kind_value(c) for c in cards)
    return (value, tie_breaker)


hands = []
for row in sys.stdin:
    hand, bid = row.strip().split(" ")
    value = hand_value(hand)
    hands.append((value, hand, int(bid)))

result = sorted(hands)

acc = 0
for index, (_, _, bid) in enumerate(result, 1):
    acc += bid * index

print(acc)
