package tokens

import (
	qt "github.com/frankban/quicktest"
	"github.com/google/uuid"
	"testing"
)

func TestToken(t *testing.T) {
	c := qt.New(t)

	c.Run("create a jwt token", func(_ *qt.C) {
		testID := uuid.New().String()

		token, err := New(testID)
		c.Check(err, qt.IsNil)

		notEmpty := token != ""
		c.Check(notEmpty, qt.IsTrue)
	})

	c.Run("parse a valid jwt token", func(_ *qt.C) {
		testID := uuid.New().String()

		token, _ := New(testID)

		payload, err := Parse(token)
		c.Check(err, qt.IsNil)
		c.Check(payload, qt.IsNotNil)
		c.Check(payload.Subject, qt.Equals, testID)
	})
}
