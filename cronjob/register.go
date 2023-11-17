package cronjob

import (
	"team-work-be/config"
	"team-work-be/model"
	"time"
)

func DeleteExpiredTemporaryEmail() {
	db := config.GetDB()

	db.Where("time_end < ?", time.Now()).Unscoped().Delete(&model.TemporaryCredential{})
}
