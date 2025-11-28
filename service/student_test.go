package service

import (
	"session-9/model"
	"session-9/repository"
	"session-9/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func newTestService(initial []model.Student) (*StudentService, *repository.MockStudentRepository) {
// 	repo := &repository.MockStudentRepository{Students: initial}
// 	svc := NewStudentService(repo)
// 	return svc, repo
// }

// func TestStudentService_Create(t *testing.T) {
// 	svc, repo := newTestService([]model.Student{})

// 	created, err := svc.Create(model.Student{
// 		Name: "Budi",
// 		Age:  20,
// 	})
// 	if err != nil {
// 		t.Fatalf("Create returned error: %v", err)
// 	}

// 	if created.ID != 1 {
// 		t.Errorf("expected ID 1, got %d", created.ID)
// 	}
// 	if created.Name != "Budi" {
// 		t.Errorf("expected Name Budi, got %s", created.Name)
// 	}

// 	if len(repo.Students) != 1 {
// 		t.Fatalf("expected repo to have 1 student, got %d", len(repo.Students))
// 	}
// }

// func TestStudentService_GetByID_Found(t *testing.T) {
// 	initial := []model.Student{
// 		{ID: 1, Name: "Andi", Age: 21},
// 		{ID: 2, Name: "Siti", Age: 22},
// 	}
// 	svc, _ := newTestService(initial)

// 	st, err := svc.GetByID(2)
// 	if err != nil {
// 		t.Fatalf("GetByID returned error: %v", err)
// 	}

// 	if st.Name != "Siti" {
// 		t.Errorf("expected Name Siti, got %s", st.Name)
// 	}
// }

// func TestStudentService_GetByID_NotFound(t *testing.T) {
// 	initial := []model.Student{
// 		{ID: 1, Name: "Andi", Age: 21},
// 		{ID: 2, Name: "Siti", Age: 22},
// 	}
// 	svc, _ := newTestService(initial)

// 	_, err := svc.GetByID(999)
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err != utils.ErrNotFound {
// 		t.Fatalf("expected ErrNotFound, got %v", err)
// 	}
// }

// func TestStudentService_GetByID_fileError(t *testing.T) {
// 	svc, _ := newTestService([]model.Student{})

// 	_, err := svc.GetByID(1)
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err != utils.ErrFile {
// 		t.Fatalf("expected error file, got %v", err)
// 	}
// }

func newTestService() (*StudentService, *repository.MockStudentRepository) {
	mokeRepo := new(repository.MockStudentRepository)
	service := NewStudentService(mokeRepo)
	return service, mokeRepo
}

// func TestStudent_Create(t *testing.T) {
// 	service, repo := newTestService([]model.Student{})

// 	created, err := service.Create(model.Student{
// 		Name: "Rudi",
// 		Age:  20,
// 	})

// 	if err != nil {
// 		t.Fatalf("Created returned error: %v", err)
// 	}

// 	if created.ID != 1 {
// 		t.Errorf("expected ID 1, got %d", created.ID)
// 	}

// 	if created.Name != "Rudi" {
// 		t.Errorf("expected Name Budi, got %s", created.Name)
// 	}

// 	if len(repo.Students) != 1 {
// 		t.Fatalf("expected repo to have 1 student, got %d", len(repo.Students))
// 	}
// }

func TestStudentService_GetByID_Found(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, nil).Once()

	st, err := svc.GetByID(2)
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}

	if st.Name != "Siti" {
		t.Errorf("expected Name Siti, got %s", st.Name)
	}
}

func TestStudentService_GetByID_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, utils.ErrFile).Once()

	_, err := svc.GetByID(999)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)

	repo.AssertExpectations(t)
}

// func TestStudentService_GetByID_NotFound(t *testing.T) {
// 	initial := []model.Student{
// 		{ID: 1, Name: "Andi", Age: 21},
// 		{ID: 2, Name: "Siti", Age: 22},
// 	}
// 	svc, repo := newTestService()
// 	repo.On("GetAll").Return(initial, utils.ErrFile).Once()

