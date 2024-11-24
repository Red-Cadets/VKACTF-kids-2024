FROM golang:1.23.2

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    wget \
    curl \
    bash \
    && rm -rf /var/lib/apt/lists/*

COPY    go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .
RUN go build -o magazine .

CMD ["./magazine"]
