package store

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/rs/zerolog/log"
)

type Penetration struct {
	ID         int
	Title      string `binding:"required,min=3,max=50"`
	Website    string `binding:"required,min=5,max=5000"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
	UserID     int `json:"-"`
}

func AddPenetration(user *User, penetration *Penetration) error {
	penetration.UserID = user.ID
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

func FetchPenetration(id int) (*Penetration, error) {
	penetration := new(Penetration)
	penetration.ID = id
	err := db.Model(penetration).WherePK().Select()
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
