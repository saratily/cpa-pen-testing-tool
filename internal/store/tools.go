package store

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/rs/zerolog/log"
)

type Tool struct {
	ID         int
	Title      string `binding:"required,min=3,max=50"`
	Content    string `binding:"required,min=5,max=5000"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	UserID     int `json:"-"`
}

func AddTool(user *User, tool *Tool) error {
	tool.UserID = user.ID
	_, err := db.Model(tool).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("Error inserting new tool")
	}
	return dbError(err)
}

func FetchUserTools(user *User) error {
	err := db.Model(user).
		WherePK().
		Relation("Tools", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("id ASC"), nil
		}).
		Select()
	if err != nil {
		log.Error().Err(err).Msg("Error fetching user's tools")
	}
	return dbError(err)
}

func FetchTool(id int) (*Tool, error) {
	tool := new(Tool)
	tool.ID = id
	err := db.Model(tool).WherePK().Select()
	if err != nil {
		log.Error().Err(err).Msg("Error fetching tool")
		return nil, dbError(err)
	}
	return tool, nil
}

func UpdateTool(tool *Tool) error {
	_, err := db.Model(tool).WherePK().UpdateNotZero()
	if err != nil {
		log.Error().Err(err).Msg("Error updating tool")
	}
	return dbError(err)
}

func DeleteTool(tool *Tool) error {
	_, err := db.Model(tool).WherePK().Delete()
	if err != nil {
		log.Error().Err(err).Msg("Error deleting tool")
	}
	return dbError(err)
}
