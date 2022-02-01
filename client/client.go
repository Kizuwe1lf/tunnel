package client

import (
	"encoding/base64"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

type Twitter struct {
	Client *anaconda.TwitterApi
}

func (t *Twitter) UploadMedia(b64media []byte) (string, error) {
	stringMedia := base64.StdEncoding.EncodeToString(b64media)
	media, err := t.Client.UploadMedia(stringMedia)

	if err != nil {
		return "", err
	}

	return media.MediaIDString, nil
}

func (t *Twitter) PostMediaTweet(mediaID string) error {
	v := url.Values{}
	v.Set("media_ids", mediaID)
	_, err := t.Client.PostTweet("", v)
	return err
}

func NewClient(token, tokenSecret, consumerKey, consumerSecret string) Twitter {
	client := anaconda.NewTwitterApiWithCredentials(
		token,
		tokenSecret,
		consumerKey,
		consumerSecret,
	)
	return Twitter{Client: client}
}
