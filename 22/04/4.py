import re
import sys

count_p1 = 0
count_p2 = 0
for line in sys.stdin:
    ranges = [int(i) for i in re.findall(r"\d+", line)]
    elf1 = set(range(ranges[0], ranges[1] + 1))
    elf2 = set(range(ranges[2], ranges[3] + 1))
    if elf1.issubset(elf2) or elf2.issubset(elf1):
        count_p1 += 1

    if elf1.intersection(elf2) or elf2.intersection(elf1):
        count_p2 += 1

print(f"Part 1: {count_p1}")
print(f"Part 2: {count_p2}")
