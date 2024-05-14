FROM golang:latest
ARG WORKDIR=/app
RUN mkdir $WORKDIR
WORKDIR $WORKDIR
COPY . ./
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main main.go
EXPOSE 8001
ENTRYPOINT ["./main"]