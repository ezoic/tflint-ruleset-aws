// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLightsailKeyPairInvalidNameRule checks the pattern is valid
type AwsLightsailKeyPairInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailKeyPairInvalidNameRule returns new rule with default attributes
func NewAwsLightsailKeyPairInvalidNameRule() *AwsLightsailKeyPairInvalidNameRule {
	return &AwsLightsailKeyPairInvalidNameRule{
		resourceType:  "aws_lightsail_key_pair",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^\w[\w\-]*\w$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailKeyPairInvalidNameRule) Name() string {
	return "aws_lightsail_key_pair_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailKeyPairInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailKeyPairInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailKeyPairInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailKeyPairInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\w[\w\-]*\w$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
