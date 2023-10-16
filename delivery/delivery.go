package delivery

import (
	"net/http"
	"time"

	httpDelivery "github.com/azinudinachzab/hukumonline/delivery/http"
	"github.com/azinudinachzab/hukumonline/service"
)

type(
	Dependency struct {
		Service  service.Service
		Timezone *time.Location
	}

	Delivery struct {
		HttpServer http.Handler
	}
)

func NewDelivery(dep Dependency) *Delivery {
	return &Delivery{
		HttpServer: httpDelivery.NewHttpServer(dep.Service),
	}
}
