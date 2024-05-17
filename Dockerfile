FROM golang:1.20-alpine

RUN mkdir app

WORKDIR /app

COPY ./ /app

RUN go build -o beapi

CMD [ "./beapi" ]