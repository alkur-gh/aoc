#!/bin/python3

import re
import numpy as np
import matplotlib.pyplot as plt

def make_rect(x0, x1, y0, y1):
    return set(
            (x, y)
            for x in range(x0, x1 + 1)
            for y in range(y0, y1 + 1)
            )

def plot_rect(rect):
    arr = np.array(list(rect))
    plt.scatter(arr[:, 0], arr[:, 1])

union = set()
regex = re.compile(
        '(?P<action>on|off) ' +
        'x=(?P<x0>-?\d+)\.\.(?P<x1>-?\d+),' +
        'y=(?P<y0>-?\d+)\.\.(?P<y1>-?\d+),' +
        'z=(?P<z0>-?\d+)\.\.(?P<z1>-?\d+)')

with open('./files/simple.txt', 'r') as f:
    for line in f.readlines()[:11]:
        m = regex.match(line)
        gs = m.groups()
        rect = make_rect(*[int(v) for v in gs[1:5]])
        if gs[0] == 'on':
            union |= rect
        else:
            union -= rect

print(len(union))
