from telethon import TelegramClient, events
import os
from dotenv import load_dotenv
# 1. Your credentials from my.telegram.org
load_dotenv()
api_id = os.getenv("telegram_api_id")
api_hash = os.getenv("telegram_api_hash")

client = TelegramClient('listener_session', api_id, api_hash)

# 2. Define the channels you want to listen to (IDs or @usernames)
# Note: Use -100 prefix for channel IDs (e.g., -100123456789)
target_channels = ['@effoyjobs', '@freelance_ethio','@geezjobs_ethiopia',
                   '@online_jobs_ethiopia01','@BeeksisaaHojii','@ethreporterjob',
                   '@vacancy8']

search=["python","java","js","nodejs",
        "flask","dotnet",".net",
        "backend","database","full stack",
        "fullstack","software","software engineer",
        "software developer","erp developer",
        "mysql","php","larvel","express",
        "flutter","Mobile Application"]

@client.on(events.NewMessage(chats=target_channels))
async def my_event_handler(event):
    # Get the text of the message
    try:
        msg_text = event.raw_text
        for i in search:
            if str(msg_text).lower().find(i)!=-1:
                # 3. Send a message to yourself (Saved Messages)
                # You can customize this to forward the message or just alert yourself
                await client.send_message(
                    'me', 
                    msg_text
                )
    except Exception as e:
        print(e)
print("Listener is active... Press Ctrl+C to stop")
client.start()
client.run_until_disconnected()