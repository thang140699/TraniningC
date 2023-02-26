package repository

import "crawl/model"

type MalshareDailyRepository interface {
	All() ([]model.MalshareDaily, error)
	Pagination(page int, limit int, condition map[string]interface{}) (int, []model.MalshareDaily, error)

	FindByID(id string) (*model.MalshareDaily, error)
	FindByMd5(Md5 string) (*model.MalshareDaily, error)
	FindBySha256(Sha256 string) (*model.MalshareDaily, error)
	FindBySha1(Sha1 string) (*model.MalshareDaily, error)
	FindByBase64(Base64 string) (*model.MalshareDaily, error)

	Save(MalshareDaily model.MalshareDaily) error

	RemoveByMd5(Md5 string) error
	RemoveBySha256(Sha256 string) error
	RemoveBySha1(Sha1 string) error
	RemoveByBase64(Base64 string) error
}
