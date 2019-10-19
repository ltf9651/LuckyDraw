package dao

import (
	"LuckyDraw/project/models"
	"github.com/go-xorm/xorm"
	"log"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

func (d *CodeDao) Get(id int) *models.Gift {
	data := &models.Gift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err != nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []models.Gift {
	datalist := make([]models.Gift, 0)
	err := d.engine.
		Asc("id").
		Find(&datalist)
	if err != nil {
		log.Println("error=", err)
		return datalist
	}
	return datalist
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Gift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *CodeDao) Delete(id int) error {
	data := &models.Gift{Id: id, SysStatus: 1}
	_, err = d.engine.Id(data.id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.Gift, columns []string) error {
	_, err = d.engine.Id(data.id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.Gift) error {
	_, err := d.engine.Insert(data)
	return err
}
