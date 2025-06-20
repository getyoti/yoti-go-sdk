package digitalidentity

// SessionTrackedDevicesResponse represents the tracked devices for a session.
type SessionTrackedDevicesResponse struct {
	Devices []TrackedDevice `json:"devices"`
}

type TrackedDevice struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	// Add more fields as needed based on the Node.js SDK response
}
