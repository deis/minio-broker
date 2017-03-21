package controller

import "errors"

type Config struct {
	Region          string `json:"region"`
	BucketPrefix    string `json:"bucket_prefix"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl"`
}

func (c Config) Validate() error {
	if c.Region == "" {
		return errors.New("Must provide a non-empty Region")
	}

	if c.BucketPrefix == "" {
		return errors.New("Must provide a non-empty BucketPrefix")
	}

	if c.Endpoint == "" {
		return errors.New("Must provide a non-empty Endpoint")
	}

	if c.AccessKeyID == "" {
		return errors.New("Must provide a non-empty AccessKeyID")
	}

	if c.SecretAccessKey == "" {
		return errors.New("Must provide a non-empty SecretAccessKey")
	}

	return nil
}
