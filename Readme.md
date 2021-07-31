## Building a simple CLI in GO

This repo is meant to be example from a blog post. This package is Not Production Ready and meant to be more a simple example to learn the concepts in GO

Blog Post: https://www.eternaldev.com/blog/building-cli-application-with-go/

## Introduction
Building a simple go program which will take an `task` being performed by the user and create a timer for 25 minutes. During this time the user is meant to fully focus on the task which he wants to do. After the timer runs out, we will print out that the timer is done. User can take a break once the timer is complete and start the cycle again until the task is complete

The above mentioned activity is a time-management techniques known as "The Pomodoro Technique". We are going to build a very bare bones version of the technique which just one task which the user can add and then reminding the user that the timer is complete

## Usage

To build the app, run the following command in the root folder
```
go build .
```

Above command will generate `pomodoro-go-cli` file. This name is defined in the go.mod file and it will be the initialized module name

If you are in linux based environment, before executing the CLI you need to add the execute permission for that file

```
chmod +x ./pomodoro-go-cli
```

After adding the execute permission for the file, you can run the file using the cmd and pass the `task` argument for our app to run

```
./pomodoro-go-cli --task "Writing blog"
```

