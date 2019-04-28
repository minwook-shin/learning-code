import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt

sns.set()

x = np.random.normal(size=100)
sns.distplot(x)

plt.show()
