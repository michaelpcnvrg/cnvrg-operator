apiVersion: v1
kind: Secret
metadata:
  name: oauth-proxy-webapp
  namespace: {{ ns . }}
  annotations:
    {{- range $k, $v := .Spec.Annotations }}
    {{$k}}: "{{$v}}"
    {{- end }}
  labels:
    {{- range $k, $v := .Spec.Labels }}
    {{$k}}: "{{$v}}"
    {{- end }}
data:
  conf: {{ oauthProxyConfig . .Spec.ControlPlane.WebApp.SvcName .Spec.ControlPlane.WebApp.OauthProxy.SkipAuthRegex .Spec.SSO.Provider .Spec.ControlPlane.WebApp.Port 3000 | b64enc }}