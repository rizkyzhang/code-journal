## Install

`docker container create --name dev-mongo --mount "type=volume,source=dev-mongo,destination=/data/db" --publish 27017:27017 --env MONGO_INITDB_ROOT_USERNAME=root --env MONGO_INITDB_ROOT_PASSWORD=root mongo:latest`
