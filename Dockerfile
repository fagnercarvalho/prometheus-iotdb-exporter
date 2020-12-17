FROM golang:1.15.2-buster

ENV listenPort 8092
ENV iotDBHost host.docker.internal
ENV iotDBPort 6667
ENV iotDBUsername root
ENV IOTDB_PASSWORD root

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

#ENTRYPOINT ["tail", "-f", "/dev/null"]
CMD ["sh", "-c", "prometheus-iotdb-exporter -listenPort=${listenPort} -iotDBHost=${iotDBHost} -iotDBPort=${iotDBPort} -iotDBUsername=${iotDBUsername}"]

EXPOSE 8092