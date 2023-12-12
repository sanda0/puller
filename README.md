# Puller - GitLab Webhook Receiver

## Overview

> This Go program serves as a webhook receiver for GitLab events. It is designed to handle incoming webhook requests, authenticate them using a secret token, and execute specified commands based on the received event type.

## Table of Contents

- [Configuration](#configuration)
- [Logging](#logging)
- [Usage](#usage)


## Configuration

The program reads its configuration from a `config.json` file. Create a `config.json` file with the required configuration. 

> Feel free to customize the configuration and commands in the `config.json` file to suit your specific project requirements.

```json
{
  "key": "your-secret-token",
  "repos": [
    {
      "name": "testrepo",
      "path": "/var/www/testrepo",
      "branch": "refs/heads/main",
      "events": [
        {
          "type": "push",
          "commands": [
            "git pull",
            "php artisan optimize:clear"
          ]
        }
      ],
      "notifications": [
        {
          "type": "email",
          "to": [
            "sandakelum@pramixit.com"
          ]
        }
      ]
    },
    {
      "name": "yoyo",
      "path": "/var/www/testrepo",
      "branch": "refs/heads/main",
      "events": [
        {
          "type": "push",
          "commands": [
            "git pull",
            "php artisan optimize:clear"
          ]
        }
      ],
      "notifications": [
        {
          "type": "email",
          "to": [
            "sandakelum@pramixit.com"
          ]
        }
      ]
    }
  ]
}

```

## Logging

The `writeLogFile` function facilitates the logging of messages to the `app.log` file based on the message type (INFO or ERROR). This log file captures execution details, providing insights into the success or failure of the executed commands.

## Usage

1. Set up a webhook in your GitLab project to communicate with the program. Point the webhook to the `/puller` endpoint on the running server where your program is listening.

2. Execute the compiled binary. 


