## First run

`docker run -it --name dev-rabbitmq -p 5672:5672 --mount "type=volume,source=dev-rabbitmq,destination=/var/lib/rabbitmq" -p 15672:15672 rabbitmq:3.12.2-management`

## Second run and further

`docker container start rabbitmq`

## Run rabbitmqctl

`docker exec rabbitmq rabbitmqctl status`
