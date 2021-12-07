#!/bin/python3

import pandas as pd
import matplotlib.pyplot as plt

if __name__ == '__main__':
    path = './data.csv'
    df = pd.read_csv(path)
    df.plot()
    plt.show()
