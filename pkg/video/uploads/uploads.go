package uploads

import "github.com/boddly/mux-golang-sdk/pkg/video/assets"

// UploadError is only set if an error occurred during asset creation.
type UploadError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// DirectUpload allows you to request a signed URL that you can PUT your content to directly instead of first needing to upload to something like S3 before creating a Mux asset.
// You can specify the settings for the upload, such as how long you want the returned URL to be valid, but also all of the settings for the new Mux asset to be created after the input is successfully uploaded.
type DirectUpload struct {
	ID               string               `json:"id"`
	Timeout          int                  `json:"timeout"`
	CORSOrigin       string               `json:"cors_origin"`
	Status           string               `json:"status"` // waiting, asset_created, errored, cancelled, timed_out
	NewAssetSettings assets.CreateOptions `json:"new_asset_settings"`
	AssetID          string               `json:"asset_id"`
	Error            UploadError          `json:"error"`
}
