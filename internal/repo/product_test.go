package repo

import (
	"database/sql"

	"github.com/knadh/koanf"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// func TestCreateProduct(t *testing.T) {
// 	conf := configs.NewConfig()
// 	type args struct {
// 	}

// 	tests := []struct {
// 		Name          string
// 		input         *types.Product
// 		args          args
// 		WantCall      bool
// 		MockDb        func(sqlmock.Sqlmock)
// 		expectedError error
// 	}{
// 		{},
// 	}
// }

func newMockDB(db *sql.DB, config *koanf.Koanf) *gorm.DB {

	gormHandler, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))
	if err != nil {
		return nil
	}

	return gormHandler
}
