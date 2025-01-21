package launchr

import (
	"dsl/api/security"
	"dsl/service"
	"errors"
	"fmt"
	"github.com/launchrctl/launchr"
)

const (
	CmdSecretExtDefault         = ".yaml"
	CmdSecretInteractiveDefault = false
	CmdSecretEncryptionDefault  = true
)

func NewCmdSecretGet(cmd *launchr.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no manifests provided")
	}

	key, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	if key == "" {
		return errors.New("no key provided")
	}

	cmd.SilenceUsage, err = cmd.Flags().GetBool("help")
	if err != nil {
		return err
	}

	out, err := cmd.Flags().GetBool("out")
	if err != nil {
		return err
	}

	ext, err := cmd.Flags().GetString("ext")
	if err != nil {
		return err
	}

	encryption, err := cmd.Flags().GetBool("encryption")
	if err != nil {
		return err
	}

	var manifests service.ManifestService

	var encrypt security.Encryption

	if encryption {
		privKey, err := cmd.Flags().GetString("privkey")
		if err != nil {
			return err
		}
		if privKey == "" {
			return errors.New("no privKey provided")
		}

		pubKey, err := cmd.Flags().GetString("pubkey")
		if err != nil {
			return err
		}
		if pubKey == "" {
			return errors.New("no pubKey provided")
		}

		encrypt, err = security.NewAgeKeyPair(privKey, pubKey)
		if err != nil {
			panic(err)
		}
	}

	if err := manifests.ReadManifestFiles(args, ext); err != nil {
		return fmt.Errorf("error reading manifests: %v\n", err)
	}

	secret, err := manifests.GetSecret(key, encrypt)
	if err != nil {
		return err
	}

	if out {
		cmd.Println(secret)
	}

	return nil
}

func NewCmdSecretSet(cmd *launchr.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no manifests provided")
	}

	key, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	if key == "" {
		return errors.New("no key provided")
	}

	secret, err := cmd.Flags().GetString("secret")
	if err != nil {
		return err
	}
	if secret == "" {
		return fmt.Errorf("no value value provided for key '%s'", key)
	}

	out, err := cmd.Flags().GetBool("out")
	if err != nil {
		return err
	}

	ext, err := cmd.Flags().GetString("ext")
	if err != nil {
		return err
	}

	cmd.SilenceUsage, err = cmd.Flags().GetBool("help")
	if err != nil {
		return err
	}

	encryption, err := cmd.Flags().GetBool("encryption")
	if err != nil {
		return err
	}

	var manifests service.ManifestService

	var encrypt security.Encryption

	if encryption {
		privKey, err := cmd.Flags().GetString("privkey")
		if err != nil {
			return err
		}
		if privKey == "" {
			return errors.New("no privKey provided")
		}

		pubKey, err := cmd.Flags().GetString("pubkey")
		if err != nil {
			return err
		}
		if pubKey == "" {
			return errors.New("no pubKey provided")
		}

		encrypt, err = security.NewAgeKeyPair(privKey, pubKey)
		if err != nil {
			panic(err)
		}
	}

	save := true

	if err := manifests.ReadManifestFiles(args, ext); err != nil {
		return fmt.Errorf("Error reading manifests: %v\n", err)
	}

	if err := manifests.SetSecret(key, secret, encrypt, save); err != nil {
		return err
	}

	if out {
		cmd.Println(secret)
	}

	return nil
}

func NewCmdSecretList(cmd *launchr.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no manifests provided")
	}

	keys, err := cmd.Flags().GetStringArray("keys")
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return errors.New("no keys provided")
	}

	cmd.SilenceUsage, err = cmd.Flags().GetBool("help")
	if err != nil {
		return err
	}

	out, err := cmd.Flags().GetBool("out")
	if err != nil {
		return err
	}

	ext, err := cmd.Flags().GetString("ext")
	if err != nil {
		return err
	}

	encryption, err := cmd.Flags().GetBool("encryption")
	if err != nil {
		return err
	}

	var manifests service.ManifestService

	var encrypt security.Encryption

	if encryption {
		privKey, err := cmd.Flags().GetString("privkey")
		if err != nil {
			return err
		}
		if privKey == "" {
			return errors.New("no privKey provided")
		}

		pubKey, err := cmd.Flags().GetString("pubkey")
		if err != nil {
			return err
		}
		if pubKey == "" {
			return errors.New("no pubKey provided")
		}

		encrypt, err = security.NewAgeKeyPair(privKey, pubKey)
		if err != nil {
			panic(err)
		}
	}

	if err := manifests.ReadManifestFiles(args, ext); err != nil {
		fmt.Printf("Error reading manifests: %v\n", err)
		return err
	}

	secrets, err := manifests.GetSecretList(keys, encrypt)
	if err != nil {
		return err
	}

	if out {
		cmd.Println(secrets)
	}

	return nil
}
