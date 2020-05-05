{{- define "resource.configmap.name" -}}
{{- if and .Values.appCatalog.config .Values.appCatalog.config.configMap -}}
{{- if eq .Values.appCatalog.config.configMap.name "" -}}
{{- .Values.appCatalog.name -}}
{{- else -}}
{{ .Values.appCatalog.config.configMap.name }}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "resource.secret.name" -}}
{{- if and .Values.appCatalog.config .Values.appCatalog.config.secret -}}
{{- if eq .Values.appCatalog.config.secret.name "" -}}
{{- .Values.appCatalog.name -}}
{{- else -}}
{{ .Values.appCatalog.config.secret.name }}
{{- end -}}
{{- end -}}
{{- end -}}
