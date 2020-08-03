package base

import (
	"github.com/gin-gonic/gin"
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	api "github.com/konveyor/virt-controller/pkg/apis/virt/v1alpha1"
	model "github.com/konveyor/virt-controller/pkg/controller/provider/model/vsphere"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strconv"
	"time"
)

//
// Root - all routes.
const (
	Root = libweb.Root
)

//
// Base handler.
type Handler struct {
	libweb.Consistent
	libweb.Paged
	// Container
	Container *libcontainer.Container
	// Provider referenced in the request.
	Provider *api.Provider
	// Reconciler responsible for the provider.
	Reconciler libcontainer.Reconciler
	// Resources include details.
	Detail bool
}

//
// Prepare to handle the request.
func (h *Handler) Prepare(ctx *gin.Context) int {
	status := h.Paged.Prepare(ctx)
	if status != http.StatusOK {
		return status
	}
	status = h.setProvider(ctx)
	if status != http.StatusOK {
		return status
	}
	status = h.setDetail(ctx)
	if status != http.StatusOK {
		return status
	}

	return http.StatusOK
}

//
// Set the provider.
func (h *Handler) setProvider(ctx *gin.Context) int {
	var found bool
	h.Provider = &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			Namespace: ctx.Param("ns1"),
			Name:      ctx.Param("provider"),
		},
	}
	if h.Provider.Name != "" {
		if h.Reconciler, found = h.Container.Get(h.Provider); !found {
			return http.StatusNotFound
		}
		h.Provider = h.Reconciler.Owner().(*api.Provider)
		status := h.EnsureConsistency(h.Reconciler, time.Second*30)
		if status != http.StatusOK {
			return status
		}
	}

	return http.StatusOK
}

//
// Set detail
func (h *Handler) setDetail(ctx *gin.Context) int {
	q := ctx.Request.URL.Query()
	pDetail := q.Get("detail")
	if len(pDetail) > 0 {
		b, err := strconv.ParseBool(pDetail)
		if err == nil {
			h.Detail = b
		} else {
			return http.StatusBadRequest
		}
	}

	return http.StatusOK
}

//
// REST Resource.
type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//
// Build the resource using the model.
func (r *Resource) With(m *model.Base) {
	r.ID = m.ID
	r.Name = m.Name
}
