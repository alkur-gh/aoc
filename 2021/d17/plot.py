#!/bin/python3

import sys
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from matplotlib.patches import Rectangle
from matplotlib.widgets import Slider

XLIM = (0, 500)
YLIM = (-250, 250)
XSTART, XEND, YSTART, YEND = 211, 232, -124, -69
#XSTART, XEND, YSTART, YEND = 20, 30, -10, -5
RECT = Rectangle((XSTART, YSTART), abs(XEND - XSTART), abs(YEND - YSTART))
LIMIT = (XEND, YSTART)

def simulate(v):
    limit = LIMIT
    p = (0, 0)
    points = [p]
    while p[0] <= limit[0] and p[1] >= limit[1]:
        p = (p[0] + v[0], p[1] + v[1])
        v = (max(v[0] - 1, 0), v[1] - 1)
        points.append(p)
    return np.array(points)

def inside_rect(data):
    return len(data[(data[:, 0] >= XSTART) & (data[:, 0] <= XEND)
        & (data[:, 1] >= YSTART) & (data[:, 1] <= YEND)]) > 0


with open('example.data', 'r') as f:
    want = set()
    for v in f.readline().split():
        a, b = map(int, v.split(','))
        want.add((a, b))

count = 0
ans = set()
for x in range(1, XEND + 1):
    for y in range(YSTART, 200):
        data = simulate((x, y))
        inside = inside_rect(data)
        if inside:
            ans.add((x, y))
            count += 1

print(count)
#print(want - ans)

sys.exit(0)

data = simulate((7, -1))

fig, ax = plt.subplots()
plt.subplots_adjust(bottom=0.25)
plt.xlim(XLIM)
plt.ylim(YLIM)
ax.add_patch(RECT)
sc = ax.scatter(data[:, 0], data[:, 1], color='red')

x_slide = plt.axes([0.25, 0.1, 0.65, 0.03])
y_slide = plt.axes([0.25, 0.05, 0.65, 0.03])
x_velocity = Slider(x_slide, 'x velocity', 0, 100, valinit=6, valstep=1)
y_velocity = Slider(y_slide, 'y velocity', -100, 100, valinit=6, valstep=1)

def update(val):
    v = (x_velocity.val, y_velocity.val)
    data = simulate(v)
    inside = inside_rect(data)
    print(v, ':', data[:, 1].max(), inside)
    sc.set_offsets(data)
    fig.canvas.draw()

x_velocity.on_changed(update)
y_velocity.on_changed(update)

plt.show()
