// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEksClusterInvalidNameRule checks the pattern is valid
type AwsEksClusterInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEksClusterInvalidNameRule returns new rule with default attributes
func NewAwsEksClusterInvalidNameRule() *AwsEksClusterInvalidNameRule {
	return &AwsEksClusterInvalidNameRule{
		resourceType:  "aws_eks_cluster",
		attributeName: "name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[0-9A-Za-z][A-Za-z0-9\-_]*`),
	}
}

// Name returns the rule name
func (r *AwsEksClusterInvalidNameRule) Name() string {
	return "aws_eks_cluster_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEksClusterInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEksClusterInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEksClusterInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEksClusterInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 100 characters or less",
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9A-Za-z][A-Za-z0-9\-_]*`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
