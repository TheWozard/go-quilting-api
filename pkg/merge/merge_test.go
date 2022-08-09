package merge_test

import (
	"go-quilting-api/pkg/merge"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		desc      string
		traversal interface{}
		documents []interface{}
		output    interface{}
	}{
		{
			desc:      "",
			traversal: nil,
			documents: []interface{}{},
			output:    nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			require.Equal(t, tC.output, merge.Merge(tC.traversal, tC.documents))
		})
	}
}
