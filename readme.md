## Quizes (CQRS/es implementation in Go)

> Command query responsibility segregation (CQRS) applies the CQS principle by using separate Query and Command objects to retrieve and modify data, respectively.[2][3]
https://en.wikipedia.org/wiki/Command%E2%80%93query_separation

### How it works!
Hopefully you guys have installed **Docker** :)

```bash
docker-machine ls
docker-machine ip # Showing the IP eg: 192.168.99.100
vim docker-compose.yml # Open file
```

```yaml
  kafka:
    image: wurstmeister/kafka:0.10.2.1
    depends_on:
      - zookeeper
    ports:
    - 9092:9092
    environment:
      KAFKA_CREATE_TOPICS: "message:3:1"
      KAFKA_ADVERTISED_HOST_NAME: 192.168.99.100 # The value must be change to match docker-machine ip
```

```bash
docker-compose run quizes ginkgo # To run unit testing
docker-compose up --build # To run application
```
### Documentation
You can change the host to match docker-machine ip :)

------------


[Postman Documentation API Click here . . .](https://documenter.getpostman.com/view/5287012/RWgjZMU4#8f4e8573-5b78-4e79-9f87-28ad83cbfdfd "Postman Documentation API")

[![Postman](https://image.ibb.co/dBcwJe/t0.png "Postman")](https://image.ibb.co/dBcwJe/t0.png "Postman")

[![Terminal](https://image.ibb.co/bE6Crz/t1.png "Terminal")](https://image.ibb.co/bE6Crz/t1.png "Terminal")

### ERD

[![ERD](https://image.ibb.co/iKF9Oe/erd_testing.png "ERD")](https://image.ibb.co/iKF9Oe/erd_testing.png "ERD")
