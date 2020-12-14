// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppsyncDatasourceInvalidTypeRule checks the pattern is valid
type AwsAppsyncDatasourceInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppsyncDatasourceInvalidTypeRule returns new rule with default attributes
func NewAwsAppsyncDatasourceInvalidTypeRule() *AwsAppsyncDatasourceInvalidTypeRule {
	return &AwsAppsyncDatasourceInvalidTypeRule{
		resourceType:  "aws_appsync_datasource",
		attributeName: "type",
		enum: []string{
			"AWS_LAMBDA",
			"AMAZON_DYNAMODB",
			"AMAZON_ELASTICSEARCH",
			"NONE",
			"HTTP",
			"RELATIONAL_DATABASE",
		},
	}
}

// Name returns the rule name
func (r *AwsAppsyncDatasourceInvalidTypeRule) Name() string {
	return "aws_appsync_datasource_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncDatasourceInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncDatasourceInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncDatasourceInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncDatasourceInvalidTypeRule) Check(runner tflint.Runner) error {
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
