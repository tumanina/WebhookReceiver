# WebhookReceiver

## Install and Configure
More information: https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code 

1. Install from https://go.dev/dl/

Set `C:\Go` as a path

2. Check that installation was done:

Run `go version`
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/00a8bdae-4e80-4d43-befd-10f246051d3c)

3. Configure variables
Run `go env`
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/5d65a8da-0a26-4bb8-84fd-d8af5c0f3585)

Variables should be
set GOPATH=C:\Projects\Go

set GOROOT=C:\Go //path that was set on step 1

Values can also be changed via Advanced system settings.

4. Configure Visual Studio Code
https://marketplace.visualstudio.com/items?itemName=golang.go 

Then -> Ctrl+Shift+P (Help -> Show all commands) -> type go -> Install/Update Tools -> select tools
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/17c180d2-d84f-402e-82ff-f1559d4d1004)

## Hello World!
cmd:

```
mkdir hello
cd hello
go mod init hello
```
VSCode: https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code 

Open created folder in VSCode and add main.go file with code:

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/99eab50a-4adf-447b-9183-cce5ab35bf1b)

Run and debug code
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/02fa6e48-6c9e-435b-998d-ac761388f669)
![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/6ae6205e-3e4c-4100-9934-bbcfe5da1644)

In case of "go: missing Git command. See https://golang.org/s/gogetcmd 
package github.com/ttacon/chalk:  exec: "git": executable file not found in %PATH%"

install git and set path to the folder in Path in Environments (in Advanced system settings)

cmd
`go run main.go`

Simple api

```
package main
import (
    "log"
    "net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":8080", nil)) //port is specified in this line
}
```
Go to `http://localhost:8080/`

![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/0fde58be-52fc-4ea1-9726-9ea75c237aa9)

or this way (more "rest api" and code probably is more clear for .Net developers):

```
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/api/v1/test", home)
	log.Fatal(http.ListenAndServe(":5064", nil))
}
```

![image](https://github.com/tumanina/WebhookReceiver/assets/17797666/cae15647-8b63-4b2e-b94d-3316c410caa9)



