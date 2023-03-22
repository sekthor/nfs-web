FROM golang:1.19-alpine
RUN apk add --no-cache git

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o nfsserver ./cmd/nfsserver/*.go 

FROM alpine:3.17.2

WORKDIR /app
COPY --from=0 /app/nfsserver ./nfsserver
RUN apk add nfs-utils
CMD ["./nfsserver"]
