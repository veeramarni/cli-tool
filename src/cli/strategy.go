package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
)

type ExtensionContext struct {
	Dir           string
	ManifestFile  string
	GraphqlClient *graphql.Client
}

type ExtensionLifecycle interface {
	Build() (bool, error)
	Pack() (string, error)
	Publish() (bool, error)
	AddToRegistry() (bool, error)
}

type AbstractExtension struct {
	ExtensionLifecycle

	Manifest *Manifest
	Context  *ExtensionContext
}

type NPMExtension struct {
	AbstractExtension
}

type RAWExtension struct {
	AbstractExtension
}

// Common Lifecycle commands

func (ae AbstractExtension) Build() (bool, error) {
	return runCommand(ae.Context.Dir, ae.Manifest.Scripts.Build, "")
}

func (ae AbstractExtension) AddToRegistry() (bool, error) {
	mutation, variables := NewPublishExtensionMutation(PublishExtensionVariables{
		force:       true,
		name:        ae.Manifest.Name,
		bundle:      ae.Manifest.Bundle,
		version:     ae.Manifest.Version,
		manifest:    ae.Manifest.String(),
		extensionID: ae.Manifest.ExtensionID,
	})

	err := ae.Context.GraphqlClient.Mutate(context.Background(), &mutation, variables)

	fmt.Printf("Mutation %s \n", ae.Manifest.String())

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ae AbstractExtension) Pack() (string, error) {
	var err error
	fmt.Println("Packing extension and preparing to publish...")

	err = ae.Manifest.ReadAssets(ae.Context.Dir)

	if err != nil {
		return "", err
	}

	err = ae.Manifest.ReadBundle(ae.Context.Dir)

	if err != nil {
		return "", err
	}

	return "", nil
}

// Build RAW Extension

func (re RAWExtension) Publish() (bool, error) {
	fmt.Println("Waiting for extension publish...")
	return true, nil
}

// Build NPM Extension

func (ne NPMExtension) Publish() (bool, error) {
	var cmd = fmt.Sprintf("npm publish --registry=%s", config.Registry)

	if ne.Manifest.Scripts.Publish != "" {
		cmd = ne.Manifest.Scripts.Publish
	}

	return runCommand(ne.Context.Dir, cmd, fmt.Sprintf("REGISTRY=%s", config.Registry))
}
