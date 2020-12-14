// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsALBInvalidIPAddressTypeRule checks the pattern is valid
type AwsALBInvalidIPAddressTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsALBInvalidIPAddressTypeRule returns new rule with default attributes
func NewAwsALBInvalidIPAddressTypeRule() *AwsALBInvalidIPAddressTypeRule {
	return &AwsALBInvalidIPAddressTypeRule{
		resourceType:  "aws_alb",
		attributeName: "ip_address_type",
		enum: []string{
			"ipv4",
			"dualstack",
		},
	}
}

// Name returns the rule name
func (r *AwsALBInvalidIPAddressTypeRule) Name() string {
	return "aws_alb_invalid_ip_address_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsALBInvalidIPAddressTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsALBInvalidIPAddressTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsALBInvalidIPAddressTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsALBInvalidIPAddressTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as ip_address_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
