package assets

import (
	"errors"
	"net/http"
)

// Delete an asset using the ID of the asset to be deleted
func (c *Client) Delete(id string) error {
	req, err := http.NewRequest(http.MethodDelete, AssetURL+"/"+id, nil)
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(resp.Status)
	}

	return nil
}
