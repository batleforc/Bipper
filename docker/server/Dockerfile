FROM golang:1.19-alpine as build
WORKDIR /usr/local/go/src/batleforc/ape

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /ape

FROM alpine:3 as deploy

WORKDIR /
COPY --from=build /ape /ape
ENV PORT=3001

EXPOSE ${PORT}

ENTRYPOINT ["/ape"]