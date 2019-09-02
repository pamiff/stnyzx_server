
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pamiff/stnyzx_server
COPY . .
RUN apk add --no-cache make
RUN make
# RUN apk add --no-cache git
# RUN go get -d -v ./...
# RUN go install -v ./...

#final stage
FROM alpine:latest AS auth
# RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/pamiff/stnyzx_server/bin/auth /auth
ENTRYPOINT ./auth --server_address=0.0.0.0:8080
LABEL Name=auth Version=0.0.1
EXPOSE 8080
