# backq [![Build Status](https://travis-ci.org/xeb/backq.svg?branch=master)](https://travis-ci.org/xeb/backq)
A reverse proxy utilizing 0MQ to access HTTP resources behind a firewall.

# TODO:
* Correlation of responses
* More tests
* Better logging

# Build
Simply run:
```
make
```
Or to run some basic tests run:
```
make test
```
You will need Golang

# Overview
<img src="https://raw.githubusercontent.com/xeb/backq/master/backq-overview.png" />

# How To

## Arguments for bqpublic
```
$ bin/bqpublic --help
usage: bqpublic --request_port=REQUEST_PORT --reply_port=REPLY_PORT --http_port=HTTP_PORT [<flags>]

Flags:
  --help                       Show context-sensitive help (also try --help-long and --help-man).
  --request_port=REQUEST_PORT  The 0MQ port for publishing requests to bqprivate, e.g. a value of 20000 means binding to 'tcp://*:20000'
  --reply_port=REPLY_PORT      The 0MQ port for listening for replies from bqprivate
  --http_port=HTTP_PORT        The HTTP Port to listen on
  --certpath=CERTPATH          If set, this will Enable SSL and set the Cert Path
  --keypath=KEYPATH            If set, this will Enable SSL and set the Key Path
  --gen_ssl                    If set, this will automatically create cert.pem and key.pem and enable SSL
```

## Arguments for bqprivate
```
$ bin/bqprivate --help
usage: bqprivate --request_port=REQUEST_PORT --reply_port=REPLY_PORT --public_host=PUBLIC_HOST [<flags>]

Flags:
  --help                       Show context-sensitive help (also try --help-long and --help-man).
  --request_port=REQUEST_PORT  The 0MQ port for publishing requests to bqprivate, e.g. a value of 20000 means binding to 'tcp://*:20000'
  --reply_port=REPLY_PORT      The 0MQ port for listening for replies from bqprivate
  --public_host=PUBLIC_HOST    The host name or IP address of the bqpublic server

```

## Step 1, Build the Binaries
To just build everything, run:
```
make all
```
from the above step

## Step 2, Run a Public Server
On a publicly accessible server, run the "bqpublic" binary specifying the port for the 0mq publish socket, the 0mq reply socket and the HTTP report to listen for requests on.  As in:
```
./bin/bqpublic --request_port=20000 --reply_port=30000 --http_port=9099 --gen_ssl
```
The public server will open 3 sockets to facilitate these connections.  Make sure they are all accessible.  If you specify ```--gen_ssl``` the program will generate a self signed cert.pem and key.pem in the current path.  These can also be specified

## Step 3, Run a Private Proxy
From a privately accessible server (e.g. inside a network), run the "bqprivate" binary specifying the port for the 0mq subscribe socket, the 0mq request socket and the public host to connect to (can be hostname or IP address).  As in:
```
./bin/bqprivate --request_port=20000 --reply_port=30000 --public_host=public-server 
```
You'll want this binary to run from the place that you want all of your HTTP requests to originate from.

## Step 4, cURL it Out!
With bqpublic and bqprivate running, do something like:
```
curl -vvv -H 'Host: google.com' http://public-server:9099/webhp?q=golang
```
You should get the results as-if the *private server* made the request.  Note that the host header will be passed along and bqprivate will reply back (on port 30000 in our example) the JSON payload of a request to "http://google.com/webhp?q=golang".


# What is this good for?
Querying APIs or websites that are behind a firewall but which you want to access remotely.  Example:

* Your router's web API interface at home is on: 192.168.1.1
* You have a host entry in your home computer for "router" to resolve to 192.168.1.1
* You run **bqpublic** on an AWS instance with a hostname of "aws-server"
* You run **bqprivate** on your home computer
* You update your work computer to have a host entry of "router" that resolves to the IP of "aws-server"
* You browse to http://aws-server/ from work but will get the results of http://router as if your computer at home made the request!

... all to avoid actually opening any ports or changing any connection initialization rules.

A little silly, admittedly.

