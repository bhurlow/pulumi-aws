// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource attaches a security group to an Elastic Network Interface (ENI).
// It can be used to attach a security group to any existing ENI, be it a
// secondary ENI or one attached as the primary interface on an instance.
// 
// > **NOTE on instances, interfaces, and security groups:** This provider currently
// provides the capability to assign security groups via the [`ec2.Instance`][1]
// and the [`ec2.NetworkInterface`][2] resources. Using this resource in
// conjunction with security groups provided in-line in those resources will cause
// conflicts, and will lead to spurious diffs and undefined behavior - please use
// one or the other.
// 
// [1]: https://www.terraform.io/docs/providers/aws/d/instance.html
// [2]: https://www.terraform.io/docs/providers/aws/r/network_interface.html
// 
// ## Output Reference
// 
// There are no outputs for this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_interface_sg_attachment.html.markdown.
type NetworkInterfaceSecurityGroupAttachment struct {
	s *pulumi.ResourceState
}

// NewNetworkInterfaceSecurityGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceSecurityGroupAttachmentArgs, opts ...pulumi.ResourceOpt) (*NetworkInterfaceSecurityGroupAttachment, error) {
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil || args.SecurityGroupId == nil {
		return nil, errors.New("missing required argument 'SecurityGroupId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["networkInterfaceId"] = nil
		inputs["securityGroupId"] = nil
	} else {
		inputs["networkInterfaceId"] = args.NetworkInterfaceId
		inputs["securityGroupId"] = args.SecurityGroupId
	}
	s, err := ctx.RegisterResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkInterfaceSecurityGroupAttachment{s: s}, nil
}

// GetNetworkInterfaceSecurityGroupAttachment gets an existing NetworkInterfaceSecurityGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkInterfaceSecurityGroupAttachmentState, opts ...pulumi.ResourceOpt) (*NetworkInterfaceSecurityGroupAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["networkInterfaceId"] = state.NetworkInterfaceId
		inputs["securityGroupId"] = state.SecurityGroupId
	}
	s, err := ctx.ReadResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkInterfaceSecurityGroupAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkInterfaceSecurityGroupAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkInterfaceSecurityGroupAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the network interface to attach to.
func (r *NetworkInterfaceSecurityGroupAttachment) NetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkInterfaceId"])
}

// The ID of the security group.
func (r *NetworkInterfaceSecurityGroupAttachment) SecurityGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["securityGroupId"])
}

// Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAttachment resources.
type NetworkInterfaceSecurityGroupAttachmentState struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId interface{}
	// The ID of the security group.
	SecurityGroupId interface{}
}

// The set of arguments for constructing a NetworkInterfaceSecurityGroupAttachment resource.
type NetworkInterfaceSecurityGroupAttachmentArgs struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId interface{}
	// The ID of the security group.
	SecurityGroupId interface{}
}
