package adalo

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// person represents a record in the Persons collection in Adalo
type person struct {
	ID int `json:"id"`
	Name string `json:"Name"`
	Age int `json:"Age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// personInput represents the schema for inputting a person in the Adalo collection
type personInput struct {
	Name string `json:"Name"`
	Age int `json:"Age"`
}

// collection is the interface for the collection we will use in this test
var collection *Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ApiKey = os.Getenv("TEST_API_KEY")
	AppID = os.Getenv("TEST_APP_ID")
	collection = NewCollection(os.Getenv("TEST_COLLECTION_ID"))
}

func TestCollection_All(t *testing.T) {
	var res []interface{}
	err := collection.All(res)
	assert.Nil(t, err)
}

func TestCollection_Insert(t *testing.T) {
	var result person
	err := collection.Insert(&personInput{
		Name:    "John",
		Age:	21,
	}, &result)
	defer collection.Delete(result.ID)

	assert.Nil(t, err)
	log.Println(result)
	assert.Equal(t, "John", result.Name)
	assert.Equal(t, 21, result.Age)
}

func TestCollection_Get(t *testing.T) {
	var createdPerson person
	if err := collection.Insert(&personInput{
		Name:    "Jane",
		Age: 28,
	}, &createdPerson); err != nil {
		t.Skip()
	}
	defer collection.Delete(createdPerson.ID)

	var result person
	err := collection.Get(createdPerson.ID, &result)
	assert.Nil(t, err)
	assert.Equal(t, "Jane", result.Name)
	assert.Equal(t, 28, result.Age)
}

func TestCollection_Update(t *testing.T) {
	var createdPerson person
	if err := collection.Insert(&personInput{
		Age: 66,
		Name: "Chocolate rain",
	}, &createdPerson); err != nil {
		t.Skip()
	}
	defer collection.Delete(createdPerson.ID)

	var updatedPerson person
	err := collection.Update(createdPerson.ID, &personInput{
		Name: "Richard Johnson",
		Age: 89,
	}, &updatedPerson)

	assert.Nil(t, err)
	assert.Equal(t, createdPerson.ID, updatedPerson.ID)
	assert.Equal(t, 89, updatedPerson.Age)
	assert.Equal(t, "Richard Johnson", updatedPerson.Name)
}

func TestCollection_Delete(t *testing.T) {
	var createdPerson person
	if err := collection.Insert(&personInput{
		Name:    "Mr. Oldman",
		Age: 321,
	}, &createdPerson); err != nil {
		t.Skip()
	}

	err := collection.Delete(createdPerson.ID)
	assert.Nil(t, err)

	var result *person
	getErr := collection.Get(createdPerson.ID, result)
	assert.Error(t, getErr)
}
