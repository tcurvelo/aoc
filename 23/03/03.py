import re
import sys
from collections import defaultdict

symbol_re = re.compile(r"[^0-9.]")
text = [row.strip() for row in sys.stdin.readlines()]

gears = defaultdict(list)


def save_gear(symbol_match, offset, row_index, number):
    if symbol_match.group() == "*":
        gears[(row_index, symbol_match.start() + offset)].append(number)


def check_horizontal(match, row, row_index):
    found = False

    if match.start() > 0:
        if symbol := symbol_re.match(row[match.start() - 1]):
            found = True
            save_gear(symbol, match.start() - 1, row_index, match.group())

    if match.end() < len(row):
        if symbol := symbol_re.match(row[match.end()]):
            found = True
            save_gear(symbol, match.end(), row_index, match.group())

    return found


def check_vertical(match, row_index):
    found = False
    if row_index > 0:
        start = max(match.start() - 1, 0)
        end = min(match.end() + 1, len(text[row_index - 1]))

        for symbol in symbol_re.finditer(text[row_index - 1][start:end]):
            found = True
            save_gear(symbol, start, row_index - 1, match.group())

    if row_index < (len(text) - 1):
        start = max(match.start() - 1, 0)
        end = min(match.end() + 1, len(text[row_index - 1]))

        for symbol in symbol_re.finditer(text[row_index + 1][start:end]):
            found = True
            save_gear(symbol, start, row_index + 1, match.group())

    return found


def main():
    acc_pt1 = acc_pt2 = 0
    for index, row in enumerate(text):
        for number in re.finditer(r"\d+", row):
            if any([check_horizontal(number, row, index), check_vertical(number, index)]):
                acc_pt1 += int(number.group())

    for gear in gears.values():
        if len(gear) == 2:
            ratio = int(gear[0]) * int(gear[1])
            acc_pt2 += ratio

    print(f"Part 1: {acc_pt1}")
    print(f"Part 2: {acc_pt2}")


if __name__ == "__main__":
    main()
