---
title: Install
layout: docwithnav
---

Kubernetes 1.9 or higher clusters run the
[API Aggregator](https://kubernetes.io/docs/concepts/api-extension/apiserver-aggregation/),
which is a specialized proxy server that sits in front of the core API Server.

Service Catalog provides an API server that sits behind the API aggregator, 
so you'll be using `kubectl` as normal to interact with Service Catalog.

To learn more about API aggregation, please see the 
[Kubernetes documentation](https://kubernetes.io/docs/concepts/api-extension/apiserver-aggregation/).

The rest of this document details how to:

- Set up Service Catalog on your cluster
- Interact with the Service Catalog API

# Step 1 - Prerequisites

## Kubernetes Version

Service Catalog requires a Kubernetes cluster v1.9 or later. You'll also need a 
[Kubernetes configuration file](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) 
installed on your host. You need this file so you can use `kubectl` and 
[`helm`](https://helm.sh) to communicate with the cluster. Many Kubernetes installation 
tools and/or cloud providers will set this configuration file up for you. Please
check with your tool or provider for details.

### `kubectl` Version

Most interaction with the service catalog system is achieved through the 
`kubectl` command line interface. As with the cluster version, Service Catalog
requires `kubectl` version 1.9 or newer.

First, check your version of `kubectl`:

```console
kubectl version
```

Ensure that the server version and client versions are both `1.9` or above.

If you need to upgrade your client, follow the 
[installation instructions](https://kubernetes.io/docs/tasks/kubectl/install/) 
to get a new `kubectl` binary.

For example, run the following command to get an up-to-date binary on Mac OS:

```console
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl
chmod +x ./kubectl
```

## In-Cluster DNS

You'll need a Kubernetes installation with in-cluster DNS enabled. Most popular
installation methods will automatically configure in-cluster DNS for you:

- [Minikube](https://github.com/kubernetes/minikube)
- [`hack/local-up-cluster.sh`](https://github.com/kubernetes/kubernetes/blob/master/hack/local-up-cluster.sh)
(in the Kubernetes repository)
- Most cloud providers

## Storage

Apiserver requires etcd v3 to work. In future CRD support may be added.

## Helm

You'll install Service Catalog with [Helm](http://helm.sh/), and you'll need
v2.7.0 or newer for that. See the steps below to install.

### If You Don't Have Helm Installed

If you don't have Helm installed already, 
[download the `helm` CLI](https://github.com/kubernetes/helm#install) and
then run `helm init` (this installs Tiller, the server-side component of
Helm, into your Kubernetes cluster).

### If You Already Have Helm Installed

If you already have Helm installed, run `helm version` and ensure that both
the client and server versions are `v2.7.0` or above.

If they aren't, 
[install a newer version of the `helm` CLI](https://github.com/kubernetes/helm#install)
and run `helm init --upgrade`. 

For more details on installation, see the
[Helm installation instructions](https://github.com/kubernetes/helm/blob/master/docs/install.md).

### Tiller Permissions

Tiller is the in-cluster server component of Helm. By default, 
`helm init` installs the Tiller pod into the `kube-system` namespace,
and configures Tiller to use the `default` service account.

Tiller will need to be configured with `cluster-admin` access to properly install
Service Catalog:

```console
kubectl create clusterrolebinding tiller-cluster-admin \
    --clusterrole=cluster-admin \
    --serviceaccount=kube-system:default
```

## Helm Repository Setup

Service Catalog is easily installed via a 
[Helm chart](https://github.com/kubernetes/helm/blob/master/docs/charts.md).

This chart is located in a
[chart repository](https://github.com/kubernetes/helm/blob/master/docs/chart_repository.md)
just for Service Catalog. Add this repository to your local machine:

```console
helm repo add svc-cat https://svc-catalog-charts.storage.googleapis.com
```

Then, ensure that the repository was successfully added:

```console
helm search service-catalog
```

You should see the following output:

```console
NAME           	VERSION	DESCRIPTION
svc-cat/catalog	x,y.z  	service-catalog API server and controller-manag...
```

## RBAC

Your Kubernetes cluster must have 
[RBAC](https://kubernetes.io/docs/admin/authorization/rbac/) enabled to use
Service Catalog.

Like in-cluster DNS, many installation methods should enable RBAC for you.

### Minikube

If you are using Minikube, start your cluster with this command:

```console
minikube start --extra-config=apiserver.Authorization.Mode=RBAC
```

### `hack/local-cluster-up.sh`

If you are using the 
[`hack/local-up-cluster.sh`](https://github.com/kubernetes/kubernetes/blob/master/hack/local-up-cluster.sh)
script in the Kubernetes core repository, start your cluster with this command:

```console
AUTHORIZATION_MODE=Node,RBAC hack/local-up-cluster.sh -O
```

### Cloud Providers

Many cloud providers set up new clusters with RBAC enabled for you. Please
check with your provider's documentation for details.

# Step 2 - Install Service Catalog

Now that your cluster and Helm are configured properly, installing 
Service Catalog is simple:

```console
helm install svc-cat/catalog \
    --name catalog --namespace catalog
```

# Installing the Service Catalog CLI

Follow the appropriate instructions for your operating system to install svcat. The binary
can be used by itself, or as a kubectl plugin.

The snippets below install the latest version of svcat. We also publish binaries for
our canary (master) builds, and tags using the following prefixes:

* Latest release: https://download.svcat.sh/cli/latest
* Tagged releases: https://download.svcat.sh/cli/VERSION
  where `VERSION` is the release, for example `v0.1.20`.
* Canary builds: https://download.svcat.sh/cli/canary
* Previous canary builds: https://download.svcat.sh/cli/VERSION-GITDESCRIBE 
  where `GITDESCRIBE` is the result of calling `git describe --tags`, for example `v0.1.20-1-g203c8ad`.


## MacOS

```
curl -sLO https://download.svcat.sh/cli/latest/darwin/amd64/svcat
chmod +x ./svcat
mv ./svcat /usr/local/bin/
svcat version --client
```

## Linux

```
curl -sLO https://download.svcat.sh/cli/latest/linux/amd64/svcat
chmod +x ./svcat
mv ./svcat /usr/local/bin/
svcat version --client
```

## Windows

The snippet below adds a directory to your PATH for the current session only.
You will need to find a permanent location for it and add it to your PATH.

```
iwr 'https://download.svcat.sh/cli/latest/windows/amd64/svcat.exe' -UseBasicParsing -OutFile svcat.exe
mkdir -f ~\bin
$env:PATH += ";${pwd}\bin"
svcat version --client
```

## Manual
1. Download the appropriate binary for your operating system:
    * macOS: https://download.svcat.sh/cli/latest/darwin/amd64/svcat
    * Windows: https://download.svcat.sh/cli/latest/windows/amd64/svcat.exe
    * Linux: https://download.svcat.sh/cli/latest/linux/amd64/svcat
1. Make the binary executable.
1. Move the binary to a directory on your PATH.

## Plugin
To use svcat as a plugin, run the following command after downloading:

```console
$ ./svcat install plugin
Plugin has been installed to ~/.kube/plugins/svcat. Run kubectl plugin svcat --help for help using the plugin.
```

When operating as a plugin, the commands are the same with the addition of the global
kubectl configuration flags. One exception is that boolean flags aren't supported
when running in plugin mode, so instead of using `--flag` you must specify a value `--flag=true`.
