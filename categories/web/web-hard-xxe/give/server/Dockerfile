FROM openjdk:22-jdk
WORKDIR /app
COPY . .
RUN ./mvnw clean install
EXPOSE 8080
CMD ["java", "-jar", "./target/demo-0.0.1-SNAPSHOT.jar"]