package example

import (
	"github.com/be-foo/adalo-sdk-go"
	"os"
)

// PersonCollection is the Collection type but customized for the Person collection
type PersonCollection struct {
	collection *adalo.Collection
}

// Person represents a CRUD interface for the Person collection
var Persons *PersonCollection

// Person represents a record in the Persons collection in Adalo
type Person struct {
	// ID in Adalo collection
	ID int `json:"id"`

	// Name of person
	Name string `json:"Name"`

	// Age of person
	Age int `json:"Age"`

	// CreatedAt date of creation of this record in the collection
	CreatedAt string `json:"created_at"`

	// CreatedAt date of last update of this record in the collection
	UpdatedAt string `json:"updated_at"`
}

// PersonInput represents the schema for inputting a person in the Adalo collection
type PersonInput struct {
	// Name of person
	Name string `json:"Name"`

	// Age of person
	Age int `json:"Age"`
}

func init() {
	if os.Getenv("ADALO_PERSON_COLLECTION_ID") == "" {
		panic("adalo person collection id not set")
	}
	Persons = &PersonCollection{collection: adalo.NewCollection(os.Getenv("ADALO_Person_COLLECTION_ID"))}
}

// All is a type specific wrapper for Collection_All
func (c *PersonCollection) All(result []*Person) error {
	return c.collection.All(result)
}

// Insert is a type specific wrapper for Collection_Insert
func (c *PersonCollection) Insert(input *PersonInput, result *Person) error {
	return c.collection.Insert(input, result)
}

// Get is a type specific wrapper for Collection_Get
func (c *PersonCollection) Get(id int, result *Person) error {
	return c.collection.Get(id, result)
}

// Delete is a type specific wrapper for Collection_Delete
func (c *PersonCollection) Delete(id int) error {
	return c.collection.Delete(id)
}

// Update is a type specific wrapper for Collection_Update
func (c *PersonCollection) Update(id int, input *PersonInput, result *Person) error {
	return c.collection.Update(id, input, result)
}
