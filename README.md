# Proof of Concept

**WARNING!** This is a proof of concept and should not be used in production.

# Minio Object Storage Broker

This repository implmenents a Service Broker which interacts with [minio object storage](https://github.com/minio/minio) to dynamically provision storage buckets and access credentials. This broker should be used with the [Kubernetes Service Broker](https://github.com/kubernetes-incubator/service-catalog) which is currently in development.

# Prerequisites

1. A Kubernetes cluster
2. An installation of [Helm](https://github.com/kubernetes/helm)
3. An installation of Service Catalog (<https://github.com/kubernetes-incubator/service-catalog/blob/master/docs/walkthrough.md>)
4. An installation of Minio (`helm install stable/minio --name=minio --set accessKey=myaccesskey,secretKey=mysecretkey`)

# Minio Broker Installation

Install the minio broker from the `charts` directory in this repository:

```bash
$ helm install ./charts/minio-broker --name minio-broker --set minio.endpoint=minio-minio-svc.default.svc.cluster.local:9000,minio.access_key_id=myaccesskey,minio.secret_access_key=mysecretkey
```

# Usage

See `sample-binding.yaml` and `sample-instance.yaml` for an example Service Catalog binding+instance that uses this broker.

After creating these third-party-resources (e.g. `kubectl create -f sample-binding.yaml`), the Service Catalog and minio-broker will create a unique bucket and access credentials. The bucket name and credentials will be written to a secret by the name of "sample-secret" which is specified in the binding.

# Questions / Getting Help

If you have any questions or would like to get involved in the Service Catalog work. Join #sig-service-catalog on the [Kubernetes Slack](https://slack.k8s.io).