package shell

import (
        "testing"

        "github.com/cheekybits/is"
)

func TestPublishAndResolve(t *testing.T) {
	is := is.New(t)
        s := NewShell(shellUrl)

	err := s.Publish(examplesHash, "24h")
	is.Nil(err)

	me_out, err := s.ID()
	is.Nil(err)

	id := me_out.ID

	hash, err := s.Resolve(id)
	is.Nil(err)

	is.Equal(hash, "/ipfs/" + examplesHash)
}
