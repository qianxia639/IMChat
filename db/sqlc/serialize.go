package db

import (
	"encoding/json"
)

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
