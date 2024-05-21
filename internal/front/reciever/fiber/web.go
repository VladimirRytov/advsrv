package front

import (
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Reciever struct {
	tlsEnabled  bool
	requestGate Requests
	frontConv   FrontConverter
	b64         Base64EncodeDecoder
	server      *fiber.App
	api         fiber.Router
	v1          fiber.Router
}

func Create(req Requests, fc FrontConverter, b64 Base64EncodeDecoder) *Reciever {
	rec := new(Reciever)
	rec.server = fiber.New(fiber.Config{
		BodyLimit: 10 << 25,
	})
	rec.server.Use(compress.New())
	rec.api = rec.server.Group("/api")
	rec.v1 = rec.api.Group("/v1")
	rec.bindClientApiv1()
	rec.bindOrderApiv1()
	rec.bindTagApiv1()
	rec.bindExtraChargeApiv1()
	rec.bindCostRateApiv1()
	rec.bindBlockAdvertisementsApiv1()
	rec.bindLineAdvertisementsApiv1()
	rec.bindAuthenticateApiv1()
	rec.bindUserApiv1()
	rec.bindSubscribeApiv1()
	rec.bindFilesApiv1()
	rec.bindActiveSubscribeApiv1()
	rec.requestGate = req
	rec.frontConv = fc
	rec.b64 = b64
	return rec
}

func (r *Reciever) bindClientApiv1() {
	r.v1.Get("/clients/", r.Clients)
	r.v1.Get("/clients/:name", r.Client)
	r.v1.Post("/clients/", r.NewClient)
	r.v1.Put("/clients/:name", r.UpdateClient)
	r.v1.Delete("/clients/:name", r.RemoveClient)
}

func (r *Reciever) bindOrderApiv1() {
	r.v1.Get("/orders/", r.Orders)
	r.v1.Get("/orders/:id", r.Order)
	r.v1.Post("/orders/", r.NewOrder)
	r.v1.Put("/orders/:id", r.UpdateOrder)
	r.v1.Delete("/orders/:id", r.DeleteOrder)
}

func (r *Reciever) bindBlockAdvertisementsApiv1() {
	r.v1.Get("/blockadvertisements/", r.BlockAdvertisements)
	r.v1.Get("/blockadvertisements/:id", r.BlockAdvertisement)
	r.v1.Post("/blockadvertisements/", r.NewBlockAdvertisement)
	r.v1.Put("/blockadvertisements/:id", r.UpdateBlockAdvertisement)
	r.v1.Delete("/blockadvertisements/:id", r.RemoveBlockAdvertisement)
}

func (r *Reciever) bindLineAdvertisementsApiv1() {
	r.v1.Get("/lineadvertisements/", r.LineAdvertisements)
	r.v1.Get("/lineadvertisements/:id", r.LineAdvertisement)
	r.v1.Post("/lineadvertisements/", r.NewLineAdvertisement)
	r.v1.Put("/lineadvertisements/:id", r.UpdateLineAdvertisement)
	r.v1.Delete("/lineadvertisements/:id", r.RemoveLineAdvertisement)
}

func (r *Reciever) bindTagApiv1() {
	r.v1.Get("/tags/", r.Tags)
	r.v1.Get("/tags/:name", r.TagByName)
	r.v1.Post("/tags/", r.NewTag)
	r.v1.Put("/tags/:name", r.UpdateTag)
	r.v1.Delete("/tags/:name", r.RemoveTag)
}

func (r *Reciever) bindExtraChargeApiv1() {
	r.v1.Get("/extracharges/", r.ExtraCharges)
	r.v1.Get("/extracharges/:name", r.ExtraChargeByName)
	r.v1.Post("/extracharges/", r.NewExtraCharge)
	r.v1.Put("/extracharges/:name", r.UpdateExtraCharge)
	r.v1.Delete("/extracharges/:name", r.RemoveExtraCharge)
}

func (r *Reciever) bindCostRateApiv1() {
	r.v1.Get("/costrates/", r.CostRates)
	r.v1.Get("/costrates/:name", r.CostRateByName)
	r.v1.Post("/costrates/", r.NewCostRate)
	r.v1.Put("/costrates/:name", r.UpdateCostRate)
	r.v1.Delete("/costrates/:name", r.RemoveCostRate)
}

func (r *Reciever) bindAuthenticateApiv1() {
	r.v1.Post("/auth", r.AuthUser).Name("basicAuth")
}

func (r *Reciever) bindUserApiv1() {
	r.v1.Get("/users/", r.Users)
	r.v1.Get("/users/:name", r.UserByName)
	r.v1.Post("/users/", r.NewUser)
	r.v1.Put("/users/:name", r.UpdateUser)
	r.v1.Delete("/users/:name", r.DeleteUser)
}

func (r *Reciever) bindSubscribeApiv1() {
	r.v1.Get("/subscribers/", r.Subscribers)
	r.v1.Get("/subscribers/:id", r.Subscriber)
	r.v1.Post("/subscribers/", r.Subscribe)
	r.v1.Delete("/subscribers/:name", r.CancelSubscribe)
}

func (r *Reciever) bindActiveSubscribeApiv1() {
	r.v1.Use("/subscribers/ws", r.WebSock)
	r.v1.Get("/subscribers/ws/:id", websocket.New(r.ActiveSubscribe))

}

func (r *Reciever) bindFilesApiv1() {
	r.v1.Get("/files/", r.Files)
	r.v1.Get("/files/:name", r.FileByName)
	r.v1.Post("/files/", r.UploadFile)
	r.v1.Delete("/files/:name", r.RemoveFile)
}

func (r *Reciever) Listen(adress string) error {
	return r.server.Listen(adress)
}

func (r *Reciever) ListenTLS(adress, certFile, keyFile string) error {
	return r.server.ListenTLS(adress, certFile, keyFile)
}

func (r *Reciever) ShutDown() error {
	return r.server.ShutdownWithTimeout(10 * time.Second)
}
