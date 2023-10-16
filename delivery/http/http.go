package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/azinudinachzab/hukumonline/model"
	"github.com/azinudinachzab/hukumonline/pkg/errs"
	"github.com/azinudinachzab/hukumonline/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/thanhhh/gin-gonic-realip"
)

type HttpServer struct {
	service service.Service
}

func NewHttpServer(svc service.Service) http.Handler {
	r := gin.Default()
	d := &HttpServer{
		service: svc,
	}

	/* ***** ***** *****
	 * init middleware
	 * ***** ***** *****/
	r.Use(cors.New(cors.Config{
		AllowOrigins:   []string{"*"},
		AllowMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodDelete},
		AllowHeaders:   []string{"*"}, // "True-Client-IP", "X-Forwarded-For", "X-Real-IP", "X-Request-Id", "Origin", "Accept", "Content-Type", "Authorization", "Token"
		AllowCredentials: true,
		MaxAge:           86400,
	}))
	r.Use(realip.RealIP())
	r.Use(requestid.New())
	rlMemory := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 60,
	})
	r.Use(ratelimit.RateLimiter(rlMemory, &ratelimit.Options{
		ErrorHandler: func (c *gin.Context, info ratelimit.Info) {
			c.String(429, "Too many requests. Try again in "+ time.Until(info.ResetTime).String())
		},
		KeyFunc: func (c *gin.Context) string {
			return c.ClientIP()
		},
	}))

	/* ***** ***** *****
	 * init custom error for 404 and 405
	 * ***** ***** *****/
	r.NoRoute(d.Custom404)
	r.NoMethod(d.Custom405)

	/* ***** ***** *****
	 * init path route
	 * ***** ***** *****/
	r.GET("/", d.Home)

	// onboarding
	r.POST("/registration", d.Registration)

	// gathering
	r.POST("/invitation", d.CreateGathering)
	r.GET("/invitation", d.ResponseInvitation)

	return r
}

func (d *HttpServer) Home(c *gin.Context) {
	responseData(c, 0, httpResponse{Message: "Hello World : " + time.Now().Format(time.RFC3339)})
}

func (d *HttpServer) Registration(c *gin.Context) {
	var req model.RegistrationRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(c, err)
		return
	}

	err := d.service.Registration(c, req)
	if err != nil {
		responseError(c, err)
		return
	}

	responseData(c, http.StatusCreated, httpResponse{Message: "Registration Success"})
}

func (d *HttpServer) CreateGathering(c *gin.Context) {
	var req model.RegistrationRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(c, err)
		return
	}

	err := d.service.Registration(c, req)
	if err != nil {
		responseError(c, err)
		return
	}

	responseData(c, http.StatusCreated, httpResponse{Message: "Registration Success"})
}

func (d *HttpServer) ResponseInvitation(c *gin.Context) {
	var (
		memberID = c.Query("id")
		gatheringID = c.Query("gathering_id")
		response = c.Query("response")
	)

	err := d.service.Registration(c, req)
	if err != nil {
		responseError(c, err)
		return
	}

	responseData(c, http.StatusCreated, httpResponse{Message: "Registration Success"})
}

func (d *HttpServer) Custom404(c *gin.Context) {
	err := errs.New(model.ECodeNotFound, "route does not exist")
	responseError(c, err)
}

func (d *HttpServer) Custom405(c *gin.Context) {
	err := errs.New(model.ECodeMethodFail, "method is not valid")
	responseError(c, err)
}
