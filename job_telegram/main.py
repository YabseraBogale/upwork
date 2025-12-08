from telethon import TelegramClient, events
import asyncio
import os

API_ID = os.getenv("API_ID") # Replace with your ID
API_HASH = os.getenv("API_HASH") # Replace with your Hash
SESSION_NAME = 'job_telegram'
CHANNEL_USERNAME = 't.me/effoyjobs' # e.g., @Telegram

# Create the client instance and connect
client = TelegramClient(SESSION_NAME, API_ID, API_HASH)

@client.on(events.NewMessage(chats=CHANNEL_USERNAME))
async def new_message_handler(event):
    print(f"New message received: {event.message.text}")
    # Process the message here

async def main():
    await client.start()
    # The first time you run this, it will ask for your phone number and login code.
    print("Client running...")

    await client.run_until_disconnected()

if __name__ == '__main__':
    with client:
        client.loop.run_until_complete(main())