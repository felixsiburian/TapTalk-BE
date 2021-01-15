package controller

import (
	"TapTalk-BE/auth"
	"TapTalk-BE/database"
	"TapTalk-BE/model/DailyEntries"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os/user"
	"strconv"
	"time"
)

var currentTime = time.Now()
var formatDate = "2006-January-02"
var updatedDate = "2006-January-02 15:04:09"

func CheckExistUser(user_id uint32, time string) bool {
	db := database.ConnDb()
	dailyEntries := DailyEntries.DailyEntries{}
	users := db.Debug().Model(DailyEntries.DailyEntries{}).Where("user_id = ? AND created_date LIKE ?", user_id, time).Find(&dailyEntries).RowsAffected
	if users == 0 {
		return true
	}
	return false
}

func AddDailyEntries(c echo.Context) error {
	db := database.ConnDb()
	entries := DailyEntries.DailyEntries{}

	err := json.NewDecoder(c.Request().Body).Decode(&entries)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if entries.UserId <= 0 {
		return c.String(http.StatusBadRequest, "Invalid User Id")
	}

	id, err := auth.ExtractTokenId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if id != entries.UserId {
		return c.JSON(http.StatusUnauthorized, "You have no permission")
	}

	entries.CreatedDate = currentTime.Format(formatDate)
	fmt.Println(entries.CreatedDate)
	users := CheckExistUser(entries.UserId, entries.CreatedDate)
	if users == true {
		entries.UpdatedDate = time.Now().Format(updatedDate)
		err := db.Debug().Create(&entries).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if entries.UserId != 0 {
			err = db.Debug().Model(&user.User{}).Where("id = ?", entries.UserId).Take(&entries.Users).Error
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusCreated, &entries)
	}

	err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("user_id = ?", entries.UserId).Take(&entries).Updates(
		map[string]interface{}{
			"updated_date": time.Now().Format(updatedDate),
		}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if entries.UserId != 0 {
		err = db.Debug().Model(&user.User{}).Where("id = ?", entries.UserId).Take(&entries.Users).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, &entries)
}

func GetDailyEntries(c echo.Context) error {
	db := database.ConnDb()
	entries := []DailyEntries.DailyEntries{}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	quarter, err := strconv.Atoi(c.Param("quarter"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if quarter == 1 {
		err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%January%").Find(&entries).Error
		if err != nil {
			err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%February%").Find(&entries).Error
			if err != nil {
				err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%March%").Find(&entries).Error
				if err != nil {
					return c.String(http.StatusInternalServerError, "No Data Found")
				}
			}
		}
	}

	if quarter == 2 {
		err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%April%").Find(&entries).Error
		if err != nil {
			err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%May%").Find(&entries).Error
			if err != nil {
				err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%June%").Find(&entries).Error
				if err != nil {
					return c.String(http.StatusInternalServerError, "No Data Found")
				}
			}
		}
	}

	if quarter == 3 {
		err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%July%").Find(&entries).Error
		if err != nil {
			err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%August%").Find(&entries).Error
			if err != nil {
				err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%September%").Find(&entries).Error
				if err != nil {
					return c.String(http.StatusInternalServerError, "No Data Found")
				}
			}
		}
	}

	if quarter == 4 {
		err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%October%").Find(&entries).Error
		if err != nil {
			err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%November%").Find(&entries).Error
			if err != nil {
				err = db.Debug().Model(DailyEntries.DailyEntries{}).Where("created_date LIKE ? AND created_date LIKE ?", strconv.Itoa(year)+"%", "%December%").Find(&entries).Error
				if err != nil {
					return c.String(http.StatusInternalServerError, "No Data Found")
				}
			}
		}
	}

	if len(entries) > 0 {
		for i, _ := range entries {
			err := db.Debug().Model(&user.User{}).Where("id = ?", entries[i].UserId).Take(&entries[i].Users).Error
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}

	return c.JSON(http.StatusOK, &entries)

}
