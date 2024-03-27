# slack-notifier-go

Golang program designed to receive standard output and send it as notifications to Slack.

## Installation

Build the executable:

```
go build
```

Create a configuration file named `config.yaml` in the same directory as the executable with the following content:

```yaml
$ cat << EOF > config.yaml
webhookURL: "YOUR_SLACK_WEBHOOK_URL_HERE"
EOF
```

Replace `YOUR_SLACK_WEBHOOK_URL_HERE` with your Slack webhook URL.

## Usage

Once you have the executable and configuration set up, you can use **slack-notifier-go** by piping the output of another command to it.

For example:

```
$ echo "Hello World" | ./slack-notifier-go
```

## Deployment to Server

To deploy **slack-notifier-go** to your server, follow these command-line steps:

1. Create a directory for the notifier application:

    ```
    # mkdir -p /opt/bin/notify
    ```

2. Copy the **slack-notifier-go** executable to the newly created directory:

    ```
    # cp slack-notifier-go /opt/bin/notify
    ```

3. Copy the configuration file (`config.yaml`) to the same directory:

    ```
    # cp config.yaml /opt/bin/notify
    ```

**slack-notifier-go** and its configuration file are both located in the `/opt/bin/notify` directory on your server.

You can execute the program from this directory and ensure it has access to the necessary configuration.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
