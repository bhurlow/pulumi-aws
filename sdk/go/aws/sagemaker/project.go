// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sagemaker Project resource.
//
//  > Note: If you are trying to use Sagemaker projects with Sagemaker studio you will need to add a tag with the key `sagemaker:studio-visibility` with value `true`. For more on requirements to use projects and permission needed see [AWS Docs](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-projects-templates-custom.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewProject(ctx, "example", &sagemaker.ProjectArgs{
// 			ProjectName: pulumi.String("example"),
// 			ServiceCatalogProvisioningDetails: &sagemaker.ProjectServiceCatalogProvisioningDetailsArgs{
// 				ProductId: pulumi.Any(aws_servicecatalog_product.Example.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Sagemaker Projects can be imported using the `project_name`, e.g.,
//
// ```sh
//  $ pulumi import aws:sagemaker/project:Project example example
// ```
type Project struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this Project.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description for the project.
	ProjectDescription pulumi.StringPtrOutput `pulumi:"projectDescription"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The name of the Project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails ProjectServiceCatalogProvisioningDetailsOutput `pulumi:"serviceCatalogProvisioningDetails"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ServiceCatalogProvisioningDetails == nil {
		return nil, errors.New("invalid value for required argument 'ServiceCatalogProvisioningDetails'")
	}
	var resource Project
	err := ctx.RegisterResource("aws:sagemaker/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("aws:sagemaker/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Project.
	Arn *string `pulumi:"arn"`
	// A description for the project.
	ProjectDescription *string `pulumi:"projectDescription"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The name of the Project.
	ProjectName *string `pulumi:"projectName"`
	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails *ProjectServiceCatalogProvisioningDetails `pulumi:"serviceCatalogProvisioningDetails"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ProjectState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Project.
	Arn pulumi.StringPtrInput
	// A description for the project.
	ProjectDescription pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The name of the Project.
	ProjectName pulumi.StringPtrInput
	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails ProjectServiceCatalogProvisioningDetailsPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// A description for the project.
	ProjectDescription *string `pulumi:"projectDescription"`
	// The name of the Project.
	ProjectName string `pulumi:"projectName"`
	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails ProjectServiceCatalogProvisioningDetails `pulumi:"serviceCatalogProvisioningDetails"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// A description for the project.
	ProjectDescription pulumi.StringPtrInput
	// The name of the Project.
	ProjectName pulumi.StringInput
	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails ProjectServiceCatalogProvisioningDetailsInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil))
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

func (i *Project) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *Project) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

type ProjectPtrInput interface {
	pulumi.Input

	ToProjectPtrOutput() ProjectPtrOutput
	ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput
}

type projectPtrType ProjectArgs

func (*projectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (i *projectPtrType) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *projectPtrType) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//          ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//          ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil))
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o.ToProjectPtrOutputWithContext(context.Background())
}

func (o ProjectOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Project) *Project {
		return &v
	}).(ProjectPtrOutput)
}

type ProjectPtrOutput struct{ *pulumi.OutputState }

func (ProjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (o ProjectPtrOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) Elem() ProjectOutput {
	return o.ApplyT(func(v *Project) Project {
		if v != nil {
			return *v
		}
		var ret Project
		return ret
	}).(ProjectOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Project)(nil))
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Project {
		return vs[0].([]Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Project)(nil))
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Project {
		return vs[0].(map[string]Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPtrInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectPtrOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
