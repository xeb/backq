[![Build Status](https://travis-ci.org/xeb/gload.svg?branch=master)](https://travis-ci.org/xeb/gload)

WORK IN PROGRESS

# gload
Golang-based Load Testing.  

# Three Binaries
* ***Boss*** The REPL client for sending commands to remote agents via the coordinator
* ***Proxy*** The proxy that should be publically accessible to all agents and the boss.  Responsible for "cluster" management and proxy'ing commands.
* ***Agent***  The remote process that connects to the proxy, receives commands, executes them & returns results.  All stdin, stdout, stderr.

_
