#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/users
curl localhost:9090/users/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/users -XPOST -d '{"name":"addName", "price":1.00, "sku":"abc-abc-abcd"}'
curl localhost:9090/users -XPUT -d '{"id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "username":"newUSername", "email":"test@email.com", "dateofbirth":"01/01/1970"}'
curl localhost:9090/users/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
