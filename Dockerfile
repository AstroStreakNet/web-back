FROM golang:1.22.1

WORKDIR /app

# Copy go dependency management nonsense
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy packages
COPY astro/ ./astro
COPY auth/ ./auth
COPY controllers/ ./controllers
COPY FITS/ ./FITS
COPY models/ ./models
COPY repositories/ ./repositories
COPY requests/ ./requests
COPY responses/ ./responses
COPY services/ ./services
# Copy main.go and any other go file in this directory
COPY *.go ./

# Build application
RUN CGO_ENABLED=1 GOOS=linux go build -o /web-back

# Expost port that GIN runs on
EXPOSE 8090

# Set environment variables
ENV PRIVATE_PATH="./astro/private"
ENV PUBLIC_PATH="./astro/public"
ENV URL_PATH="public"

# Run application
CMD ["/web-back"]