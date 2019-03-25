package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	models "github.com/port-project/api-service/port"
	parser "github.com/port-project/api-service/port/parser"
	sender "github.com/port-project/api-service/port/sender"
)

type handler struct {
	mux *http.ServeMux
}

// NewHandler create new instance of gRPC data sender
func NewHandler() Handler {
	h := http.NewServeMux()
	handler := handler{h}
	handler.setHandlers()
	return &handler
}

func portHandler(port *models.Port) {
	sender := sender.NewGrpcSender()
	sender.StorePort(port)
}

// Response contains response data
type Response struct {
	http.ResponseWriter
}

// Text is used to produce text response
func (r *Response) Text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)

	io.WriteString(r, fmt.Sprintf("%s\n", body))
}

// Handler is used to handle available routes
func (h *handler) setHandlers() {
	h.mux.HandleFunc("/parse", func(w http.ResponseWriter, r *http.Request) {

		filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		filePath += "/resources/ports.json"
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(filePath)

		resp := Response{w}
		resp.Text(http.StatusOK, "Start file parsing")

		log.Println("start file parsing")

		parser := parser.NewPortParser(portHandler)
		go parser.Parse(filePath)
	})

	h.mux.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request) {
		sender := sender.NewGrpcSender()
		ports := sender.GetAll()

		json, err := json.Marshal(ports)
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}

		resp := Response{w}
		resp.Text(http.StatusOK, string(json))
	})

	h.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{w}
		resp.Text(http.StatusNotFound, "Not found")
	})
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.mux.ServeHTTP(w, req)
}
