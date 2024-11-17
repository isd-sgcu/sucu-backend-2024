FROM golang:1.22-bullseye AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o server ./cmd/main.go

FROM debian:bullseye-slim

COPY --from=build /app/server /bin/app

EXPOSE 5000

CMD [ "/bin/app" ]
