// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchLogStreamInvalidNameRule checks the pattern is valid
type AwsCloudwatchLogStreamInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchLogStreamInvalidNameRule returns new rule with default attributes
func NewAwsCloudwatchLogStreamInvalidNameRule() *AwsCloudwatchLogStreamInvalidNameRule {
	return &AwsCloudwatchLogStreamInvalidNameRule{
		resourceType:  "aws_cloudwatch_log_stream",
		attributeName: "name",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^[^:*]*$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogStreamInvalidNameRule) Name() string {
	return "aws_cloudwatch_log_stream_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogStreamInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogStreamInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogStreamInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogStreamInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 512 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^:*]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
