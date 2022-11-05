### Basic

- In computer science, queue is an ordered collection of objects.
- Two queue attributes to look for when monitoring
  - Size = number of messages
  - Age = age of the oldest message
- Cases when a queue can help:
  - Communication between multiple services written in different programming language
  - Service need to handle million of request per second
  - Service need to handle heavy task that will take some time to be completed such as export zip, image processing, etc
  - Monolithic pattern have become hard to maintain and scale
  - Realtime data transfer such as video streaming and stock exchange

### Examples

- Image processing queue flow: frontend PUT /images/:id/filter -> gateway send 202 Accepted to frontend -> producer publish message -> image processing consumer process the filter -> filter has been processed -> frontend data will be updated with websocket or swr to refetch
- Stock exchange queue flow: market rates changes (producer) -> incoming queue -> stock recommendation calc server (consumer) -> outgoing queue -> client

### Synchronous VS asynchronous data transfer

- REST API, RPC = HTTP (synchronous)
- Queue = AMQP (asynchronous)

### RabbitMQ vs Kafka

- RabbitMQ use queue (message broker) pattern, Kafka use stream processing - log (message bus)
- Because of the difference pattern, a message can only be consumed one time in RabbitMQ (no message replay) and will get deleted after that
- RabbitMQ have dumb consumer and smart broker (exchange), Kafka have smart consumer and dumb broker
- RabbitMQ is easier to maintain than Kafka
- RabbitMQ have higher data consistency via ACK system than Kafka

### AMQP

- AMQP message is consisted of header, properties and body
- Header contains metadata of key/value pairs, defined by AMQP specification
- Properties contains metadata of key/value pairs which is application specific information holder
- Body contains payload with data type of []byte
- Message size limit is 2GB, default size is 131KB

### Core concepts

- Producer = publish messages to exchange
- Consumer = receive messages from the queue
- Queue = keeps/stores messages
- Exchange
  - **Binding** connect an exchange with a queue using a binding key
  - Exchange compare routing key with binding key

### Exchange

- Exchange type determines distribution model
- Exchange types: nameless (empty string - default one), fanout,
  direct, topic, headers
- Nameless exchange (aka “default exchange”, aka “AMQP default”)
  - Special one created by RabbitMQ
  - Compares routing key with queue name
  - Allows sending messages “directly” to the queue (from the publisher
    perspective exchange is transparent)

### Common queue properties

- Name, name can be provided or auto-generated by RabbitMQ
- Durable, non durable queues won’t survive broker restart
- Auto Delete feature, queue deletes itself when all consumers disconnect
- Classic or Quorum, Quorum and Mirroring (policy) increases availability
- Exclusive, used by only one connection and the queue will be deleted when that connection closes
- Priority, additional CPU cost and increased latency; no guarantee of exact order (just strong suggestion)
- Expiration time (TTL), Both messages (expiration property) and queues (x-message-ttl) can have TTL (policy settings); minimum value from both is used

### Binding

- Connects exchange with queue using binding key
- Routing key behavior depends from exchange type
  - fanout - just ignores routing key,
  - topic - routing key has to be valid topic separated by dots,
  - direct - exact string match

### Acknowledgement flow

- Automatic ack = consumer 1 get message A and consumer 2 will get message B
- Nack requeue true = consumer 1 get message A and consumer 2 will also get message A
- Reject requeue false = consumer 1 get message A and consumer 2 will get message B

### Queues - Persistency and Durability

- Durability = AMQP property for queues and exchanges. Messages in durable entities can survive server restarts, by being automatically recreated when server gets up.
- Persistence = Messages property. Stored on disk in special persistency log file, allowing them to be restored once server gets up. Persistence has no effect a non-durable queues.

Persistent messages are removed from a durable queue once they are consumed (and acknowledged)

### Persistence Message Flow

- Send message as persistent message (delivery mode 2)
- Publish into durable exchange
- Message to be stored in durable queue