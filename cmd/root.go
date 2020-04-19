package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var (
	version bool
	cfgFile string
)

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(home)

		viper.AddConfigPath(home + "/.fleepr")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		//clientConf = clientConf.GetConf(viper.ConfigFileUsed())

		//fmt.Printf("TOKEN: %s", clientConf.Im.Token)
		//if clientConf.im.Password == "" {
		//	fmt.Println("No password")
		//}
	}

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fleepr",
	Short: "Fleepr deployments",
	Long: `Fleepr deployments.
Default configuration file searched in $HOME/.fleepr.yaml

Usage examples:
"""
# CREATE A CLUSTER FROM TEMPLATE
fleepr create --template my_tosca_template.yml

# VALIDATE TOSCA TEMPLATE
fleepr validate --template my_tosca_template.yml
"""`,

	Run: func(cmd *cobra.Command, args []string) {
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// VersionString ..
var VersionString string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VersionString = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildDoc ...
func BuildDoc() {
	err := doc.GenMarkdownTree(rootCmd, "docs")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "fleepr config file (default is $HOME/.fleepr/config.yaml)")
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "DODAS client version")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
