package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	type given struct {
		data     []Node
		tryShift bool
	}
	type expected struct {
		hasError     bool
		length       int
		queueUpdated *Queue
		nodeShifted  Node
	}
	var cases = []struct {
		name     string
		given    given
		expected expected
	}{
		{
			name: "Should return empty queue",
			expected: expected{
				length: 0,
			},
		},
		{
			name: "Should be able to push data into the queue",
			given: given{
				data: []Node{
					{
						Name:     "A",
						Distance: 1,
					},
					{
						Name:     "B",
						Distance: 2,
					},
				},
			},
			expected: expected{
				length: 2,
			},
		},
		{
			name: "Should be able to shift element from the queue",
			given: given{
				data: []Node{
					{
						Name:     "A",
						Distance: 1,
					},
					{
						Name:     "B",
						Distance: 2,
					},
					{
						Name:     "C",
						Distance: 3,
					},
				},
				tryShift: true,
			},
			expected: expected{
				length: 3,
				nodeShifted: Node{
					Name:     "A",
					Distance: 1,
				},
				queueUpdated: &Queue{
					{
						Name:     "B",
						Distance: 2,
					},
					{
						Name:     "C",
						Distance: 3,
					},
				},
			},
		},
	}
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue()
			for _, node := range tc.given.data {
				q.Push(node)
			}
			assert.Equal(t, tc.expected.length, q.Size())

			if len(tc.given.data) == 0 {
				assert.True(t, q.IsEmpty())
			}

			if tc.given.tryShift {
				node := q.Shift()
				assert.Equal(t, tc.expected.nodeShifted, node)
				assert.Equal(t, tc.expected.queueUpdated, q)
			}
		})
	}
}
