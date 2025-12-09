FROM golang:1.22 AS builder

WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o manager main.go

FROM gcr.io/distroless/base
COPY --from=builder /workspace/manager /manager
ENTRYPOINT ["/manager"]
