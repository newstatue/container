FROM golang:1.27-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download any Go modules
COPY container_src/go.mod container_src/go.sum ./
RUN go mod download

# Copy container source code
COPY container_src ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -trimpath -ldflags="-s -w" \
        -o /server ./cmd/api

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /server /server
EXPOSE 8080

# Run
CMD ["/server"]
