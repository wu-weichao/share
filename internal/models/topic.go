package models

type Topic struct {
	Model

	Title  string `gorm:"size:255" json:"title,omitempty"`
	Url    string `gorm:"size:512" json:"url,omitempty"`
	Sort   int    `gorm:"default:9999" json:"sort,omitempty"`
	Status int    `gorm:"default:1;comment: 1 enable -1 disable" json:"status,omitempty"`
}

func TopicGetAll(pageNum, pageSize int, maps map[string]interface{}) ([]*Topic, error) {
	var topics []*Topic
	offset := (pageNum - 1) * pageSize
	topicDb := db.Model(&Topic{})
	for query, args := range maps {
		topicDb.Where(query, args)
	}
	err := topicDb.Offset(offset).Limit(pageSize).Order("sort").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func TopicGetSimpleAll(maps map[string]interface{}) ([]*Topic, error) {
	var topics []*Topic
	topicDb := db.Model(&Topic{})
	for query, args := range maps {
		topicDb.Where(query, args)
	}
	err := topicDb.Select("id, title, url").Order("sort").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func TopicGetTotal(maps map[string]interface{}) (int, error) {
	var total int64
	topicDb := db.Model(&Topic{})
	for query, args := range maps {
		topicDb.Where(query, args)
	}
	err := topicDb.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func TopicGetById(id int) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func TopicAdd(data map[string]interface{}) (*Topic, error) {
	topic := &Topic{
		Title:  data["title"].(string),
		Url:    data["url"].(string),
		Sort:   data["sort"].(int),
		Status: data["status"].(int),
	}
	err := db.Create(topic).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func TopicUpdate(id int, data map[string]interface{}) (*Topic, error) {
	topic, _ := TopicGetById(id)
	err := db.Model(&topic).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func TopicDelete(id int) error {
	return db.Delete(&Topic{}, id).Error
}
