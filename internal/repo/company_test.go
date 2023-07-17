package repo

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// func TestCreateProduct(t *testing.T) {
// 	//conf := configs.NewConfig()

// 	input := types.Product{
// 	//	ID:              1,
// 		Name:            "test_product",
// 		Price:           100,
// 		Description:     "this is a test product",
// 		Type:            "test_Type",
// 		SpecificCountry: "test_Country",
// 	}
// 	tests := []struct {
// 		Name string

// 		WantCall      bool
// 		MockDb        func(sqlmock.Sqlmock)
// 		expectedError error
// 	}{
// 		{
// 			Name: "error inserting",

// 			MockDb: func(sqlmock sqlmock.Sqlmock) {
// 				sqlmock.ExpectBegin()
// 				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "products" ("name","price","description","type","specific_country","id")
// 				VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
// 					WithArgs(input.Name, input.Price, input.Description, input.Type, input.SpecificCountry, input.ID).
// 					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(input.ID))
// 				sqlmock.ExpectRollback()
// 			},
// 			expectedError: errors.New("something went wrong"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {

// 			mockDB, mock, err := sqlmock.New()
// 			require.NoError(t, err)
// 			db := newMockDB(mockDB)
// 			rep := NewRepo(db)

// 			tt.MockDb(mock)
// 			_, err = rep.Create(&input)

// 			if err != nil {
// 				require.EqualError(t, err, tt.expectedError.Error())
// 			}

// 		})
// 	}
// }

func newMockDB(db *sql.DB) *gorm.DB {

	gormHandler, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	fmt.Println(err)
	if err != nil {
		return nil
	}

	return gormHandler
}
