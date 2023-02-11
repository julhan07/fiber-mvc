package model

import (
	"app-test-fiber/entity"
	"app-test-fiber/infra"
)

func Create(userRequest *entity.UserEntity) (user *entity.UserEntity, err error) {
	db := infra.GetDB()

	if err := db.Model(&user).Create(&userRequest).Error; err != nil {
		return nil, err
	}

	return userRequest, nil
}
