// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ARN of a State Machine in AWS Step
// Function (SFN). By using this data source, you can reference a
// state machine without having to hard code the ARNs as input.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sfn.LookupStateMachine(ctx, &sfn.LookupStateMachineArgs{
//				Name: "an_example_sfn_name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupStateMachine(ctx *pulumi.Context, args *LookupStateMachineArgs, opts ...pulumi.InvokeOption) (*LookupStateMachineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStateMachineResult
	err := ctx.Invoke("aws:sfn/getStateMachine:getStateMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStateMachine.
type LookupStateMachineArgs struct {
	// Friendly name of the state machine to match.
	Name string `pulumi:"name"`
}

// A collection of values returned by getStateMachine.
type LookupStateMachineResult struct {
	// Set to the arn of the state function.
	Arn string `pulumi:"arn"`
	// Date the state machine was created.
	CreationDate string `pulumi:"creationDate"`
	// Set to the state machine definition.
	Definition  string `pulumi:"definition"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The revision identifier for the state machine.
	RevisionId string `pulumi:"revisionId"`
	// Set to the roleArn used by the state function.
	RoleArn string `pulumi:"roleArn"`
	// Set to the current status of the state machine.
	Status string `pulumi:"status"`
}

func LookupStateMachineOutput(ctx *pulumi.Context, args LookupStateMachineOutputArgs, opts ...pulumi.InvokeOption) LookupStateMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStateMachineResult, error) {
			args := v.(LookupStateMachineArgs)
			r, err := LookupStateMachine(ctx, &args, opts...)
			var s LookupStateMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStateMachineResultOutput)
}

// A collection of arguments for invoking getStateMachine.
type LookupStateMachineOutputArgs struct {
	// Friendly name of the state machine to match.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupStateMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineArgs)(nil)).Elem()
}

// A collection of values returned by getStateMachine.
type LookupStateMachineResultOutput struct{ *pulumi.OutputState }

func (LookupStateMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineResult)(nil)).Elem()
}

func (o LookupStateMachineResultOutput) ToLookupStateMachineResultOutput() LookupStateMachineResultOutput {
	return o
}

func (o LookupStateMachineResultOutput) ToLookupStateMachineResultOutputWithContext(ctx context.Context) LookupStateMachineResultOutput {
	return o
}

// Set to the arn of the state function.
func (o LookupStateMachineResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Date the state machine was created.
func (o LookupStateMachineResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Set to the state machine definition.
func (o LookupStateMachineResultOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Definition }).(pulumi.StringOutput)
}

func (o LookupStateMachineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStateMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStateMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The revision identifier for the state machine.
func (o LookupStateMachineResultOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.RevisionId }).(pulumi.StringOutput)
}

// Set to the roleArn used by the state function.
func (o LookupStateMachineResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// Set to the current status of the state machine.
func (o LookupStateMachineResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStateMachineResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStateMachineResultOutput{})
}
