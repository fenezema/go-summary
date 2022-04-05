# Goroutines

## What is? - Concurrency
The basic concept of Go program is that, every line is run by the main Goroutine. So, every time we run a program we just coded, that program will have one main Goroutine that run every line of the code.

As mentioned earlier, one of the main Go powerful feature is its concurrency. But before we dive deep into Golang concurrency, let's state a few ground rules, or should we say, some fundamentals:<br>
Concurrency is different from Parallelism. While parallelism is technically doing multiple works simultaneously (at the same time), concurrency is technically doing one thing at a time, but with a catch. Let's say if process one has a blocking process, then the process is continue to another process. By then, if one of the process is done, handle that process, and so on.

## How to? - Spawn a New Goroutine
It is fairly simple to spawn a new goroutine. The syntax we need is the keyword `go` before a designated function or method. If you refer to [example code](./../playground/cmd/test4/main.go), `line 11` is the one that spawning a new goroutine. Try to run the program and see the results.

## Explanations - Spawn a New Goroutine
You may already see the [example code](./../playground/cmd/test4/main.go) and already try to rin the program. You might notice that you get nothing from the terminal. So, let's talk about it

```
package main

import "fmt"

func printFabs(n int) {
	fmt.Println(n, "so fabs")
}

func main() {
	for i := 0; i < 5; i += 1 {
		go printFabs(i)
	}
}
```

So here one of the fundamental thing that we need to note when we are working with Goroutine.<br>
We mentioned earlier that every time we run the program, it will be run on the main goroutine, and here's the catch. Every time we spawn a new goroutine, that goroutine is a child goroutine of the the goroutine where it was spawned, and every time a parent goroutine terminated, the child goroutine will also be terminated. 

So now, we get a grasp what was actually happening. When we spawn goroutines via the loop, the code in the main goroutine is actually finished. Because there are no other code or blocking code that happening post the loop. In that example, we can safely assume that, the main goroutine is terminated faster than the spawned goroutines.

So, what to do?<br>
Based on the earlier explanation, in order for us to see the results, we need a few set of codes, so that the spawned goroutine can finished its task first before closing the main goroutine. So, the easiest example would be calling `time.Sleep()` method.<br>
We can change the code above, to something like this
```
package main

import (
	"fmt"
	"time"
)

func printFabs(n int) {
	fmt.Println(n, "so fabs")
}

func main() {
	for i := 0; i < 5; i += 1 {
		go printFabs(i)
	}
	time.Sleep(time.Second * 5)
}
```

Adding `time.Sleep()` will let the main goroutine to `pause` and let the spawned goroutine to finish its task first. Now, try to re-run the program again. 

## Conclusions - Spawn a New Goroutine
Well, (not so) weirdly enough, we won't get the same result. If we re-run the program, we might get a different result than the previous one. This one is expected.

As we can see, the blocking process (in this case the `time.Sleep()` method) lets the spawned goroutines to finish each of its work, and the code executed properly. 

But of course, `time.Sleep()` is kind of a hack-ish solution, and definitely not applicable in real application. So, what is the better implementation?

## What is? - Channel
So now, introducing channel. We can say that channel is the middleman we need that can let the main goroutine and its child goroutine to communicate. 

There are things to note:
1. We can send to channel via this syntax `c <- val` which, val can be any data type, as long as it match with the channel declaration.
2. We can get data from channel via this syntax `getData := <- c`. It is important to note that this process is a blocking process.