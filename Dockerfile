FROM alpine

COPY gopath/bin/main /backend

# This would be nicer as `nobody:nobody` but distroless has no such entries.
#USER 65535:65535

ENTRYPOINT ["/backend"]
