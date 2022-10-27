apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Spec.SSO.Jwks.Name }}
  namespace: {{.Namespace }}
  annotations:
    mlops.cnvrg.io/default-loader: "true"
    mlops.cnvrg.io/own: "true"
    mlops.cnvrg.io/updatable: "true"
    {{- range $k, $v := .Spec.Annotations }}
    {{$k}}: "{{$v}}"
    {{- end }}
  labels:
    app: cnvrg-jwks
    {{- range $k, $v := .Spec.Labels }}
    {{$k}}: "{{$v}}"
    {{- end }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Spec.SSO.Jwks.Name }}
  template:
    metadata:
      labels:
        app: {{ .Spec.SSO.Jwks.Name }}
        {{- range $k, $v := .Spec.Labels }}
        {{$k}}: "{{$v}}"
        {{- end }}
    spec:
      priorityClassName: {{ .Spec.PriorityClass.AppClassRef }}
      serviceAccountName: cnvrg-jwks
      {{- if isTrue .Spec.Tenancy.Enabled }}
      nodeSelector:
        "{{ .Spec.Tenancy.Key }}": "{{ .Spec.Tenancy.Value }}"
      tolerations:
        - operator: "Exists"
      {{- end }}
      containers:
      - name: cnvrg-jwks
        command:
          - /usr/bin/cnvrg-jwks
          - start
        image: {{ image .Spec.ImageHub .Spec.SSO.Jwks.Image }}
        imagePullPolicy: Always
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
          limits:
            cpu: 1000m
            memory: 1Gi
        volumeMounts:
          - mountPath: /opt/app-root/config
            name: cnvrg-jwks
        ports:
          - containerPort: 8080
        livenessProbe:
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          httpGet:
            port: 8080
            path: /healthz
        readinessProbe:
          successThreshold: 1
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
          httpGet:
            port: 8080
            path: /healthz
      - name: redis-cache
        image: {{ image .Spec.ImageHub .Spec.SSO.Jwks.Cache.Image }}
        resources:
          requests:
            cpu: 200m
            memory: 200Mi
          limits:
            cpu: 1000m
            memory: 1Gi
        livenessProbe:
          successThreshold: 1
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          exec:
            command: ["redis-cli", "ping"]
        readinessProbe:
          successThreshold: 1
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          exec:
            command: [ "redis-cli", "ping" ]
      volumes:
        - name: cnvrg-jwks
          configMap:
            name: cnvrg-jwks
