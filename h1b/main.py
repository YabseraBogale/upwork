import pandas as  pd

df = pd.read_html("")

df = df[0]

df =df[df["h_1b_dependent"]=="No"]
print("Number of h1b dependent saying NO is:",len(df))

print(df["job_title"].unique)
