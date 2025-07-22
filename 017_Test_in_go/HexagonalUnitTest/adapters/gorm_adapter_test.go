package adapters

import (
	"errors"
	"mymodule/017_Test_in_go/HexagonalUnitTest/core"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormOrderRepoitory_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}),&gorm.Config{})
	if err != nil {
    t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	repo := NewGormOrderRepository(gormDB)

	t.Run("success",func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "orders"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		
		err := repo.Save(core.Order{Total: 100})
		assert.NoError(t,err)
		assert.NoError(t,mock.ExpectationsWereMet())
	})
	
	t.Run("Failure",func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "orders"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()
		
		err := repo.Save(core.Order{Total: 100})
		assert.Error(t,err)
		assert.NoError(t,mock.ExpectationsWereMet())


	})
		
}