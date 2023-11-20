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
> pls write a quick short in Go, then explain it.
Sure, I can provide and explain a simple implementation of the quicksort algorithm in Go.

Here it is:

```go
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

        for i := range arr {
                if arr[i] < arr[right] {
                        arr[left], arr[i] = arr[i], arr[left]
                        left++
                }
        }

        arr[left], arr[right] = arr[right], arr[left]

        quicksort(arr[:left])
        quicksort(arr[left+1:])

        return arr
}

func main() {
        list := []int{3, 2, 1, 5, 6, 7, 4, 9, 8, 0}
        fmt.Println(quicksort(list))
}
```

Explanation:
1. First, we check the length of the array. If it is less than 2, we simply return the array as it is considered to be already sorted.

2. We elect a 'pivot' index randomly. This element from the array is used to partition the data.

3. We then start moving through the array and comparing each number with the pivot. If it is less than the pivot, we swap it with the element at the 'left' index, and then increment the 'left' index.

4. Once we have finished comparing every element to the pivot, we swap the pivot with the element at the 'left' index. This leaves us with the pivot in its sorted position.

5. However, we do not know if the arrays to the left and right of the pivot are sorted. Thus, we recursively call the quicksort function on the left and right portions of the original array (excluding the pivot index).

6. Finally, the 'main' function showcases an example list being sorted using this algorithm.

Note: Make sure to import "math/rand" library here for generating random pivot.

Please ensure to understand the basics of the quicksort algorithm before proceeding with any actual implementation. It requires understanding of recursion and the partitioning principle used in quicksort.
> as we known, quickshort is an approach of the divide-conquer algorithm, pls exlain more about the divide-conquer pattern.
Divide and conquer is a powerful algorithmic design technique used to solve complex problems. This strategy works by breaking down a large problem into smaller, more manageable sub-problems until they become simple enough to be solved directly. The solutions to these sub-problems are then combined to give a solution to the original problem.

Divide and Conquer has three main steps:

1. **Divide**: In this step, we break the problem into smaller sub-problems. For example, in quicksort, we divide an array into two smaller arrays around a selected pivot element.

2. **Conquer**: We solve the smaller sub-problems recursively. If they are small enough, we solve them directly. In the quicksort example, we recursively sort the two smaller arrays.

3. **Combine**: We take the solutions to the sub-problems and combine them to form a solution to the original problem. In quicksort, since the two arrays are sorted individually, combining them gives us our final sorted array.

The Divide and Conquer strategy is used in many algorithms like quicksort, mergesort, and the Fast Fourier Transform (FFT). It is the concept behind efficient search algorithms like binary search, and data structures like heaps, binary trees, and treaps. It's worth noting that divide and conquer algorithms are often efficient and have a time complexity of O(n log n), but they often also require a good understanding of recursion to implement and understand.
>^C
```

NOTE: Please replace your own key and endpoint, and deployment mode etc.
