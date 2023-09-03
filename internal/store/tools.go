package store

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

type Tool struct {
	ID            int
	Unique_ID     uuid.UUID `json:"uuid"`
	Type          string    `binding:"required,min=3,max=50"`
	Category      string    `binding:"required,min=3,max=50"`
	Options       string
	Command       string
	Output        string
	CanChange     bool
	Selected      bool
	CreatedAt     time.Time
	ModifiedAt    time.Time
	PenetrationID int `json:"-"`
}

func AddTool(pen *Penetration, tool *Tool) error {
	tool.PenetrationID = pen.ID
	tool.Unique_ID = uuid.NewV4()
	_, err := db.Model(tool).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("Error inserting new tool")
	}
	return dbError(err)
}

func FetchUserTools(pen *Penetration) error {
	err := db.Model(pen).
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
