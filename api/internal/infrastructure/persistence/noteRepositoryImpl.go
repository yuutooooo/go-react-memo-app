package persistence

import (
	"log"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	gormmodel "github.com/yourusername/go-react-memo-app/internal/domain/model/gormModel"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) repository.NoteRepository {
	return &NoteRepositoryImpl{
		db: db,
	}
}

func (r *NoteRepositoryImpl) CreateNote(note *model.Note) (*model.Note, error) {
	gormNote := convertToGormModel(note)
	if err := r.db.Create(gormNote).Error; err != nil {
		return nil, err
	}
	noteModel := convertToDomainModel(gormNote)
	return noteModel, nil
}

func (r *NoteRepositoryImpl) GetNoteByID(id string) (*model.Note, error) {
	var gormNote gormmodel.Note
	if err := r.db.First(&gormNote, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return convertToDomainModel(&gormNote), nil
}

func (r *NoteRepositoryImpl) UpdateNote(note *model.Note) (*model.Note, error) {
	gormNote := convertToGormModel(note)
	log.Println("gormNote")
	log.Println(gormNote.ID, gormNote.Title, gormNote.Content, gormNote.FolderID, gormNote.UserID, gormNote.CreatedAt, gormNote.UpdatedAt)
	if err := r.db.Save(gormNote).Error; err != nil {
		return nil, err
	}
	return convertToDomainModel(gormNote), nil
}

func (r *NoteRepositoryImpl) DeleteNote(id string) error {
	if err := r.db.Delete(&gormmodel.Note{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *NoteRepositoryImpl) GetNotesByFolderID(folderID string) ([]*model.Note, error) {
	var gormNotes []gormmodel.Note
	if err := r.db.Where("folder_id = ?", folderID).Find(&gormNotes).Error; err != nil {
		return nil, err
	}
	return convertToDomainModelMany(gormNotes), nil
}

func (r *NoteRepositoryImpl) GetNotesByUserID(userID string) ([]*model.Note, error) {
	var gormNotes []gormmodel.Note
	if err := r.db.Where("user_id = ?", userID).Find(&gormNotes).Error; err != nil {
		return nil, err
	}
	return convertToDomainModelMany(gormNotes), nil
}

func convertToGormModel(note *model.Note) *gormmodel.Note {
	return &gormmodel.Note{
		ID:        note.ID(),
		Title:     note.Title(),
		Content:   note.Content(),
		FolderID:  note.FolderID(),
		UserID:    note.UserID(),
		CreatedAt: note.CreatedAt(),
		UpdatedAt: note.UpdatedAt(),
	}
}

func convertToDomainModel(gormNote *gormmodel.Note) *model.Note {
	var noteModel model.Note
	noteModel.SetID(gormNote.ID)
	noteModel.SetTitle(gormNote.Title)
	noteModel.SetContent(gormNote.Content)
	noteModel.SetFolderID(gormNote.FolderID)
	noteModel.SetUserID(gormNote.UserID)
	noteModel.SetCreatedAt(gormNote.CreatedAt)
	noteModel.SetUpdatedAt(gormNote.UpdatedAt)
	return &noteModel
}

func convertToDomainModelMany(gormNotes []gormmodel.Note) []*model.Note {
	noteModels := make([]*model.Note, len(gormNotes))
	for i, gormNote := range gormNotes {
		noteModels[i] = convertToDomainModel(&gormNote)
	}
	return noteModels
}
