package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func readManifest(manifestFlag string) (*Manifest, error) {
	manifest := Manifest{}
	data, err := ioutil.ReadFile(manifestFlag)
	if err != nil {
		return nil, fmt.Errorf("%s\n\nRun this command in a directory with a %s file for an extension.\n\n", err, manifestFlag)
	}

	if err = json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

type Manifest struct {
	Name        string `json:"name"`
	Main        string `json:"main"`
	Readme      string `json:"readme"`
	Version     string `json:"version"`
	Publisher   string `json:"publisher"`
	ExtensionID string `json:"extensionID"`
	Scripts     struct {
		Prepublish string `json:"cdebase:prepublish"`
	} `json:"scripts"`
}

func (m *Manifest) String() string {
	var str string

	data, _ := json.Marshal(m)
	str = string(data)

	return str
}

func (m *Manifest) readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func (m *Manifest) Prepublish(dir string) error {
	if m.Scripts.Prepublish == "" {
		return nil
	}

	cmd := exec.Command("bash", "-c", m.Scripts.Prepublish)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s:%s", filepath.Join(dir, "node_modules", ".bin"), os.Getenv("PATH")))
	cmd.Dir = dir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	fmt.Fprintf(os.Stderr, "# cdebase:prepublish: %s\n", m.Scripts.Prepublish)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cdebase:prepublish script failed: %s (see output above)", err)
	}

	fmt.Fprintln(os.Stderr)
	return nil
}

func (m *Manifest) ReadBundle() (string, error) {
	bundle, err := m.readFile(m.Main)

	if err != nil {
		return "", err
	}

	return string(bundle), nil
}

func (m *Manifest) ReadArtifacts(dir string) error {
	m.Readme = m.GetReadme(dir)

	if m.Name == "" && m.Publisher == "" {
		return errors.New(`extension manifest must contain "name" and "publisher" string properties (the extension ID is of the form "publisher/name" and uses these values)`)
	}

	if m.Name == "" {
		return fmt.Errorf(`extension manifest must contain a "name" string property for the extension name (the extension ID will be %q)`, m.Publisher+"/name")
	}

	if m.Publisher == "" {
		return fmt.Errorf(`extension manifest must contain a "publisher" string property referring to a username or organization name on Sourcegraph (the extension ID will be %q)`, "publisher/"+m.Name)
	}

	//m.ExtensionID = fmt.Sprintf("%s/%s", m.Publisher, m.Name)

	return nil
}

func (m *Manifest) GetReadme(dir string) string {
	var readme string
	filenames := []string{"readme.md", "README.txt", "README", "readme.md", "readme.txt", "readme", "Readme.md", "Readme.txt", "Readme"}
	for _, f := range filenames {
		data, err := ioutil.ReadFile(filepath.Join(dir, f))
		if err != nil {
			continue
		}
		readme = string(data)
		break
	}

	return readme
}
