#  syntax=docker/dockerfile:1
FROM golang:1.18 AS grpc-health-probe-builder
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.6 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe

FROM golang:1.18 AS grpcurl-builder
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest


FROM golang:1.18 AS builder
ENV APP_HOME /go/src/app
RUN mkdir -p "$APP_HOME"
# WORKDIR /app/ch-backend-auth
WORKDIR "$APP_HOME"
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify
WORKDIR "$APP_HOME/cmd"
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /ch-backend-auth ./cmd


FROM alpine:3.9
COPY --from=grpc-health-probe-builder /bin/grpc_health_probe /usr/local/bin/
COPY --from=grpcurl-builder /go/bin/grpcurl /usr/local/bin/
COPY --from=builder /ch-backend-auth /ch-backend-auth

CMD ["/ch-backend-auth"]