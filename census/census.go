// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	// panic("Please implement NewResident.")
	r := &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return r
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// panic("Please implement HasRequiredInfo.")
	if r.Name == "" {
		return false
	}
	street, ok := r.Address["street"]
	return ok && street != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	// panic("Please implement Delete.")
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	// panic("Please implement Count.")
	var count int
	for _, i := range residents {
		if i.HasRequiredInfo() {
			count += 1
		}
	}
	return count
}
