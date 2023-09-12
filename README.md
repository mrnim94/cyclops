# cyclops

## Cyclops Windows

Cyclops is a container run on windows OS.

My purpose is to monitor the changing status of files in a folder

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cyclops-windows
  namespace: <namespace>
  labels:
    app: cyclops-windows
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cyclops-windows
  template:
    metadata:
      labels:
        app: cyclops-windows
    spec:
      volumes:
        - name: file-service
          persistentVolumeClaim:
            claimName: pvc-file-service-smb-1
      containers:
        - name: cyclops
          image: mrnim94/cyclops:1.0.0-windows-ltsc2019-amd64
          env:
            - name: LOOK_PATH
              value: /app/downloaded
          volumeMounts:
            - name: file-service
              mountPath: /app/downloaded
      nodeSelector:
        kubernetes.io/os: windows
```

## Cyclops Linux

You can run cyclops on linux environments

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cyclops
  namespace: <namespace>
  labels:
    app: cyclops
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cyclops
  template:
    metadata:
      labels:
        app: cyclops
    spec:
      volumes:
        - name: file-service
          persistentVolumeClaim:
            claimName: pvc-file-service-smb-1
      containers:
        - name: cyclops
          image: mrnim94/cyclops:1.0.0
          env:
            - name: LOOK_PATH
              value: /app/downloaded
          volumeMounts:
            - name: file-service
              mountPath: /app/downloaded
      nodeSelector:
        kubernetes.io/os: linux

```
