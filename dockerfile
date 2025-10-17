FROM golang

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o task-list .
ENTRYPOINT [ "bash" ]