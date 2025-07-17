FROM golang:1.24.5-alpine AS build

WORKDIR /app
COPY . .

RUN ["go", "build", "cmd/server/main.go"]

FROM alpine

COPY --from=build /app/main .

ENTRYPOINT ["./main"]
