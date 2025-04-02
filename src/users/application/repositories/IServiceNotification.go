package repositories

import "Multidiciplinario/src/users/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.User) error
}
