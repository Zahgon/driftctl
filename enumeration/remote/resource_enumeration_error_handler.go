package remote

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	remoteerror "github.com/snyk/driftctl/enumeration/remote/error"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"google.golang.org/grpc/status"
)

func HandleResourceEnumerationError(err error, alerter alerter.AlerterInterface) error {
	_ = "STUB: not implemented"
	return nil
}

// We cannot use the status.FromError() method because AWS errors are not well-formed.
// Indeed, they compose the error interface without implementing the Error() method and thus triggering a nil panic
// when returning an unknown error from status.FromError()
// As a workaround we duplicated the logic from status.FromError here

// at least for storage api google sdk does not return grpc error so we parse the error message.

// This handles access denied errors like the following:
// aws_s3_bucket_policy: AccessDenied: Error listing bucket policy <policy_name>

func handleAWSError(alerter alerter.AlerterInterface, listError *remoteerror.ResourceScanningError, reqerr awserr.RequestFailure) error {
	_ = "STUB: not implemented"
	return nil
}

func handleGoogleEnumerationError(alerter alerter.AlerterInterface, err *remoteerror.ResourceScanningError, st *status.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldHandleGoogleForbiddenError(err *remoteerror.ResourceScanningError) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if this is a Google related error
