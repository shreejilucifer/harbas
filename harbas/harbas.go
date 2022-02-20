package harbas

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/shreejilucifer/harbas/render"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Harbas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config   config
}

type config struct {
	port     string
	renderer string
}

func (h *Harbas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := h.Init(pathConfig)
	if err != nil {
		return err
	}

	err = h.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	infoLog, errorLog := h.startLoggers()
	h.InfoLog = infoLog
	h.ErrorLog = errorLog
	h.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	h.Version = version
	h.RootPath = rootPath
	h.Routes = h.routes().(*chi.Mux)

	h.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	h.createRenderer()

	return nil
}

func (h *Harbas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := h.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Harbas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     h.ErrorLog,
		Handler:      h.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	h.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	h.ErrorLog.Fatal(err)
}

func (h *Harbas) checkDotEnv(path string) error {
	err := h.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (h *Harbas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (h *Harbas) createRenderer() {
	myRenderer := render.Render{
		Renderer: h.config.renderer,
		RootPath: h.RootPath,
		Port:     h.config.port,
	}
	h.Render = &myRenderer
}
