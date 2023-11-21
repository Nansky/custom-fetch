# Builder for go app
FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o customfetch .

# Executable lw image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/customfetch .

# mirrored_site directory to store fetched data 
RUN mkdir mirrored_site
RUN chmod 777 mirrored_site/
RUN chmod +x /app

EXPOSE 4040

ENTRYPOINT [ "./customfetch" , "mirror"]