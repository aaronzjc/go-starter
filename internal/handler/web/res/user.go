package res

import "go-starter/internal/service/dto"

type UserList struct {
	List []dto.User `json:"list"`
}
