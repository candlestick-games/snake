package space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectI_Inside(t *testing.T) {
	testcases := map[string]struct {
		r1 RectI
		r2 RectI
		i  bool
	}{
		"equal": {
			r1: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			r2: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			i: true,
		},
		"fully-outside": {
			r1: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			r2: RectI{
				Pos:  Vec2I{X: 10, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			i: false,
		},
		"touching": {
			r1: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			r2: RectI{
				Pos:  Vec2I{X: 3, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			i: false,
		},
		"intercept": {
			r1: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			r2: RectI{
				Pos:  Vec2I{X: 2, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			i: false,
		},
		"inside": {
			r1: RectI{
				Pos:  Vec2I{X: 1, Y: 1},
				Size: Vec2I{X: 2, Y: 2},
			},
			r2: RectI{
				Pos:  Vec2I{X: 2, Y: 2},
				Size: Vec2I{X: 1, Y: 1},
			},
			i: true,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.i, tc.r2.Inside(tc.r1))
		})
	}
}
