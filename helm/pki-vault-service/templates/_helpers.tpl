{{- define "pkivaultservice.name" -}}
{{- default "pki-vault-service" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Helm required labels */}}
{{- define "pkivaultservice.labels" -}}
heritage: {{ .Release.Service }}
release: {{ .Release.Name }}
chart: {{ .Chart.Name }}
app: "{{ template "pkivaultservice.name" . }}"
{{- end -}}

{{/* matchLabels */}}
{{- define "pkivaultservice.matchLabels" -}}
release: {{ .Release.Name }}
app: "{{ template "pkivaultservice.name" . }}"
{{- end -}}