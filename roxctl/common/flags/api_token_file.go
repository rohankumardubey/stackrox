package flags

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	apiTokenFile string
)

// AddAPITokenFile adds the token-file flag to the base command.
func AddAPITokenFile(c *cobra.Command) {
	c.PersistentFlags().StringVarP(&apiTokenFile,
		"token-file",
		"",
		"",
		"Use the API token in the provided file to authenticate")
}

// APITokenFile returns the currently specified API token file name.
func APITokenFile() string {
	return apiTokenFile
}

// ReadTokenFromFile attempts to retrieve a token from the currently specified file.
func ReadTokenFromFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to retrieve token from file %q", fileName)
	}
	token := strings.TrimSpace(string(content))
	if token != "" {
		return token, nil
	}
	return "", errors.Errorf("failed to retrieve token from file %q: file is empty", fileName)
}
