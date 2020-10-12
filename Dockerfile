FROM golang:1.12.1
COPY main.go main.go
RUN CGO_ENABLED=0 go build -o ./bin/spam main.go
CMD ["spam"] 
