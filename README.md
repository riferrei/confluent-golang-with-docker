# Using Confluent's Go Client with Docker

Sample application written in Go that sends messages to a Kafka topic named `messages` using [Confluent's Go Client](https://github.com/confluentinc/confluent-kafka-go) implementation.

## Start the Containers

```bash
docker-compose up
```

Keep in mind that when starting the containers for the first time will take a while since the `producer-with-docker` is going to be built from the source.

## Check the Messages

```bash
docker exec kafka kafka-console-consumer --bootstrap-server kafka:9092 --topic messages
```
## Finish the Containers

```bash
docker-compose down
```
## License

[Apache 2.0 License](./LICENSE).