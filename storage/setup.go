package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Phrase struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Text string
}

var (
	Session *mgo.Session
)

func Setup(dbServer string) error {
	var err error
	Session, err = mgo.Dial(dbServer)
	if err != nil {
		return err
	}
	Session.SetMode(mgo.Monotonic, true)
	c, s := getCollection()
	defer s.Close()

	count, err := c.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		err = c.Insert(&Phrase{Text: "al que madruga dios lo ayuda from db"})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRandom() (string, error) {
	c, s := getCollection()
	defer s.Close()
	result := Phrase{}
	err := c.Find(bson.M{}).One(&result)
	if err != nil {
		return "", err
	}
	return result.Text, nil
}

func GetAll() ([]Phrase, error) {
	c, s := getCollection()
	defer s.Close()
	var result []Phrase
	err := c.Find(bson.M{}).All(&result)
	return result, err
}

func Add(text string) error {
	c, s := getCollection()
	defer s.Close()
	return c.Insert(&Phrase{Text: text})
}

func getCollection() (*mgo.Collection, *mgo.Session) {
	s := Session.New()
	return s.DB("phrase_service").C("phrases"), s
}
