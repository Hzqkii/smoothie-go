package portable

import (
	"log"
	"os"
	"path/filepath"
)

func GetExecutableDirectory() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exe)
}

func dropFileAtPath(path string, contents []byte) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		log.Fatal("DEV Error: must check if stuff exists before writing")
	}

	err := os.WriteFile(path, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func IsPortable() bool {
	if _, err := os.Stat(filepath.Join(GetExecutableDirectory(), "smoothie-go.portable")); err == nil {
		return true
	}
	return false
}

func GetUserHome() string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDirectory
}

func GetConfigDirectory() string {
	if IsPortable() {
		return GetExecutableDirectory()
	}

	configDirectory := filepath.Join(GetUserHome(), ".config", "smoothie-go")
	if err := os.MkdirAll(configDirectory, 0755); err != nil {
		log.Fatal(err)
	}
	return configDirectory
}

func GetLocalDirectory() string {
	localDirectory := filepath.Join(GetUserHome(), ".local", "share", "smoothie-go")
	if err := os.MkdirAll(localDirectory, 0755); err != nil {
		log.Fatal(err)
	}
	return localDirectory
}

func GetRecipePathCustom(name string) string {
	return filepath.Join(GetConfigDirectory(), name)
}

func GetRecipePath() string {
	recipePath := GetRecipePathCustom("recipe.ini")
	if _, err := os.Stat(recipePath); os.IsNotExist(err) {
		dropFileAtPath(recipePath, []byte(recipe_ini))
	}
	return recipePath
}

func GetDefaultRecipePath() string {
	defaultsPath := filepath.Join(GetExecutableDirectory(), "defaults.ini")

	if _, err := os.Stat(defaultsPath); os.IsNotExist(err) {
		dropFileAtPath(defaultsPath, []byte(defaults_ini))
	}
	return defaultsPath
}

func GetEncodingPresetsPath() string {
	encodingPresetsPath := filepath.Join(GetConfigDirectory(), "encoding_presets.ini")
	if _, err := os.Stat(encodingPresetsPath); os.IsNotExist(err) {
		dropFileAtPath(encodingPresetsPath, []byte(encoding_resets_ini))
	}
	return encodingPresetsPath
}

func GetModelsPath() string {
	modelsPath := filepath.Join(GetLocalDirectory(), "models")
	if _, err := os.Stat(modelsPath); os.IsNotExist(err) {
		err := os.MkdirAll(modelsPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
		writeEmbeddedFiles(models, modelsPath, "assets/models")
	}
	return modelsPath
}
