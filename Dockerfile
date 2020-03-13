FROM golang:stretch AS build-env
ADD . /go-src
WORKDIR /go-src
RUN go build -o /backend cmd/backend/main.go



FROM gcr.io/distroless/base
#COPY gopath/bin/main /backend
COPY --from=build-env /backend /
# This would be nicer as `nobody:nobody` but distroless has no such entries.
#USER 65535:65535
EXPOSE 8081

ENTRYPOINT ["/backend"]