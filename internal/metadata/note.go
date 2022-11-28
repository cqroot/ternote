package metadata

import (
	"gorm.io/gorm/clause"

	"github.com/cqroot/ternote/pkg/types"
)

func Update(note *types.Note) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{"category", "title", "mod_time"},
		),
	}).Create(note)
}

func RemoveNote(id string) {
	db.Delete(&types.Note{}, id)
}

func NoteCategoryFromId(id string) string {
	var note types.Note
	db.First(&note, id)
	return note.Category
}

func Notes() []types.Note {
	var notes []types.Note
	db.Order("category, title").Find(&notes)
	return notes
}

func Categories() []string {
	var notes []types.Note
	var categories []string

	db.Distinct("category").Find(&notes)

	for _, note := range notes {
		categories = append(categories, note.Category)
	}
	return categories
}
