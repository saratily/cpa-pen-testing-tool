package store

import (
	"time"

	"github.com/rs/zerolog/log"
)

type DefaultTool struct {
	ID         int
	Type       string `binding:"required,min=3,max=50"`
	Category   string `binding:"required,min=3,max=50"`
	Options    string
	Format     string
	Active     bool
	CanChange  bool
	Selected   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func FetchDefaultTool() ([]DefaultTool, error) {
	var defaultTools []DefaultTool
	err := db.Model(&defaultTools).Select()

	if err != nil {
		log.Error().Err(err).Msg("Error fetching tool")
		return nil, dbError(err)
	}
	return defaultTools, nil
}
