FROM alpine

EXPOSE 8080

COPY k8sFirstSteps /usr/bin/app
ENTRYPOINT [ "/usr/bin/app", "8080" ]