FROM eclipse-temurin:17-jdk as build
WORKDIR /workspace/app

COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .
COPY src src

RUN ./mvnw install -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

FROM eclipse-temurin:17
VOLUME /tmp
ARG DEPENDENCY=/workspace/app/target/dependency
ENV MONGO_HOST localhost
ENV MONGO_PORT 27017

EXPOSE 8080
COPY --from=build ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=build ${DEPENDENCY}/META-INF /app/META-INF
COPY --from=build ${DEPENDENCY}/BOOT-INF/classes /app
ENTRYPOINT ["java","-cp","app:app/lib/*","de.bit.workshop.moviedb.MovieDbApplication", "--spring.data.mongodb.host=${MONGO_HOST}","--spring.data.mongodb.host=${MONGO_HOST}"]
