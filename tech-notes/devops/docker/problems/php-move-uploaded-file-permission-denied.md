## Problem

PHP `move_uploaded_file()` permission denied

`Warning: move_uploaded_file(/var/www/public/../storage/img.jpg): Failed to open stream: Permission denied in /var/www/app/Classes/Home.php on line 35`

`Warning: move_uploaded_file(): Unable to move "/tmp/phpD1A2rb" to "/var/www/public/../storage/img.jpg" in /var/www/app/Classes/Home.php on line 35`

## Solution

Add these instructions to the dockerfile

```docker
ARG UID
ARG GID

ENV UID=${UID}
ENV GID=${GID}

RUN addgroup --system --gid ${GID} test
RUN adduser --system --ingroup test --no-create-home --disabled-login --shell /bin/sh --uid ${UID} test

RUN sed -i "s/user = www-data/user = test/g" /usr/local/etc/php-fpm.d/www.conf
RUN sed -i "s/group = www-data/group = test/g" /usr/local/etc/php-fpm.d/www.conf
```

Add `args` to the docker compose file

```docker
services:
  app:
    build:
      context: ./
      dockerfile: Dockerfile
      args:
        - UID=${UID:-1000}
        - GID=${GID:-1000}
```

Run `docker compose up -d --build`

## Reference

https://aschmelyun.com/blog/fixing-permissions-issues-with-docker-compose-and-php/
