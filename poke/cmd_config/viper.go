package cmd_config

const (
	DefaultConfigFileName string = "cmd_config"
	DefaultConfigFileType string = "json"
)

const (
	OutputPathField string = "output-path"
)

type CMDConfig struct {
	OutputPath string
}
