package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/messaging"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/observability"
	"net/http"
	"time"
)

type SSEHandlerFactory struct {
	subscriber messaging.Subscriber
	observer   observability.SubmitLog
}
type ServeNotificationsOverSSE func(w http.ResponseWriter, r *http.Request)

func (f *SSEHandlerFactory) GetSSEHandler() ServeNotificationsOverSSE {
	cOrder := f.subscriber.GetOrderChannel()
	return func(w http.ResponseWriter, r *http.Request) {
		// Set http headers required for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// You may need this locally for CORS requests
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Create a channel for client disconnection
		clientGone := r.Context().Done()

		rc := http.NewResponseController(w)
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-clientGone:
				fmt.Println("Client disconnected")
				return
			case <-t.C:
				// Send an event to the client
				// Here we send only the "data" field, but there are few others
				_, err := fmt.Fprintf(w, "data: The time is %s\n\n", time.Now().Format(time.UnixDate))
				if err != nil {
					return
				}
				err = rc.Flush()
				if err != nil {
					return
				}
			case rOrder := <-cOrder:
				err := f.observer("Received message from Messaging ")
				if err != nil {
				}
				jOrder, err := json.Marshal(rOrder)
				if err != nil {
				}
				_, err = fmt.Fprintf(w, "\n%s\n\n", string(jOrder))
				if err != nil {
					return
				}
				err = rc.Flush()
				if err != nil {
					return
				}
				err = f.observer("Sent notification to user")
				if err != nil {
				}
			}
		}
	}
}
