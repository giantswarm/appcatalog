{{- define "resource.configmap.name" -}}
{{- if and .Values.appCatalog.config .Values.appCatalog.config.configMap -}}
{{- if eq .Values.appCatalog.config.configMap.name "" -}}
{{- .Values.appCatalog.name -}}-catalog
{{- else -}}
{{ .Values.appCatalog.config.configMap.name }}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "resource.secret.name" -}}
{{- if and .Values.appCatalog.config .Values.appCatalog.config.secret -}}
{{- if eq .Values.appCatalog.config.secret.name "" -}}
{{- .Values.appCatalog.name -}}-catalog
{{- else -}}
{{ .Values.appCatalog.config.secret.name }}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "configMapExists" -}}
{{- or ( and .Values.appCatalog.config .Values.appCatalog.config.configMap .Values.appCatalog.config.configMap.values) ( not .Values.appCatalog.config.configMap.managed ) }}
{{- end -}}

{{- define "secretExists" -}}
{{- or ( and .Values.appCatalog.config .Values.appCatalog.config.secret .Values.appCatalog.config.secret.values) ( not .Values.appCatalog.config.secret.managed ) }}
{{- end -}}
