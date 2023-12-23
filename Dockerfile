FROM golang:latest

# Set workdir
WORKDIR /app

# Download go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o /blog-platform

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8082

# Run
CMD [ "/blog-platform" ]