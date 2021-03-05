package did

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURI(t *testing.T) {

	t.Run("for VC types", func(t *testing.T) {
		u, err := ParseURI("SomeType")

		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "SomeType", u.String())
	})

	t.Run("for URI types", func(t *testing.T) {
		u, err := ParseURI("https://example.com/context/v1")

		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "https://example.com/context/v1", u.String())
	})

	t.Run("malformed input", func(t *testing.T) {
		_, err := ParseURI(string([]byte{0}))

		assert.Error(t, err)
	})
}
