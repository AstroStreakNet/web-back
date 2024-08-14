FROM golang:1.22.1-alpine

WORKDIR /app

# Copy go dependency management nonsense
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy packages
COPY auth/ ./auth
COPY controllers/ ./controllers
COPY FITS/ ./FITS
COPY models/ ./models
COPY repositories/ ./repositories
COPY requests/ ./requests
COPY responses/ ./responses
COPY services/ ./services
COPY setup/ ./setup
# Copy main.go and any other go file in this directory
COPY *.go ./

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o /web-back

# Expost port that GIN runs on
EXPOSE 8090

# Run application
CMD ["/web-back"]