package config

import "errors"

var gitTagVersion string

type Version struct {
	GitTagVersion string
}

func NewVersion() *Version {
	return &Version{
		GitTagVersion: gitTagVersion,
	}
}

var (
	ErrVersionTypeAssertion = errors.New("cannot type assert version interface{} -> *config.Version")
)
