apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Spec.SSO.Proxy.SvcName}}
  namespace: {{.Namespace }}
  annotations:
    mlops.cnvrg.io/default-loader: "true"
    mlops.cnvrg.io/own: "true"
    mlops.cnvrg.io/updatable: "true"
  labels:
    app: {{.Spec.SSO.Proxy.SvcName}}
spec:
  selector:
    matchLabels:
      app: {{.Spec.SSO.Proxy.SvcName}}
  template:
    metadata:
      labels:
        app: {{.Spec.SSO.Proxy.SvcName}}
    spec:
      priorityClassName: {{ .Spec.PriorityClass.AppClassRef }}
      serviceAccountName: {{ .Spec.SSO.Proxy.SvcName}}
      containers:
      - name: proxy-central
        imagePullPolicy: Always
        image: {{image .Spec.ImageHub .Spec.SSO.Proxy.Image}}
        command:
          - /opt/app-root/proxy
          - --authz-addr=127.0.0.1:50052
          - --ingress-type={{.Spec.Networking.Ingress.Type}}
        ports:
          - containerPort: 8888
      - name: authz
        imagePullPolicy: Always
        image: {{  image .Spec.ImageHub .Spec.SSO.Proxy.Image }}
        command:
          - /opt/app-root/authz
          - --ingress-type={{.Spec.Networking.Ingress.Type}}