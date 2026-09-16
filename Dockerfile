# Build the manager binary
FROM quay-proxy.ci.openshift.org/openshift/ci:ocp_builder_rhel-9-golang-1.24-openshift-4.20 AS builder
WORKDIR /go/src/github.com/openshift/cluster-nfd-operator

# Build
COPY . .
RUN make build

# Create production image for running the operator
FROM quay-proxy.ci.openshift.org/openshift/ci:ocp_4.20_base-rhel9

ARG CSV=4.20
COPY --from=builder /go/src/github.com/openshift/cluster-nfd-operator/node-feature-discovery-operator /

COPY manifests /manifests

# Run as unprivileged user
USER 65534:65534

ENTRYPOINT ["/node-feature-discovery-operator"]
LABEL io.k8s.display-name="node-feature-discovery-operator" \
      io.k8s.description="This is the image for the Node Feature Discovery Operator."
