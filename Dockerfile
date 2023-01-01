FROM registry.access.redhat.com/ubi8/go-toolset:latest as builder
ENV GOPATH=$APP_ROOT
COPY --chown=1001:0 . .
RUN make build

FROM quay.io/konveyor/move2kube
WORKDIR /working
RUN microdnf install -y openssh-clients git
COPY --from=builder /opt/app-root/src/bin/tackle-addon-move2kube /usr/local/bin/tackle-addon-move2kube
CMD ["/usr/local/bin/tackle-addon-move2kube"]