// 	_, err := svc.GetByID(999)
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err != utils.ErrNotFound {
// 		t.Fatalf("expected ErrNotFound, got %v", err)
// 	}
// }

func TestStudentService_GetByID_FileError(t *testing.T) {
	svc, repo := newTestService()
	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.GetByID(1)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != utils.ErrFile {
		t.Fatalf("expected error file, got %v", err)
	}
}

/*
    MY TASK START FROM HERE
*/

func TestStudentService_Create_Success(t *testing.T) {

	svc, repo := newTestService()

	Data := []model.Student{
		{ID: 1, Name: "jane Margolis", Age: 26},
	}

	// Input data for creating a new student (without ID)
	Input := model.Student{
		Name: "Azwin",
		Age:  15,
	}

	repo.On("GetAll").Return(Data, nil)
	repo.On("SaveAll", mock.MatchedBy(func(students []model.Student) bool {
        // Validasi bahwa slice berisi 2 student
        if len(students) != 2 {
            return false
        }
        // Validasi student baru ada di index 1
        return students[1].ID == 2 && students[1].Name == "Azwin"
    })).Return(nil)

	created, err := svc.Create(Input)

	assert.Nil(t, err)
	assert.Equal(t, 2, created.ID)
	assert.Equal(t, "Azwin", created.Name)
	assert.Equal(t, 21, created.Age)
}

func TestStudentService_Update_Success(t *testing.T) {
    svc, repo := newTestService()

    existingData := []model.Student{
        {ID: 1, Name: "Jane Margolis", Age: 21},
        {ID: 2, Name: "Azwin", Age: 22},
    }

    updatedInput := model.Student{
        Name: "Azwin Updated",
        Age:  25,
    }

    repo.On("GetAll").Return(existingData, nil)
    repo.On("SaveAll", mock.MatchedBy(func(students []model.Student) bool {
        if len(students) != 2 {
            return false
        }
        // Validasi student dengan ID 2 sudah diupdate
        return students[1].ID == 2 && students[1].Name == "Azwin Updated" && students[1].Age == 25
    })).Return(nil)

    updated, err := svc.Update(2, updatedInput)

    assert.Nil(t, err)
    assert.Equal(t, 2, updated.ID)
    assert.Equal(t, "Azwin Updated", updated.Name)
    assert.Equal(t, 25, updated.Age)

    repo.AssertExpectations(t)
}

func TestStudentService_Update_NotFound(t *testing.T) {
    svc, repo := newTestService()

    existingData := []model.Student{
        {ID: 1, Name: "Jane Margolis", Age: 21},
        {ID: 2, Name: "Azwin", Age: 22},
    }

    updatedInput := model.Student{
        Name: "Unknown",
        Age:  30,
    }

    repo.On("GetAll").Return(existingData, nil)

    _, err := svc.Update(999, updatedInput)

    assert.Error(t, err)
    assert.Equal(t, utils.ErrNotFound, err)

    repo.AssertExpectations(t)
}

func TestStudentService_Update_GetAllError(t *testing.T) {
    svc, repo := newTestService()

    updatedInput := model.Student{
        Name: "Test",
        Age:  20,
    }

    repo.On("GetAll").Return([]model.Student{}, utils.ErrFile)

    _, err := svc.Update(1, updatedInput)

    assert.Error(t, err)
    assert.Equal(t, utils.ErrFile, err)

    repo.AssertExpectations(t)
}

func TestStudentService_Update_SaveAllError(t *testing.T) {
    svc, repo := newTestService()

    existingData := []model.Student{
        {ID: 1, Name: "Azwin", Age: 21},
    }

    updatedInput := model.Student{
        Name: "Azwin Updated",
        Age:  22,
    }

    repo.On("GetAll").Return(existingData, nil)
    repo.On("SaveAll", mock.Anything).Return(utils.ErrFile)

    _, err := svc.Update(1, updatedInput)

    assert.Error(t, err)
    assert.Equal(t, utils.ErrFile, err)

    repo.AssertExpectations(t)
}