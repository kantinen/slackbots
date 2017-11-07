package main

import (
    "fmt"
    "log"
    "strings"
    "os"
    "os/exec"
    "strconv"
    "github.com/nlopes/slack"
)

func main() {
    pullDB()

    api := slack.New(os.Getenv("SLACK_API_TOKEN"))

    if strings.ToLower(os.Getenv("GO_ENV")) != "production" {
        logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
        slack.SetLogger(logger)
        api.SetDebug(true)
    }

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

        case *slack.MessageEvent:
            if len(ev.Text) > 9 && ev.Text[0:9] == "!set cost" {
                err := handleSet(ev, api)
                if err != nil {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                        fmt.Sprintf("'%s'\n%v\nUsage: !set cost [cost] [name]",
                            ev.Text[10:],
                            err),
                        ev.Msg.Channel))
                } else {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                        fmt.Sprintf("'%s'\nProduct price sucessfully set", ev.Text[10:]),
                        ev.Msg.Channel))
                }
            } else if len(ev.Text) > 11 && ev.Text[0:11] == "!get prices" {
                costs, err := handleGet(ev, api)
                if err != nil {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                        fmt.Sprintf("'%s'\n%v\nUsage: !get prices [name]",
                            ev.Text[12:],
                            err),
                        ev.Msg.Channel))
                } else {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                        fmt.Sprintf("'%s'\n%+v\n", ev.Text[12:], costs),
                        ev.Msg.Channel))
                }
            } else if len(ev.Text) > 15 && ev.Text[0:15] == "!create product" {
                err := handleCreate(ev, api)
                if err != nil {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                            fmt.Sprintf("'%s'\n%v\nUsage: !create product [cost] [name]",
                                ev.Text[16:],
                                err),
                            ev.Msg.Channel))
                } else {
                    rtm.SendMessage(rtm.NewOutgoingMessage(
                            fmt.Sprintf("'%s'\nCreated the new product", ev.Text[16:]),
                            ev.Msg.Channel))
                }
            } else if len(ev.Text) == 5 && ev.Text == "!help" {
                rtm.SendMessage(rtm.NewOutgoingMessage(
                            "Usage: !create product [cost] [name]\n" +
                            "Usage: !set cost [cost] [name]\n" +
                            "Usage: !get prices [name]\n" +
                            "Usage: !help - display this help and exit",
                            ev.Msg.Channel))
            }

        case *slack.InvalidAuthEvent:
            fmt.Printf("Invalid credentials")
            return

        default:

            // Ignore other events..
            // fmt.Printf("Unexpected: %v\n", msg.Data)
        }
    }
}

func handleSet(ev *slack.MessageEvent, api *slack.Client) error{
    err := pullDB()
    if err != nil {
        return err
    }

    text := ev.Text[10:]

    commands := strings.Split(text, " ")

    cost, err := strconv.Atoi(commands[0])
    if err != nil {
        return err
    }

    name := strings.Join(commands[1:], " ")

    err = setProductPrice(name, cost)
    if err != nil {
        return err
    }

    err = updateDB(ev, api)
    if err != nil {
        return err
    }
    return nil
}

func handleGet(ev *slack.MessageEvent, api *slack.Client) (Costs, error){
    err := pullDB()
    if err != nil {
        return Costs{}, err
    }

    name := ev.Text[12:]

    costs, err := getProductPrice(name)
    if err != nil {
        return Costs{}, err
    }

    err = updateDB(ev, api)
    if err != nil {
        return Costs{}, err
    }
    return costs, nil
}

func handleCreate(ev *slack.MessageEvent, api *slack.Client) error {
    err := pullDB()
    if err != nil {
        return err
    }

    text := ev.Text[16:]

    commands := strings.Split(text, " ")

    cost, err := strconv.Atoi(commands[0])
    if err != nil {
        return err
    }

    name := strings.Join(commands[1:], " ")

    err = createProduct(name, cost)
    if err != nil {
        return err
    }

    err = updateDB(ev, api)
    if err != nil {
        return err
    }
    return nil
}

func updateDB(ev *slack.MessageEvent, api *slack.Client) error{
    // To run the update script
    cmd := exec.Command("./updateDB.sh", getUsername(ev, api), ev.Text)
    return cmd.Run()
}

func pullDB() error{
    cmd := exec.Command("./pullDB.sh")
    return cmd.Run()
}

func getUsername(ev *slack.MessageEvent, api *slack.Client) string{
    user, _ := api.GetUserInfo(ev.Msg.User)
    return user.RealName
}