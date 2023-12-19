# Puller - GitLab Webhook Receiver

## Overview

> This Go program serves as a webhook receiver for GitLab events. It is designed to handle incoming webhook requests, authenticate them using a secret token, and execute specified commands based on the received event type.

## Table of Contents

- [Configuration](#configuration)
- [Logging](#logging)
- [Usage](#usage)
- [Todo](#todo)


## Configuration

The program reads its configuration from a `config.json` file. Create a `config.json` file with the required configuration. 

> Feel free to customize the configuration and commands in the `config.json` file to suit your specific project requirements.

```json
{
    "email": "email@example.com",
    "email_password": "password",
    "key": "<key>",
    "repos": [
        {
            "branch": "branch-name",
            "events": [
                {
                    "commands": [
                        "git pull"
                    ],
                    "type": "push"
                }
            ],
            "name": "repo-name",
            "path": "local-path"
        }
    ],
    "smtp_host": "smtp.pramixit.com",
    "smtp_port": "587"
}


```

## Logging

The `writeLogFile` function facilitates the logging of messages to the `app.log` file based on the message type (INFO or ERROR). This log file captures execution details, providing insights into the success or failure of the executed commands.

## Usage

```bash
#download
wget https://github.com/sanda0/puller/releases/download/v1.6/puller

##
sudo chmod +x puller

### init
sudo ./puller -i




```

## TODO List

- [x] Logging the success or failure of the executed commands.
- [x] E-mail notifications.
- [ ] Discord notifications.


