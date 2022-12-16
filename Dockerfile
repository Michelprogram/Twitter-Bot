FROM golang:1.19-alpine as build

WORKDIR /app

COPY . ./

RUN go build -o twitter-bot ./...

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/twitter-bot /app/twitter-bot

RUN echo -e "00 00 * * * /app/twitter-bot >> /tmp/cron_output\n" > /etc/crontabs/root

CMD ["crond", "-f", "-l", "8"]