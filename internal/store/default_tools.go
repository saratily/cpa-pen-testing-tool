package store

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type DefaultTool struct {
	ID         int
	Type       string `binding:"required,min=3,max=50"`
	Category   string `binding:"required,min=3,max=50"`
	Options    string
	Format     string
	Active     int
	CanChange  int
	Selected   int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func FetchDefaultTool() ([]DefaultTool, error) {
	var defaultTools []DefaultTool
	err := db.Model(&defaultTools).
		Where("active=1").
		Select()

	fmt.Println(defaultTools)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching tool")
		return nil, dbError(err)
	}
	return defaultTools, nil
}
