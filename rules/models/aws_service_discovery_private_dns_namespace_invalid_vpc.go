// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule checks the pattern is valid
type AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule returns new rule with default attributes
func NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule() *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule {
	return &AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule{
		resourceType:  "aws_service_discovery_private_dns_namespace",
		attributeName: "vpc",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule) Name() string {
	return "aws_service_discovery_private_dns_namespace_invalid_vpc"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"vpc must be 64 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
