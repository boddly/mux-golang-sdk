// +build integration

package assets

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	secret      = os.Getenv("MUX_SECRET_KEY")
	accessToken = os.Getenv("MUX_ACCESS_TOKEN")
)

func TestAssetLifeCycle(t *testing.T) {
	c := Client{AccessToken: accessToken, SecretKey: secret}
	if secret == "" || accessToken == "" {
		t.FailNow()
	}

	cases := []struct {
		opts          CreateOptions
		expected      CreateResponse
		expectedError error
	}{
		{
			opts: CreateOptions{
				Input: []Input{
					Input{
						URL: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
					},
				},
			},
			expected: CreateResponse{
				Data: CreateData{
					Status:       "preparing",
					MP4Support:   "none",
					MasterAccess: "none",
				},
			},
			expectedError: nil,
		},
	}

	for _, v := range cases {
		// Create
		asset, err := c.Create(v.opts)
		require.NoError(t, err)
		assert.Equal(t, v.expected.Data.Status, asset.Data.Status)
		assert.Equal(t, v.expected.Data.MP4Support, asset.Data.MP4Support)
		assert.Equal(t, v.expected.Data.MasterAccess, asset.Data.MasterAccess)
		// List
		assets, err := c.List(25, 0)
		require.NoError(t, err)
		assert.Len(t, assets, 1)
		// InputInfo
		time.Sleep(3 * time.Second)
		ii, err := c.InputInfo(assets[0].ID)
		require.NoError(t, err)
		assert.Len(t, ii, 1)
		// Update MP4 support
		_, err = c.UpdateMP4(assets[0].ID, "none")
		require.NoError(t, err)
		// Update Master
		_, err = c.UpdateMaster(assets[0].ID, "none")
		require.NoError(t, err)
		// Delete
		for _, v := range assets {
			require.NoError(t, c.Delete(v.ID))
		}
	}
}
