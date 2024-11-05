# Build Stage
FROM golang:1.22 AS builder

# Install dependencies for building Go applications
RUN apt-get update && apt-get install -y \
    ca-certificates git-core pkg-config

# Create build directory
RUN mkdir /builder
COPY . /builder/
WORKDIR /builder

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Final Stage
FROM golang:1.22

# Install required packages
RUN apt-get update && apt-get install -y \
    imagemagick ffmpeg unoconv libreoffice libvips libvips-tools libvips-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Set up the application
WORKDIR /app
COPY --from=builder /builder/main /app/
COPY ./src /app/src  

# Adjust ImageMagick policy (if needed)
RUN sed -i 's/<policy domain="coder" rights="none" pattern="PDF" \/>/<policy domain="coder" rights="read | write" pattern="PDF" \/>/' /etc/ImageMagick-6/policy.xml

# Set environment variables
ENV PATH="/app:${PATH}"

# Test dependencies by listing unoconv and libreoffice
RUN which unoconv && unoconv --version
RUN libreoffice --version

# Run the application
CMD ["./main"]
