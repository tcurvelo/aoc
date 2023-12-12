import re
import sys


def interval_intersect(a, b):
    if b[0] > a[1] or a[0] > b[1]:
        return None
    else:
        return max(a[0], b[0]), min(a[1], b[1])


def interval_fillings(main: tuple, sub: list):
    pieces = []
    last = main[0]
    for s in sub:
        if last < s[0]:
            pieces.append((last, s[0] - 1))
        last = s[1] + 1
    if last < main[1]:
        pieces.append((last, main[1]))

    return pieces


def get_location(seed, maps):
    intervals = [seed]

    for _, rules in maps.items():
        next_interval = []
        matched = []
        for interval in intervals:
            for rule in rules:
                if inter := interval_intersect(interval, (rule[1], rule[1] + rule[2])):
                    matched.append(inter)

            filled = interval_fillings(interval, matched)
            next_interval.append(sorted(filled + matched))

        intervals = next_interval

    print(intervals)


def parse_input():
    seeds = [int(s) for s in re.findall(r"(\d+)", sys.stdin.readline())]

    maps = {}
    while row := sys.stdin.readline():
        if not (title := row.strip().strip(":")):
            continue
        title = re.findall(r"-to-(.*?) map", title).pop()
        maps[title] = []
        while row := sys.stdin.readline().strip():
            maps[title].append(tuple(int(i) for i in row.split()))

    return seeds, maps


def part1(seeds, maps):
    locations = []
    for s in seeds:
        cursor = s
        for title, rules in maps.items():
            value = None
            for rule in rules:
                if rule[1] <= cursor <= rule[1] + rule[2]:
                    value = cursor - rule[1] + rule[0]
                    break

            if value is not None:
                cursor = value

        locations.append(cursor)
    return locations


def expand_seeds(seeds):
    return [(start, start + length - 1) for start, length in zip(seeds[::2], seeds[1::2])]


# def part2(seeds, maps):
#     for seed_range in seeds:
#         get_location(seed_range, maps)


if __name__ == "__main__":
    seeds, maps = parse_input()
    print(f"Part 1: {min(part1(seeds, maps))}")

    # expanded_seeds = expand_seeds(seeds)
    # print(f"Part 2: {min(part2(expanded_seeds, maps))}")
    # print(f"Part 2: {part2(expanded_seeds, maps)}")
