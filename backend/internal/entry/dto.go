package entry

type AddEntryDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint   `json:"user_id"`
}

func (e AddEntryDto) MaptoEntryModel() EntryModel {
	return NewEntry(e.Title, e.Content, e.UserId)
}

type UpdateEntryDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
