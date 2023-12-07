import re
import sys

acc = 0
total_power = 0
bag = {"red": 12, "green": 13, "blue": 14}

for row in sys.stdin:
    game, sets = re.match(r"Game (\d+):(.+)", row.strip()).groups()

    mins = {"red": 0, "green": 0, "blue": 0}
    for set_ in sets.split(";"):
        cubes = re.findall(r"(\d+) (\w+)", set_)
        for cube in cubes:
            if bag[cube[1]] < int(cube[0]):
                acc += int(game)

            if mins[cube[1]] < int(cube[0]):
                mins[cube[1]] = int(cube[0])

    power = mins["red"] * mins["green"] * mins["blue"]
    total_power += power

print(f"Part 1: {acc}")
print(f"Part 2: {total_power}")
