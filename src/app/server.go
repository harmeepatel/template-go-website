package app

// 	app.Static("/static/", "./web/static/", fiber.Static{
// 		Compress:      true,
// 		ByteRange:     true,
// 		Browse:        true,
// 		Index:         "./web/static/",
// 		CacheDuration: 7 * 24 * time.Hour,
// 		MaxAge:        3600,
// 	})
import (
	"errors"
	"log"
	"net/http"

	handler "makeyourcareer.net/app/handlers"
)

// add database
// const dev_db = "./src/data/test.db"
type app struct {
	server http.Server
}

func NewApp(address string) (a *app) {
	a = &app{
		server: http.Server{
			Addr: ":" + address,
		},
	}
	return
}

func (a *app) InitStaticServer() {
	static := http.FileServer(http.Dir("./web/static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", static))
}

func (a *app) InitRoutes() {
	http.Handle("GET /", handler.Index())
	http.Handle("GET /link_2", handler.Page2())
	http.Handle("GET /link_3", handler.Page3())
}

func (a *app) Run() {
	log.Println(" ----- starting ----- ")
	defer log.Println(" ----- stopped -----")

	if err := http.ListenAndServe(a.server.Addr, nil); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v\n", err)
	}
}
