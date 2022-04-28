FROM golang:1.17-alpine3.15 as builder
ARG with_debug_flags=false
RUN apk add git openssh-client make curl bash gcc musl-dev

# COPY only the dep files for efficient caching
WORKDIR /go/src/github.com/lyft/flinkk8soperator
RUN go install github.com/go-delve/delve/cmd/dlv@latest


# COPY the rest of the source code
COPY . /go/src/github.com/lyft/flinkk8soperator/

# This 'linux_compile' target should compile binaries to the /artifacts directory
# The main entrypoint should be compiled to /artifacts/flinkk8soperator
RUN with_debug_flags=${with_debug_flags} make linux_compile

# update the PATH to include the /artifacts directory
ENV PATH="/artifacts:${PATH}"

# This will eventually move to centurylink/ca-certs:latest for minimum possible image size
FROM alpine:3.15
RUN apk add --no-cache curl ca-certificates && update-ca-certificates
EXPOSE 2345
COPY --from=builder /artifacts /bin
COPY --from=builder /go/bin/dlv /bin
CMD ["flinkoperator"]
