{{- define "operator" }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnvrg-operator
  namespace: {{ template "spec.cnvrgNs" . }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: cnvrg-operator-role
rules:
  - apiGroups:
      - '*'
    resources:
      - '*'
    verbs:
      - '*'
  - apiGroups:
      - mlops.cnvrg.io
    resources:
      - cnvrgapps
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
  - apiGroups:
      - mlops.cnvrg.io
    resources:
      - cnvrgapps/status
    verbs:
      - get
      - patch
      - update
  - apiGroups:
      - mlops.cnvrg.io
    resources:
      - cnvrginfras
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
  - apiGroups:
      - mlops.cnvrg.io
    resources:
      - cnvrginfras/status
    verbs:
      - get
      - patch
      - update
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cnvrg-operator-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnvrg-operator-role
subjects:
  - kind: ServiceAccount
    name: cnvrg-operator
    namespace: {{ template "spec.cnvrgNs" . }}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: cnvrg-operator
  name: cnvrg-operator
  namespace: {{ template "spec.cnvrgNs" . }}
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: cnvrg-operator
  template:
    metadata:
      labels:
        control-plane: cnvrg-operator
    spec:
      containers:
        - command:
            - /opt/app-root/cnvrg-operator
            - run
            - --max-concurrent-reconciles
            - "3"
          image: "docker.io/cnvrg/cnvrg-operator:{{ .Chart.Version }}"
          imagePullPolicy: Always
          name: cnvrg-operator
          resources:
            limits:
              cpu: 1000m
              memory: 1000Mi
            requests:
              cpu: 500m
              memory: 200Mi
      serviceAccountName: cnvrg-operator
      terminationGracePeriodSeconds: 10
---
{{- end }}