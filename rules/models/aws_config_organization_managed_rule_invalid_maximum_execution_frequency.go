// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule() *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule {
	return &AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "maximum_execution_frequency",
		enum: []string{
			"One_Hour",
			"Three_Hours",
			"Six_Hours",
			"Twelve_Hours",
			"TwentyFour_Hours",
		},
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_maximum_execution_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as maximum_execution_frequency`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
