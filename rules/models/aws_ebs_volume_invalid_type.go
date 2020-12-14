// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEbsVolumeInvalidTypeRule checks the pattern is valid
type AwsEbsVolumeInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEbsVolumeInvalidTypeRule returns new rule with default attributes
func NewAwsEbsVolumeInvalidTypeRule() *AwsEbsVolumeInvalidTypeRule {
	return &AwsEbsVolumeInvalidTypeRule{
		resourceType:  "aws_ebs_volume",
		attributeName: "type",
		enum: []string{
			"standard",
			"io1",
			"io2",
			"gp2",
			"sc1",
			"st1",
		},
	}
}

// Name returns the rule name
func (r *AwsEbsVolumeInvalidTypeRule) Name() string {
	return "aws_ebs_volume_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEbsVolumeInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEbsVolumeInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEbsVolumeInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEbsVolumeInvalidTypeRule) Check(runner tflint.Runner) error {
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
