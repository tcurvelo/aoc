import re
import sys

lookup = {"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
acc_pt1 = acc_pt2 = 0

for row in sys.stdin:
    # Part 1
    digits_pt1 = re.findall(r"(\d)", row)
    acc_pt1 += int(digits_pt1[0] + digits_pt1[-1])

    # Part 2
    digits_pt2 = re.findall(r"(?=(\d|one|two|three|four|five|six|seven|eight|nine))", row)
    first = digits_pt2[0] if digits_pt2[0].isdigit() else lookup[digits_pt2[0]]
    last = digits_pt2[-1] if digits_pt2[-1].isdigit() else lookup[digits_pt2[-1]]
    acc_pt2 += int(f"{first}{last}")


print(f"Part 1: {acc_pt1}")
print(f"Part 2: {acc_pt2}")
