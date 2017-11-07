# Installation
First clone the repo and run the following commands:
```bash
# The Slack API
$ go get -u github.com/nlopes/slack

# The YAML package
$ go get -u gopkg.in/yaml.v2
```

Now you can build the bot by running the command: `$ go build`\
This will create a file called `slackbots` which can be run from the command line. Be sure
to check out the [Environment variables](#environment-variables) section.

# Usage
First add the bot to your workspace, afterwards invite the bot to the relevant channels.
The bot reacts to three commands:
```
!create product [name]
!set cost [cost] [name]
!get prices [name]
```

The first command creates a product with the given name. The name may contain spaces.\
The second command sets the price of a given product to the cost in ører.\
The third command makes the bot send a message to the channel with the prices for the
product.

# Environment variables
- `KANTINE_DB` - must be set for bot to work
- `SLACK_API_TOKEN` - must be set for Slack to work
- `GO_ENV` - should be set to 'production' in production mode to disable logging

# APIs
The api we are using:

https://github.com/nlopes/slack
https://github.com/go-yaml/yaml/tree/v2