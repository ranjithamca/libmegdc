/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package ubuntu


import (
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

const (
	nilavuConf = `sed -i 's/^[ \t]*bundle exec passenger start -a 127.0.0.1 -p 8080 -d -e production.*/    bundle exec passenger start -a %s -p 8080 -d -e production/' /etc/init/verticenilavu.conf`
)

var ubuntunilavuinstall *UbuntuNilavuInstall

func init() {
	ubuntunilavuinstall = &UbuntuNilavuInstall{}
	templates.Register("UbuntuNilavuInstall", ubuntunilavuinstall)
}

type UbuntuNilavuInstall struct{
	hostip string
}

func (tpl *UbuntuNilavuInstall) Render(p urknall.Package) {
	p.AddTemplate("nilavu", &UbuntuNilavuInstallTemplate{
		hostip: tpl.hostip,
	})
}

func (tpl *UbuntuNilavuInstall) Options(t *templates.Template) {
	if host,ok := t.Options[HOST]; ok {
		tpl.hostip = host
	}
}

func (tpl *UbuntuNilavuInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, tpl,inputs)
}

type UbuntuNilavuInstallTemplate struct{
		hostip string
}

func (m *UbuntuNilavuInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("verticenilavu",
		InstallPackages("verticenilavu"),
	)
	pkg.AddCommands("conf",
	Shell(fmt.Sprintf(nilavuConf, m.hostip)),
	Shell("sudo restart verticenilavu"),
	)
}
