// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmAssociationInvalidComplianceSeverityRule checks the pattern is valid
type AwsSsmAssociationInvalidComplianceSeverityRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmAssociationInvalidComplianceSeverityRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidComplianceSeverityRule() *AwsSsmAssociationInvalidComplianceSeverityRule {
	return &AwsSsmAssociationInvalidComplianceSeverityRule{
		resourceType:  "aws_ssm_association",
		attributeName: "compliance_severity",
		enum: []string{
			"CRITICAL",
			"HIGH",
			"MEDIUM",
			"LOW",
			"UNSPECIFIED",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Name() string {
	return "aws_ssm_association_invalid_compliance_severity"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as compliance_severity`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
