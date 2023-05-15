// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeBuild Project resource. See also the `codebuild.Webhook` resource, which manages the webhook to the source (e.g., the "rebuild every time a code change is pushed" option in the CodeBuild web console).
//
// ## Import
//
// CodeBuild Project can be imported using the `name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:codebuild/project:Project name project-name
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// ARN of the CodeBuild project.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block. Detailed below.
	Artifacts ProjectArtifactsOutput `pulumi:"artifacts"`
	// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
	BadgeEnabled pulumi.BoolPtrOutput `pulumi:"badgeEnabled"`
	// URL of the build badge when `badgeEnabled` is enabled.
	BadgeUrl pulumi.StringOutput `pulumi:"badgeUrl"`
	// Defines the batch build options for the project.
	BuildBatchConfig ProjectBuildBatchConfigPtrOutput `pulumi:"buildBatchConfig"`
	// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
	BuildTimeout pulumi.IntPtrOutput `pulumi:"buildTimeout"`
	// Configuration block. Detailed below.
	Cache ProjectCachePtrOutput `pulumi:"cache"`
	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit pulumi.IntPtrOutput `pulumi:"concurrentBuildLimit"`
	// Short description of the project.
	Description pulumi.StringOutput `pulumi:"description"`
	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey pulumi.StringOutput `pulumi:"encryptionKey"`
	// Configuration block. Detailed below.
	Environment ProjectEnvironmentOutput `pulumi:"environment"`
	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations ProjectFileSystemLocationArrayOutput `pulumi:"fileSystemLocations"`
	// Configuration block. Detailed below.
	LogsConfig ProjectLogsConfigPtrOutput `pulumi:"logsConfig"`
	// Project's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility pulumi.StringPtrOutput `pulumi:"projectVisibility"`
	// The project identifier used with the public build APIs.
	PublicProjectAlias pulumi.StringOutput `pulumi:"publicProjectAlias"`
	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
	QueuedTimeout pulumi.IntPtrOutput `pulumi:"queuedTimeout"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole pulumi.StringPtrOutput `pulumi:"resourceAccessRole"`
	// Configuration block. Detailed below.
	SecondaryArtifacts ProjectSecondaryArtifactArrayOutput `pulumi:"secondaryArtifacts"`
	// Configuration block. Detailed below.
	SecondarySourceVersions ProjectSecondarySourceVersionArrayOutput `pulumi:"secondarySourceVersions"`
	// Configuration block. Detailed below.
	SecondarySources ProjectSecondarySourceArrayOutput `pulumi:"secondarySources"`
	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole pulumi.StringOutput `pulumi:"serviceRole"`
	// Configuration block. Detailed below.
	Source ProjectSourceOutput `pulumi:"source"`
	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion pulumi.StringPtrOutput `pulumi:"sourceVersion"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block. Detailed below.
	VpcConfig ProjectVpcConfigPtrOutput `pulumi:"vpcConfig"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Artifacts == nil {
		return nil, errors.New("invalid value for required argument 'Artifacts'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.ServiceRole == nil {
		return nil, errors.New("invalid value for required argument 'ServiceRole'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource Project
	err := ctx.RegisterResource("aws:codebuild/project:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:codebuild/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// ARN of the CodeBuild project.
	Arn *string `pulumi:"arn"`
	// Configuration block. Detailed below.
	Artifacts *ProjectArtifacts `pulumi:"artifacts"`
	// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
	BadgeEnabled *bool `pulumi:"badgeEnabled"`
	// URL of the build badge when `badgeEnabled` is enabled.
	BadgeUrl *string `pulumi:"badgeUrl"`
	// Defines the batch build options for the project.
	BuildBatchConfig *ProjectBuildBatchConfig `pulumi:"buildBatchConfig"`
	// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
	BuildTimeout *int `pulumi:"buildTimeout"`
	// Configuration block. Detailed below.
	Cache *ProjectCache `pulumi:"cache"`
	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit *int `pulumi:"concurrentBuildLimit"`
	// Short description of the project.
	Description *string `pulumi:"description"`
	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Configuration block. Detailed below.
	Environment *ProjectEnvironment `pulumi:"environment"`
	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations []ProjectFileSystemLocation `pulumi:"fileSystemLocations"`
	// Configuration block. Detailed below.
	LogsConfig *ProjectLogsConfig `pulumi:"logsConfig"`
	// Project's name.
	Name *string `pulumi:"name"`
	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility *string `pulumi:"projectVisibility"`
	// The project identifier used with the public build APIs.
	PublicProjectAlias *string `pulumi:"publicProjectAlias"`
	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
	QueuedTimeout *int `pulumi:"queuedTimeout"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole *string `pulumi:"resourceAccessRole"`
	// Configuration block. Detailed below.
	SecondaryArtifacts []ProjectSecondaryArtifact `pulumi:"secondaryArtifacts"`
	// Configuration block. Detailed below.
	SecondarySourceVersions []ProjectSecondarySourceVersion `pulumi:"secondarySourceVersions"`
	// Configuration block. Detailed below.
	SecondarySources []ProjectSecondarySource `pulumi:"secondarySources"`
	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole *string `pulumi:"serviceRole"`
	// Configuration block. Detailed below.
	Source *ProjectSource `pulumi:"source"`
	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion *string `pulumi:"sourceVersion"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block. Detailed below.
	VpcConfig *ProjectVpcConfig `pulumi:"vpcConfig"`
}

type ProjectState struct {
	// ARN of the CodeBuild project.
	Arn pulumi.StringPtrInput
	// Configuration block. Detailed below.
	Artifacts ProjectArtifactsPtrInput
	// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
	BadgeEnabled pulumi.BoolPtrInput
	// URL of the build badge when `badgeEnabled` is enabled.
	BadgeUrl pulumi.StringPtrInput
	// Defines the batch build options for the project.
	BuildBatchConfig ProjectBuildBatchConfigPtrInput
	// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
	BuildTimeout pulumi.IntPtrInput
	// Configuration block. Detailed below.
	Cache ProjectCachePtrInput
	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit pulumi.IntPtrInput
	// Short description of the project.
	Description pulumi.StringPtrInput
	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey pulumi.StringPtrInput
	// Configuration block. Detailed below.
	Environment ProjectEnvironmentPtrInput
	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations ProjectFileSystemLocationArrayInput
	// Configuration block. Detailed below.
	LogsConfig ProjectLogsConfigPtrInput
	// Project's name.
	Name pulumi.StringPtrInput
	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility pulumi.StringPtrInput
	// The project identifier used with the public build APIs.
	PublicProjectAlias pulumi.StringPtrInput
	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
	QueuedTimeout pulumi.IntPtrInput
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole pulumi.StringPtrInput
	// Configuration block. Detailed below.
	SecondaryArtifacts ProjectSecondaryArtifactArrayInput
	// Configuration block. Detailed below.
	SecondarySourceVersions ProjectSecondarySourceVersionArrayInput
	// Configuration block. Detailed below.
	SecondarySources ProjectSecondarySourceArrayInput
	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole pulumi.StringPtrInput
	// Configuration block. Detailed below.
	Source ProjectSourcePtrInput
	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Configuration block. Detailed below.
	VpcConfig ProjectVpcConfigPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Configuration block. Detailed below.
	Artifacts ProjectArtifacts `pulumi:"artifacts"`
	// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
	BadgeEnabled *bool `pulumi:"badgeEnabled"`
	// Defines the batch build options for the project.
	BuildBatchConfig *ProjectBuildBatchConfig `pulumi:"buildBatchConfig"`
	// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
	BuildTimeout *int `pulumi:"buildTimeout"`
	// Configuration block. Detailed below.
	Cache *ProjectCache `pulumi:"cache"`
	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit *int `pulumi:"concurrentBuildLimit"`
	// Short description of the project.
	Description *string `pulumi:"description"`
	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Configuration block. Detailed below.
	Environment ProjectEnvironment `pulumi:"environment"`
	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations []ProjectFileSystemLocation `pulumi:"fileSystemLocations"`
	// Configuration block. Detailed below.
	LogsConfig *ProjectLogsConfig `pulumi:"logsConfig"`
	// Project's name.
	Name *string `pulumi:"name"`
	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility *string `pulumi:"projectVisibility"`
	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
	QueuedTimeout *int `pulumi:"queuedTimeout"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole *string `pulumi:"resourceAccessRole"`
	// Configuration block. Detailed below.
	SecondaryArtifacts []ProjectSecondaryArtifact `pulumi:"secondaryArtifacts"`
	// Configuration block. Detailed below.
	SecondarySourceVersions []ProjectSecondarySourceVersion `pulumi:"secondarySourceVersions"`
	// Configuration block. Detailed below.
	SecondarySources []ProjectSecondarySource `pulumi:"secondarySources"`
	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole string `pulumi:"serviceRole"`
	// Configuration block. Detailed below.
	Source ProjectSource `pulumi:"source"`
	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion *string `pulumi:"sourceVersion"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block. Detailed below.
	VpcConfig *ProjectVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Configuration block. Detailed below.
	Artifacts ProjectArtifactsInput
	// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
	BadgeEnabled pulumi.BoolPtrInput
	// Defines the batch build options for the project.
	BuildBatchConfig ProjectBuildBatchConfigPtrInput
	// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
	BuildTimeout pulumi.IntPtrInput
	// Configuration block. Detailed below.
	Cache ProjectCachePtrInput
	// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
	ConcurrentBuildLimit pulumi.IntPtrInput
	// Short description of the project.
	Description pulumi.StringPtrInput
	// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
	EncryptionKey pulumi.StringPtrInput
	// Configuration block. Detailed below.
	Environment ProjectEnvironmentInput
	// A set of file system locations to mount inside the build. File system locations are documented below.
	FileSystemLocations ProjectFileSystemLocationArrayInput
	// Configuration block. Detailed below.
	LogsConfig ProjectLogsConfigPtrInput
	// Project's name.
	Name pulumi.StringPtrInput
	// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
	ProjectVisibility pulumi.StringPtrInput
	// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
	QueuedTimeout pulumi.IntPtrInput
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole pulumi.StringPtrInput
	// Configuration block. Detailed below.
	SecondaryArtifacts ProjectSecondaryArtifactArrayInput
	// Configuration block. Detailed below.
	SecondarySourceVersions ProjectSecondarySourceVersionArrayInput
	// Configuration block. Detailed below.
	SecondarySources ProjectSecondarySourceArrayInput
	// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole pulumi.StringInput
	// Configuration block. Detailed below.
	Source ProjectSourceInput
	// Version of the build input to be built for this project. If not specified, the latest version is used.
	SourceVersion pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Configuration block. Detailed below.
	VpcConfig ProjectVpcConfigPtrInput
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
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
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
//	ProjectMap{ "key": ProjectArgs{...} }
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
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// ARN of the CodeBuild project.
func (o ProjectOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) Artifacts() ProjectArtifactsOutput {
	return o.ApplyT(func(v *Project) ProjectArtifactsOutput { return v.Artifacts }).(ProjectArtifactsOutput)
}

// Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
func (o ProjectOutput) BadgeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.BadgeEnabled }).(pulumi.BoolPtrOutput)
}

// URL of the build badge when `badgeEnabled` is enabled.
func (o ProjectOutput) BadgeUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.BadgeUrl }).(pulumi.StringOutput)
}

// Defines the batch build options for the project.
func (o ProjectOutput) BuildBatchConfig() ProjectBuildBatchConfigPtrOutput {
	return o.ApplyT(func(v *Project) ProjectBuildBatchConfigPtrOutput { return v.BuildBatchConfig }).(ProjectBuildBatchConfigPtrOutput)
}

// Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
func (o ProjectOutput) BuildTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.BuildTimeout }).(pulumi.IntPtrOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) Cache() ProjectCachePtrOutput {
	return o.ApplyT(func(v *Project) ProjectCachePtrOutput { return v.Cache }).(ProjectCachePtrOutput)
}

// Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
func (o ProjectOutput) ConcurrentBuildLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.ConcurrentBuildLimit }).(pulumi.IntPtrOutput)
}

// Short description of the project.
func (o ProjectOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
func (o ProjectOutput) EncryptionKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.EncryptionKey }).(pulumi.StringOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) Environment() ProjectEnvironmentOutput {
	return o.ApplyT(func(v *Project) ProjectEnvironmentOutput { return v.Environment }).(ProjectEnvironmentOutput)
}

// A set of file system locations to mount inside the build. File system locations are documented below.
func (o ProjectOutput) FileSystemLocations() ProjectFileSystemLocationArrayOutput {
	return o.ApplyT(func(v *Project) ProjectFileSystemLocationArrayOutput { return v.FileSystemLocations }).(ProjectFileSystemLocationArrayOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) LogsConfig() ProjectLogsConfigPtrOutput {
	return o.ApplyT(func(v *Project) ProjectLogsConfigPtrOutput { return v.LogsConfig }).(ProjectLogsConfigPtrOutput)
}

// Project's name.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the visibility of the project's builds. Possible values are: `PUBLIC_READ` and `PRIVATE`. Default value is `PRIVATE`.
func (o ProjectOutput) ProjectVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ProjectVisibility }).(pulumi.StringPtrOutput)
}

// The project identifier used with the public build APIs.
func (o ProjectOutput) PublicProjectAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.PublicProjectAlias }).(pulumi.StringOutput)
}

// Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
func (o ProjectOutput) QueuedTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.QueuedTimeout }).(pulumi.IntPtrOutput)
}

// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
func (o ProjectOutput) ResourceAccessRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ResourceAccessRole }).(pulumi.StringPtrOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) SecondaryArtifacts() ProjectSecondaryArtifactArrayOutput {
	return o.ApplyT(func(v *Project) ProjectSecondaryArtifactArrayOutput { return v.SecondaryArtifacts }).(ProjectSecondaryArtifactArrayOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) SecondarySourceVersions() ProjectSecondarySourceVersionArrayOutput {
	return o.ApplyT(func(v *Project) ProjectSecondarySourceVersionArrayOutput { return v.SecondarySourceVersions }).(ProjectSecondarySourceVersionArrayOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) SecondarySources() ProjectSecondarySourceArrayOutput {
	return o.ApplyT(func(v *Project) ProjectSecondarySourceArrayOutput { return v.SecondarySources }).(ProjectSecondarySourceArrayOutput)
}

// Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
func (o ProjectOutput) ServiceRole() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ServiceRole }).(pulumi.StringOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) Source() ProjectSourceOutput {
	return o.ApplyT(func(v *Project) ProjectSourceOutput { return v.Source }).(ProjectSourceOutput)
}

// Version of the build input to be built for this project. If not specified, the latest version is used.
func (o ProjectOutput) SourceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.SourceVersion }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ProjectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ProjectOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Configuration block. Detailed below.
func (o ProjectOutput) VpcConfig() ProjectVpcConfigPtrOutput {
	return o.ApplyT(func(v *Project) ProjectVpcConfigPtrOutput { return v.VpcConfig }).(ProjectVpcConfigPtrOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
