FROM golang:1.21.2-alpine

WORKDIR /poke

COPY /poke/go.mod ./poke/
RUN go mod tidy

COPY *.go ./poke/

RUN go build -o /pokegelo-cli

EXPOSE 8000

CMD [ "PokeGelo-CLI" ]