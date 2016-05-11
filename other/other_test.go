package other

import (
	"testing"

	"github.com/daniel-garcia/withcontext/client"
)

func TestOther(t *testing.T) {

	c := NewContextToBarClientAdapter(client.NewClient())
	other := NewOther(c)
	t.Logf("built other: %s", other)
}
