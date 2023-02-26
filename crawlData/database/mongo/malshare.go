package mongo

import (
	"crawl/model"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MalshareDailyMongoCollection = "MalshareDaily"

type MalshareDailyRepository struct {
	provider       *mongo.MongoProvider
	collectionName string
}

// func NewMalshareMongoRepository(provider provider.MongoProvider) *MalshareDailyRepository {
// 	repo := &MalshareDailyRepository{provider, MalshareDailyMongoCollection}
// 	collection, close := repo.collection()
// 	defer close()
// 	collection.EnsureIndex(mgo.Index{
// 		key: []string{
// 			"Md5",
// 		},
// 		Unique: true,
// 	})
// 	return repo
// }

func (repo *MalshareDailyRepository) collection() (collection *mgo.Collection, close func()) {
	session := repo.provider.MongoClient().GetCopySession()
	close = session.Close

	return session.DB(repo.provider.MongoClient().Database()).C(repo.collectionName), close
}

func (repo *MalshareDailyRepository) All() ([]model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	result := make([]model.MalshareDaily, 0)
	err := collection.Find(nil).All(&result)
	return result, repo.provider.NewError(err)
}
func (repo *MalshareDailyRepository) FindByID(id string) (*model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid id")
	}

	var user model.MalshareDaily
	err := collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return &user, repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) FindByMd5(Md5 string) (*model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	var MalshareDaily model.MalshareDaily
	err := collection.Find(bson.M{"md5": Md5}).One(&MalshareDaily)
	return &MalshareDaily, repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) FindBySha256(Sha256 string) (*model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	var MalshareDaily model.MalshareDaily
	err := collection.Find(bson.M{"Sha256": Sha256}).One(&MalshareDaily)
	return &MalshareDaily, repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) FindBySha1(Sha1 string) (*model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	var MalshareDaily model.MalshareDaily
	err := collection.Find(bson.M{"Sha1": Sha1}).One(&MalshareDaily)
	return &MalshareDaily, repo.provider.NewError(err)
}
func (repo *MalshareDailyRepository) FindByBase64(Base64 string) (*model.MalshareDaily, error) {
	collection, close := repo.collection()
	defer close()

	var MalshareDaily model.MalshareDaily
	err := collection.Find(bson.M{"Base64": Base64}).One(&MalshareDaily)
	return &MalshareDaily, repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) Save(MalshareDaily model.MalshareDaily) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Insert(MalshareDaily)
	return repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) RemoveByID(id string) error {
	collection, close := repo.collection()
	defer close()

	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}

	err := collection.RemoveId(bson.ObjectIdHex(id))
	return repo.provider.NewError(err)
}
func (repo *MalshareDailyRepository) RemoveByMd5(Md5 string) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Remove(bson.M{"Md5": Md5})
	return repo.provider.NewError(err)
}

func (repo *MalshareDailyRepository) RemoveBySha256(Sha256 string) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Remove(bson.M{"Sha256": Sha256})
	return repo.provider.NewError(err)
}
func (repo *MalshareDailyRepository) RemoveBySha1(Sha1 string) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Remove(bson.M{"Sha1": Sha1})
	return repo.provider.NewError(err)
}
func (repo *MalshareDailyRepository) RemoveByBase64(Base64 string) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Remove(bson.M{"Base64": Base64})
	return repo.provider.NewError(err)
}
