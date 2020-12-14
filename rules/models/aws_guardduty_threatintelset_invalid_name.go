// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyThreatintelsetInvalidNameRule checks the pattern is valid
type AwsGuarddutyThreatintelsetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyThreatintelsetInvalidNameRule returns new rule with default attributes
func NewAwsGuarddutyThreatintelsetInvalidNameRule() *AwsGuarddutyThreatintelsetInvalidNameRule {
	return &AwsGuarddutyThreatintelsetInvalidNameRule{
		resourceType:  "aws_guardduty_threatintelset",
		attributeName: "name",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyThreatintelsetInvalidNameRule) Name() string {
	return "aws_guardduty_threatintelset_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyThreatintelsetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyThreatintelsetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyThreatintelsetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyThreatintelsetInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 300 characters or less",
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
			return nil
		})
	})
}
