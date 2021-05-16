{{- define "spec.networking_app" }}
networking:
  https:
    enabled: {{ .Values.networking.https.enabled }}
    certSecret: "{{ .Values.networking.https.certSecret }}"
  ingress:
    type: {{ .Values.networking.ingress.type }}
    {{- if and (eq .Values.networking.ingress.type "istio") (eq .Values.spec "allinone") }}
    istioGwName: "do-not-deploy"
    {{- end }}
{{- end }}
