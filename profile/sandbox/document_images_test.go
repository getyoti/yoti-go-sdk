package sandbox

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/media"
	"gotest.tools/v3/assert"
)

const expectedBase64Content = "3q2+7w=="

func TestShouldAddJpegImage(t *testing.T) {
	documentImages := DocumentImages{}.WithJpegImage([]byte{0xDE, 0xAD, 0xBE, 0xEF})

	documentImage := documentImages.Images[0]
	assert.Equal(t, media.ImageTypeJPEG, documentImage.MIMEType())
	assert.Equal(
		t,
		documentImages.getValue(),
		fmt.Sprintf("data:image/jpeg;base64,%s", expectedBase64Content))
}

func TestShouldAddPngImage(t *testing.T) {
	documentImages := DocumentImages{}.WithPngImage([]byte{0xDE, 0xAD, 0xBE, 0xEF})

	documentImage := documentImages.Images[0]
	assert.Equal(t, media.ImageTypePNG, documentImage.MIMEType())
	assert.Equal(
		t,
		documentImages.getValue(),
		fmt.Sprintf("data:image/png;base64,%s", expectedBase64Content))
}

func TestShouldAddMultipleImage(t *testing.T) {
	documentImages := DocumentImages{}.
		WithPngImage([]byte{0xDE, 0xAD, 0xBE, 0xEF}).
		WithPngImage([]byte{0xDE, 0xAD, 0xBE, 0xEF}).
		WithJpegImage([]byte{0xDE, 0xAD, 0xBE, 0xEF})

	assert.Equal(t, 3, len(documentImages.Images))

	assert.Equal(
		t,
		documentImages.getValue(),
		fmt.Sprintf("data:image/png;base64,%[1]s&data:image/png;base64,%[1]s&data:image/jpeg;base64,%[1]s", expectedBase64Content))
}
