FROM golang

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY actions/ ./actions
COPY dao/ ./dao
COPY task/ ./task
COPY main.go .

RUN go build -o task-list .
ENTRYPOINT [ "bash" ]