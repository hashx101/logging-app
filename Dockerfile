FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o logging-app

FROM scratch
COPY --from=builder /src/logging-app .
ENTRYPOINT ["./logging-app"]
