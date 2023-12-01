package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/spf13/pflag"
)

var (
	keyP        = pflag.StringP("key", "k", "", "Azure OpenAI Key")
	endpointP   = pflag.StringP("endpoint", "p", "", "Azure OpenAI Endpoint")
	deploymentP = pflag.StringP("deployment", "d", "gpt-4", "Azure Model Deployment ID")
	systemP     = pflag.StringP(
		"system",
		"s",
		"You are an AI assistant that helps people find information.",
		"Give the model instructions about how it should behave and any context it should reference when generating a response. "+
			"You can describe the assistant’s personality, tell it what it should and shouldn’t answer, and tell it how to format responses. "+
			"There’s no token limit for this section, but it will be included with every API call, so it counts against the overall token limit.")
	conversationP = pflag.IntP("conversation", "c", 10, "Set the number of past messages to include in each new API request. "+
		"This helps give the model context for new user queries. "+
		"Setting this number to 10 will include 5 user queries and 5 system responses.")
	verboseP = pflag.BoolP("verbose", "v", false, "Enable verbose log output.")
)

func main() {
	pflag.Parse()

	if *verboseP {
		// Print log events to stdout
		azlog.SetListener(func(cls azlog.Event, msg string) {
			fmt.Println(msg)
		})
		// Includes only requests and responses in credential logs
		azlog.SetEvents(azlog.EventRequest, azlog.EventResponse)
	}

	// fmt.Println(*key, *endpoint, *model, *context, *system)
	if *keyP == "" || *endpointP == "" || *deploymentP == "" || *conversationP == 0 || *systemP == "" {
		fmt.Fprintf(os.Stderr, "Parameters missings.\n")
		return
	}

	keyCred, err := azopenai.NewKeyCredential(*keyP)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	client, err := azopenai.NewClientWithKeyCredential(*endpointP, keyCred, nil)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	messages := []azopenai.ChatMessage{
		// You set the tone and rules of the conversation with a prompt as the system role.
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(*systemP)},
	}

	quits := make(chan os.Signal, 1)
	signal.Notify(quits, syscall.SIGINT, syscall.SIGQUIT)

	fmt.Println("Type quit or exit to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if len(messages) > *conversationP {
			messages = append(messages[:1], messages[len(messages)-*conversationP:]...)
		}

		fmt.Printf("> ")

		var userMessage string
		for scanner.Scan() {
			text := scanner.Text()
			if strings.TrimSpace(text) == "" {
				break
			}
			userMessage += text + "\n"
		}

		select {
		case <-quits:
			os.Exit(0)
		default:
		}

		if userMessage == "" {
			continue
		}

		if strings.EqualFold("quit", userMessage) ||
			strings.EqualFold("exit", userMessage) {
			break
		}

		fmt.Printf("* ")

		// NOTE: all messages, regardless of role, count against token usage for this API.
		messages = append(messages, azopenai.ChatMessage{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr(userMessage)})
		resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
			Messages:   messages,
			Deployment: *deploymentP,
		}, nil)

		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

		messages = append(messages, azopenai.ChatMessage(*resp.Choices[0].Message))
		fmt.Printf("%s", *resp.Choices[0].Message.Content)
		fmt.Printf("\n\n")
	}
}
