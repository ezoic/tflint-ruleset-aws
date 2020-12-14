// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53HealthCheckInvalidReferenceNameRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidReferenceNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53HealthCheckInvalidReferenceNameRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidReferenceNameRule() *AwsRoute53HealthCheckInvalidReferenceNameRule {
	return &AwsRoute53HealthCheckInvalidReferenceNameRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "reference_name",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidReferenceNameRule) Name() string {
	return "aws_route53_health_check_invalid_reference_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidReferenceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53HealthCheckInvalidReferenceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidReferenceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidReferenceNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"reference_name must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"reference_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
