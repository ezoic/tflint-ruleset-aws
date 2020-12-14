// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppautoscalingPolicyInvalidServiceNamespaceRule checks the pattern is valid
type AwsAppautoscalingPolicyInvalidServiceNamespaceRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppautoscalingPolicyInvalidServiceNamespaceRule returns new rule with default attributes
func NewAwsAppautoscalingPolicyInvalidServiceNamespaceRule() *AwsAppautoscalingPolicyInvalidServiceNamespaceRule {
	return &AwsAppautoscalingPolicyInvalidServiceNamespaceRule{
		resourceType:  "aws_appautoscaling_policy",
		attributeName: "service_namespace",
		enum: []string{
			"ecs",
			"elasticmapreduce",
			"ec2",
			"appstream",
			"dynamodb",
			"rds",
			"sagemaker",
			"custom-resource",
			"comprehend",
			"lambda",
			"cassandra",
			"kafka",
		},
	}
}

// Name returns the rule name
func (r *AwsAppautoscalingPolicyInvalidServiceNamespaceRule) Name() string {
	return "aws_appautoscaling_policy_invalid_service_namespace"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppautoscalingPolicyInvalidServiceNamespaceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppautoscalingPolicyInvalidServiceNamespaceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppautoscalingPolicyInvalidServiceNamespaceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppautoscalingPolicyInvalidServiceNamespaceRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as service_namespace`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
