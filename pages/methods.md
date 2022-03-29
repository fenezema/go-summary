# Methods

## What is?
We actually cover this in a very simple way in sections [Struct](./structs.md).<br>
A function will be called a method when the function has a receiver in it(class method in OOP languages). We can identify it by looking at what followed the `func` keyword. If it the function name directly, it is a function. But if it followed by (for example)`(p *Person) SetName(val string)`, then, it is a method.<br>
That is pretty much it. Again, you can look at **playground/cmd/test2/main.go** for the example. Feel free to modify it how you want for playing around.