package arguments_test

import (
	"fmt"
	"os"
	"testing"

	"filter/arguments"

	"github.com/stretchr/testify/assert"
)

func setupArgs(args []string) []string {
	oldArgs := os.Args
	os.Args = args
	return oldArgs
}

func TestParseWhenKeyword(t *testing.T) {
	keyword := "testKeyword"
	oldArgs := setupArgs([]string{"cmd", fmt.Sprintf("-keyword=%s", keyword)})
	defer func() {
		os.Args = oldArgs
	}()

	actual := arguments.Parse()
	assert.Equal(t, keyword, actual, "wrong keyword")
}
