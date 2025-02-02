FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o auth-app ./cmd/main.go

CMD ["./auth-app"]