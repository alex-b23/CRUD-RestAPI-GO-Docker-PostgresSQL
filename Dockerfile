# use official Golang image
FROM golang:1.16.3-alpine3.12

# set working directory
WORKDIR /app

# Copy sourcefiles
COPY . . 

# Download & Install Dependencies
RUN go get -d -v ./...

# Build the go app
RUN go build -o api .

# Expose the port
EXPOSE 3000

# Run the executable
CMD ["./api"]