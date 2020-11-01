package business

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*pB := personaBusiness{
				personaRepo: tt.fields.personaRepo,
			}*/
			var db *mgo.Database
			tt.fields.personaRepo = repository.NewPersonaRepository(db)
			pB := NewPersonaBusiness(tt.fields.personaRepo)
			if got := pB.CreaPersona(tt.args.nom, tt.args.ape, tt.args.fechaNacimiento); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("personaBusiness.CreaPersona() = %v, want %v", got, tt.want)
			}
		})
	}
}
