package models

import "sample/pkg/app"


func GetAll(i interface{}) error {
	return app.Db.Find(i).Error
}

func GetOne(i interface{}, id string) error {
	return app.Db.First(i, id).Error
}

func Create(i interface{}) error {
	return app.Db.Create(i).Error
}

func Save(i interface{}) error {
	return app.Db.Save(i).Error
}

func Delete(i interface{}, id string) error{
	app.Db.Where("id = ?", id).Delete(i)
	return nil
}
