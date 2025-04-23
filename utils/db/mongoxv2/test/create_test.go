package mongoxv2_test

import (
	"awesome-util/utils/db/mongoxv2"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	*mongoxv2.ObjectID

	Name  string
	Email string
}

func (u User) CollName() string {
	return "users"
}

// MockRepo is a mock implementation of the Repoer interface.
type MockRepo[T mongoxv2.IDModeler] struct {
	mock.Mock
}

func (m *MockRepo[T]) CreateOne(ctx context.Context, data T, opts ...*options.InsertOneOptions) error {
	args := m.Called(ctx, data, opts)
	return args.Error(0)
}

func (m *MockRepo[T]) CreateMany(
	ctx context.Context,
	data []T,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0).(*mongo.InsertManyResult), args.Error(1)
}

// Add similar methods for all other functions in the Repoer interface.

func TestCreateOne(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(MockRepo[User])
	testData := User{
		Name:  "Tuna",
		Email: "tuna@example.com",
	}

	mockRepo.On("CreateOne", ctx, testData, mock.Anything).Return(nil)

	err := mockRepo.CreateOne(ctx, testData)
	require.NoError(t, err)

	fmt.Println(testData.ObjectID)

	mockRepo.AssertExpectations(t)
}

func TestCreateMany(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(MockRepo[User]) // Replace `YourModelType` with the actual type.
	testData := []User{}            // Replace with actual data slice.

	mockResult := &mongo.InsertManyResult{InsertedIDs: []interface{}{"id1", "id2"}}
	mockRepo.On("CreateMany", ctx, testData, mock.Anything).Return(mockResult, nil)

	result, err := mockRepo.CreateMany(ctx, testData)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.InsertedIDs, 2)

	mockRepo.AssertExpectations(t)
}
