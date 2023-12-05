import sys
from functools import reduce


def priority(c):
    offset = ord("a") - 1 if c.islower() else ord("A") - 27
    return ord(c) - offset


acc = 0
elf_count = 0
group = []
for line in sys.stdin:
    group.append(set(line.strip()))
    elf_count += 1
    if elf_count < 3:
        continue
    else:
        common = reduce(lambda e, i: e.intersection(i), group).pop()
        acc += priority(common)
        group = []
        elf_count = 0

print(acc)
