/*
Copyright 2017 congcongke authors. All rights reserved.
*/

package v1

import (
	"github.com/congcongke/helm-registry/pkg/api/models"
	"github.com/congcongke/helm-registry/pkg/storage"
)

// StringCollectionResult describes a collection of []string
type StringCollectionResult struct {
	Metadata models.Metadata `json:"metadata"`
	Items    []string        `json:"items"`
}

// MetadataCollectionResult describes a collection of []*chart.Metadata
type MetadataCollectionResult struct {
	Metadata models.Metadata     `json:"metadata"`
	Items    []*storage.Metadata `json:"items"`
}
