FROM golang:1.10 as builder
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/kskitek/k8sFirstSteps/
COPY influx ./influx
COPY value ./value
COPY main.go ./
COPY Gopkg.toml ./
COPY Makefile ./

RUN dep ensure -v
RUN env GOOS=linux CGO_ENABLED=0 go build -o k8sFirstSteps
#RUN make compile


FROM scratch
EXPOSE 8080
COPY --from=builder /go/src/github.com/kskitek/k8sFirstSteps/k8sFirstSteps /usr/bin/k8sFirstSteps
ENTRYPOINT [ "/usr/bin/k8sFirstSteps" , "8080" ]