package adalo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// person represents a record in the Persons collection in Adalo
type person struct {
	ID        int    `json:"id"`
	Name      string `json:"Name"`
	Age       int    `json:"Age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// personInput represents the schema for inputting a person in the Adalo collection
type personInput struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

// collection is the interface for the collection we will use in this test
var collection *Collection

// invalidID is just a random high number that we use to test for non-existent IDs.
// We must ensure that a record with this ID will never exist in the Adalo app this project is tested with.
const invalidID = 8834

func init() {
	if os.Getenv("TEST_COLLECTION_ID") == "" {
		panic("environment variable TEST_COLLECTION_ID is not set")
	}
	collection = NewCollection(os.Getenv("TEST_COLLECTION_ID"))
}

func TestCollection_All(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		setup()
		var res []interface{}
		err := collection.All(res)
		assert.Nil(t, err)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)
		var res []interface{}
		err := collection.All(res)
		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("app mismatch", func(t *testing.T) {
		setup(invalidApp)
		var res []interface{}
		err := collection.All(res)
		assert.Equal(t, ErrorAppMismatch, err)
	})
}

func TestCollection_Insert(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		setup()

		var result person
		err := collection.Insert(&personInput{
			Name: "John",
			Age:  21,
		}, &result)
		defer collection.Delete(result.ID)

		assert.Nil(t, err)
		assert.Equal(t, "John", result.Name)
		assert.Equal(t, 21, result.Age)
	})

	t.Run("with invalid input", func(t *testing.T) {
		setup()
		err := collection.Insert("this cannot be marshaled as a json", nil)
		assert.Error(t, err)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)
		err := collection.Insert(&personInput{
			Name: "John",
			Age:  21,
		}, nil)
		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("app mismatch", func(t *testing.T) {
		setup(unauthorized)
		err := collection.Insert(&personInput{
			Name: "John",
			Age:  21,
		}, nil)
		assert.Equal(t, ErrorUnauthorized, err)
	})
}

func TestCollection_Get(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		setup()

		var createdPerson person
		if err := collection.Insert(&personInput{
			Name: "Jane",
			Age:  28,
		}, &createdPerson); err != nil {
			t.Skip()
		}
		defer collection.Delete(createdPerson.ID)

		var result person
		err := collection.Get(createdPerson.ID, &result)
		assert.Nil(t, err)
		assert.Equal(t, "Jane", result.Name)
		assert.Equal(t, 28, result.Age)
	})

	t.Run("with id that does not exist", func(t *testing.T) {
		setup()
		err := collection.Get(invalidID, nil)
		assert.Error(t, err)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)
		err := collection.Get(1, nil)
		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("app mismatch", func(t *testing.T) {
		setup(invalidApp)
		err := collection.Get(1, nil)
		assert.Equal(t, ErrorAppMismatch, err)
	})
}

func TestCollection_Update(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		setup()

		var createdPerson person
		if err := collection.Insert(&personInput{
			Age:  66,
			Name: "Chocolate rain",
		}, &createdPerson); err != nil {
			t.Skip()
		}
		defer collection.Delete(createdPerson.ID)

		var updatedPerson person
		err := collection.Update(createdPerson.ID, &personInput{
			Name: "Richard Johnson",
			Age:  89,
		}, &updatedPerson)

		assert.Nil(t, err)
		assert.Equal(t, createdPerson.ID, updatedPerson.ID)
		assert.Equal(t, 89, updatedPerson.Age)
		assert.Equal(t, "Richard Johnson", updatedPerson.Name)
	})

	t.Run("with id that does not exist", func(t *testing.T) {
		setup()
		err := collection.Update(invalidID, &personInput{
			Name: "Richard Johnson",
			Age:  89,
		}, nil)
		assert.Error(t, err)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)
		err := collection.Update(1, &personInput{
			Name: "Richard Johnson",
			Age:  89,
		}, nil)
		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("app mismatch", func(t *testing.T) {
		setup(invalidApp)
		err := collection.Update(1, &personInput{
			Name: "Richard Johnson",
			Age:  89,
		}, nil)
		assert.Equal(t, ErrorAppMismatch, err)
	})
}

func TestCollection_Delete(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		setup()

		var createdPerson person
		if err := collection.Insert(&personInput{
			Name: "Mr. Oldman",
			Age:  321,
		}, &createdPerson); err != nil {
			t.Skip()
		}

		err := collection.Delete(createdPerson.ID)
		assert.Nil(t, err)

		// ensure record is really deleted
		var result *person
		getErr := collection.Get(createdPerson.ID, result)
		assert.Error(t, getErr)
	})

	t.Run("with id that does not exist", func(t *testing.T) {
		setup()
		err := collection.Delete(invalidID)
		assert.Equal(t, ErrorResourceNotFound, err)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)
		err := collection.Delete(1)
		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("app mismatch", func(t *testing.T) {
		setup(invalidApp)
		err := collection.Delete(1)
		assert.Equal(t, ErrorAppMismatch, err)
	})
}
