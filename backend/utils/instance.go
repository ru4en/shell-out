package shellout

// Instance struct
type Instance struct {
	ID           string     `json:"id"`
	BaseTemplate Template   `json:"baseTemplate"`
	IsTemporary  bool       `json:"isTemporary"`
	Status       string     `json:"status"`
	DockerHost   DockerHost `json:"dockerHost"`
}

func GetInstanceList() {
	// Get all instances
}

func GetInstance(id string) {
	// Get an instance
}

func CreateInstance() {
	// Create an instance
}

func UpdateInstance(id string) {
	// Update an instance
}

func DeleteInstance(id string) {
	// Delete an instance
}

func StartInstance(id string) {
	// Start an instance
}

func StopInstance(id string) {
	// Stop an instance
}

func RestartInstance(id string) {
	// Restart an instance
}
