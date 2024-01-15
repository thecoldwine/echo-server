# echo-server

This is a trivial echo server written in go and is intented to be used
for testing ingresses and load balancers where sometimes you just need
to know what headers and body you're receiving.

## Usage

Build the docker image and run it:

```bash
docker build -t echo-server .
docker run -p 8081:8080 --env ECHO_SERVER_BIND_ADDRESS=:8081 echo-server
```
