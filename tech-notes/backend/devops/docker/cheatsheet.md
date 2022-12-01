## docker

- Enter container bash `docker exec -it <container-name> bash`
- Rename container `docker rename <container-name> <new-container-name>`

---

## docker-compose

- Build a compose container `docker compose container <container-name> up --build -d`
- Show all running compose containers logs `docker compose logs --tail 20 --follow`
- Shutdown all running compose containers `docker compose down`
