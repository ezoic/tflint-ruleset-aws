// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackSetInstanceInvalidAccountIDRule checks the pattern is valid
type AwsCloudformationStackSetInstanceInvalidAccountIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudformationStackSetInstanceInvalidAccountIDRule returns new rule with default attributes
func NewAwsCloudformationStackSetInstanceInvalidAccountIDRule() *AwsCloudformationStackSetInstanceInvalidAccountIDRule {
	return &AwsCloudformationStackSetInstanceInvalidAccountIDRule{
		resourceType:  "aws_cloudformation_stack_set_instance",
		attributeName: "account_id",
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Name() string {
	return "aws_cloudformation_stack_set_instance_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
