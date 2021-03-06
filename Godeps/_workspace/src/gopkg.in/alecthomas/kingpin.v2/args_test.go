package kingpin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgRemainder(t *testing.T) {
	app := New("test", "")
	v := app.Arg("test", "").Strings()
	args := []string{"hello", "world"}
	_, err := app.Parse(args)
	assert.NoError(t, err)
	assert.Equal(t, args, *v)
}

func TestArgRemainderErrorsWhenNotLast(t *testing.T) {
	a := newArgGroup()
	a.Arg("test", "").Strings()
	a.Arg("test2", "").String()
	assert.Error(t, a.init())
}

func TestArgMultipleRequired(t *testing.T) {
	app := New("test", "")
	app.Arg("a", "").Required().String()
	app.Arg("b", "").Required().String()

	_, err := app.Parse([]string{})
	assert.Error(t, err)
	_, err = app.Parse([]string{"A"})
	assert.Error(t, err)
	_, err = app.Parse([]string{"A", "B"})
	assert.NoError(t, err)
}

func TestInvalidArgsDefaultCanBeOverridden(t *testing.T) {
	app := New("test", "")
	app.Arg("a", "").Default("invalid").Bool()
	_, err := app.Parse([]string{})
	assert.Error(t, err)
}
