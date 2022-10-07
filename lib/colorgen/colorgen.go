package colorgen

import (
	"math"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

// 1/golden ratio (i.e 1/Φ)
const goldenRatioConjugate = 0.618033988749895

const hsvSaturation = 0.6273
const hsvValue = 0.7725

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generator can generate distinct colors for any number of new keys.
//
// Inspired by: [https://martin.ankerl.com/2009/12/09/how-to-create-random-colors-programmatically/]
type Generator struct {
	hue       float64
	knownKeys map[string]float64
}

// NewGenerator returns a new initialized generator, with a randomly selected
// seed value.
func NewGenerator() *Generator {
	return &Generator{
		hue:       math.Mod(math.Abs(random.Float64()), 1),
		knownKeys: make(map[string]float64),
	}
}

func (gen *Generator) hueForKey(key string) float64 {
	if v, ok := gen.knownKeys[key]; ok {
		return v
	}
	gen.hue = math.Mod(gen.hue+goldenRatioConjugate, 1)
	gen.knownKeys[key] = gen.hue
	return gen.hue
}

// GenerateRGB returns R, G, and B values, spanning 0-255, for a given key.
//
// - If a key is never seen before, then a new color is generated.
// - If a key is reused, then the previously generated color for that key is used.
func (gen *Generator) GenerateRGB(key string) (r, g, b uint8) {
	hue := gen.hueForKey(key)
	color := colorful.Hsv(360*hue, hsvSaturation, hsvValue)
	return color.RGB255()
}
