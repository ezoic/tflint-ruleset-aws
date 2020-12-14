// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSesReceiptRuleInvalidTLSPolicyRule checks the pattern is valid
type AwsSesReceiptRuleInvalidTLSPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSesReceiptRuleInvalidTLSPolicyRule returns new rule with default attributes
func NewAwsSesReceiptRuleInvalidTLSPolicyRule() *AwsSesReceiptRuleInvalidTLSPolicyRule {
	return &AwsSesReceiptRuleInvalidTLSPolicyRule{
		resourceType:  "aws_ses_receipt_rule",
		attributeName: "tls_policy",
		enum: []string{
			"Require",
			"Optional",
		},
	}
}

// Name returns the rule name
func (r *AwsSesReceiptRuleInvalidTLSPolicyRule) Name() string {
	return "aws_ses_receipt_rule_invalid_tls_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSesReceiptRuleInvalidTLSPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSesReceiptRuleInvalidTLSPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSesReceiptRuleInvalidTLSPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSesReceiptRuleInvalidTLSPolicyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as tls_policy`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
