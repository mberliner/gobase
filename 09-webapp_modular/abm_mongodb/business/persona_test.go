package business

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	gomock "github.com/golang/mock/gomock"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/repository"
)

func Test_personaBusiness_CreaPersona(t *testing.T) {
	type fields struct {
		personaRepo repository.PersonaRepository
	}
	type args struct {
		nom             string
		ape             string
		fechaNacimiento string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.Personas
	}{
		{
			name:   "Sin BD",
			fields: fields{},
			args: args{nom: "J",
				ape:             "J",
				fechaNacimiento: "2010-04-02"},
			want: model.Personas{
				Error:   nil,
				Mensaje: "",
			},
		},
	}
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	//m := NewMockSession(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*pB := personaBusiness{
				personaRepo: tt.fields.personaRepo,
			}*/

			var db = &mgo.Database{}
			fDB := MongoDatabase{
				Database: db,
			}
			fmt.Println(fDB)

			tt.fields.personaRepo = repository.NewPersonaRepository(db)
			pB := NewPersonaBusiness(tt.fields.personaRepo)
			if got := pB.CreaPersona(tt.args.nom, tt.args.ape, tt.args.fechaNacimiento); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("personaBusiness.CreaPersona() = %v, want %v", got, tt.want)
			}
		})
	}
}

type Session interface {
	DB(name string) DataLayer
	Close()
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) DataLayer {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// MockSession satisfies Session and act as a mock of *mgo.session.
type MockSession struct{}

// NewMockSession mock NewSession.
func NewMockSession() Session {
	return MockSession{}
}

// Close mocks mgo.Session.Close().
func (fs MockSession) Close() {}

// DB mocks mgo.Session.DB().
func (fs MockSession) DB(name string) DataLayer {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

// NewSession returns a new Mongo Session.
func NewSession() Session {
	mgoSession, err := mgo.Dial("<MONGO_URI>")
	if err != nil {
		panic(err)
	}
	return MongoSession{mgoSession}
}

// OR
// NewMockSession mock NewSession.
//func NewMockSession() Session {
//return MockSession{}
//}

type MongoCollection struct {
	*mgo.Collection
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) *mgo.Query
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// MockDatabase satisfies DataLayer and act as a mock.
type MockDatabase struct{}

// MockCollection satisfies Collection and act as a mock.
type MockCollection struct{}

// Find mock.
func (fc MockCollection) Find(query interface{}) *mgo.Query {
	return nil
}

// Count mock.
func (fc MockCollection) Count() (n int, err error) {
	return 10, nil
}

// Insert mock.
func (fc MockCollection) Insert(docs ...interface{}) error {
	return nil
}

// Remove mock.
func (fc MockCollection) Remove(selector interface{}) error {
	return nil
}

// Update mock.
func (fc MockCollection) Update(selector interface{}, update interface{}) error {
	return nil
}

// C mocks mgo.Database(name).Collection(name).
func (db MockDatabase) C(name string) Collection {
	return MockCollection{}
}
