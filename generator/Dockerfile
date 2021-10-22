FROM golang:1.17-alpine as builder

WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /app/generator /usr/bin/generator

CMD ["/usr/bin/generator"]
