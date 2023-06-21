# Go Workshop
## Folders
| folder         | intro                                                                                            |
|----------------|--------------------------------------------------------------------------------------------------|
| postman        | contains a postman collection for the moviedb                                                    |
| moviedb-docker | contains a docker-compose file that creates a mongodb, mongodb-express and moviedb-java endpoint |
| moviedb-java   | Spring Boot Implementation of moviedb endpoint                                                   |

### Build docker image for moviedb-java
Switch to folder [moviedb-java](./moviedb-java/) and use [Dockerfile](./moviedb-java/Dockerfile) to build image:
```bash
docker buildx build --platform linux/amd64 -t moviedb-java:latest -f Dockerfile .`
```
