package ports

import "github.com/MehdiEidi/pubsub/internal/core/domain"

type Publisher interface {
	Publish(domain.Message)
}

type Subscriber interface {
	Subscribe(string)
}
