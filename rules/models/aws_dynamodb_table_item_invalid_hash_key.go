// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDynamoDBTableItemInvalidHashKeyRule checks the pattern is valid
type AwsDynamoDBTableItemInvalidHashKeyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBTableItemInvalidHashKeyRule returns new rule with default attributes
func NewAwsDynamoDBTableItemInvalidHashKeyRule() *AwsDynamoDBTableItemInvalidHashKeyRule {
	return &AwsDynamoDBTableItemInvalidHashKeyRule{
		resourceType:  "aws_dynamodb_table_item",
		attributeName: "hash_key",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Name() string {
	return "aws_dynamodb_table_item_invalid_hash_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"hash_key must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"hash_key must be 1 characters or higher",
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
