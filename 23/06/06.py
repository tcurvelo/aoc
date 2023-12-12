import re
import sys

times = [int(t) for t in re.findall(r"\d+", sys.stdin.readline())]
distances = [int(t) for t in re.findall(r"\d+", sys.stdin.readline())]


def get_ways(times, distances):
    total = 1
    for time, distance in zip(times, distances):
        ways = 0
        for i in range(time):
            if i * (time - i) > distance:
                ways += 1
        total *= ways

    return total


print(f"Part 1: {get_ways(times, distances)}")

times_pt2 = [int("".join(str(t) for t in times))]
distances_pt2 = [int("".join(str(d) for d in distances))]
print(f"Part 2: {get_ways(times_pt2, distances_pt2)}")
