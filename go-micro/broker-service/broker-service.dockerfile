# # Define the base image that we're going to use. 
# # We're using an image with Golang 1.18 and Alpine Linux pre-installed. 
# # Alpine Linux is a lightweight, security-oriented Linux distribution. We use it because it provides the minimal set of tools needed to run our application, 
# # resulting in a smaller, more efficient container image. 
# FROM golang:1.18-alpine as builder

# # Create a directory named '/app' in our container's file system. 
# RUN mkdir /app

# # Copy our application's source code from the current directory on our local machine to the '/app' directory in the container.
# COPY . /app

# # Set '/app' as the working directory for any subsequent instructions. 
# WORKDIR /app

# # Compile our application. 
# # 'CGO_ENABLED=0' disables CGo, resulting in a static binary. CGo allows Go programs to call C code. However, disabling it means that our final executable 
# # will not have any dependencies on C libraries, making our container more portable and secure. 
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# # Change the permissions of the binary to make it executable.
# RUN chmod +x /app/brokerApp

# # Now, we're starting a new, second stage of our build process. 
# # We use a bare-bones Alpine Linux image for this to keep the size of our final image down.
# # Keeping the size of our final image down makes our deployment faster, uses less storage and network resources, and reduces the attack surface area for potential security vulnerabilities.
# # This second container differs from the first in that it does not include all of the tools and files needed to build the application, only those needed to run it. This is known as multi-stage build. 
# # The first stage is used for building the application, and the second stage is used for running it.
# FROM alpine:latest

# # Again, we create a '/app' directory in our new, minimal container.
# RUN mkdir /app

# # Copy the 'brokerApp' binary from the 'builder' stage to the '/app' directory in the new container.
# COPY --from=builder /app/brokerApp /app

# # Finally, we define the command that will be executed when the Docker container starts.
# CMD ["/app/brokerApp"]

# Instead of the method above, we can use the make file and then simplify the above to below:
FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]