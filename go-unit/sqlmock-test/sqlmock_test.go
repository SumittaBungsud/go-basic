package sqlmock

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	// "gorm.io/driver/sqlite"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Unit test with gorm-postgres to AddUser
func TestAddUser(t *testing.T) {
  db, mock, err := sqlmock.New()
  if err != nil {
    t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
  }
  defer db.Close()

  // Connect postgres and existing db
  gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
  if err != nil {
    t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
  }

  t.Run("Case #1 Successfully add user", func(t *testing.T) {
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
      WithArgs("john.doe@example.com").
      WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

      // Define your expectations for SQL operations
    mock.ExpectBegin()
    mock.ExpectQuery("^INSERT INTO \"users\" (.+)$").
      WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
    mock.ExpectCommit()

    err := AddUser(gormDB, "John Doe", "john.doe@example.com", 30)
    assert.NoError(t, err)

    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("Case #2 Fail to add user with existing email", func(t *testing.T) {
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
      WithArgs("john.doe@example.com").
      WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

    err := AddUser(gormDB, "John Doe", "john.doe@example.com", 30)
    assert.EqualError(t, err, "email already exists")

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

// Unit test with gorm-sqlite to AddUser
// func setupTestDB() *gorm.DB {
//   db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{}) // SQLite temporary memory
//   if err != nil {
//     panic(fmt.Sprintf("Failed to open database: %v", err))
//   }
//   db.AutoMigrate(&User{})
//   return db
// }

// func TestAddUser(t *testing.T) {
//   db := setupTestDB()
//   t.Run("Case #1 Successfully add user", func(t *testing.T) {
//     err := AddUser(db, "John Doe", "john.doe@example.com", 30)
//     assert.NoError(t, err)
//
//     var user User
//     db.First(&user, "email = ?", "john.doe@example.com")
//     assert.Equal(t, "John Doe", user.Fullname)
//   })
//
//   t.Run("Case #2 Fail to add user with existing email", func(t *testing.T) {
//     err := AddUser(db, "Jane Doe", "john.doe@example.com", 28)
//     assert.EqualError(t, err, "email already exists")
//   })
// }