FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o logging-test

FROM scratch
COPY --from=builder /src/logging-test .
ENTRYPOINT ["./logging-test"]
