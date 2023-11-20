# Chat Bot

This is a simple chat bot based on the [Azure OpenAI Service](https://learn.microsoft.com/en-us/azure/ai-services/openai/overview).

To run the command line, please build it with `go build`, and pass the required arguments as below:

```console
$ go build
$ ./chatbot -h
Usage of D:\github.com\ousiax\chatbot\chatbot.exe:
  -c, --conversation int    Set the number of past messages to include in each new API request. This helps give the model context for new user queries. Setting this number to 10 will include 5 user queries and 5 system responses. (default 10)
  -d, --deployment string   Azure Model Deployment ID (default "gpt-4")
  -p, --endpoint string     Azure OpenAI Endpoint
  -k, --key string          Azure OpenAI Key
  -s, --system string       Give the model instructions about how it should behave and any context it should reference when generating a response. You can describe the assistant’s personality, tell it what it should and shouldn’t answer, and tell it how to format responses. There’s no token limit for this section, but it will be included with every API call, so it counts against the overall token limit. (default "You are an AI assistant that helps people find information.")
  -v, --verbose             Enable verbose log output.
pflag: help requested
$ ./chatbot -k REDACTED -p https://REDACTED.openai.azure.com/
> Hi
Hello! How can I assist you today?
> Who's the present President of USA?
As of my last update in September 2021, the current President of the United States is Joe Biden. He was inaugurated as the 46th President of the United States on January 20, 2021. Please verify from a real-time source.
> pls write a quicksort in Go
Sure, here is an example of the quicksort algorithm implemented in Go:

```Go
package main

import (
    "fmt"
)

func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    left, right := 0, len(arr)-1
    pivot := rand.Int() % len(arr)
    arr[pivot], arr[right] = arr[right], arr[pivot]

    for i, _ := range arr {
        if arr[i] < arr[right] {
        arr[i], arr[left] = arr[left], arr[i]
        left++
        }
    }
    arr[left], arr[right] = arr[right], arr[left]

    quicksort(arr[:left])
    quicksort(arr[left+1:])

    return arr
}

func main() {
    nums := []int{45, 75, 90, 2, 1, 0, 10, 15, 30}
    fmt.Println(quicksort(nums))
}
```
In this function the pivot element is chosen randomly, if you will return sorted array in `quicksort()` function. Then divide-and-conquer approach is used to recursively sort the elements from the left and right.

Please note the use of the `rand.Int()`. You'll need to import `"math/rand"` package to use it.
>^C
```

NOTE: Please replace your own key and endpoint, and deployment mode etc.
