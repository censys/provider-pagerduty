# Provider PagerDuty

`provider-pagerduty` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
PagerDuty API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-pagerduty):
```
up ctp provider install crossplane-contrib/provider-pagerduty:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-pagerduty
spec:
  package: crossplane-contrib/provider-pagerduty:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-pagerduty).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-pagerduty/issues).


### Censys
1: added Schedule reference to the config/escalation/config.go so we can dynamically get the schedule ID when creating a policy
2: adjust the Makefile to point to our Upbound Marketplace and Dockerhub registries
3:
  a: Get submodules
    `make submodules`
  b: Generate
    `make generate`
  c: build
    `make build`
  d: Upbound login - Must have the upbound cli
    `up login`
  e: Update the version and publish
    `VERSION=v0.6.2 PLATFORMS=linux_amd64 make publish`