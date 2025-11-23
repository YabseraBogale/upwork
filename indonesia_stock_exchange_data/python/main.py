import json
import requests
import time
from urllib.request import Request,urlopen
count=0
headers = {
       'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0'
   }

url=f"https://idx.co.id/primary/TradingSummary/GetStockSummary?length=9999&start=0&date=20200106"
data=Request(url=url,headers=headers)
print(urlopen(url))
