package definition

import (
	"dsl/api/fs"
	"dsl/api/reader"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const (
	TestdataPath = "testdata"
	TestfileExt  = ".yaml"
)

var Testdata []string

func init() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to determine current work dir")
	}

	testPath := filepath.Join(workDir, TestdataPath)

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		log.Fatalf("testdata path %s is not exists", testPath)
	}

	Testdata, err = fs.DiscoverFiles([]string{testPath}, TestfileExt)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestKindDetection(t *testing.T) {
	for _, file := range Testdata {
		t.Logf("testing of kind detection of file '%s'\n", file)
		yamlData, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("unable to extract schema from yaml %s: %s", string(yamlData), err.Error())
		}

		component, err := reader.ParseComponent(yamlData)

		if err != nil {
			t.Fatalf("error while creating component from '%s': '%s'", file, err.Error())
		}
		if component == nil {
			t.Fatalf("component created from '%s' is nil", file)
		}
	}
}

func TestValidate(t *testing.T) {
	for _, file := range Testdata {
		t.Logf("testing deserizlization of %s", file)
		yamlData, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("unable to extract schema from yaml %s: %s", string(yamlData), err.Error())
		}

		component, err := reader.ParseComponent(yamlData)
		if err != nil {
			t.Fatalf("error when creating component: '%s'", err.Error())
		}

		if err := component.Unmarshal(yamlData); err != nil {
			t.Fatalf("unable to deserizlize %s: %s", file, err.Error())
		}

		if err := component.Validate(); err != nil {
			t.Fatalf("unable to deserizlize %s: %s", file, err.Error())
		}
	}
}
