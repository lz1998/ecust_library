package admin

import (
	"fmt"
	"time"

	"github.com/lz1998/ecust_library/model"
)

type EcustAdmin struct {
	ID        int64     `gorm:"column:id" json:"id" form:"id"`
	Username  string    `gorm:"column:username" json:"username" form:"username"`
	Password  string    `gorm:"column:password" json:"password" form:"password"`
	Status    int32     `gorm:"column:status" json:"status" form:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

func init() {
	if err := model.Db.AutoMigrate(&EcustAdmin{}); err != nil {
		panic(err)
	}
}

func CreateAdmin(username string, password string) error {
	admins, err := ListAdmin([]string{username})
	if err != nil {
		return err
	}
	if len(admins) != 0 {
		return fmt.Errorf("username exists")
	}
	return model.Db.Save(&EcustAdmin{
		Username: username, // TODO 加密保存
		Password: password,
	}).Error
}

func ListAdmin(usernames []string) ([]*EcustAdmin, error) {
	var admins []*EcustAdmin

	q := model.Db.Model(&EcustAdmin{})
	q = q.Where("status = 0")
	if len(usernames) != 0 {
		q = q.Where("username in ?", usernames)
	}

	if err := q.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func GetAdminByUsername(username string) (*EcustAdmin, error) {
	var admin EcustAdmin
	if err := model.Db.Model(&EcustAdmin{}).Where("username = ?", username).Where("status = 0").First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func UpdateAdmin(username string, password string, status int32) error {
	var admin EcustAdmin
	if err := model.Db.Model(&EcustAdmin{}).Where("username = ?", username).First(&admin).Error; err != nil {
		return err
	}
	admin.Password = password // TODO 加密保存
	admin.Status = status
	return model.Db.Save(admin).Error
}
