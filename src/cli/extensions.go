package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
	"github.com/urfave/cli"
	"log"
	"path/filepath"
)

type Extension struct {
	dir      string
	entry    string
	bundle   string
	manifest *Manifest
}

func (e *Extension) GetManifest() (*Manifest, error) {
	manifest, err := readManifest(e.entry)

	if err != nil {
		log.Fatal("Cannot read extension manifest. Check directory and files!")
	}
	return manifest, nil
}

func (e *Extension) PrepareManifest() {
	if e.manifest == nil {
		log.Fatal("Cannot build extension without manifest!")
	}

	// Add files to manifest
	e.manifest.ReadArtifacts(e.dir)

	// Prepare files - run build script
	e.manifest.Prepublish(e.dir)
}

func NewExtension(dir string, entry string) (*Extension, error) {
	extension := Extension{
		dir:   dir,
		entry: entry,
	}

	manifest, err := extension.GetManifest()
	if err != nil {
		return nil, err
	}

	extension.manifest = manifest
	extension.PrepareManifest()

	bundle, _ := extension.manifest.ReadBundle()
	extension.bundle = bundle

	return &extension, nil
}

func init() {
	fmt.Println(">>>")
	commands = append(commands, cli.Command{
		Name:        "extension",
		Aliases:     []string{"e"},
		Description: "Extensions management",
		Subcommands: []cli.Command{
			{
				Name:        "publish",
				Aliases:     []string{"p"},
				Description: "Publish Extension to CDEBase Repository",
				Action: func(c *cli.Context) error {

					fmt.Println("Publishing extension...")
					config, _ := loadConfig(flags.ConfigPath, &flags)
					client := graphql.NewClient(config.Endpoint, nil)

					extension, err := NewExtension(filepath.Dir(flags.Manifest), flags.Manifest)
					if err != nil {
						log.Fatal(err)
					}

					mutation, variables := NewPublishExtensionMutation(PublishExtensionVariables{
						force:       true,
						bundle:      extension.bundle,
						name:        extension.manifest.Name,
						version:     extension.manifest.Version,
						manifest:    extension.manifest.String(),
						extensionID: extension.manifest.ExtensionID,
					})

					// TODO: Add log file.
					err = client.Mutate(context.Background(), &mutation, variables)

					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Printf("Extension %s has been published! \n", extension.manifest.ExtensionID)
					}

					return err
				},
			},
		},
	})
}
