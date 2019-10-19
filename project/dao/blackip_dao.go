package dao

import (
	"LuckyDraw/project/models"
	"github.com/go-xorm/xorm"
	"log"
)

type BliackIpDao struct {
	engine *xorm.Engine
}

func NewBliackIpDao(engine *xorm.Engine) *BliackIpDao {
	return &BliackIpDao{
		engine: engine,
	}
}

func (d *BliackIpDao) Get(id int) *models.Gift {
	data := &models.Gift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err != nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *BliackIpDao) GetAll() []models.Gift {
	datalist := make([]models.Gift, 0)
	err := d.engine.
		Asc("sys_status").
		Asc("displaycode").
		Find(&datalist)
	if err != nil {
		log.Println("error=", err)
		return datalist
	}
	return datalist
}

func (d *BliackIpDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Gift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *BliackIpDao) Delete(id int) error {
	data := &models.Gift{Id: id, SysStatus: 1}
	_, err = d.engine.Id(data.id).Update(data)
	return err
}

func (d *BliackIpDao) Update(data *models.Gift, columns []string) error {
	_, err = d.engine.Id(data.id).MustCols(columns...).Update(data)
	return err
}

func (d *BliackIpDao) Create(data *models.Gift) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BliackIpDao) GetByIp(ip string) *models.BlackIpList {
	datalist := make([]models.BlackIpList, 0)
	err := d.engine.
		Where("ip=?", ip).
		Asc("id").
		Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	}
	return &datalist[0]
}
