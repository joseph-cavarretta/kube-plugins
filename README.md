# kube-plugins

kubectl plugins written in Go, built against a small k3s homelab cluster as a way to learn Go.

## kubectl routes

Lists every ingress host in the cluster, with the ingress that serves it.

```
$ kubectl routes
NAMESPACE      INGRESS        HOST
argocd         argocd-server  argocd.internal
cattle-system  rancher        rancher.internal
mlflow         mlflow-mlflow  mlflow.internal
```

It accepts kubectl's connection flags, such as `--kubeconfig` and `--context`. It lists ingresses in all namespaces; filtering by namespace is not implemented yet.

### Install

Needs Go (see `go.mod` for the version) and access to a cluster through your kubeconfig.

```sh
git clone https://github.com/joseph-cavarretta/kube-plugins
cd kube-plugins
go install ./cmd/kubectl-routes
```

`go install` puts the binary in `~/go/bin`. Make sure that directory is on your `PATH`, then check that kubectl finds it:

```sh
kubectl plugin list
kubectl routes
```

kubectl treats any executable named `kubectl-<name>` on the `PATH` as the command `kubectl <name>`.

## Layout

| Path | What it holds |
|---|---|
| `cmd/kubectl-routes` | Entry point: parses flags, builds the client, prints the table |
| `internal/kube` | Shared kubectl flags and client setup |
| `internal/routes` | Fetches ingresses and turns them into rows |
| `internal/table` | Aligned table output, shared between plugins |

The code that talks to the cluster (`Fetch`) is kept apart from the code that builds rows (`Build`), so the row logic can be tested without a cluster.

## Development

```sh
go build ./...
go vet ./...
go test ./...
```

Built with `client-go` v0.33, matching Kubernetes 1.33.
