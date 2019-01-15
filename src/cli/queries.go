package main

import "github.com/shurcooL/graphql"

// Publish Extension
type PublishExtensionVariables struct {
	force       bool
	bundle      string
	extensionID string
	manifest    string
	name        string
	version     string
	sourceMap   string
}

type PublishExtensionMutation struct {
	PublishExtension struct {
		Extension struct {
			Id   graphql.ID
			Url  graphql.String
			Name graphql.String
		}
	} `graphql:"publishExtension(force: $force, bundle: $bundle, name: $name, version: $version, manifest: $manifest, sourceMap: $sourceMap, extensionID: $extensionID)"`
}

func NewPublishExtensionMutation(vars PublishExtensionVariables) (PublishExtensionMutation, map[string]interface{}) {
	m := PublishExtensionMutation{}
	variables := map[string]interface{}{
		"name":        graphql.String(vars.name),
		"bundle":      graphql.String(vars.bundle),
		"force":       graphql.Boolean(vars.force),
		"version":     graphql.String(vars.version),
		"manifest":    graphql.String(vars.manifest),
		"sourceMap":   graphql.String(vars.sourceMap),
		"extensionID": graphql.String(vars.extensionID),
	}

	return m, variables
}
