import pandas as pd
import matplotlib.pyplot as plt

data = pd.read_csv("data4.csv")

mean_data = data.groupby("Iterations")["Time"].mean().reset_index()

plt.plot(mean_data["Time"], mean_data["Iterations"])
plt.xlabel("Time (seconds)")
plt.xlim(0, 0.019)
plt.ylabel("Iterations (i)")
plt.title("Process Load With HT On")
plt.grid(True)

plt.show()