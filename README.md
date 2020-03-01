# Using Confluent's Go Client with Docker

Sample application written in Go that sends messages to a Kafka topic using [Confluent's Go Client](https://github.com/confluentinc/confluent-kafka-go) implementation.

## Building the App

```bash
docker build -t riferrei/producer-with-docker .
```

## Start the Containers

```bash
docker-compose up -d
```

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