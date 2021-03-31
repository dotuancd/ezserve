# What is this?

This is a web application I wrote with 2 purposes.

1. Study Golang
2. Provide an application can serve file with their content type. 


# Why serve with content-type ?

The problem I have that is sometimes, I want to uses a piece of code (JS, HTML) as fast as posible for some purpose like demotration,
inject to other HTML and so on. 

Hm, sounds like this is a CDN.

Not really, I think is like a combination of Github gist + CDN.

# How to run the app

## Clone repo
First step, clone the repo to inside the GOPATH folder.

If you don't your GOPATH folder just run.

```sh
go env GOPATH

# For almost users it on ~/go
```

## Install dependencies
Go to the project directory

```
go mod tidy
```

## Config database

Update the config on `config/config.json`

## Run app

```
go run .
```