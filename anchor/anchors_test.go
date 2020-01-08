package anchor

import (
	"testing"

	"gotest.tools/assert"
)

func TestFilterAnchors_FilterSources(t *testing.T) {
	anchorSlice := []*Anchor{
		{subtype: "a", anchorType: AnchorTypeSource},
		{subtype: "b", anchorType: AnchorTypeVerifier},
		{subtype: "c", anchorType: AnchorTypeSource},
	}
	sources := filterAnchors(anchorSlice, AnchorTypeSource)
	assert.Equal(t, len(sources), 2)
	assert.Equal(t, sources[0].subtype, "a")
	assert.Equal(t, sources[1].subtype, "c")

}
