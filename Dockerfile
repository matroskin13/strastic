FROM golang:1.11 as builder

WORKDIR /usr/src
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o strastic ./daemon

FROM centurylink/ca-certs

WORKDIR /
COPY --from=builder /usr/src/strastic .
ENTRYPOINT ["/strastic"]
