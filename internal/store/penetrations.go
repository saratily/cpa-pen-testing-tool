package store

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

type Penetration struct {
	ID         int
	Unique_ID  uuid.UUID `json:"uuid"`
	Title      string    `binding:"required,min=3,max=50"`
	Website    string    `binding:"required,min=5,max=5000"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
	UserID     int     `json:"-"`
	Tools      []*Tool `json:"-" pg:"fk:user_id,rel:has-many,on_delete:CASCADE"`
}

var _ pg.AfterSelectHook = (*User)(nil)

func (penetration *Penetration) AfterSelect(ctx context.Context) error {
	if penetration.Tools == nil {
		penetration.Tools = []*Tool{}
	}
	return nil
}

func AddPenetration(user *User, penetration *Penetration) error {
	penetration.UserID = user.ID
	penetration.Unique_ID = uuid.NewV4()
	_, err := db.Model(penetration).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("Error inserting new penetration")
	}
	return dbError(err)
}

func FetchUserPenetrations(user *User) error {
	err := db.Model(user).
		WherePK().
		Relation("Penetrations", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("id ASC"), nil
		}).
		Select()
	if err != nil {
		log.Error().Err(err).Msg("Error fetching user's penetrations")
	}
	return dbError(err)
}

func FetchPenetration(id string) (*Penetration, error) {
	penetration := new(Penetration)
	penetration.Unique_ID = uuid.Must(uuid.FromString(id))
	err := db.Model(penetration).
		Where("unique_id = ?", penetration.Unique_ID).
		Select()

	if err != nil {
		log.Error().Err(err).Msg("Error fetching penetration")
		return nil, dbError(err)
	}
	return penetration, nil
}

func UpdatePenetration(penetration *Penetration) error {
	_, err := db.Model(penetration).WherePK().UpdateNotZero()
	if err != nil {
		log.Error().Err(err).Msg("Error updating penetration")
	}
	return dbError(err)
}

func DeletePenetration(penetration *Penetration) error {
	_, err := db.Model(penetration).WherePK().Delete()
	if err != nil {
		log.Error().Err(err).Msg("Error deleting penetration")
	}
	return dbError(err)
}
