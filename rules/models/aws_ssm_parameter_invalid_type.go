// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmParameterInvalidTypeRule checks the pattern is valid
type AwsSsmParameterInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmParameterInvalidTypeRule returns new rule with default attributes
func NewAwsSsmParameterInvalidTypeRule() *AwsSsmParameterInvalidTypeRule {
	return &AwsSsmParameterInvalidTypeRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "type",
		enum: []string{
			"String",
			"StringList",
			"SecureString",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidTypeRule) Name() string {
	return "aws_ssm_parameter_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
