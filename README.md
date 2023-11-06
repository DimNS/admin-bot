# Geeksonator Bot for PanteleevGroup chats

[![Go Report Card](https://goreportcard.com/badge/github.com/DimNS/admin-bot?style=flat-square)](https://goreportcard.com/report/github.com/DimNS/admin-bot)
[![Release](https://img.shields.io/github/release/DimNS/admin-bot.svg?style=flat-square)](https://github.com/DimNS/admin-bot/releases/latest)
[![Audit](https://github.com/DimNS/admin-bot/workflows/audit/badge.svg?branch=master)](https://github.com/DimNS/admin-bot/actions?query=workflow%3A"audit")
[![License](https://img.shields.io/github/license/DimNS/admin-bot.svg)](https://github.com/DimNS/admin-bot)

## Install

```
echo "GEEKSONATOR_TELEGRAM_BOT_TOKEN=\"bot_token_here\"" >> /etc/geeksonator.conf
echo "GEEKSONATOR_TELEGRAM_TIMEOUT_SECONDS=\"15\"" >> /etc/geeksonator.conf
echo "GEEKSONATOR_DEBUG_MODE=\"false\"" >> /etc/geeksonator.conf
sudo chmod 755 /opt/geeksonator
sudo cp geeksonator.service /lib/systemd/user
sudo systemctl enable /lib/systemd/user/geeksonator.service
sudo service geeksonator start
```

Also, the bot must disable Privacy mode (in BotFather) before being included in groups (otherwise it will not have access to messages to do reply)

#### Defaults
- `GEEKSONATOR_TELEGRAM_TIMEOUT_SECONDS` = `15`
- `GEEKSONATOR_DEBUG_MODE` = `false`

## Run in debug mode

```
export GEEKSONATOR_TELEGRAM_BOT_TOKEN="bot_token_here"
export GEEKSONATOR_TELEGRAM_TIMEOUT_SECONDS="5"
export GEEKSONATOR_DEBUG_MODE="true"
/opt/geeksonator
```
