package bitwormhole

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/bitwormhole"
	myModuleVer  = "v0.0.1"
	myModuleRev  = 1
)

// Module 导出模块【github.com/bitwormhole/bitwormhole】
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(myModuleRes())

	return mb.Create()
}

//go:embed src/main/resources
var theModuleResFS embed.FS

func myModuleRes() collection.Resources {
	return collection.LoadEmbedResources(&theModuleResFS, "src/main/resources")
}
