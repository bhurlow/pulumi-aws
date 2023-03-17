// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on an EC2 Transit Gateway's attachment to a resource.
func GetAttachment(ctx *pulumi.Context, args *GetAttachmentArgs, opts ...pulumi.InvokeOption) (*GetAttachmentResult, error) {
	var rv GetAttachmentResult
	err := ctx.Invoke("aws:ec2transitgateway/getAttachment:getAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAttachment.
type GetAttachmentArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters []GetAttachmentFilter `pulumi:"filters"`
	// Key-value tags for the attachment.
	Tags map[string]string `pulumi:"tags"`
	// ID of the attachment.
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
}

// A collection of values returned by getAttachment.
type GetAttachmentResult struct {
	// ARN of the attachment.
	Arn string `pulumi:"arn"`
	// The state of the association (see [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachmentAssociation.html) for valid values).
	AssociationState string `pulumi:"associationState"`
	// The ID of the route table for the transit gateway.
	AssociationTransitGatewayRouteTableId string                `pulumi:"associationTransitGatewayRouteTableId"`
	Filters                               []GetAttachmentFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the resource.
	ResourceId string `pulumi:"resourceId"`
	// ID of the AWS account that owns the resource.
	ResourceOwnerId string `pulumi:"resourceOwnerId"`
	// Resource type.
	ResourceType string `pulumi:"resourceType"`
	// Attachment state.
	State string `pulumi:"state"`
	// Key-value tags for the attachment.
	Tags                       map[string]string `pulumi:"tags"`
	TransitGatewayAttachmentId string            `pulumi:"transitGatewayAttachmentId"`
	// ID of the transit gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
	// The ID of the AWS account that owns the transit gateway.
	TransitGatewayOwnerId string `pulumi:"transitGatewayOwnerId"`
}

func GetAttachmentOutput(ctx *pulumi.Context, args GetAttachmentOutputArgs, opts ...pulumi.InvokeOption) GetAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAttachmentResult, error) {
			args := v.(GetAttachmentArgs)
			r, err := GetAttachment(ctx, &args, opts...)
			var s GetAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAttachmentResultOutput)
}

// A collection of arguments for invoking getAttachment.
type GetAttachmentOutputArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters GetAttachmentFilterArrayInput `pulumi:"filters"`
	// Key-value tags for the attachment.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// ID of the attachment.
	TransitGatewayAttachmentId pulumi.StringPtrInput `pulumi:"transitGatewayAttachmentId"`
}

func (GetAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachmentArgs)(nil)).Elem()
}

// A collection of values returned by getAttachment.
type GetAttachmentResultOutput struct{ *pulumi.OutputState }

func (GetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachmentResult)(nil)).Elem()
}

func (o GetAttachmentResultOutput) ToGetAttachmentResultOutput() GetAttachmentResultOutput {
	return o
}

func (o GetAttachmentResultOutput) ToGetAttachmentResultOutputWithContext(ctx context.Context) GetAttachmentResultOutput {
	return o
}

// ARN of the attachment.
func (o GetAttachmentResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The state of the association (see [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachmentAssociation.html) for valid values).
func (o GetAttachmentResultOutput) AssociationState() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.AssociationState }).(pulumi.StringOutput)
}

// The ID of the route table for the transit gateway.
func (o GetAttachmentResultOutput) AssociationTransitGatewayRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.AssociationTransitGatewayRouteTableId }).(pulumi.StringOutput)
}

func (o GetAttachmentResultOutput) Filters() GetAttachmentFilterArrayOutput {
	return o.ApplyT(func(v GetAttachmentResult) []GetAttachmentFilter { return v.Filters }).(GetAttachmentFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAttachmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the resource.
func (o GetAttachmentResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// ID of the AWS account that owns the resource.
func (o GetAttachmentResultOutput) ResourceOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.ResourceOwnerId }).(pulumi.StringOutput)
}

// Resource type.
func (o GetAttachmentResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

// Attachment state.
func (o GetAttachmentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.State }).(pulumi.StringOutput)
}

// Key-value tags for the attachment.
func (o GetAttachmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAttachmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetAttachmentResultOutput) TransitGatewayAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.TransitGatewayAttachmentId }).(pulumi.StringOutput)
}

// ID of the transit gateway.
func (o GetAttachmentResultOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.TransitGatewayId }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the transit gateway.
func (o GetAttachmentResultOutput) TransitGatewayOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachmentResult) string { return v.TransitGatewayOwnerId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAttachmentResultOutput{})
}
