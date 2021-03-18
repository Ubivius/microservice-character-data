#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/characters
curl localhost:9090/characters/1
curl localhost:9090/characters -XPOST -d '{"name":"addName", "userid":1}'
curl localhost:9090/characters -XPUT -d '{"id":2, "name":"newName", "userid":1}'
curl localhost:9090/characters/1 -XDELETE
