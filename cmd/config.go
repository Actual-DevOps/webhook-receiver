package cmd

var configFile string

func init() {
	defaultConfigPath := "config/config.yaml"
	rootCmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigPath, "Config path")
}
