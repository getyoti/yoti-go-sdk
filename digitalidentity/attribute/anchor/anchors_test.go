package anchor

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFilterAnchors_FilterSources(t *testing.T) {
	anchorSlice := []*Anchor{
		{subtype: "a", anchorType: TypeSource},
		{subtype: "b", anchorType: TypeVerifier},
		{subtype: "c", anchorType: TypeSource},
	}
	sources := filterAnchors(anchorSlice, TypeSource)
	assert.Equal(t, len(sources), 2)
	assert.Equal(t, sources[0].subtype, "a")
	assert.Equal(t, sources[1].subtype, "c")

}
