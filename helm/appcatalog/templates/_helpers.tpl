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
