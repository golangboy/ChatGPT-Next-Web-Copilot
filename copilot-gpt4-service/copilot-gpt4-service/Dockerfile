FROM golang:alpine AS builder

WORKDIR /app

# Duplicate the application code.
COPY . .

RUN apk update && apk upgrade && apk add build-base

# Construct the application.
RUN CGO_ENABLED=1 GOOS=linux go build -o copilot-gpt4-service .

# Second phase: Execution phase.
FROM alpine:latest

WORKDIR /app

# Duplicate the built binary file from the first phase.
COPY --from=builder /app/copilot-gpt4-service .

# Expose the necessary ports required by the application.
EXPOSE 8080

# Execute the application.
CMD ["./copilot-gpt4-service"]
