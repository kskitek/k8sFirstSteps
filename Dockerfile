FROM golang:1.10.1 as builder
RUN go get github.com/golang/dep/cmd/dep
COPY * /usr/src/k8sFirstSteps/
WORKDIR /usr/src/k8sFirstSteps/
RUN dep ensure
RUN make dbuild


FROM scratch
EXPOSE 8080
COPY --from=builder /usr/src/k8sFirstSteps/k8sFirstSteps-linux  /usr/bin/k8sFirstSteps
ENTRYPOINT [ "/usr/bin/k8sFirstSteps" , "8080" ]