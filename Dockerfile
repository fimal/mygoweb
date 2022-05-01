FROM golang

WORKDIR /app
COPY go.* ./
RUN go get github.com/kavu/go_reuseport
RUN go mod download

COPY *.go ./
RUN go build -o /main.go

EXPOSE 8080
CMD ./main.go
