package res

import "go-starter/internal/application/dto"

type UserList struct {
	List []dto.User `json:"list"`
}
