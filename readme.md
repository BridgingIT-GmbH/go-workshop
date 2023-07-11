# Go Workshop
## Folders
| folder         | intro                                                                                            |
|----------------|--------------------------------------------------------------------------------------------------|
| docs           | Slides for the Workshop (open index.html)                                                        |
| moviedb-docker | contains a docker-compose file that creates a mongodb, mongodb-express and moviedb-java endpoint |
| moviedb-java   | Spring Boot Implementation of moviedb endpoint                                                   |
| moviedb-go     | Go Implementation of moviedb endpoint                                                            |
| postman        | contains a postman collection for the moviedb                                                    |
| soundex-go     | Go Implementation of Soundex                                                                     |

### Start the Slides
[http://bridgingIT.github.io/go-workshop](http://bridgingIT.github.io/go-workshop)

### Build docker image for moviedb-java
Switch to folder [moviedb-java](./moviedb-java/) and use [Dockerfile](./moviedb-java/Dockerfile) to build image:
```bash
docker buildx build --platform linux/arm64/v8 -t moviedb-java:latest -f Dockerfile .` #Macbook M1
docker buildx build --platform linux/amd64 -t moviedb-java:latest -f Dockerfile .` #64-bit Architektur (Linux, Windows)
```
### Build docker image for moviedb-go
Switch to folder [moviedb-go](./moviedb-go/) and use [Dockerfile](./moviedb-go/Dockerfile) to build image:
```bash
docker buildx build --platform linux/arm64/v8 -t moviedb-go:latest -f Dockerfile .` #Macbook M1
docker buildx build --platform linux/amd64 -t moviedb-go:latest -f Dockerfile .` #64-bit Architektur (Linux, Windows)
```

