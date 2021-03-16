package models

type Tag struct {
	Model

	Name        string `gorm:"size:100" json:"name,omitempty"`
	Flag        string `gorm:"index;size:50" json:"flag,omitempty"`
	Icon        string `gorm:"size:512" json:"icon,omitempty"`
	Description string `gorm:"size:512" json:"description,omitempty"`
	Status      int    `gorm:"default:1;comment:1 enable -1 disable" json:"status,omitempty"`
}

func TagGetAll(pageNum, pageSize int, maps map[string]interface{}) ([]*Tag, error) {
	var tags []*Tag
	offset := (pageNum - 1) * pageSize
	tagDb := db.Model(&Tag{})
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err := tagDb.Offset(offset).Limit(pageSize).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func TagGetSimpleAll(maps map[string]interface{}) ([]*Tag, error) {
	var tags []*Tag
	tagDb := db.Model(&Tag{})
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err := tagDb.Select("id, name, flag").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func TagGetTotal(maps map[string]interface{}) (int, error) {
	var total int64
	tagDb := db.Model(&Tag{})
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err := tagDb.Count(&total).Error
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

func TagGetByIds(ids []int) ([]*Tag, error) {
	var tags []*Tag
	err := db.Select("id", "name", "flag").Where("id IN ?", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
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
	tag, _ := TagGetById(id)
	err := db.Model(&tag).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}
