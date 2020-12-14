// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackInvalidIAMRoleArnRule checks the pattern is valid
type AwsCloudformationStackInvalidIAMRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackInvalidIAMRoleArnRule returns new rule with default attributes
func NewAwsCloudformationStackInvalidIAMRoleArnRule() *AwsCloudformationStackInvalidIAMRoleArnRule {
	return &AwsCloudformationStackInvalidIAMRoleArnRule{
		resourceType:  "aws_cloudformation_stack",
		attributeName: "iam_role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackInvalidIAMRoleArnRule) Name() string {
	return "aws_cloudformation_stack_invalid_iam_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackInvalidIAMRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackInvalidIAMRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackInvalidIAMRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackInvalidIAMRoleArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"iam_role_arn must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"iam_role_arn must be 20 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
