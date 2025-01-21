package launchr

import (
	"dsl/service"
	"errors"
	"fmt"
	"github.com/launchrctl/launchr"
)

func NewCmdValidate(cmd *launchr.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no manifests provided")
	}

	ext, err := cmd.Flags().GetString("ext")
	if err != nil {
		return err
	}

	out, err := cmd.Flags().GetBool("out")
	if err != nil {
		return err
	}

	cmd.SilenceUsage, err = cmd.Flags().GetBool("help")
	if err != nil {
		return err
	}

	var manifests service.ManifestService

	if err := manifests.ReadManifestFiles(args, ext); err != nil {
		return err
	}

	if out {
		for file, manifest := range manifests.GetManifests() {
			serialized, err := manifest.Marshal()
			if err != nil {
				return fmt.Errorf("error '%s' while serializing '%s'", err.Error(), file)
			}
			cmd.Printf("# Content of '%s'\n%s\n", file, string(serialized))
		}
	}

	return nil
}
