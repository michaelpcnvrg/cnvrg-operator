{{- define "ccp" }}
---
apiVersion: mlops.cnvrg.io/v1
kind: CnvrgApp
metadata:
  name: cnvrg-app
  namespace: {{ template "spec.cnvrgNs" . }}
spec:
  clusterDomain: {{ .Values.clusterDomain }}
  {{- include "spec.controlPlane" . | indent 2 }}
  {{- include "spec.registry" . | indent 2 }}
  {{- include "spec.app_dbs" . | indent 2 }}
  {{- include "spec.logging_app" . | indent 2 }}
  {{- include "spec.monitoring_app" . | indent 2 }}
  {{- include "spec.networking_app" . | indent 2 }}
  {{- include "spec.sso" . | indent 2 }}
  {{- include "spec.tenancy" . | indent 2 }}
---
{{- end }}