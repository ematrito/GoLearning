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
	resident := Resident{
        Name: name,
        Age: age,
        Address: address,
    }

    return &resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    nameValidation := r.Name != ""
    // A baby could have 0 years
    ageValidation := r.Age > -1
    addressValidation := r.Address != nil && r.Address["street"] != ""
    
	return nameValidation && ageValidation && addressValidation
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var counter int

    for _, resident := range residents {
        if resident.HasRequiredInfo() {
            counter ++
        }
    }

    return counter
}
