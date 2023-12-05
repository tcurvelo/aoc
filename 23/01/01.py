import sys
import re

filename = sys.argv[1]

babel = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

with open(filename, "r") as f:
    acc1 = acc2 = 0
    for row in f.readlines():
        current = row.strip()

        # Part 1
        digits1 = re.findall(r"(\d)", current)
        acc1 += int(digits1[0] + digits1[-1])

        # Part 2
        digits2 = re.findall(
            r"(?=(\d|one|two|three|four|five|six|seven|eight|nine))",
            current,
        )
        first = digits2[0] if digits2[0].isdigit() else babel[digits2[0]]
        last = digits2[-1] if digits2[-1].isdigit() else babel[digits2[-1]]
        acc2 += int(f"{first}{last}")


print(f"Part 1: {acc1}")
print(f"Part 2: {acc2}")
