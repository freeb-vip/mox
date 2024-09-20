{{/*
Expand the name of the chart.
*/}}
{{- define "roundcubemail.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "roundcubemail.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "roundcubemail.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "roundcubemail.labels" -}}
helm.sh/chart: {{ include "roundcubemail.chart" . }}
{{ include "roundcubemail.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "roundcubemail.selectorLabels" -}}
app.kubernetes.io/name: {{ include "roundcubemail.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "roundcubemail.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "roundcubemail.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "roundcube.encryption" -}}
{{- if not (has . (list "none" "starttls" "ssltls")) -}}
{{ required (printf "invalid value for encryption: %s" .) nil -}}
{{- else if  eq . "starttls" -}}
tls://
{{- else if  eq . "ssltls" -}}
ssl://
{{- end -}}
{{- end -}}

# inspired by https://stackoverflow.com/a/67523275/2400785
{{- define "roundcube.desKey" -}}
{{- if .Values.config.desKey }}
{{- .Values.config.desKey -}}
{{- else -}}
{{- $lookupResult := (lookup "v1" "Secret" .Release.Namespace (include "roundcubemail.fullname" . )).data -}}
{{- if $lookupResult -}}
{{- (index $lookupResult "desKey" | b64dec) | default (randAlphaNum 64) -}}
{{- else -}}
{{- randAlphaNum 64 -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "roundcube.plugins.list" -}}
{{- $pluginList := list -}}
{{- range $plugin, $settings := .Values.config.plugins -}}
{{- if $settings }}
{{- if $settings.enabled -}}
{{- $pluginList = append $pluginList $plugin -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- $pluginList | join "," -}}
{{- end -}}

{{- define "roundcube.skins.requirements" -}}
  {{- $skinList := list -}}
  {{- range $skin, $settings := .Values.config.skins -}}
    {{- if $settings -}}
      {{- if $settings.enabled -}}
        {{- $composerPackage := printf "roundcube/%s" $skin -}}
        {{- $composerVersion := $.Chart.AppVersion -}}
        {{- if $settings.composerPackage -}}
          {{- if $settings.composerPackage.name -}}
            {{- $composerPackage = $settings.composerPackage.name -}}
          {{- end -}}
          {{- if $settings.composerPackage.version -}}
            {{- $composerVersion = $settings.composerPackage.version -}}
          {{- end -}}
        {{- end -}}
        {{- $skinRequirement := printf "%s:%s" $composerPackage $composerVersion -}}
        {{- $skinList = append $skinList $skinRequirement -}}
      {{- end -}}
    {{- end -}}
  {{- end -}}
  {{- $skinList | join " " -}}
{{- end -}}

{{- define "roundcube.plugins.requirements" -}}
{{- $pluginList := list -}}
{{- range $plugin, $settings := .Values.config.plugins -}}
{{- if $settings }}
{{- if and $settings.enabled $settings.composerPackage -}}
{{- $pluginRequirement := "" -}}
{{- if $settings.composerPackage.version -}}
{{- $pluginRequirement = printf "%s:%s" $settings.composerPackage.name $settings.composerPackage.version -}}
{{- else -}}
{{- $pluginRequirement = $settings.composerPackage.name -}}
{{- end -}}
{{- $pluginList = append $pluginList $pluginRequirement -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- $pluginList | join " " -}}
{{- end -}}


{{- define "roundcube.deployment.command" -}}
(cd /usr/src/roundcubemail && composer require roundcube/plugin-installer>=0.3.5) &&
{{- if (include "roundcube.plugins.requirements" .) -}}
(cd /usr/src/roundcubemail && composer require {{ include "roundcube.plugins.requirements" . }}) &&
{{- end -}}
{{- if (include "roundcube.skins.requirements" .) -}}
(cd /usr/src/roundcubemail && composer require {{ include "roundcube.skins.requirements" . }}) &&
{{- end -}}
/docker-entrypoint.sh php-fpm
{{- end -}}


{{- define "roundcube.helm2php" -}}
{{- $kind := kindOf . -}}
{{- if has $kind (list "slice" "map") -}}
json_decode({{- . | toJson | quote -}}, true)
{{- else if has $kind (list "string") -}}
{{- . | quote -}}
{{- else -}}
{{- . -}}
{{- end -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "roundcube.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "common.names.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}