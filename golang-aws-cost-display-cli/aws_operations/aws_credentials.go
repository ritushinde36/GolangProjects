package aws_operations

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

func CheckAwsCredentials(ctx context.Context) error {
	configuration, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("AWS Credentials not found : %w", err)
	}

	_, err = configuration.Credentials.Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("AWS Cresentials are not found or are invalid : %w", err)
	}

	return nil
}
