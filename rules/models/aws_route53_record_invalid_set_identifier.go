// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53RecordInvalidSetIdentifierRule checks the pattern is valid
type AwsRoute53RecordInvalidSetIdentifierRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53RecordInvalidSetIdentifierRule returns new rule with default attributes
func NewAwsRoute53RecordInvalidSetIdentifierRule() *AwsRoute53RecordInvalidSetIdentifierRule {
	return &AwsRoute53RecordInvalidSetIdentifierRule{
		resourceType:  "aws_route53_record",
		attributeName: "set_identifier",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53RecordInvalidSetIdentifierRule) Name() string {
	return "aws_route53_record_invalid_set_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53RecordInvalidSetIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53RecordInvalidSetIdentifierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53RecordInvalidSetIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53RecordInvalidSetIdentifierRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"set_identifier must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"set_identifier must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
