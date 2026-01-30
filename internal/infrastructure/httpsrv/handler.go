package httpsrv

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
)

type WebhookHandler struct {
	updateHandler ihandler.UpdateHandler
	webhookPath   string
}

func NewWebhookHandler(uh ihandler.UpdateHandler, wp string) *WebhookHandler {
	return &WebhookHandler{
		updateHandler: uh,
		webhookPath:   wp,
	}
}

func (wh *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != wh.webhookPath {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("failed to read body: %v", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("failed to close io reader: %v", err)
			return
		}
	}(r.Body)

	var update tgapi.Update
	if err = json.Unmarshal(body, &update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid JSON: %v", err)
		return
	}

	if err = wh.updateHandler.Handle(&update); err != nil {
		log.Printf("[handle] failed to handle update: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *WebhookHandler) Handler() http.Handler {
	return Recovery(Logger(Method(http.MethodPost, http.HandlerFunc(wh.ServeHTTP))))
}
