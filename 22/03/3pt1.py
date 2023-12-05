import sys


def priority(c):
    offset = ord("a") - 1 if c.islower() else ord("A") - 27
    return ord(c) - offset


acc = 0
for line in sys.stdin:
    rucksack = line.strip()
    half = len(rucksack) // 2

    comp1, comp2 = (
        set(rucksack[0:half]),
        set(rucksack[half:]),
    )
    for item in comp1.intersection(comp2):
        acc += priority(item)

print(acc)
