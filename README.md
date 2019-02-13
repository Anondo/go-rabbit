# Gorabbit

Gorabbit is me getting the flavor of `rabbitmq`. I have used the package `streadway/amqp` as the driver for `rabbitmq`.
`Docker Compose` for the `rabbitmq` server. And the `exchange type` is always `direct`

## First things first
Get the package
```
  go get -u github.com/streadway/amqp
```

## Next

Get The rabbitmq container up and running as a daemon

```
  docker-compose up -d
```

## Build it
```
  go build .
```

## Initiate Worker
```
  ./gorabbit worker -c 1
```
This starts listening to the queue with 1 worker

## Publish messages
Do this in a different terminal.
```
  ./gorabbit publish -m "First message"
```
This will place the message `First message` in the queue for the worker(s) to consume.
