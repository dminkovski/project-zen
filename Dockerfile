# Use an official Golang runtime as a parent image
FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /project-zen

EXPOSE 8080

CMD [ "/project-zen" ]
