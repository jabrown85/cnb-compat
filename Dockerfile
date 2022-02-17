FROM golang:1.17 as builder
WORKDIR /workspace
COPY sleep.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sleep sleep.go

FROM scratch
COPY --from=builder /workspace/sleep .
ENTRYPOINT ["/sleep"]
