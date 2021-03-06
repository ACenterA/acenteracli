/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"text/template"

	"github.com/wallix/awless/gen/aws"
)

func generateGlobal() {
	templ, err := template.New("global").Parse(globalTempl)
	if err != nil {
		panic(err)
	}

	writeTemplateToFile(templ, aws.GlobalDefinitions, GLOBAL_DIR, "gen_global.go")
}

const globalTempl = `/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package global

var (
  {{- range $, $prop := . }}
  {{ $prop.AwlessLabel }} = IfThenElse(os.Getenv("ACENTERA_{{ $prop.AwlessLabel }}") != "", os.Getenv("ACENTERA_{{ $prop.AwlessLabel }}"), "{{ $prop.Value }}")
  {{- end }}
)

`
