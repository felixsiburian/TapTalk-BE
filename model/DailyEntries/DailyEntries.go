package DailyEntries

import (
	"TapTalk-BE/model/User"
)

type DailyEntries struct {
	Id          uint32    `json:"id" gorm:"primary_key;auto_increment"`
	UserId      uint32    `json:"user_id"`
	Users       User.User `json:"users"`
	CreatedDate string    `json:"created_date"`
	UpdatedDate string    `json:"updated_date"`
}
