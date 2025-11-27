package gateaway

import (
	"encoding/json"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/model"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/messaging"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/observability"
	"io"
	"net/http"
)

type GatewayHandlerFactory struct {
	publisher messaging.Publisher
	observer  observability.SubmitLog
}

type GatewayHttpHandler func(w http.ResponseWriter, r *http.Request)

func (f *GatewayHandlerFactory) SubmitResponse(w http.ResponseWriter, statusCode int, response []byte) error {
	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	return err
}

func (f *GatewayHandlerFactory) ProduceHandler() GatewayHttpHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		var err = f.observer("Received user request") // Submit observability log at start
		defer func(err error) {
			if err != nil {
				err = f.observer("User request processed without error") // Submit observability log at end
				w.WriteHeader(http.StatusCreated)
				_, err = w.Write([]byte("OK")) //TODO change response later
			} else {
				err = f.observer("User request processed with error")                    // Submit observability log at end
				err = f.SubmitResponse(w, http.StatusInternalServerError, []byte("NOK")) //TODO change response later
			}
			return
		}(err)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		var o model.Order
		err = json.Unmarshal(body, &o)
		if err != nil {

		}
		err = f.publisher.Publish(o)
		if err != nil {

		}
		err = f.observer("Sent user request to messaging")
		if err != nil {

		}
	}
}
