package keybase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

type Client struct {
	client        *resty.Client
	picturesCache *cache.Cache
}

func NewClient() *Client {
	return &Client{
		client:        resty.New(),
		picturesCache: cache.New(time.Hour, time.Hour),
	}
}

type lookupPayload struct {
	Them []struct {
		Pictures *struct {
			Primary *struct {
				URL *string
			}
		}
	}
}

func (c *Client) LookupPicture(ctx context.Context, kid string) (*url.URL, error) {
	if picture, ok := c.picturesCache.Get(kid); ok {
		return picture.(*url.URL), nil
	}

	lookup, err := c.lookup(ctx, kid)
	if err != nil {
		return nil, err
	}

	var picture *url.URL
	if len(lookup.Them) > 0 &&
		lookup.Them[0].Pictures != nil &&
		lookup.Them[0].Pictures.Primary != nil &&
		lookup.Them[0].Pictures.Primary.URL != nil {
		picture, err = url.Parse(*lookup.Them[0].Pictures.Primary.URL)
		if err != nil {
			return nil, err
		}
		c.picturesCache.SetDefault(kid, picture)
	}
	return picture, err
}

func (c *Client) lookup(ctx context.Context, kid string) (*lookupPayload, error) {
	var lookup lookupPayload
	res, err := c.client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetQueryParam("key_suffix", kid).
		SetQueryParam("fields", "pictures").
		SetResult(&lookup).
		Get("https://keybase.io/_/api/1.0/user/lookup.json")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("got non 200 http code: '%d'", res.StatusCode())
	}
	return &lookup, nil
}
