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
	s := Session.New()
	defer s.Close()
	c := s.DB("phrase_service").C("phrases")
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
	s := Session.New()
	defer s.Close()
	c := s.DB("phrase_service").C("phrases")
	result := Phrase{}
	err := c.Find(bson.M{}).One(&result)
	if err != nil {
		return "", err
	}
	return result.Text, nil
}

func GetAll() ([]Phrase, error) {
	s := Session.New()
	defer s.Close()
	c := s.DB("phrase_service").C("phrases")
	var result []Phrase
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
