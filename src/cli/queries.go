package main

import "github.com/shurcooL/graphql"

// Publish Extension
type PublishExtensionVariables struct {
	force bool
	bundle string
	extensionID string
	manifest string
	sourceMap string
}

type PublishExtensionMutation struct {
	PublishExtension struct {
		url graphql.String
		name graphql.String
	} `graphql:"publishExtension(force: $force, bundle: $bundle, manifest: $manifest, sourceMap: $sourceMap, extensionID: $extensionID)"`
}

func NewPublishExtensionMutation(vars PublishExtensionVariables) (PublishExtensionMutation, map[string]interface{}) {
	m := PublishExtensionMutation{}
	variables := map[string]interface{}{
		"force": graphql.Boolean(vars.force),
		"bundle": graphql.String(vars.bundle),
		"manifest": graphql.String(vars.manifest),
		"sourceMap": graphql.String(vars.sourceMap),
		"extensionID": graphql.String(vars.extensionID),
	}

	return m, variables
}

