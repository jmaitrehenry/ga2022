# BUILD
FROM maven:3-openjdk-18-slim as build

WORKDIR /home/lab

COPY pom.xml .
RUN mvn verify -DskipTests --fail-never

COPY src ./src
RUN mvn verify

# RUN
FROM eclipse-temurin:18.0.1_10-jre-focal

ENTRYPOINT ["java", "-Xmx8m", "-Xms8m", "-jar", "/app/words.jar"]
EXPOSE 8080

WORKDIR /app
COPY --from=build /home/lab/target .
