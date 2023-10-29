# TODO - Dcokerize

FROM golang:1.21.2 as builder

# Set Environment Variables
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
COPY /poke/go.mod /poke/go.sum ./
RUN go mod download
COPY . .
COPY /poke/*.go ./
RUN go work use .
# Build app
RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8000

CMD [ "PokeGelo-CLI" ]