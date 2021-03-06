package shell

import (
	"context"
	"encoding/json"
)

// Publish updates a mutable name to point to a given value
func (s *Shell) Publish(value string, lifetime string) error {
	req := NewRequest(context.Background(), s.url, "name/publish")
	req.Opts["arg"] = value
	req.Opts["lifetime"] = lifetime

	resp, err := req.Send(s.httpcli)
	if err != nil {
		return err
	}
	defer resp.Close()

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

// Resolve gets resolves the string provided to an /ipfs/[hash]. If asked to
// resolve an empty string, resolve instead resolves the node's own /ipns value.
func (s *Shell) Resolve(id string) (string, error) {
	var resp *Response
	var err error
	if id != "" {
		resp, err = s.newRequest(context.Background(), "name/resolve", id).Send(s.httpcli)
	} else {
		resp, err = s.newRequest(context.Background(), "name/resolve").Send(s.httpcli)
	}
	if err != nil {
		return "", err
	}
	defer resp.Close()

	if resp.Error != nil {
		return "", resp.Error
	}

	var out struct{ Path string }
	err = json.NewDecoder(resp.Output).Decode(&out)
	if err != nil {
		return "", err
	}

	return out.Path, nil
}
