# Stage 1: Build natively (cross-compile for target platform)
FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server ./cmd/server

# Stage 2: Runtime
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./server"]
