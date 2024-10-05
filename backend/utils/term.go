package shellout

type Term struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Instance Instance `json:"instance"`
}

func GetTermList() {
	// Get all terms
}

func GetTerm(id string) {
	// Get a term
}

func CreateTerm() {
	// Create a term
}

func UpdateTerm(id string) {
	// Update a term
}

func DeleteTerm(id string) {
	// Delete a term
}
