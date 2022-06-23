package tokens

import (
	"github.com/Piramind/Testforgo/internal/users"
)

type Session struct {
	RefreshToken string
	User         users.User
}
