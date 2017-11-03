package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"

    "github.com/nlopes/slack"
)

func main() {
    api := slack.New(os.Getenv("SLACK_API_TOKEN"))
    logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
    slack.SetLogger(logger)
    api.SetDebug(true)

    rtm := api.NewRTM()
    go rtm.ManageConnection()

    for msg := range rtm.IncomingEvents {
//      fmt.Print("Event Received: ")
        switch ev := msg.Data.(type) {
        case *slack.HelloEvent:
            // Ignore hello

        case *slack.ConnectedEvent:
            fmt.Println("Infos:", ev.Info)
            fmt.Println("Connection counter:", ev.ConnectionCount)
            // Replace #general with your Channel ID
            rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "gogadget-test"))

        case *slack.MessageEvent:
            if len(ev.Text) >= 9 && ev.Text[0:9] == "!set cost" {
                    handleSet(ev.Text[10:])
            } else if len(ev.Text) >= 11 && ev.Text[0:11] == "!get prices" {
                    handleGet(ev.Text[11:])
            }

        case *slack.InvalidAuthEvent:
            fmt.Printf("Invalid credentials")
            return

        default:

            // Ignore other events..
            // fmt.Printf("Unexpected: %v\n", msg.Data)
        }
    }
    // To run the update script
    cmd := exec.Command("./update.sh", "username", "command")
    cmd.Run()
}

func handleSet(a string) {
    fmt.Printf("Ran handleSet %s\n\n", a)
    return
}

func handleGet(a string) {
    fmt.Printf("Ran handleGet %s\n\n", a)
    return
}
