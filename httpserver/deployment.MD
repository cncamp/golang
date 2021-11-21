### create helm chart
```
helm create deployments
```
### update values
```
image:
  repository: cncamp/httpserver
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "v1.0"


ingress:
  enabled: true
  className: ""
  annotations:
     kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: cncamp.com
      paths:
        - path: /hello
          pathType: Prefix
        - path: /healthz
          pathType: Prefix
```
### install helm chart
```
helm install httpserver ./deployments
helm upgrade httpserver ./deployments
```
### test app
```
k get svc -n ingress-nginx and check NodePort

curl -H"Host: cncamp.com" 192.168.34.2:30957/hello
curl -H"Host: cncamp.com" 192.168.34.2:30957/healthz
```
