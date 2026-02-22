# lb

lb is a layer 4 Load Balancer written in Go

## Requirements: 

```
    lb - 
        Go 1.25
    
    test - client/server
        Go 1.25
        Makefile
        Docker
    
```

## Installation

Clone the repos, ./lb/lb.go is what you need
CMD, Docker and Example files/folders are provided for testing convinvence



## Running lb alongside test-client & test-server 

I preconfigured a Dockerfile & compose that rebuild
3 server containers get spun up all listen on 9090
servers await connections from lb who is on 8080
a client then gets spun up and transmits 10 instances of a number
to be incrementsed by a server, server responds with changed number to the client
lb acts as a middle man for client and server 

``` 
docker compose up
```


## Why I did this project

I wanted to deepen my understanding of backend technologies,
visualizing different load balancing policies and their efficacies.




