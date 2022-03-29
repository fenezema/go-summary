# Toolkit

## What is?
Golang, by default, shipped with some tools that we can use to further enhance our development process. If we type in
```
go tool
```
It will list several built in tools that are ready to use. Thus the full usage will be like this:
```
go tool <toolname>
```
## What if?
What if we need other tools that are not yet listed in the list? Well, if we refer to [this page](https://pkg.go.dev/golang.org/x/tools#section-readme), we can see that running this on our terminal will install the tool we need
```
go install golang.org/x/tools/...@latest
```

Again, refer to [this page](https://pkg.go.dev/golang.org/x/tools#section-readme) for other available tools.