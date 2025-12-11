{{/*
Common labels
*/}}
{{- define "labels.common" -}}
application.giantswarm.io/branch: {{ .Chart.AppVersion | replace "#" "-" | replace "/" "-" | replace "." "-" | trunc 63 | trimSuffix "-" | quote }}
application.giantswarm.io/commit: {{ .Chart.AppVersion | quote }}
application.kubernetes.io/managed-by: {{ .Release.Service | quote }}
application.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
application.giantswarm.io/team: {{ index .Chart.Annotations "io.giantswarm.application.team" | quote }}
giantswarm.io/managed-by: {{ .Release.Name | quote }}
{{- end -}}

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
{{- or ( and .Values.appCatalog.config .Values.appCatalog.config.secret .Values.appCatalog.config.secret.values (not .Values.appCatalog.config.secret.mergeIntoCM)) ( not .Values.appCatalog.config.secret.managed ) }}
{{- end -}}

{{- define "configMapValues" -}}
{{- if .Values.appCatalog.config.secret.mergeIntoCM -}}
{{- mergeOverwrite .Values.appCatalog.config.configMap.values .Values.appCatalog.config.secret.values | toYaml | nindent 4 }}
{{- else -}}
{{ .Values.appCatalog.config.configMap.values | toYaml | nindent 4 }}
{{- end -}}
{{- end -}}
