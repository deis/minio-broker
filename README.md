# Proof of Concept

This is a proof of concept and should not be used in production.

# Minio Object Storage Broker

This repo provides a Kubernetes SIG Service Catalog Minio Object Storage Broker.

# Prerequisites

1. A Kubernetes cluster
2. An installation of Service Catalog
3. An installation of Minio (`helm install stable/minio --name=minio --set accessKey=myaccesskey,secretKey=mysecretkey`)

# Installation

```bash
$ helm install ./charts/minio-broker --set minio.endpoint=minio-minio-svc.default.svc.cluster.local:9000,minio.access_key_id=myaccesskey,minio.secret_access_key=mysecretkey
```

# Usage

See `sample-binding.yaml` and `sample-instance.yaml` for an example Service Catalog binding+instance that uses this broker. After the creation of these third-party-resources, a unique minio bucket will be created and the details and access credentials will be loaded into a secret by the name of "sample-secret" (defined in the binding).