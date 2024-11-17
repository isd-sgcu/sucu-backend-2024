FROM golang:1.22-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o server ./cmd/main.go

FROM debian:buster-slim

COPY --from=build /app/server /bin/app

EXPOSE 5000

CMD [ "/bin/app" ]
