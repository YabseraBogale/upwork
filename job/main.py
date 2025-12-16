from telethon import TelegramClient, events

# 1. Your credentials from my.telegram.org
api_id = 1234567
api_hash = 'your_api_hash_here'

client = TelegramClient('listener_session', api_id, api_hash)

# 2. Define the channels you want to listen to (IDs or @usernames)
# Note: Use -100 prefix for channel IDs (e.g., -100123456789)
target_channels = ['@tech_news_channel', -1001987654321]

@client.on(events.NewMessage(chats=target_channels))
async def my_event_handler(event):
    # Get the text of the message
    msg_text = event.raw_text
    channel_name = event.chat.title if event.chat else "Unknown Channel"
    
    print(f"New message from {channel_name}: {msg_text}")

    # 3. Send a message to yourself (Saved Messages)
    # You can customize this to forward the message or just alert yourself
    await client.send_message(
        'me', 
        f"📢 **New Post in {channel_name}**\n\n{msg_text}"
    )

print("Listener is active... Press Ctrl+C to stop.")
client.start()
client.run_until_disconnected()