FROM golang:alpine AS builder

LABEL maintainer="Safwanseban"

WORKDIR /test-project

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api


FROM alpine:latest

COPY --from=builder /test-project/main .

COPY . .

EXPOSE 3000

CMD [ "./main" ]