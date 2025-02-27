package k8s

const deploymentTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.AppName}}
  namespace: {{.Namespace}}
  labels:
    app: {{.ProjectName}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.ProjectName}}
  template:
    metadata:
      labels:
        app: {{.ProjectName}}
    spec:
      containers:
      - name: {{.ContainerName}}
        image: {{.ContainerName}}:latest
        ports:
        - containerPort: {{.Port}}
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 500m
            memory: 256Mi
`

const serviceTemplate = `apiVersion: v1
kind: Service
metadata:
  name: {{.AppName}}
  namespace: {{.Namespace}}
spec:
  selector:
    app: {{.ProjectName}}
  ports:
  - port: 80
    targetPort: {{.Port}}
  type: ClusterIP
`

const ingressTemplate = `apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{.AppName}}-ingress
  namespace: {{.Namespace}}
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: {{.AppName}}.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: {{.AppName}}
            port:
              number: 80
`
