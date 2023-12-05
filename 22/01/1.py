import sys

filename = sys.argv[1]

calories = 0
calories_count = []
with open(filename, "r") as f:
    for row in f.readlines():
        current = row.strip()
        if not current:  # empty line
            calories_count.append(calories)
            calories = 0
        else:
            calories += int(current)


sorted_results = sorted(calories_count, reverse=True)
print(f"Part 1: {sorted_results[0]}")
print(f"Part 2: {sum(sorted_results[:3])}")
