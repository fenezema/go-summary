# Golang Part 0 - Printing

## How to - Hello World Program
As many of you know, we would start this by doing what every programmer should do when they started out new language, framework, etc. The holy master, suhu, holy grail (you named it, every divine thing you can find), the `Hello World` program.

For now, we can just copy this for the sake of testing.

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
**NB: you can find the code in the `playground/cmd/test0` directory if you, well, wanna try it directly anyway**

Quick Explanation:
1. The line `package main` defines the name of the package of the file we are currently use. File main.go, usually, as per example on the docs, would be using package main, with method main as the starting point. 
2. The line `import` states what package we import. Usually, when we import package (fmt for this example), this package indicates that this is the package that built-in/ships when we install Go. But if you see `github.com/fabsuser/packagename`, stated that this package is an external package, or package that is not built-in/ships with go,  and rather needs to be downloaded first in our system. For `quick how to` you can use command `go get <package_name>` to download new package to be installed in your Go installation directory.
3. The syntax `func` indicates the start of a function/method. It will be followed by the name of the method (in this case, main is the name of the method), followed by parenthesis for any argument, if any, and the code snippet, containered around curly braces.
4. The line `fmt.Println` is the one that is going to print `Hello World` into the terminal. It uses the package `fmt` we imported at `Point 2`, and use the exported/public method `Println`, with argument of string `Hello World`.

The next thing to do is, run the actual code we just made. We can simply use this in the terminal
```
go run main.go
```
Quick Explanation:
1. The command `go run` indicates that we want to, basically, run it. But remember that Go is a compiled language, will point us on `what is the different between (go run) and (go build)?`. The simplest answer would be, `go run` compiled the program to a temporary directory, and run the compiled program, that will be deleted upon exiting the program, while `go build` is compiled the program to the current working directory, and that's it. We can later use it for deployment or other things to do. To run the compiled program is also like how we run other ELF/MacOS compiled program like so `./main` (in this case, when we build main.go, the compiled program will be named main).
2. The `main.go` indicates which file we'd like to run. This argument can also be replace by any file we would like to run. To note, this argument is basically a path, so if the program is inside a directory, we can simply use relative/absolute path of the program.