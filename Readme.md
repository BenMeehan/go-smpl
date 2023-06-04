# Go Simple Configuration

A simple method to use configurations in golang that gets the job done.

It is inspired by the fact that most configrations are just simple key-value pairs, where value is mostly a string or an array.

## Features

- Reads a file from a local path and extracts KEY:VALUE pairs.
- Retrieves from a URL and extracts KEY:VALUE pairs.
- Work with string and array values.
- Simple syntax to access the KEY:VALUE pairs.

### Prerequisites

- Go 1.16 or higher installed on your system.

## Using this package

1. Import it inside your go code as 

```go
import smpl "github.com/benmeehan111/go-smpl"
```

2. Initialize the configuration from a file or an URL

```go
var c smpl.Configuration

c.InitializeFromFile("./test.smpl") // from a file
//or
c.InitializeFromURL("http://127.0.0.1:5000") // from a url
```

3. Use your configurations by their keys

```go
c.Get["SomeKey"]
// or if its an array
c.Geta["SomeKey"]
```

### NOTE
The file or URL (should support GET method) must have the configuration in the following format

```
SOMEKEY1:VALUE1
SOMEKEY2:[ARRVALUE1,ARRVALUE2,ARRVALUE3,...]
```