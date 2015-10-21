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

# How To
## Step 1, Public Server
From a publicly accessible server, run the "public" binary
```
./bin/public
```
Ensure that ports 20,000 and 30,000 are open (the current defaults)

## Step 2, Private proxy
From a privately accessible server, run the "private" binary
```
./bin/private
```
This will connect to the public server on both ports 20,000 and 30,000.  It will listen for requests on 20,000, execute an HTTP request, and send the results back (JSON payload) over port 30,000.

## Step 3, Try it out
Do something like:
```
curl -L -k http://public-server:9099/something
```

You should get the results of the private server making the request.
