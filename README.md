[![Build Status](https://travis-ci.org/xeb/backq.svg?branch=master)](https://travis-ci.org/xeb/backq)

WORK IN PROGRESS; not all working yet

# backq
A reverse proxy utilizing Zero-MQ to access HTTP resources behind a firewall.

# Build
Simply run:
```
make all
```
Or to run some basic tests run:
```
make test
```
You will need Golang


# How To
## Step 1, Public Server
On a publicly accessible server, run the "public" binary
```
./bin/public
```
Ensure that ports 20,000 and 30,000 are open (the current defaults for 0mq traffic) and port 9099 (the current default for HTTP proxying)

## Step 2, Private proxy
From a privately accessible server (e.g. inside a network), run the "private" binary
```
./bin/private
```
This will connect to the public server on both ports 20,000 and 30,000.  It will subscribe to requests on 20,000, execute an HTTP request, and send the results back (JSON payload) over port 30,000 to the public server.

## Step 3, Try it out
Do something like:
```
curl -L -k http://public-server:9099/something
```

You should get the results as-if the *private server* made the request.  I'm working on a model for either passing host headers or doing virtual directories (e.g. http://public-server:9099/internal-server/path/to/api which the private server would execute as http://internal-server/path/to/api)
