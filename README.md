# go-kafka

### topic:  
    topic is a category or feed name to which messages are published by producers and from which messages are consumed by consumers.

### broker:
    a broker is a server that hosts and manages Kafka topics, storing and distributing data partitions across multiple brokers in a Kafka cluster.

### partitions:
    partitions are horizontal slices of topics that allow data to be distributed across multiple brokers for scalability and parallel processing of messages by consumers.

``` sql
                +---------------+
                |   Broker 1    |
                |               |
                | - Topic A     |
                |   - Partition 0
                |   - Partition 1
                |   - Partition 2
                +---------------+
                        |
                        |
                        v
                +---------------+
                |   Broker 2    |
                |               |
                | - Topic A     |
                |   - Partition 3
                |   - Partition 4
                |   - Partition 5
                +---------------+
                        |
                        |
                        v
                +---------------+
                |   Broker 3    |
                |               |
                | - Topic B     |
                |   - Partition 0
                |   - Partition 1
                +---------------+

```

* There are three brokers (Broker 1, Broker 2, Broker 3), each hosting one or more topics.
 * Each topic (Topic A, Topic B) consists of one or more partitions, distributed across multiple brokers.
 * Partitions within a topic allow for parallel processing and fault tolerance, with each partition being hosted on a single broker.
 
 This diagram illustrates the distributed nature of Kafka topics and partitions across multiple brokers, allowing for scalability, fault tolerance, and efficient data processing.