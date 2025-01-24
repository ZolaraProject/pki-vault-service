{{- define "pkiVaultService.name" -}}
{{- default "pki-vault-service" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Helm required labels */}}
{{- define "pkiVaultService.labels" -}}
heritage: {{ .Release.Service }}
release: {{ .Release.Name }}
chart: {{ .Chart.Name }}
app: "{{ template "pkiVaultService.name" . }}"
{{- end -}}

{{/* matchLabels */}}
{{- define "pkiVaultService.matchLabels" -}}
release: {{ .Release.Name }}
app: "{{ template "pkiVaultService.name" . }}"
{{- end -}}