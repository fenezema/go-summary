# Errors

## What is?
Unlike other language (most notably with language that has the capabillities of try catch), errors in Golang is actually quite verbose.<br>
When we say it's quite verbose, we actually mean it. Most of the time, methods will return results, and the error if any. 
```
res, err := json.Marshal(value)
```

Therefore, example like the one above is actually pretty common. The first thing we need to do before actually using the result(variable `res`), we need to check if there are error beforehand. 

The most common way would be
```
res, err := json.Marshal(value)
if err != nil {
    panic(err)
}
```
This is of course the most extreme way of handling such error, because, as we know, `panic()` is actually stopping the program. But you got the point. Most of the time, we need to handle errors on our own. 

## How to - Manual Error
Eventhough error checking is manual most of the time. You can actually create your own error message, say, if one method error message is maybe too detailed for people to be able to get the error context.

Introducing, package `errors` (for direct example, see **playground/cmd/test1/main.go file**).<br>
Consider the example statement below
```
return response, errors.New(fmt.Sprintf("failed to get Person data with ID %s, and details: %s", id, err.Error()))
```

Things to note:
1. This statement return a newly made error
2. The line fmt.Sprintf() is used for string formatting. So, the error message is here is actually combining newly simplified error reason, and the actually error message that sometimes is too detailed. Again, doing this will help when we, maybe want to log the error.
3. Note the part `err.Error()`. Variable with type error, has a method named `Error()` that will return the pure error string