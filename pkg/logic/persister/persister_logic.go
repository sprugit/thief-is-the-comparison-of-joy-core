package persister

import (
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/messaging"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/observability"
	"github.com/sprugit/thief-is-the-comparison-of-joy-core.git/pkg/stubs/persistance"
)

type PersistFuncFactory struct {
	subscriber messaging.Subscriber
	publisher  messaging.Publisher
	persist    persistance.Persist
	observer   observability.SubmitLog
}
type PersistFunc func()

func (f *PersistFuncFactory) getPersistLogic() PersistFunc {
	oChannel := f.subscriber.GetOrderChannel()
	return func() {
		select {
		case o := <-oChannel:
			{
				err := f.observer("Received message over Messaging")
				if err != nil {
					return
				}
				err = f.persist(o)
				if err != nil {
					return
				}
				err = f.publisher.Publish(o)
				if err != nil {
					return
				}
				err = f.observer("Published message over Messaging")
				if err != nil {
					return
				}
			}
		}
	}

}
