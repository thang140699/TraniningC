package download

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/mgo.v2"
)

const (
	DB_SERVER     = "DB_SERVER"
	DB_TIMEOUT    = "DB_TIMEOUT"
	DB_DATABASE   = "DB_DATABASE"
	DB_COLLECTION = "DB_COLLECTION"

	DB_DEFAULT_TIMEOUT = 10000
)

type MongoDB struct {
	timeout    int
	URI        string
	database   string
	collection string

	session *mgo.Session
	url     string
}

var isLocked = false

func NewMongoDBFromURL(u string) *MongoDB {
	timeout := DB_DEFAULT_TIMEOUT

	instance := MongoDB{
		url:     u,
		timeout: timeout,
	}

	err := instance.Init()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &instance
}
func NewMongoDB(config map[string]string) *MongoDB {
	timeout, err := strconv.Atoi(config[DB_TIMEOUT])
	if err != nil {
		timeout = DB_DEFAULT_TIMEOUT
	}
	instance := MongoDB{
		database:   config[DB_DATABASE],
		collection: config[DB_COLLECTION],
		timeout:    timeout,
	}
	err = instance.Init()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &instance
}
func (db *MongoDB) Init() error {

	db.session.SetSafe(&mgo.Safe{})
	db.session.SetMode(mgo.Monotonic, true)
	db.session.SetSocketTimeout(1 * time.Hour)

	return nil
}

func (db *MongoDB) Database() string {
	return db.database
}

func (db *MongoDB) Collection() string {
	return db.collection
}

func (db *MongoDB) GetSession() *mgo.Session {
	return db.session
}

func (db *MongoDB) GetCollection() *mgo.Collection {
	return db.GetDatabase().C(db.collection)
}

func (db *MongoDB) GetDatabase() *mgo.Database {
	return db.GetSession().DB(db.database)
}

func LoadEnvFromFile(config interface{}, configPrefix, envPath string) (err error) {
	godotenv.Load(envPath)
	err = envconfig.Process(configPrefix, config)
	return
}
