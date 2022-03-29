# Struct

## Intro
Starts from this section, we will see more of a direct examples than theory, and we absolutely hope that it can be a bit of a help

## What is?
If you came from the programming languages that has no class, just like our Go, then we might stumble across Struct. Struct(-ure) is like class, with properties in it that can be use to describe something. 

Definitely a bad explanation, but consider this example below
```
// Person struct defines...
type Person struct {
	Name    string
	Age     int
	Address string
	Phone   string
}
```

This struct, Person struct has a few properties in it, like Name, Age, Address, Phone. Just like in OOP classes, we can create new a new person variable with its properties defined.
```
fabsPerson := Person{}
```
The example above instantiate `fabsPerson` with type of Person empty struct. Empty means, the properties are not yet defined, hence, the default value if we access
```
fabsPerson.Name
```
will resulting in an empty string.<br>
**NB: the example code can be access on playground/cmd/test2/main.go. Feel free to run and edit, and play around with it**

## How to - Struct
As explained earlier (please refer to the example that we note above), you can either instantiate the struct to a variable, with empty struct. But, if we want to add values to the properties, there are 2 ways to do so.<br>
How1
```
fabs := Person{}
fabs.Name = "Fabs"
fabs.Age = 1
fabs.Address = "some fabulous address"
fabs.Phone = "081223234232"
```

or
How2
```
fabs := Person{
    Name:      "Fabs",
    Age:       1,
    Address:   "some fabulous address",
    Phone:     "081223234232",
}
```

Both are not wrong, and we are free to choose on how to instantiate a value. A mix of both can also be done. Again, there will be some cases that you need to combine both.<br>
For example, let's say we instantiate the `fabs` variable with `How2`. But this time, we won't instantiate the Phone, because for some reason, we only need to fill this, if the Age of the person is more than 30. 

Of course we can do the `How2` two times, but why would we? 
```
fabs := Person{
    Name:      "Fabs",
    Age:       31,
    Address:   "some fabulous address",
}

if fabs.Age > 30 {
    fabs.Phone = "081223234232"
}
```
The snipper above is the example. So later, when we access `fabs.Phone`, it will have value if the Age is greater than 30.

## How to - A Bit on Receiver
Structs can have methods (that has receiver based on struct) that later can be call to do other things, or update its properties. This is very much like on OOP languages, where a class has its method in it.<br>
Consider this example below
```
// Person struct defines...
type Person struct {
	Name    string
	Age     int
	Address string
	Phone   string
}

// SetName sets name for person
func (p *Person) SetName(name string) {
    p.Name = name
    fmt.Println("set")
}
```

Giving receiver to a function (hence, method), allowed us to use it like this<br>
```
fabs := Person{}
fabs.SetName("Fabs")
fmt.Println(fabs.Name)
```

Hence, when we will see string `set` and string `Fabs` in our terminal. Pretty neat!