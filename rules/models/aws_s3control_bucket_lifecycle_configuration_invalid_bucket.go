// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3controlBucketLifecycleConfigurationInvalidBucketRule checks the pattern is valid
type AwsS3controlBucketLifecycleConfigurationInvalidBucketRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsS3controlBucketLifecycleConfigurationInvalidBucketRule returns new rule with default attributes
func NewAwsS3controlBucketLifecycleConfigurationInvalidBucketRule() *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule {
	return &AwsS3controlBucketLifecycleConfigurationInvalidBucketRule{
		resourceType:  "aws_s3control_bucket_lifecycle_configuration",
		attributeName: "bucket",
		max:           255,
		min:           3,
	}
}

// Name returns the rule name
func (r *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule) Name() string {
	return "aws_s3control_bucket_lifecycle_configuration_invalid_bucket"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3controlBucketLifecycleConfigurationInvalidBucketRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"bucket must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"bucket must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
