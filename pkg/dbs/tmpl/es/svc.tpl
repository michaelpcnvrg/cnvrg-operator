
apiVersion: v1
kind: Service
metadata:
  name: {{ .Spec.Dbs.Es.SvcName }}
  namespace: {{ ns . }}
  labels:
    app: {{ .Spec.Dbs.Es.SvcName }}
spec:
  {{- if eq .Spec.Networking.Ingress.IngressType "nodeport" }}
  type: NodePort
  {{- end }}
  ports:
  - port: {{ .Spec.Dbs.Es.Port}}
    {{- if eq .Spec.Networking.Ingress.IngressType "nodeport" }}
    nodePort: {{ .Spec.Dbs.Es.NodePort }}
    {{- end }}
  selector:
    app: {{ .Spec.Dbs.Es.SvcName }}