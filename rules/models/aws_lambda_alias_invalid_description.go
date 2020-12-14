// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaAliasInvalidDescriptionRule checks the pattern is valid
type AwsLambdaAliasInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsLambdaAliasInvalidDescriptionRule returns new rule with default attributes
func NewAwsLambdaAliasInvalidDescriptionRule() *AwsLambdaAliasInvalidDescriptionRule {
	return &AwsLambdaAliasInvalidDescriptionRule{
		resourceType:  "aws_lambda_alias",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsLambdaAliasInvalidDescriptionRule) Name() string {
	return "aws_lambda_alias_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaAliasInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaAliasInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaAliasInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaAliasInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
