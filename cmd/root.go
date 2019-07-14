package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"

  homedir "github.com/mitchellh/go-homedir"
  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var project = "fundamentals"
var (
  rootCmd = &cobra.Command{
    Use: project,
    Short: "Fundamental computer science concepts implemented in go",
    PersistentPreRun: preRunFunction,
  }
  cfgDir string
)

func setRootFlags() {
  rootCmd.PersistentFlags().StringVar(&cfgDir, "config", "", fmt.Sprintf("config directory (default is $HOME/.%s)", project))
  rootCmd.PersistentFlags().Bool("debug", false, "Debug flag")
}

func init() {
  setRootFlags()
  cobra.OnInitialize(initConfig)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func preRunFunction(cmd *cobra.Command, args []string) {
  viper.BindPFlags(cmd.PersistentFlags())
  viper.BindPFlags(cmd.Flags())
  
}

func initConfig() {
  viper.BindPFlags(rootCmd.PersistentFlags())

  if cfgDir != "" {
    viper.AddConfigPath(cfgDir)
    viper.Set("configPath", filepath.Join(cfgDir, "config.yml"))
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".cobra" (without extension).
    viper.AddConfigPath(filepath.Join(home, fmt.Sprintf(".%s", project)))
    viper.SetConfigName(fmt.Sprintf(".%s", project))
    viper.Set("configPath", filepath.Join(home, "config.yml"))
  }

  viper.SetConfigName("config")
  viper.SetEnvPrefix(fmt.Sprintf(".%s", project))
  viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err != nil {
    viper.Set("configPath", viper.ConfigFileUsed())
  }

  // set up logging
  log.SetOutput(os.Stderr)
  if (viper.GetBool("debug")) {
    log.SetLevel(log.DebugLevel)
  } else {
    log.SetLevel(log.InfoLevel)
  }
}
