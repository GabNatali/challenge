package entry

type EntryUsesCases interface {
	Create(dto AddEntryDto) (*EntryModel, error)
	Update(id uint, entry UpdateEntryDto) (*EntryModel, error)
	GetAllEntriesByUserId(id uint, limit int, offset int) ([]EntryModel, error)
	GetEntryById(id uint) (*EntryModel, error)
	Delete(id uint) (uint, error)
}

type entryUsesCases struct {
	repository EntryRepository
}

func NewEntryUsesCases(repository EntryRepository) EntryUsesCases {
	return &entryUsesCases{repository: repository}
}

func (e *entryUsesCases) Create(dto AddEntryDto) (*EntryModel, error) {
	entry := dto.MaptoEntryModel()

	return e.repository.CreateEntry(&entry)
}

func (e *entryUsesCases) Delete(id uint) (uint, error) {
	return e.repository.DeleteEntry(id)
}

func (e *entryUsesCases) GetAllEntriesByUserId(id uint, limit int, offset int) ([]EntryModel, error) {
	return e.repository.GetAllEntriesById(id, limit, offset)
}

func (e *entryUsesCases) GetEntryById(id uint) (*EntryModel, error) {
	return e.repository.GetEntryById(id)
}

func (e *entryUsesCases) Update(id uint, data UpdateEntryDto) (*EntryModel, error) {
	return e.repository.UpdateEntry(id, data)
}
