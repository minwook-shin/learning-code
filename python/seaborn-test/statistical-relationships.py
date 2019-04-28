import matplotlib.pyplot as plt
import seaborn as sns

sns.set()

tips = sns.load_dataset("tips")
sns.relplot(x="total_bill", y="tip", data=tips)

plt.show()
