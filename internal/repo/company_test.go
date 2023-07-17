package repo

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/Safwanseban/voixme-project/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestFindUsingCountry(t *testing.T) {
	//conf := configs.NewConfig()
	columns := []string{"o_offer_id", "o_client_id", "o_country", "o_image"}

	input := types.OfferCompany{
		OfferID:  1,
		ClientID: 1,
		Country:  "US",
		Image:    "test_image",
	}
	tests := []struct {
		Name string

		WantCall      bool
		MockDb        func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			Name: "no db handler",

			MockDb: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectBegin()
				query := fmt.Sprintf(regexp.QuoteMeta("^SELECT (.+) FROM `" + "offer_companies" + "`(.*)"))
				sqlmock.ExpectQuery(query).WithArgs(1).WillReturnError(errors.New("something wrong"))
				sqlmock.ExpectRollback()
			},
			WantCall:      false,
			expectedError: errors.New("something went wrong"),
		},
		{

			Name: "success",
			MockDb: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectBegin()
				sqlmock.ExpectQuery("^SELECT (.+) FROM `" + "offer_companies" + "`(.*)").WithArgs(1).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, "US", "test_image"))

				sqlmock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			db := utils.NewMockDB(mockDB)
			repo := NewRepo(db)

			tt.MockDb(mock)
			if tt.WantCall {
				_, err = repo.FindUsingCountry(&input)
			}
			if err != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			}

		})
	}
}
