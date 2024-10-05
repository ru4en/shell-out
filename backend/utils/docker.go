package shellout

type DockerHost struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
}

func CreateDockerHost() {
	// Create a docker host
}

func ConnectDockerHost(id string) {
	// Connect to a docker host
}

func DisconnectDockerHost(id string) {
	// Disconnect from a docker host
}

func DeleteDockerHost(id string) {
	// Delete a docker host
}

func GetDockerHostList() {
	// Get all docker hosts
}
