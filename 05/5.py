import re
import sys
from collections import defaultdict

stacks = defaultdict(list)
stacks_read = False

for line in sys.stdin:
    row = line.rstrip()

    if not row:
        # just read stacks
        stacks_read = True
        continue

    if row and not stacks_read:
        # reading stacks
        for index, stack in enumerate(range(0, len(line), 4)):
            if match := re.findall(r"\[(\w+)\]", row[stack : stack + 3]):
                stacks[index + 1].insert(0, match.pop())

    else:
        # # reading commands
        qty, src, dst = [
            int(i) for i in re.search(r"move (\d+) from (\d+) to (\d+)", row).groups()
        ]
        if sys.argv[1] == "pt2":
            items = stacks[src][-qty:]
            del stacks[src][-qty:]
            stacks[dst] += items
        else:
            for op in range(qty):
                stacks[dst].append(stacks[src].pop())


print(stacks)
print("".join(stacks[idx][-1] for idx in range(1, len(stacks) + 1)))
