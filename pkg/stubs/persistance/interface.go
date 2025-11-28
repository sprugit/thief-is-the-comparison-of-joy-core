package persistance

import "github.com/sprugit/thief-is-the-comparison-of-joy-core/pkg/model"

type Persist func(model.Order) error

type Fetch func(model.Order) (model.Order, error)
