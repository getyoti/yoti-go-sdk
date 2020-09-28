package docscan

import "fmt"

func createSessionPath() string {
	return "/sessions"
}

func getSessionPath(sessionID string) string {
	return fmt.Sprintf("/sessions/%s", sessionID)
}

func deleteSessionPath(sessionID string) string {
	return getSessionPath(sessionID)
}

func getMediaContentPath(sessionID string, mediaID string) string {
	return fmt.Sprintf("/sessions/%s/media/%s/content", sessionID, mediaID)
}

func deleteMediaPath(sessionID string, mediaID string) string {
	return getMediaContentPath(sessionID, mediaID)
}

func getSupportedDocumentsPath() string {
	return "/supported-documents"
}
