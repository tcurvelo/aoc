import sys

LENGTH = 14 if sys.argv[1] == "pt2" else 4

for line in sys.stdin:
    for i in range(len(line.strip())):
        if len(set(line[i : i + LENGTH])) == LENGTH:
            print(i + LENGTH)
            break
