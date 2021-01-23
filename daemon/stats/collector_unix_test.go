// +build linux

package stats

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollector_cpuNanoSeconds(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		expectedVal uint64
	}{
		{
			name:        "large",
			filename:    "largeProcStat",
			expectedVal: 234436556630000000,
		},
	}

	collector := NewCollector()

	for _, tt := range tests {
		f, err := os.Open("./testdata/" + tt.filename)
		assert.NoError(t, err)

		t.Run(tt.name, func(t *testing.T) {
			actualVal, err := collector.cpuNanoSeconds(f)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, cpuTime, "expect")
		})
	}
}
