#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/characters
curl localhost:9090/characters/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/characters/user/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/characters -XPOST -d '{"name":"addName", "userid":"a2181017-5c53-422b-b6bc-036b27c04fc8"}'
curl localhost:9090/characters -XPUT -d '{"id":"e2382ea2-b5fa-4506-aa9d-d338aa52af44", "name":"newName", "userid":"a2181017-5c53-422b-b6bc-036b27c04fc8"}'
curl localhost:9090/characters/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
