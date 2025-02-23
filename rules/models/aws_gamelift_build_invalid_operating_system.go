// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGameliftBuildInvalidOperatingSystemRule checks the pattern is valid
type AwsGameliftBuildInvalidOperatingSystemRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGameliftBuildInvalidOperatingSystemRule returns new rule with default attributes
func NewAwsGameliftBuildInvalidOperatingSystemRule() *AwsGameliftBuildInvalidOperatingSystemRule {
	return &AwsGameliftBuildInvalidOperatingSystemRule{
		resourceType:  "aws_gamelift_build",
		attributeName: "operating_system",
		enum: []string{
			"WINDOWS_2012",
			"AMAZON_LINUX",
			"AMAZON_LINUX_2",
		},
	}
}

// Name returns the rule name
func (r *AwsGameliftBuildInvalidOperatingSystemRule) Name() string {
	return "aws_gamelift_build_invalid_operating_system"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftBuildInvalidOperatingSystemRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftBuildInvalidOperatingSystemRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftBuildInvalidOperatingSystemRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftBuildInvalidOperatingSystemRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as operating_system`, truncateLongMessage(val)),
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
