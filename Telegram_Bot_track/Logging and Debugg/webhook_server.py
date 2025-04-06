from flask import Flask, request
import logging
from logging.handlers import RotatingFileHandler
import requests
import os

os.makedirs("logs", exist_ok=True)

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

file_handler = RotatingFileHandler("logs/bot.log", maxBytes=1000000, backupCount=5)
file_formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')
file_handler.setFormatter(file_formatter)
logger.addHandler(file_handler)

console_handler = logging.StreamHandler()
console_handler.setFormatter(file_formatter)
logger.addHandler(console_handler)

BOT_TOKEN = "THE_TOKEN_HERE"
WEBHOOK_URL = "YOUR_WEBHOOK_URL"

if BOT_TOKEN and WEBHOOK_URL:
    try:
        response = requests.post(
            f"https://api.telegram.org/bot{BOT_TOKEN}/setWebhook",
            data={"url": WEBHOOK_URL}
        )
        logger.info("Webhook set: %s", response.json())
    except Exception as e:
        logger.error("Error setting webhook: %s", e)
else:
    logger.warning("BOT_TOKEN or WEBHOOK_URL not set!")

app = Flask(__name__)
TELEGRAM_API_URL = f"https://api.telegram.org/bot{BOT_TOKEN}"

@app.route('/', methods=['GET'])
def home():
    return "Webhook server running!"

@app.route('/webhook', methods=['POST'])
def telegram_webhook():
    data = request.get_json()
    logger.info("Received update: %s", data)

    try:
        if "message" in data:
            chat_id = data["message"]["chat"]["id"]
            text = data["message"].get("text", "")
            reply_text = f"You said: {text}"
            
            logger.info("User %s said: %s", chat_id, text)
            logger.info("Replied to %s with: %s", chat_id, reply_text)

            requests.post(f"{TELEGRAM_API_URL}/sendMessage", json={
                "chat_id": chat_id,
                "text": reply_text
            })
    except Exception as e:
        logger.error("Error processing update: %s", e)

    return 'OK', 200

if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    app.run(host="0.0.0.0", port=port)
