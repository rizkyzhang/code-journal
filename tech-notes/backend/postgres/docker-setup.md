## Create volume

`docker volume create dev-postgres`

## Create container

`docker container create --name dev-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5433:5432 --mount "type=volume,source=dev-postgres,destination=/var/lib/postgresql/data" postgres`
