// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaPermissionInvalidSourceAccountRule checks the pattern is valid
type AwsLambdaPermissionInvalidSourceAccountRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidSourceAccountRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidSourceAccountRule() *AwsLambdaPermissionInvalidSourceAccountRule {
	return &AwsLambdaPermissionInvalidSourceAccountRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "source_account",
		pattern:       regexp.MustCompile(`^\d{12}$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidSourceAccountRule) Name() string {
	return "aws_lambda_permission_invalid_source_account"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidSourceAccountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidSourceAccountRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidSourceAccountRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidSourceAccountRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
