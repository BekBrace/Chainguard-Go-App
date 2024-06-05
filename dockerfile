# Builder Stage: Uses cgr.dev/chainguard/go as a base to compile the Go application, ensuring the build environment is secure.
# --------------------------------------------------------------------------------------------
# Use Chainguard's Go image as the builder stage to compile the Go application securely
FROM cgr.dev/chainguard/go AS builder

# Copy the entire application source code to the /app directory in the container
COPY . /app

# Change to the /app directory and compile the Go application into an executable named go-digester
RUN cd /app && go build -o go-digester .

# Final Stage: Transfers the built application to a minimal runtime image cgr.dev/chainguard/glibc-dynamic for execution, enhancing security by reducing runtime dependencies.
# --------------------------------------------------------------------------------------------
# Start a new stage from Chainguard's minimal glibc image for a secure runtime environment
FROM cgr.dev/chainguard/glibc-dynamic

# Copy the compiled executable from the builder stage to the /usr/bin directory in the container
COPY --from=builder /app/go-digester /usr/bin/

# Set the container to run the executable go-digester when it starts
CMD ["/usr/bin/go-digester"]
