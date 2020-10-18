# employment-increase-bot

When I was working at nata.house, everytime we had a new hire, this person was added to our Telegram group and then everyone sent a welcome message.

One day, I sent a sticker from a funny bird pack that I had found. It was a bird wearing a tie saying "Employment Increase". This became a internal joke and and still, a year later, whenever someone is hired, this sticker is sent.

I created this bot to start to learn Golang and make life easier for anyone who wants to continue this joke 😂  

Built using only Golang standard library and relying on Telegram Bots webhook. 

## How to run
```
TELEGRAM_TOKEN=YOUR_BOT_TELEGRAM_TOKEN_HERE go run pkg/handler/main.go
```

It will start on port 3000, you can use a tool like [ngrok](https://ngrok.com/) to serve this port to an external URL, so Telegram can access it!
