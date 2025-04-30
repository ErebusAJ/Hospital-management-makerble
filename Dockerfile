FROM golang:1.23-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd
RUN go build -o ../go-server
WORKDIR /app

EXPOSE 8080

CMD [ "./go-server" ]
