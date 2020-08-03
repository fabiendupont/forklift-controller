package web

import (
	"github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	"github.com/konveyor/controller/pkg/logging"
	vsphere "github.com/konveyor/virt-controller/pkg/controller/provider/web/vsphere"
)

//
// Shared logger.
var Log *logging.Logger

func init() {
	log := logging.WithName("web")
	Log = &log
}

//
// All handlers.
func All(container *container.Container) (all []libweb.RequestHandler) {
	vsphere.Log = Log
	all = append(
		all,
		vsphere.Handlers(container)...)

	return
}
