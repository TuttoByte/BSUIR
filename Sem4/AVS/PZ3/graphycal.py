import sys
import os
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker


def load_bench_data(path):
    df = pd.read_csv(path, comment="#")

    df = df.sort_values("Iteration")

    iters = df["Iteration"].to_numpy()
    times = df["Time_us"].to_numpy()
    cum_time = np.cumsum(times)

    return cum_time, iters



file = "data_gpu.csv"

cum, iters = load_bench_data(file)

fig, ax = plt.subplots(figsize=(12, 7))

ax.set_title("Calculating Macloren sequnce with GPU RTX 3050")

ax.plot(cum, iters, color="red", label=f"GPU")

ax.set_xlabel("Time (microseconds)")
ax.set_ylabel("Iterations ")
plt.grid(True, linestyle='--', alpha=0.7)
plt.legend()


plt.tight_layout()
plt.savefig("gpu_result.png")
plt.show()