package clean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {
	t.Run("Admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Username("Admin "))
	})
	t.Run(" Admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Username(" Admin "))
	})
	t.Run(" admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Username(" admin "))
	})
}

func TestEmail(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		assert.Equal(t, "hello@photoprism.app", Email("hello@photoprism.app"))
	})
	t.Run("Whitespace", func(t *testing.T) {
		assert.Equal(t, "hello@photoprism.app", Email(" hello@photoprism.app "))
	})
	t.Run("Invalid", func(t *testing.T) {
		assert.Equal(t, "", Email(" hello-photoprism "))
	})
}

func TestRole(t *testing.T) {
	t.Run("Admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Role("Admin "))
	})
	t.Run(" Admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Role(" Admin "))
	})
	t.Run(" admin ", func(t *testing.T) {
		assert.Equal(t, "admin", Role(" admin "))
	})
}

func TestFlags(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := ""
		assert.Equal(t, s, Attr(s))
	})
	t.Run("SlackScope", func(t *testing.T) {
		s := "admin.conversations.removeCustomRetention admin.usergroups:read"
		assert.Equal(t, s, Attr(s))
	})
	t.Run("Random", func(t *testing.T) {
		s := "  admin.conversations.removeCustomRetention   admin.usergroups:read  me:yes FOOt0-2U	6VU #$#%$ cm,Nu"
		cleaned := "6VU FOOt0-2U admin.conversations.removeCustomRetention admin.usergroups:read cmNu me"
		assert.Equal(t, cleaned, Attr(s))
	})
}

func TestPassword(t *testing.T) {
	t.Run("Alnum", func(t *testing.T) {
		assert.Equal(t, "fgdg5yw4y", Password("fgdg5yw4y "))
	})
	t.Run("Upper", func(t *testing.T) {
		assert.Equal(t, "AABDF24245vgfrg", Password(" AABDF24245vgfrg "))
	})
	t.Run("Special", func(t *testing.T) {
		assert.Equal(t, "!#$T#)$%I#J$I", Password("!#$T#)$%I#J$I"))
	})
}
