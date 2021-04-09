package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twpayne/go-vfs/v2"
	"github.com/twpayne/go-vfs/v2/vfst"

	"github.com/twpayne/chezmoi/internal/chezmoitest"
)

func TestVerifyCmd(t *testing.T) {
	for _, tc := range []struct {
		name        string
		root        interface{}
		expectedErr error
	}{
		{
			name: "empty",
			root: map[string]interface{}{
				"/home/user": &vfst.Dir{
					Perm: 0o700,
				},
			},
		},
		{
			name: "file",
			root: map[string]interface{}{
				"/home/user": map[string]interface{}{
					".bashrc": &vfst.File{
						Contents: []byte("# contents of .bashrc\n"),
						Perm:     0o666 &^ chezmoitest.Umask,
					},
					".local/share/chezmoi/dot_bashrc": "# contents of .bashrc\n",
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			chezmoitest.WithTestFS(t, tc.root, func(fs vfs.FS) {
				assert.Equal(t, tc.expectedErr, newTestConfig(t, fs).execute([]string{"verify"}))
			})
		})
	}
}
