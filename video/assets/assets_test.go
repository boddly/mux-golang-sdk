// +build integration

package assets

import (
	"fmt"
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
		fmt.Printf("secret == %v\naccessToken == %v\n", secret, accessToken)
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

/*&assets.InputInfoData{
	Settings:assets.Input{
		URL:"https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
		OverlaySetting:
		assets.OverlaySetting{
			VerticalAlign:"",
			VerticalMargin:"",
			HorizontalAlign:"",
			HorizontalMargin:"",
			Width:"",
			Height:"",
			Opacity:""
		}
	},
	File:assets.InputDetails{
		ContainerFormat:"mov,mp4,m4a,3gp,3g2,mj2",
		Tracks:[]assets.InputTracks{
			assets.InputTracks{
				Type:"video",
				Duration:23.8238,
				Width:1920,
				Height:1080,
				FrameRate:29.97,
				Encoding:"h264",
				SampleRate:0,
				SampleSize:0
			},
			assets.InputTracks{
				Type:"audio",
				Duration:23.823792,
				Width:0,
				Height:0,
				FrameRate:0,
				Encoding:"aac",
				SampleRate:48000,
				SampleSize:0
			}
		}
	}
}*/
