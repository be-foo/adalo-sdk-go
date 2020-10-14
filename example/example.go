// Package example demonstrates the use of an extended Collection designed
// for a certain Adalo collection to achieve type safety.
package example

func init() {
	var winifred *Person

	err := Persons.Insert(&PersonInput{
		Name: "Winifred",
		Age:  34,
	}, winifred)

	if err != nil {
		panic(err)
	}
}
