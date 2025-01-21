package dsl

import (
	cli "dsl/api/launchr"
	"github.com/launchrctl/launchr"
)

func init() {
	launchr.RegisterPlugin(&Plugin{})
}

// Plugin is [launchr.Plugin] plugin providing compose.
type Plugin struct {
	wd string
}

// PluginInfo implements [launchr.Plugin] interface.
func (p *Plugin) PluginInfo() launchr.PluginInfo {
	return launchr.PluginInfo{Weight: 10}
}

// OnAppInit implements [launchr.OnAppInitPlugin] interface.
func (p *Plugin) OnAppInit(app launchr.App) error {
	//app.GetService(&p.k)
	p.wd = app.GetWD()
	//buildDir := filepath.Join(p.wd, compose.BuildDir)
	//app.RegisterFS(action.NewDiscoveryFS(os.DirFS(BuildDir), p.wd))
	return nil
}

// CobraAddCommands implements [launchr.CobraPlugin] interface to provide compose functionality.
func (p *Plugin) CobraAddCommands(rootCmd *launchr.Command) error {
	var flagValidateExt string
	var flagValidateOut bool
	var flagValidateSilent bool

	var validateCmd = &launchr.Command{
		Use:   "validate [...manifests] --ext --out",
		Short: "validation utility for schema and components according to schema",
		Long:  `command to perform Plasma Component yaml schema validation`,
		RunE: func(cmd *launchr.Command, args []string) error {
			return cli.NewCmdValidate(cmd, args)
		},
	}
	validateCmd.Flags().StringVarP(&flagValidateExt, "ext", "x", cli.CmdManifestExtDefault, "file extension regex to read plasma manifests")
	validateCmd.Flags().BoolVarP(&flagValidateOut, "out", "o", !cli.CmdOutDefault, "stdout validated manifests")
	validateCmd.Flags().BoolVarP(&flagValidateSilent, "help", "h", cli.CmdSilentDefault, "show help on failure")
	rootCmd.AddCommand(validateCmd)

	//var flagInteractive bool
	var secretCmd = &launchr.Command{
		Use:   "secret [...manifests]",
		Short: "secret management utility for secret containing yaml fields",
		Long:  `command to set encrypted secret value by key`,
	}
	rootCmd.AddCommand(secretCmd)

	var flagSecretGetExt string
	var flagSecretGetOut bool
	var flagSecretGetSilent bool
	var flagSecretGetKey string
	var flagSecretGetEncryption bool
	var flagSecretGetPrivateKey string
	var flagSecretGetPublicKey string

	var secretGetCmd = &launchr.Command{
		Use:   "get --key [KEY] --privkey [PRIVKEY] --encryption --out [...MANIFESTS]",
		Short: "subcommand to get secret value by key",
		Long:  `command to get decrypted secret value by key in kv and user private key`,
		RunE: func(cmd *launchr.Command, args []string) error {
			return cli.NewCmdSecretGet(cmd, args)
		},
	}
	secretGetCmd.Flags().StringVarP(&flagSecretGetExt, "ext", "x", cli.CmdSecretExtDefault, "file extensions to read plasma components")
	secretGetCmd.Flags().BoolVarP(&flagSecretGetOut, "out", "o", cli.CmdOutDefault, "stdout secret value by key")
	secretGetCmd.Flags().BoolVarP(&flagSecretGetSilent, "help", "h", cli.CmdSilentDefault, "show help on failure")
	secretGetCmd.Flags().StringVarP(&flagSecretGetKey, "key", "k", "", "key to get value by")
	secretGetCmd.Flags().BoolVarP(&flagSecretGetEncryption, "encryption", "c", cli.CmdSecretEncryptionDefault, "enable encrypted mode. Will decrypted secret for output. It will out encrypted value otherwise")
	secretGetCmd.Flags().StringVarP(&flagSecretGetPrivateKey, "privkey", "d", "", "provide private (decrypt) key of key pair")
	secretGetCmd.Flags().StringVarP(&flagSecretGetPublicKey, "pubkey", "e", "", "provide public (encrypt) key of key pair")
	secretCmd.AddCommand(secretGetCmd)

	var flagSecretSetOut bool
	var flagSecretSetExt string
	var flagSecretSetSilent bool
	var flagSecretSetKey string
	var flagSecretSetSecret string
	var flagSecretSetEncryption bool
	var flagSecretSetPrivateKey string
	var flagSecretSetPublicKey string

	var secretSetCmd = &launchr.Command{
		Use:   "set --key [KEY] --value [VALUE] [...manifests]",
		Short: "subcommand to set secret value by key and update manifest secrets",
		Long:  `setting of empty string deletes key-value. Setting of unexistent key creates new key-value. All [...manifests] provided are affected by this utility`,
		RunE: func(cmd *launchr.Command, args []string) error {
			return cli.NewCmdSecretSet(cmd, args)
		},
	}
	secretSetCmd.Flags().StringVarP(&flagSecretSetExt, "ext", "x", cli.CmdSecretExtDefault, "file extensions to read plasma manifests")
	secretSetCmd.Flags().BoolVarP(&flagSecretSetOut, "out", "o", cli.CmdOutDefault, "stdout updated secret")
	secretSetCmd.Flags().BoolVarP(&flagSecretSetSilent, "help", "h", cli.CmdSilentDefault, "show help on failure")
	secretSetCmd.Flags().StringVarP(&flagSecretSetKey, "key", "k", "", "key to set value")
	secretSetCmd.Flags().StringVarP(&flagSecretSetSecret, "secret", "s", "", "plain value to update secret by key")
	secretSetCmd.Flags().BoolVarP(&flagSecretSetEncryption, "encryption", "c", cli.CmdSecretEncryptionDefault, "enable encrypted mode. Consider secret fields are encrypted values. The mode performs encryption of new secret before appending (if it's not encrypted already). Otherwise, secrets considered plain and no encryption enforced. Recommended if you are updating secret with already encrypted value")
	secretSetCmd.Flags().StringVarP(&flagSecretSetPrivateKey, "privkey", "d", "", "provide private (decryption) key")
	secretSetCmd.Flags().StringVarP(&flagSecretSetPublicKey, "pubkey", "e", "", "provide public (encryption) key")
	secretCmd.AddCommand(secretSetCmd)

	var flagSecretListExt string
	var flagSecretListEncryption bool
	var flagSecretListOut bool
	var flagSecretListSilent bool
	var flagSecretListKeys []string
	var flagSecretListPrivateKey string
	var flagSecretListPublicKey string

	var secretListCmd = &launchr.Command{
		Use:   "list [...manifests]",
		Short: "secret management utility for secret containing yaml fields",
		Long:  `command to list secret KV in encrypted and plain modes`,
		RunE: func(cmd *launchr.Command, args []string) error {
			return cli.NewCmdSecretList(cmd, args)
		},
	}
	secretListCmd.Flags().StringVarP(&flagSecretListExt, "ext", "x", cli.CmdSecretExtDefault, "file extensions to read plasma components")
	secretListCmd.Flags().BoolVarP(&flagSecretListOut, "out", "o", cli.CmdOutDefault, "stdout queried secrets as list")
	secretListCmd.Flags().BoolVarP(&flagSecretListSilent, "help", "h", cli.CmdSilentDefault, "show help on failure")
	secretListCmd.Flags().StringArrayVarP(&flagSecretListKeys, "keys", "k", []string{}, "list of keys to lookup")
	secretListCmd.Flags().BoolVarP(&flagSecretListEncryption, "encryption", "c", cli.CmdSecretEncryptionDefault, "enable encrypted mode. Will decrypted secret for output. It will out encrypted value otherwise")
	secretListCmd.Flags().StringVarP(&flagSecretListPrivateKey, "privkey", "d", "", "provide private (decryption) key")
	secretListCmd.Flags().StringVarP(&flagSecretListPublicKey, "pubkey", "e", "", "provide public (encryption) key")
	secretCmd.AddCommand(secretListCmd)

	return nil
}
