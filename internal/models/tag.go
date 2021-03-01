package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model

	Name        string `gorm:"size:100" json:"name"`
	Flag        string `gorm:"index;size:50" json:"flag"`
	Icon        string `gorm:"size:512" json:"icon"`
	Description string `gorm:"size:512" json:"description"`
	Status      int    `gorm:"default:1;comment:1 enable 0 disable" json:"status"`
}

func TagGetAll(pageNum, pageSize int, maps interface{}) ([]*Tag, error) {
	var tags []*Tag
	offset := (pageNum - 1) * pageSize
	err := db.Where(maps).Offset(offset).Limit(pageSize).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func TagGetTotal(maps interface{}) (int, error) {
	var total int64
	err := db.Model(&Tag{}).Where(maps).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func TagGetById(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func TagGetByFlag(flag string) (*Tag, error) {
	var tag Tag
	err := db.Where("flag = ?", flag).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func TagAdd(data map[string]interface{}) (*Tag, error) {
	tag := &Tag{
		Name:        data["name"].(string),
		Flag:        data["flag"].(string),
		Icon:        data["icon"].(string),
		Description: data["description"].(string),
	}
	err := db.Create(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func TagUpdate(id int, data map[string]interface{}) (*Tag, error) {
	tag := Tag{
		Name:        data["name"].(string),
		Flag:        data["flag"].(string),
		Icon:        data["icon"].(string),
		Description: data["description"].(string),
	}
	err := db.Model(&Tag{}).Select("name", "flag", "icon", "description", "status", "updated_at").Where("id = ?", id).Updates(tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
