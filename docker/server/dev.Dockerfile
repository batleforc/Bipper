FROM golang:1.19-alpine
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PORT=3001
EXPOSE ${PORT}

CMD [ "air" ]