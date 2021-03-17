package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model
	Title       string `gorm:"size:255;index;" json:"title,omitempty"`
	Cover       string `gorm:"size:512;" json:"cover,omitempty"`
	Keywords    string `gorm:"size:255;" json:"keywords,omitempty"`
	Description string `gorm:"size:512;" json:"description,omitempty"`
	Content     string `gorm:"type:text;" json:"content,omitempty"`
	Type        int    `gorm:"default:1;comment: 1 markdown 2 html" json:"type,omitempty"`
	View        int    `gorm:"default:0;" json:"view,omitempty"`
	Status      int    `gorm:"default:1;comment: 1 enable 0 disable" json:"status,omitempty"`
	Published   int    `gorm:"default:0;" json:"published,omitempty"`
	PublishedAt int    `gorm:"" json:"published_at,omitempty"`
}

type ArticleTag struct {
	ID        uint `gorm:"primarykey" json:"id"`
	ArticleId uint `gorm:"index" json:"article_id"`
	TagId     uint `gorm:"index" json:"tag_id"`
}

const (
	ArticleTypeDefault = iota
	ArticleTypeMarkdown
	ArticleTypeHtml
)

func ArticleGetAll(pageNum, pageSize int, maps map[string]interface{}, sort string) ([]*Article, error) {
	var articles []*Article
	offset := (pageNum - 1) * pageSize
	artDb := db.Model(&Article{})
	for query, args := range maps {
		artDb.Where(query, args)
	}
	if sort == "" {
		sort = ""
	}
	err := artDb.Offset(offset).Limit(pageSize).Order(sort).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func ArticleGetTotal(maps map[string]interface{}) (int, error) {
	var total int64
	artDb := db.Model(&Article{})
	for query, args := range maps {
		artDb.Where(query, args)
	}
	err := artDb.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func ArticleGetById(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).Find(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func ArticleAdd(data map[string]interface{}) (*Article, error) {
	article := &Article{
		Title:       data["title"].(string),
		Cover:       data["cover"].(string),
		Keywords:    data["keywords"].(string),
		Description: data["description"].(string),
		Content:     data["content"].(string),
		Type:        data["type"].(int),
		Published:   data["published"].(int),
	}
	if article.Published == 1 {
		article.PublishedAt = int(time.Now().Unix())
	}
	db.Transaction(func(tx *gorm.DB) error {
		// add article
		if err := tx.Create(article).Error; err != nil {
			return err
		}
		// add article tag relation
		var tags []ArticleTag
		for _, tagId := range data["tags"].([]int) {
			tags = append(tags, ArticleTag{
				ArticleId: article.ID,
				TagId:     uint(tagId),
			})
		}
		if err := tx.Create(&tags).Error; err != nil {
			return err
		}
		return nil
	})

	return article, nil
}

func ArticleUpdate(id int, data map[string]interface{}) (*Article, error) {
	var tags []ArticleTag
	for _, tagId := range data["tags"].([]int) {
		tags = append(tags, ArticleTag{
			ArticleId: uint(id),
			TagId:     uint(tagId),
		})
	}
	delete(data, "tags")
	article, err := ArticleGetById(id)
	if err != nil {
		return nil, err
	}
	if data["published"].(int) == 1 && article.PublishedAt == 0 {
		article.PublishedAt = int(time.Now().Unix())
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := db.Where("article_id = ?", id).Delete(&ArticleTag{}).Error; err != nil {
			return err
		}
		if err := tx.Model(&article).Updates(data).Error; err != nil {
			return err
		}
		if err := tx.Create(&tags).Error; err != nil {
			return err
		}
		return nil
	})
	return article, nil
}

func ArticleTagGetByArticleId(id int) ([]*Tag, error) {
	var tags []*Tag
	err := db.Model(&ArticleTag{}).Select("tags.id, tags.name, tags.flag").Joins("left join tags on article_tags.tag_id = tags.id").Where("article_tags.article_id = ?", id).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func ArticleTagGetByArticleIds(ids []int) (map[int][]*Tag, error) {
	rows, err := db.Model(&ArticleTag{}).Select("article_tags.article_id, tags.id, tags.name, tags.flag").Joins("left join tags on article_tags.tag_id = tags.id").Where("article_tags.article_id IN ?", ids).Rows()
	fmt.Printf("rows: %+v\n", rows)
	if err != nil {
		return nil, err
	}
	row := map[string]interface{}{
		"article_id": 0,
	}
	result := make(map[int][]*Tag)
	for rows.Next() {
		err = db.ScanRows(rows, row)
		if err != nil {
			rows.Close()
			return nil, err
		}
		tag := Tag{}
		err = db.ScanRows(rows, &tag)
		if err != nil {
			rows.Close()
			return nil, err
		}
		artId := int(row["article_id"].(int64))
		result[artId] = append(result[artId], &tag)
	}
	return result, nil
}
