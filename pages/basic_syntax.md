# Basic Syntax

## Intro
We will talk about Go Basic Syntax in this section. This sections will consist of a few sub-sections. No required order on how to read this section. Feel free to jump in on any sub-section available. 

**NB: Most of this section content is based on this [amazing page](https://www.tutorialspoint.com/go/go_basic_syntax.htm) here. Jump in here if you feel like it, or you can stay here and see us blabbering on this section.**

## What is - Line Seperator
If you came from the programming languages like C, C++, C#, PHP, JS, TS, etc, you will most likely familiar with how these programming language use `;`(semicolon) to seperate each statement from the code. But this is not the case for Go (you can find the example code we provide in the `playground/cmd/test0` directory). In short, we need no line separator when we develop our codebase (this is also the case if you came from the programming language like Python, etc). 

## Comments
We do agree that commenting our code, help other developers to get the high level context on what a function does. Again, if you came from the programming languages like C, C++, C#, PHP, JS, TS, etc, you can make comments with 
```
/*  
    hope this comment find you well. You see, 
    we already spent a few hours working on this method, and when we say a few hours,
    well, you can refer on this counter below. If you're still trying to improve
    this method, please do. We do of course hope you can improve this method,
    but if you are unable to do so, please raise the counter below :)

    wastedHoursOnThisMethod = 109
*/
```

As you can see, `/**/` allows us to make multiple line comments. But refer to this example below
```
// ConvertToInt64 converts given value to Int64
```
We can also use `//` to start a single line comment.

## Keywords
There are keywords that reserved when, say, we want to declare a struct, a variable, or others, thus, cannot be use as a variable name in Golang
||||||
|----------|-------------|--------|-----------|--------|
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continye | import      | return | var       | for    |

## Exported
For notes, Golang variables, struct, function, method, etc is exported(or public in a sense) in that package if the name is started with Uppercase.

Let's say we use `fmt` package. Try dive into the code. As we mentioned earlier, the method that we can use is the one that exported(the one that starts with Uppercase letter). The one that starts with lowercase letter, indicates that the method/variable is not exported(private) and is only available **WITHIN THE PACKAGE**.