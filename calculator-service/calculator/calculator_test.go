package calculator
 
import (
	"fmt"
	"context"
    "testing"
    "time"

	"github.com/hashicorp/go-hclog"
)

// utils
func parseDate(date string) time.Time {
	parsedDate, _ := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", date))
	return parsedDate
}

// User repository mock
type userRepoMock struct {}

func newUserRepoMock() UserRespository {
	return &userRepoMock{}
}

var getUserMock func(id string) (*User, error)
func (repo *userRepoMock) GetUser(id string) (*User, error) {
	return getUserMock(id)
}

// Products repository mock
type productRepoMock struct {}

func newProductRepoMock() ProductRepository {
	return &productRepoMock{}
}

var getProductMock func(id string) (*Product, error)
func (repo *productRepoMock) GetProduct(id string) (*Product, error) {
	return getProductMock(id)
}

func TestCalculator(t *testing.T) {
	testCases := []struct {
		name 		  string
        userBirthdate time.Time
		today		  time.Time
		expectedPct   float64
		expectedValue int32
	} { 
		{ 
			name: "birthday discount",
            userBirthdate: parseDate("1992-01-04"),
			today: parseDate("2020-01-04"),
			expectedPct: 0.05,
            expectedValue: 5,
		},
		{ 
			name: "blackfriday discount",
            userBirthdate: parseDate("1992-01-04"),
			today: parseDate("2020-11-25"),
			expectedPct: 0.1,
            expectedValue: 10,
		},
		{ 
			name: "max percetange discount",
            userBirthdate: parseDate("1992-11-25"),
			today: parseDate("2020-11-25"),
			expectedPct: 0.1,
            expectedValue: 10,
		},
		{
			name: "no discount at all",
            userBirthdate: parseDate("1992-01-04"),
			today: parseDate("2020-10-17"),
			expectedPct: 0.0,
            expectedValue: 0,
		},
    }

	for _, tc := range testCases {
		testCase := tc
		ctx := context.Background()
		logger := hclog.NewNullLogger()

		// mock product repository to inject
		mockedProductRepository := newProductRepoMock()
		getProductMock = func(id string) (*Product, error) {
			product := &Product{
				Title: "Some Product",
				Description: "Testing",
				PriceInCents: 100,
			}
			return product, nil
		}

		// mock user repository to inject
		mockedUserRepository := newUserRepoMock()
		getUserMock = func(id string) (*User, error) {
			user := &User{
				FirstName: "Jane",
				LastName: "Doe",
				DateOfBirth: testCase.userBirthdate,
			}

			return user, nil
		}

		// mock now
		now = func() time.Time { return testCase.today }

		t.Run(testCase.name, func(t *testing.T) {
			calculator := NewCalculator(logger, mockedUserRepository, mockedProductRepository)
			request := &ProductDiscountRequest{ProductId: "test", UserId: "test"}
			response, _ := calculator.GetProductDiscount(ctx, request)

			if response.GetPct() != testCase.expectedPct {
				t.Fatalf("expected pct %v, got %v", testCase.expectedPct, response.GetPct())
			}

			if response.GetValueInCents() != testCase.expectedValue {
				t.Fatalf("expected value %v, got %v", testCase.expectedValue, response.GetValueInCents())
			}

		},)
	}
}
