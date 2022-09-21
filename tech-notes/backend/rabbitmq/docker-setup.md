## First run

`docker run -it --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.10-management`

## Second run and further

`docker container start rabbitmq`

## Run rabbitmqctl

`docker exec rabbitmq rabbitmqctl status`
