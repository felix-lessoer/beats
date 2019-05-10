// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package ibmmq

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "ibmmq", asset.ModuleFieldsPri, AssetIbmmq); err != nil {
		panic(err)
	}
}

// AssetIbmmq returns asset data.
// This is the base64 encoded gzipped contents of module/ibmmq.
func AssetIbmmq() string {
	return "eJzMlkFuszAQhfecYpR9pD9/papi0UV3XURqT1C5eEKs2IwznqTh9hWkNCYBtd1gWIH9ePPxbGCWsMM6B/Pu3D4DECMWc1i014sMQGMo2HgxVOXwmAHAWQtr0geLGcDGoNUhb6eWUCmHF7vmkNpjDiXTwX+NDHj2bWIrZCa2VH5PXN3//LSG9etZBpbKEAnxpJxvnufFogoISutuLFLdAg7xXIg8U4EhrnOFFFWLxnv635F1bIInGQA5BOQJKXZYfxDr4URKVm4WLFsK0pylXx5TBVHWqhv3VNEc2W3Sx7J35ZS7dmxxAj3c/1u9iXEYRDk/AVJXm9GRYLNR0+eg2MjWVAFZVrOi+Z+epiDnsJLZpNPjmVs+d+l52h6gID2Db3+LMlYhDdHJW1VN/TMae8+LlBwdRcAjspF6rLX8keIv/WVL8BkAAP//bdE+mQ=="
}
