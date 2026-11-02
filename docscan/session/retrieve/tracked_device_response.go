package retrieve

import "time"

// TrackedDeviceResponse represents a single tracked-device event returned by the
// GET /sessions/{sessionId}/tracked-devices endpoint.
type TrackedDeviceResponse struct {
	// Event is the type of device event (e.g. "SESSION_CREATED", "RESOURCE_CREATED").
	Event string `json:"event"`
	// ResourceID is only present for RESOURCE_CREATED events.
	ResourceID *string `json:"resource_id,omitempty"`
	// Created is the timestamp when this event occurred.
	Created time.Time `json:"created"`
	// Device contains metadata about the device.
	Device *TrackedDeviceDeviceResponse `json:"device"`
}

// TrackedDeviceDeviceResponse contains metadata about the device associated with a tracked-device event.
type TrackedDeviceDeviceResponse struct {
	IPAddress        *string `json:"ip_address,omitempty"`
	IPISOCountryCode *string `json:"ip_iso_country_code,omitempty"`
	ManufactureName  *string `json:"manufacture_name,omitempty"`
	ModelName        *string `json:"model_name,omitempty"`
	OSName           *string `json:"os_name,omitempty"`
	OSVersion        *string `json:"os_version,omitempty"`
	BrowserName      *string `json:"browser_name,omitempty"`
	BrowserVersion   *string `json:"browser_version,omitempty"`
	Locale           *string `json:"locale,omitempty"`
	ClientVersion    *string `json:"client_version,omitempty"`
}
