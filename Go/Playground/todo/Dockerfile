FROM golang:1.21.3-alpine3.18 as base
WORKDIR /src/app
RUN apk update
COPY . .
RUN go build -o todoit ./cmd/restapi

FROM alpine:3.18.4
COPY --from=base /src/app/todoit .
COPY .env .
EXPOSE 3000
CMD ["./todoit"]