import sys
from itertools import permutations
import numpy as np
import matplotlib.pyplot as plt

with open('./files/handout.txt', 'r') as f:
    scanners = []
    beacons = []
    for line in f:
        if line.startswith('---'):
            beacons = []
            continue
        elif line.strip() == '':
            scanners.append(beacons.copy())
            continue

        x, y, z = map(int, line.strip().split(','))
        beacons.append((x, y, z))
    
    scanners.append(beacons)

basic = np.array([
    [[1, 0, 0], [0, 1, 0], [0, 0, 1]],      # i + j + k
    [[-1, 0, 0], [0, 1, 0], [0, 0, 1]],     # -i + j + k
    [[1, 0, 0], [0, -1, 0], [0, 0, 1]],     # i - j + k
    [[1, 0, 0], [0, 1, 0], [0, 0, -1]],     # i + j - k
#    [[-1, 0, 0], [0, -1, 0], [0, 0, 1]],    # -i - j + k
#    [[-1, 0, 0], [0, 1, 0], [0, 0, -1]],    # -i + j - k
#    [[1, 0, 0], [0, -1, 0], [0, 0, -1]],    # i - j - k
#    [[-1, 0, 0], [0, -1, 0], [0, 0, -1]],   # -i - j - k
])

v0 = np.array(sorted(scanners[0]))
v1 = np.array(sorted(scanners[1]))
v = v1 @ basic[0]

ts = np.concatenate(tuple(basic[:, p] for p in permutations([0, 1, 2])))
#print(ts)
#print(len(ts))

#print(v0[0] == v1[1])

sys.exit(0)

fig = plt.figure()
ax = fig.add_subplot(projection='3d')
ax.scatter(v0[:, 0], v0[:, 1], v0[:, 2], label='0')
ax.scatter(v1[:, 0], v1[:, 1], v1[:, 2], label='1')
ax.scatter(v[:, 0], v[:, 1], v[:, 2], label='v')
ax.legend()

plt.show()
