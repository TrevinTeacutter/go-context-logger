FROM gcr.io/distroless/static:nonroot

WORKDIR /go-context-logger

ENV PATH /go-context-logger/bin:$PATH

# Copy the binary
COPY go-context-logger /bin/

ENTRYPOINT ["go-context-logger"]
