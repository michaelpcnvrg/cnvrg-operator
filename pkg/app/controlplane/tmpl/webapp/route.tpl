apiVersion: route.openshift.io/v1
kind: Route
metadata:
  annotations:
    mlops.cnvrg.io/default-loader: "true"
    mlops.cnvrg.io/own: "true"
    mlops.cnvrg.io/updatable: "true"
    haproxy.router.openshift.io/timeout: {{ .Spec.Networking.Ingress.PerTryTimeout }}
    {{- if isTrue .Spec.SSO.Enabled }}
    sso.cnvrg.io/enabled: "true"
    sso.cnvrg.io/skipAuthRoutes: \/assets \/healthz \/public \/pack \/vscode.tar.gz \/jupyter.vsix \/gitlens.vsix \/ms-python-release.vsix \/webhooks \/api/v2/metrics \/api/v1/events/endpoint_rule_alert
    sso.cnvrg.io/central: "{{ .Spec.SSO.Central.PublicUrl }}"
    {{- end }}
    {{- range $k, $v := .Spec.Annotations }}
    {{$k}}: "{{$v}}"
    {{- end }}
  name: {{ .Spec.ControlPlane.WebApp.SvcName }}
  namespace: {{ ns . }}
  labels:
    app: {{ ns . }}
    {{- range $k, $v := .Spec.Labels }}
    {{$k}}: "{{$v}}"
    {{- end }}
spec:
  host: "{{ .Spec.ControlPlane.WebApp.SvcName }}.{{ .Spec.ClusterDomain }}"
  port:
    targetPort: {{ .Spec.ControlPlane.WebApp.Port }}
  to:
    kind: Service
    name: {{ .Spec.ControlPlane.WebApp.SvcName }}
    weight: 100
  {{- if isTrue .Spec.Networking.HTTPS.Enabled  }}
  tls:
    termination: edge
    insecureEdgeTerminationPolicy: Redirect
  {{- end }}