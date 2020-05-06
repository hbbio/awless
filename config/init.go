package config

import (
	"fmt"
	"os"
	"path/filepath"

	"strconv"

	awsservices "github.com/hbbio/awless/aws/services"
	"github.com/hbbio/awless/database"
)

var (
	AwlessHome         = filepath.Join(os.Getenv("HOME"), ".awless")
	DBPath             = filepath.Join(AwlessHome, database.Filename)
	Dir                = filepath.Join(AwlessHome, "aws")
	KeysDir            = filepath.Join(AwlessHome, "keys")
	AwlessFirstInstall bool
)

func init() {
	os.Setenv("__AWLESS_HOME", AwlessHome)
	os.Setenv("__AWLESS_CACHE", filepath.Join(AwlessHome, "cache"))
	os.Setenv("__AWLESS_KEYS_DIR", KeysDir)
}

func InitAwlessEnv() error {
	_, err := os.Stat(DBPath)

	AwlessFirstInstall = os.IsNotExist(err)
	os.Setenv("__AWLESS_FIRST_INSTALL", strconv.FormatBool(AwlessFirstInstall))

	os.MkdirAll(KeysDir, 0700)

	if AwlessFirstInstall {
		fmt.Fprintln(os.Stderr, AWLESS_ASCII_LOGO)
		fmt.Fprintln(os.Stderr, "Welcome! Resolving environment data...")
		fmt.Fprintln(os.Stderr)

		if err = InitConfig(resolveRequiredConfigFromEnv()); err != nil {
			return err
		}

		err = database.Execute(func(db *database.DB) error {
			return db.SetStringValue("current.version", Version)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot store current version in db: %s\n", err)
		}
	}

	if err = LoadConfig(); err != nil {
		return err
	}

	return nil
}

func resolveRequiredConfigFromEnv() map[string]string {
	region := awsservices.ResolveRegionFromEnv()

	resolved := make(map[string]string)
	if region != "" {
		resolved[RegionConfigKey] = region
	}

	return resolved
}
