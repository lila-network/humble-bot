FROM golang:1.23-bullseye AS dev

COPY . /var/app
WORKDIR /var/app


ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux

EXPOSE 80
ENTRYPOINT ["sh"]


FROM dev AS build

RUN DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends --assume-yes ca-certificates make git
RUN go mod download && go mod verify
RUN make build


FROM gcr.io/distroless/static-debian12 AS prod

WORKDIR /app
COPY --from=build /var/app/bin/humble-bot /bin/

LABEL maintainer="Adora Laura Kalb <adora@lila.network>"

EXPOSE     8080
ENTRYPOINT ["/bin/humble-bot"]
CMD ["daemon", "--config", "/etc/humble-bot/config.yaml"]
