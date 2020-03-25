package filesystem

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configDirPath     = ".bpaas/config/clusters"
	userConfigDirPath = ".bpaas/config/user"
)

//CreateConfigDirIfNotExists
func CreateDirIfNotExists(path string) error {

	if IsNotExistsConfigDirectory(path) {
		return CreateDirectory(path)
	}

	return nil
}

func IsNotExistsConfigDirectory(configPath string) bool {

	_, err := os.Stat(configPath)

	if err != nil && os.IsNotExist(err) {
		return true
	}

	return false
}

func GetClusterConfigPath() (string, error) {

	home, err := GetUserHomeDirectory()

	if err != nil {
		fmt.Println("...error retrieving home path", err)
		return "", err
	}

	path := fmt.Sprintf("%s/%s", home, configDirPath)

	return path, nil
}

func GetUserConfigPath() (string, error) {

	home, err := GetUserHomeDirectory()

	if err != nil {
		fmt.Println("...error retrieving home path", err)
		return "", err
	}

	path := fmt.Sprintf("%s/%s", home, userConfigDirPath)

	return path, nil
}

func GetUserHomeDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir, nil
}

func CreateDirectory(path string) error {

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	return nil

}

func WriteToYAMLFile(data string, fileName string, configDir string) error {

	if err := CreateDirIfNotExists(configDir); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%v/%v.yaml", configDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	fmt.Println("..writing config to ", filePath)
	_, err = file.WriteString(data)

	if err != nil {
		file.Close()
		return err
	}
	return nil
}

func CreateYAML(data interface{}) (string, error) {

	ym, err := yaml.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(ym), nil
}
