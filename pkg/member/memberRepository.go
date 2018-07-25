package member

import (
	"github.com/BrunoDM2943/church-members-api/pkg/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MemberRepository struct {
	col *mgo.Collection
}

func NewMemberRepository(session mgo.Session) *MemberRepository {
	return &MemberRepository{
		col: session.DB("disciples").C("membros"),
	}
}

func (repo *MemberRepository) FindAll() ([]*entity.Membro, error) {
	var result []*entity.Membro
	err := repo.col.Find(nil).Select(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
