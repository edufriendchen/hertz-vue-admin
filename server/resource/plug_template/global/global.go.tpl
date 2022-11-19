package global

{{- if .HasGlobal }}

import "github.com/edufriendchen/hertz-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}