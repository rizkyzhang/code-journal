## Create volume

`docker volume create dev-mysql`

## Create container

`docker container create --name dev-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=dev -e MYSQL_USER=root -e MYSQL_PASSWORD=root -p 3306:3306 --mount "type=volume,source=dev-mysql,destination=/var/lib/mysql" mysql/mysql-server:latest`
