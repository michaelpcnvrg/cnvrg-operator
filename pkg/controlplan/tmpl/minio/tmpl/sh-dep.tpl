apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Spec.Minio.SvcName }}
  namespace: {{ ns . }}
  labels:
    app: {{ .Spec.Minio.SvcName }}
spec:
  selector:
    matchLabels:
      app: {{ .Spec.Minio.SvcName }}
  replicas: {{ .Spec.Minio.Replicas }}
  template:
    metadata:
      labels:
        app: {{ .Spec.Minio.SvcName }}
    spec:
      securityContext:
        runAsUser: 1000
        fsGroup: 1000
      serviceAccountName: {{ .Spec.ControlPlan.Rbac.ServiceAccountName }}
      {{- if and (ne .Spec.ControlPlan.BaseConfig.HostpathNode "") (eq .Spec.ControlPlan.Tenancy.Enabled "false") }}
      nodeSelector:
        kubernetes.io/hostname: "{{ .Spec.ControlPlan.BaseConfig.HostpathNode }}"
      {{- else if and (eq .Spec.ControlPlan.BaseConfig.HostpathNode "") (eq .Spec.ControlPlan.Tenancy.Enabled "true") }}
      nodeSelector:
        {{ .Spec.ControlPlan.Tenancy.Key }}: "{{ .Spec.ControlPlan.Tenancy.Value }}"
      {{- else if and (ne .Spec.ControlPlan.BaseConfig.HostpathNode "") (eq .Spec.ControlPlan.Tenancy.Enabled "true") }}
      nodeSelector:
        kubernetes.io/hostname: "{{ .Spec.ControlPlan.BaseConfig.HostpathNode }}"
        {{ .Spec.ControlPlan.Tenancy.Key }}: "{{ .Spec.ControlPlan.Tenancy.Value }}"
      {{- end }}
      tolerations:
        - key: {{ .Spec.ControlPlan.Tenancy.Key }}
          operator: Equal
          value: "{{ .Spec.ControlPlan.Tenancy.Value }}"
          effect: "NoSchedule"
      containers:
        - args:
            - proxy
            - sidecar
            - --domain
            - $(POD_NAMESPACE).svc.cluster.local
            - --serviceCluster
            - minio.$(POD_NAMESPACE)
            - --proxyLogLevel=warning
            - --proxyComponentLogLevel=misc:error
            - --concurrency
            - "2"
          env:
            - name: JWT_POLICY
              value: first-party-jwt
            - name: PILOT_CERT_PROVIDER
              value: istiod
            - name: CA_ADDR
              value: "istiod.{{ ansible_operator_meta.namespace }}.svc:15012"
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: INSTANCE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.podIP
            - name: SERVICE_ACCOUNT
              valueFrom:
                fieldRef:
                  fieldPath: spec.serviceAccountName
            - name: HOST_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: CANONICAL_SERVICE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.labels['service.networking.istio.io/canonical-name']
            - name: CANONICAL_REVISION
              valueFrom:
                fieldRef:
                  fieldPath: metadata.labels['service.networking.istio.io/canonical-revision']
            - name: PROXY_CONFIG
              value: |
                {"discoveryAddress":"istiod.{{ ansible_operator_meta.namespace }}.svc:15012","tracing":{"zipkin":{"address":"zipkin.{{ ansible_operator_meta.namespace }}:9411"}},"proxyMetadata":{"DNS_AGENT":""}}
            - name: ISTIO_META_POD_PORTS
              value: |-
                [
                    {"containerPort":9000}
                ]
            - name: ISTIO_META_APP_CONTAINERS
              value: minio
            - name: ISTIO_META_CLUSTER_ID
              value: Kubernetes
            - name: ISTIO_META_INTERCEPTION_MODE
              value: REDIRECT
            - name: ISTIO_META_WORKLOAD_NAME
              value: minio
            - name: ISTIO_META_OWNER
              value: "kubernetes://apis/apps/v1/namespaces/{{ ansible_operator_meta.namespace }}/deployments/minio"
            - name: ISTIO_META_MESH_ID
              value: cluster.local
            - name: TRUST_DOMAIN
              value: cluster.local
            - name: DNS_AGENT
            - name: ISTIO_KUBE_APP_PROBERS
              value: '{"/app-health/minio/livez":{"httpGet":{"path":"/minio/health/live","port":9000}},"/app-health/minio/readyz":{"httpGet":{"path":"/minio/health/ready","port":9000}}}'
          image: "{{ .Spec.Ingress.Istio.Hub }}/{{ .Spec.Ingress.Istio.ProxyImage }}:{{ .Spec.Ingress.Istio.Tag }}"
          imagePullPolicy: Always
          name: istio-proxy
          ports:
            - containerPort: 15090
              name: http-envoy-prom
              protocol: TCP
          readinessProbe:
            failureThreshold: 30
            httpGet:
              path: /healthz/ready
              port: 15021
            initialDelaySeconds: 1
            periodSeconds: 2
            timeoutSeconds: 3
          resources:
            limits:
              cpu: "2"
              memory: 1Gi
            requests:
              cpu: 100m
              memory: 128Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
                - ALL
            privileged: false
            readOnlyRootFilesystem: true
            runAsGroup: 1337
            runAsNonRoot: true
            runAsUser: 1337
          volumeMounts:
            - mountPath: /var/run/secrets/istio
              name: istiod-ca-cert
            - mountPath: /var/lib/istio/data
              name: istio-data
            - mountPath: /etc/istio/proxy
              name: istio-envoy
            - mountPath: /etc/istio/pod
              name: istio-podinfo
        - name: minio
          image: {{.Spec.Minio.Image}}
          args:
            - gateway
            - nas
            - /data
          env:
            - name: MINIO_SSE_MASTER_KEY
              valueFrom:
                secretKeyRef:
                  name: cp-object-storage
                  key: MINIO_SSE_MASTER_KEY
            - name: MINIO_ACCESS_KEY
              valueFrom:
                secretKeyRef:
                  name: cp-object-storage
                  key: CNVRG_STORAGE_ACCESS_KEY
            - name: MINIO_SECRET_KEY
              valueFrom:
                secretKeyRef:
                  name: cp-object-storage
                  key: CNVRG_STORAGE_SECRET_KEY
          ports:
            - containerPort: {{ .Spec.Minio.Port }}
          volumeMounts:
            - name: minio-storage
              mountPath: /data
          readinessProbe:
            httpGet:
              path: /minio/health/ready
              port: {{ .Spec.Minio.Port }}
            initialDelaySeconds: 5
            periodSeconds: 5
          livenessProbe:
            httpGet:
              path: /minio/health/live
              port: {{ .Spec.Minio.Port }}
            initialDelaySeconds: 60
            periodSeconds: 20
          resources:
            requests:
              cpu: {{ .Spec.Minio.CPURequest }}
              memory: {{ .Spec.Minio.MemoryRequest }}
      initContainers:
        - args:
            - istio-iptables
            - -p
            - "15001"
            - -z
            - "15006"
            - -u
            - "1337"
            - -m
            - REDIRECT
            - -i
            - '*'
            - -x
            - ""
            - -b
            - '*'
            - -d
            - 15090,15021,15020
          env:
            - name: DNS_AGENT
          image: "{{ .Spec.Ingress.Istio.Hub }}/{{ .Spec.Ingress.Istio.ProxyImage }}:{{ .Spec.Ingress.Istio.Tag }}"
          imagePullPolicy: Always
          name: istio-init
          resources:
            limits:
              cpu: "2"
              memory: 1Gi
            requests:
              cpu: 100m
              memory: 128Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              add:
                - NET_ADMIN
                - NET_RAW
              drop:
                - ALL
            privileged: false
            readOnlyRootFilesystem: false
            runAsGroup: 0
            runAsNonRoot: false
            runAsUser: 0
      volumes:
        - name: minio-storage
          persistentVolumeClaim:
            {{- if eq .Spec.Minio.SharedStorage.UseExistingClaim "" }}
            claimName: {{Minio.SvcName}}
            {{- else }}
            claimName: {{ Minio.SharedStorage.SseExistingClaim }}
            {{- end }}
        - emptyDir:
            medium: Memory
          name: istio-envoy
        - emptyDir: {}
          name: istio-data
        - downwardAPI:
            items:
            - fieldRef:
                fieldPath: metadata.labels
              path: labels
            - fieldRef:
                fieldPath: metadata.annotations
              path: annotations
          name: istio-podinfo
        - configMap:
            name: istio-ca-root-cert
          name: istiod-ca-cert