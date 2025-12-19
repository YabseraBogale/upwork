import pandas as  pd

df = pd.read_html("/home/yabsera/Documents/h1b_data/Alabama_H1B_Jobs.html")

df = df[0]

print(df.columns)