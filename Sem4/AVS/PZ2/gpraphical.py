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

    threads = df["Thread"].nunique()

    return cum_time, iters, threads



file_on  = "data_ht_on.csv"
file_off = "data_ht_off.csv"

cum_on,  iters_on,  threads_on  = load_bench_data(file_on)
cum_off, iters_off, threads_off = load_bench_data(file_off)

fig, ax = plt.subplots(figsize=(12, 7))

ax.set_title("Calculating Macloren sequnce SMT ON and SMT OFF")

ax.plot(cum_on, iters_on, color="blue", label=f"SMT ON ({threads_on} threads)")
ax.plot(cum_off, iters_off, color="red", label=f"SMT OFF ({threads_off} threads)")

ax.set_xlabel("Time (microseconds)")
ax.set_ylabel("Iterations ")
plt.grid(True, linestyle='--', alpha=0.7)
plt.legend()


plt.tight_layout()
plt.savefig("smt_comparison_result.png")
# print("График сохранен как smt_comparison_result.png")
plt.show()