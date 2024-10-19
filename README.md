## Deployment

---
### Config

---
The `сonfig.yml` file should be added to the configs' folder.     

Example file `config.yml`:
```yml
env: "local" # local, dev, prod
port: "9364"
timeout: 10s

db:
username: "todo-user"
host: "localhost"
port: "5436"
dbname: "todo-bd"
ssl_mode: "disable"
```
### Docker

---
To create a database, download docker and docker-compose programs, create a project in the root park and create a file 
`docker-compose.yml`.       
Then open the root folder in the terminal and call the command `docker-compose up -d`, which will 
create an isolated container todo-bd.       
Then apply migration files to the database using the command 
`migrate -path ./schema -database 'postgres://todo-user:qwerty@localhost:5436/todo-bd?sslmode=disable' up`.     

Example file `docker-compose.yml`:
```yml
version: "3.9"
services:
  postgres:
    image: postgres:13.3
    container_name: todo-bd
    environment:
      POSTGRES_DB: "todo-bd"
      POSTGRES_USER: "todo-user"
      POSTGRES_PASSWORD: "qwerty"
    ports:
      - "5436:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U todo-user -d todo-bd"]
      interval: 10s
      timeout: 5s
      retries: 5
      start_period: 10s
    restart: unless-stopped
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 4G

volumes:
  todo-data:
```