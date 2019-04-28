import seaborn as sns
import matplotlib.pyplot as plt

sns.set()

tips = sns.load_dataset("tips")
sns.regplot(x="total_bill", y="tip", data=tips)

plt.show()
