// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloud9EnvironmentEc2InvalidOwnerArnRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidOwnerArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloud9EnvironmentEc2InvalidOwnerArnRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidOwnerArnRule() *AwsCloud9EnvironmentEc2InvalidOwnerArnRule {
	return &AwsCloud9EnvironmentEc2InvalidOwnerArnRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "owner_arn",
		pattern:       regexp.MustCompile(`^arn:aws:(iam|sts)::\d+:(root|(user\/[\w+=/:,.@-]{1,64}|federated-user\/[\w+=/:,.@-]{2,32}|assumed-role\/[\w+=:,.@-]{1,64}\/[\w+=,.@-]{1,64}))$`),
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidOwnerArnRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_owner_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidOwnerArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidOwnerArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidOwnerArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidOwnerArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws:(iam|sts)::\d+:(root|(user\/[\w+=/:,.@-]{1,64}|federated-user\/[\w+=/:,.@-]{2,32}|assumed-role\/[\w+=:,.@-]{1,64}\/[\w+=,.@-]{1,64}))$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
