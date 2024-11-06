package entry

import (
	"gorm.io/gorm"
)

type EntryRepository interface {
	CreateEntry(*EntryModel) (*EntryModel, error)
	UpdateEntry(id uint, entry UpdateEntryDto) (*EntryModel, error)
	GetEntryById(id uint) (*EntryModel, error)
	GetAllEntriesById(id uint, limit int, offset int) ([]EntryModel, error)
	DeleteEntry(id uint) (uint, error)
}

type repositoryEntry struct {
	db *gorm.DB
}

func NewEntryRepository(db *gorm.DB) EntryRepository {
	return &repositoryEntry{
		db: db,
	}
}

// CreateEntry implements Repository.
func (r *repositoryEntry) CreateEntry(entry *EntryModel) (*EntryModel, error) {

	tx := r.db.Create(&entry)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return entry, nil
}

func (r *repositoryEntry) UpdateEntry(id uint, entry UpdateEntryDto) (*EntryModel, error) {

	foundEntry, err := r.GetEntryById(id)

	if err != nil {
		return nil, err
	}

	foundEntry.Title = entry.Title
	foundEntry.Content = entry.Content
	foundEntry.StatusCode = 2
	foundEntry.Status = "UPDATED"

	tx := r.db.Model(&foundEntry).Updates(foundEntry)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return foundEntry, nil
}

func (r *repositoryEntry) GetEntryById(id uint) (*EntryModel, error) {

	var entry EntryModel

	tx := r.db.First(&entry, "id = ?", id)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &entry, nil
}

func (r *repositoryEntry) DeleteEntry(id uint) (uint, error) {

	var entry EntryModel
	tx := r.db.Where("id = ?", id).Delete(&entry)

	if err := tx.Error; err != nil {
		return 0, err
	}

	return entry.Id, nil

}

func (r *repositoryEntry) GetAllEntriesById(id uint, limit int, offset int) ([]EntryModel, error) {
	var entries []EntryModel

	tx := r.db.Model(&EntryModel{}).Limit(limit).Offset(offset).Find(&entries, "user_id = ?", id)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return entries, nil
}
