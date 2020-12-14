// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrLifecyclePolicyInvalidPolicyRule checks the pattern is valid
type AwsEcrLifecyclePolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsEcrLifecyclePolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsEcrLifecyclePolicyInvalidPolicyRule() *AwsEcrLifecyclePolicyInvalidPolicyRule {
	return &AwsEcrLifecyclePolicyInvalidPolicyRule{
		resourceType:  "aws_ecr_lifecycle_policy",
		attributeName: "policy",
		max:           30720,
		min:           100,
	}
}

// Name returns the rule name
func (r *AwsEcrLifecyclePolicyInvalidPolicyRule) Name() string {
	return "aws_ecr_lifecycle_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrLifecyclePolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrLifecyclePolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrLifecyclePolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrLifecyclePolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 30720 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 100 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
