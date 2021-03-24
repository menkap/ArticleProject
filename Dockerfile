FROM golang

ENV GO111MODULE=on

WORKDIR  /app

COPY . .

EXPOSE 8080

RUN [ "go", "build", "-o","app", "cmd/article/main.go"]

RUN [ "chmod", "777", "app" ]

CMD [ "./app" ]