package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Config holds all application configuration
type Config struct {
	GoogleAnalyticsID  string
	SecurityTxtContent string

	// Server settings
	Port       string
	SiteNameEn string
	SiteNameFr string

	// WordPress API settings
	WordPressBaseURL  string
	WordPressUsername string
	WordPressPassword string
	WordPressMenuIdEn string
	WordPressMenuIdFr string
}

// Load reads configuration and resolves Lambda credentials from SSM at startup.
func Load() (*Config, error) {
	cfg := &Config{Port: os.Getenv("PORT")}

	requiredVars := map[string]*string{
		"SECURITY_TXT_CONTENT": &cfg.SecurityTxtContent,
		"SITE_NAME_EN":         &cfg.SiteNameEn,
		"SITE_NAME_FR":         &cfg.SiteNameFr,
		"WORDPRESS_URL":        &cfg.WordPressBaseURL,
		"WORDPRESS_MENU_ID_EN": &cfg.WordPressMenuIdEn,
		"WORDPRESS_MENU_ID_FR": &cfg.WordPressMenuIdFr,
	}

	var usernameParam, passwordParam string
	if cfg.Port == "" {
		requiredVars["WORDPRESS_USERNAME_PARAM"] = &usernameParam
		requiredVars["WORDPRESS_PASSWORD_PARAM"] = &passwordParam
	} else {
		requiredVars["WORDPRESS_USERNAME"] = &cfg.WordPressUsername
		requiredVars["WORDPRESS_PASSWORD"] = &cfg.WordPressPassword
	}

	// Check all required variables
	var missingVars []string
	for name, ptr := range requiredVars {
		val := os.Getenv(name)
		if val == "" {
			missingVars = append(missingVars, name)
		} else {
			*ptr = val
		}
	}

	// Return error if any required variables are missing
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	if cfg.Port == "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		awsCfg, err := awsconfig.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("load AWS configuration: %w", err)
		}

		result, err := ssm.NewFromConfig(awsCfg).GetParameters(ctx, &ssm.GetParametersInput{
			Names:          []string{usernameParam, passwordParam},
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			return nil, fmt.Errorf("load WordPress credentials from SSM: %w", err)
		}
		if len(result.InvalidParameters) > 0 {
			return nil, fmt.Errorf("SSM parameters not found: %v", result.InvalidParameters)
		}

		values := make(map[string]string, len(result.Parameters))
		for _, parameter := range result.Parameters {
			values[aws.ToString(parameter.Name)] = aws.ToString(parameter.Value)
		}
		if values[usernameParam] == "" {
			return nil, fmt.Errorf("SSM parameter %q is missing or empty", usernameParam)
		}
		if values[passwordParam] == "" {
			return nil, fmt.Errorf("SSM parameter %q is missing or empty", passwordParam)
		}
		cfg.WordPressUsername = values[usernameParam]
		cfg.WordPressPassword = values[passwordParam]
	}

	// Set optional variables
	cfg.GoogleAnalyticsID = os.Getenv("GOOGLE_ANALYTICS_ID")

	return cfg, nil
}
