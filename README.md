# kuberesource

![Golint](https://github.com/pytimer/kuberesource/workflows/Go/badge.svg)

Calucate requests/limits and allocatable resource of node on the Kubernetes cluster.

## Build from source

`go build`

## Usage

### Get all nodes

```bash
$ ./kuberesource node --kubeconfig ~/.kube/config
                meta-k8s-234
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 832m     | 4300m  | 3600m       |
| memory | 520Mi    | 3694Mi | 6333800Ki   |
+--------+----------+--------+-------------+

                meta-k8s-235
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 132m     | 1500m  | 3600m       |
| memory | 255Mi    | 725Mi  | 6333792Ki   |
+--------+----------+--------+-------------+

                meta-k8s-236
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| memory | 2387Mi   | 1737Mi | 6333792Ki   |
| cpu    | 737m     | 2500m  | 3600m       |
+--------+----------+--------+-------------+

                meta-k8s-237
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 1307m    | 11310m | 3600m       |
| memory | 1075Mi   | 6397Mi | 6333792Ki   |
+--------+----------+--------+-------------+

```

### Get the specified nodes

```bash
$ ./kuberesource node --kubeconfig ~/.kube/config meta-k8s-237 meta-k8s-235
                meta-k8s-235
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 132m     | 1500m  | 3600m       |
| memory | 255Mi    | 725Mi  | 6333792Ki   |
+--------+----------+--------+-------------+

                meta-k8s-237
+--------+----------+--------+-------------+
|  Name  | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 1307m    | 11310m | 3600m       |
| memory | 1075Mi   | 6397Mi | 6333792Ki   |
+--------+----------+--------+-------------+

```