package store

import (
	"fmt"
	"time"

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
	CanChange     int
	Selected      int
	CreatedAt     time.Time
	ModifiedAt    time.Time
	PenetrationID int `json:"-"`
}

func AddTool(pen *Penetration, tool *Tool) error {
	tool.PenetrationID = pen.ID
	tool.Unique_ID = uuid.NewV4()
	fmt.Println(tool)
	_, err := db.Model(tool).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("Error inserting new tool")
	}
	return dbError(err)
}

func FetchPenTools(pen *Penetration, toolType string) ([]Tool, error) {

	var tools []Tool

	err := db.Model(&tools).
		Where("penetration_id = ?", pen.ID).
		Where("type = ?", toolType).
		Select()

	if err != nil {
		log.Error().Err(err).Msg("Error fetching tool")
		return nil, dbError(err)
	}
	return tools, nil
}

func FetchAllTools(pen *Penetration) ([]Tool, error) {

	var tools []Tool

	err := db.Model(&tools).
		Where("penetration_id = ?", pen.ID).
		Select()

	if err != nil {
		log.Error().Err(err).Msg("Error fetching tool")
		return nil, dbError(err)
	}
	return tools, nil
}

func UpdateTool(tool *Tool) error {
	_, err := db.Model(tool).WherePK().UpdateNotZero()
	if err != nil {
		log.Error().Err(err).Msg("Error updating tool")
	}
	return dbError(err)
}

/*
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


func DeleteTool(tool *Tool) error {
	_, err := db.Model(tool).WherePK().Delete()
	if err != nil {
		log.Error().Err(err).Msg("Error deleting tool")
	}
	return dbError(err)
}
*/
