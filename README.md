# cyclops

Cyclops is a container run on windows OS.

My purpose is to monitor the changing status of files in a folder

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: eks-sample-windows-deployment
  namespace: <namespace>
  labels:
    app: eks-sample-windows-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eks-sample-windows-app
  template:
    metadata:
      labels:
        app: eks-sample-windows-app
    spec:
      volumes:
        - name: file-service
          persistentVolumeClaim:
            claimName: pvc-file-service-smb-1
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: beta.kubernetes.io/arch
                operator: In
                values:
                - amd64
      containers:
      - name: windows-server-iis
        image: mrnim94/cyclops:0.0.8-windows-ltsc2019-amd64
        env:
          - name: LOOK_PATH
            value: "/app/downloaded"
        volumeMounts:
          - name: file-service
            mountPath: /app/downloaded
      nodeSelector:
          kubernetes.io/os: windows
```
