// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an SSM Document resource
//
// > **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
// or greater can update their content once created, see [SSM Schema Features](http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html#document-schemas-features). To update a document with an older schema version you must recreate the resource. Not all document types support a schema version of 2.0 or greater. Refer to [SSM document schema features and examples](https://docs.aws.amazon.com/systems-manager/latest/userguide/document-schemas-features.html) for information about which schema versions are supported for the respective `documentType`.
//
// ## Example Usage
// ### Create an ssm document in JSON format
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
//				Content: pulumi.String(`  {
//	    "schemaVersion": "1.2",
//	    "description": "Check ip configuration of a Linux instance.",
//	    "parameters": {
//
//	    },
//	    "runtimeConfig": {
//	      "aws:runShellScript": {
//	        "properties": [
//	          {
//	            "id": "0.aws:runShellScript",
//	            "runCommand": ["ifconfig"]
//	          }
//	        ]
//	      }
//	    }
//	  }
//
// `),
//
//				DocumentType: pulumi.String("Command"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create an ssm document in YAML format
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
//				Content: pulumi.String(`schemaVersion: '1.2'
//
// description: Check ip configuration of a Linux instance.
// parameters: {}
// runtimeConfig:
//
//	'aws:runShellScript':
//	  properties:
//	    - id: '0.aws:runShellScript'
//	      runCommand:
//	        - ifconfig
//
// `),
//
//				DocumentFormat: pulumi.String("YAML"),
//				DocumentType:   pulumi.String("Command"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Permissions
//
// The permissions attribute specifies how you want to share the document. If you share a document privately,
// you must specify the AWS user account IDs for those people who can use the document. If you share a document
// publicly, you must specify All as the account ID.
//
// The permissions mapping supports the following:
//
// * `type` - The permission type for the document. The permission type can be `Share`.
// * `accountIds` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.
//
// ## Import
//
// Using `pulumi import`, import SSM Documents using the name. For example:
//
// ```sh
//
//	$ pulumi import aws:ssm/document:Document example example
//
// ```
//
//	The `attachments_source` argument does not have an SSM API method for reading the attachment information detail after creation. If the argument is set in the TODO configuration on an imported resource, TODO will always show a difference. To workaround this behavior, either omit the argument from the TODO configuration or use `ignore_changes` to hide the difference. For example:
type Document struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayOutput `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content pulumi.StringOutput `pulumi:"content"`
	// The date the document was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The default version of the document.
	DefaultVersion pulumi.StringOutput `pulumi:"defaultVersion"`
	// The description of the document.
	Description pulumi.StringOutput `pulumi:"description"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrOutput `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringOutput `pulumi:"documentType"`
	// The document version.
	DocumentVersion pulumi.StringOutput `pulumi:"documentVersion"`
	// The sha1 or sha256 of the document content
	Hash pulumi.StringOutput `pulumi:"hash"`
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType pulumi.StringOutput `pulumi:"hashType"`
	// The latest version of the document.
	LatestVersion pulumi.StringOutput `pulumi:"latestVersion"`
	// The name of the document.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS user account of the person who created the document.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The parameters that are available to this document.
	Parameters DocumentParameterArrayOutput `pulumi:"parameters"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes pulumi.StringArrayOutput `pulumi:"platformTypes"`
	// The schema version of the document.
	SchemaVersion pulumi.StringOutput `pulumi:"schemaVersion"`
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrOutput `pulumi:"targetType"`
	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName pulumi.StringPtrOutput `pulumi:"versionName"`
}

// NewDocument registers a new resource with the given unique name, arguments, and options.
func NewDocument(ctx *pulumi.Context,
	name string, args *DocumentArgs, opts ...pulumi.ResourceOption) (*Document, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.DocumentType == nil {
		return nil, errors.New("invalid value for required argument 'DocumentType'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Document
	err := ctx.RegisterResource("aws:ssm/document:Document", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocument gets an existing Document resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocument(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentState, opts ...pulumi.ResourceOption) (*Document, error) {
	var resource Document
	err := ctx.ReadResource("aws:ssm/document:Document", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Document resources.
type documentState struct {
	Arn *string `pulumi:"arn"`
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources []DocumentAttachmentsSource `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content *string `pulumi:"content"`
	// The date the document was created.
	CreatedDate *string `pulumi:"createdDate"`
	// The default version of the document.
	DefaultVersion *string `pulumi:"defaultVersion"`
	// The description of the document.
	Description *string `pulumi:"description"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat *string `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType *string `pulumi:"documentType"`
	// The document version.
	DocumentVersion *string `pulumi:"documentVersion"`
	// The sha1 or sha256 of the document content
	Hash *string `pulumi:"hash"`
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType *string `pulumi:"hashType"`
	// The latest version of the document.
	LatestVersion *string `pulumi:"latestVersion"`
	// The name of the document.
	Name *string `pulumi:"name"`
	// The AWS user account of the person who created the document.
	Owner *string `pulumi:"owner"`
	// The parameters that are available to this document.
	Parameters []DocumentParameter `pulumi:"parameters"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `pulumi:"permissions"`
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes []string `pulumi:"platformTypes"`
	// The schema version of the document.
	SchemaVersion *string `pulumi:"schemaVersion"`
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `pulumi:"targetType"`
	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName *string `pulumi:"versionName"`
}

type DocumentState struct {
	Arn pulumi.StringPtrInput
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayInput
	// The JSON or YAML content of the document.
	Content pulumi.StringPtrInput
	// The date the document was created.
	CreatedDate pulumi.StringPtrInput
	// The default version of the document.
	DefaultVersion pulumi.StringPtrInput
	// The description of the document.
	Description pulumi.StringPtrInput
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrInput
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringPtrInput
	// The document version.
	DocumentVersion pulumi.StringPtrInput
	// The sha1 or sha256 of the document content
	Hash pulumi.StringPtrInput
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType pulumi.StringPtrInput
	// The latest version of the document.
	LatestVersion pulumi.StringPtrInput
	// The name of the document.
	Name pulumi.StringPtrInput
	// The AWS user account of the person who created the document.
	Owner pulumi.StringPtrInput
	// The parameters that are available to this document.
	Parameters DocumentParameterArrayInput
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapInput
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes pulumi.StringArrayInput
	// The schema version of the document.
	SchemaVersion pulumi.StringPtrInput
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrInput
	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName pulumi.StringPtrInput
}

func (DocumentState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentState)(nil)).Elem()
}

type documentArgs struct {
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources []DocumentAttachmentsSource `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content string `pulumi:"content"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat *string `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType string `pulumi:"documentType"`
	// The name of the document.
	Name *string `pulumi:"name"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `pulumi:"permissions"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `pulumi:"targetType"`
	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName *string `pulumi:"versionName"`
}

// The set of arguments for constructing a Document resource.
type DocumentArgs struct {
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayInput
	// The JSON or YAML content of the document.
	Content pulumi.StringInput
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrInput
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringInput
	// The name of the document.
	Name pulumi.StringPtrInput
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapInput
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrInput
	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName pulumi.StringPtrInput
}

func (DocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentArgs)(nil)).Elem()
}

type DocumentInput interface {
	pulumi.Input

	ToDocumentOutput() DocumentOutput
	ToDocumentOutputWithContext(ctx context.Context) DocumentOutput
}

func (*Document) ElementType() reflect.Type {
	return reflect.TypeOf((**Document)(nil)).Elem()
}

func (i *Document) ToDocumentOutput() DocumentOutput {
	return i.ToDocumentOutputWithContext(context.Background())
}

func (i *Document) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentOutput)
}

func (i *Document) ToOutput(ctx context.Context) pulumix.Output[*Document] {
	return pulumix.Output[*Document]{
		OutputState: i.ToDocumentOutputWithContext(ctx).OutputState,
	}
}

// DocumentArrayInput is an input type that accepts DocumentArray and DocumentArrayOutput values.
// You can construct a concrete instance of `DocumentArrayInput` via:
//
//	DocumentArray{ DocumentArgs{...} }
type DocumentArrayInput interface {
	pulumi.Input

	ToDocumentArrayOutput() DocumentArrayOutput
	ToDocumentArrayOutputWithContext(context.Context) DocumentArrayOutput
}

type DocumentArray []DocumentInput

func (DocumentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Document)(nil)).Elem()
}

func (i DocumentArray) ToDocumentArrayOutput() DocumentArrayOutput {
	return i.ToDocumentArrayOutputWithContext(context.Background())
}

func (i DocumentArray) ToDocumentArrayOutputWithContext(ctx context.Context) DocumentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentArrayOutput)
}

func (i DocumentArray) ToOutput(ctx context.Context) pulumix.Output[[]*Document] {
	return pulumix.Output[[]*Document]{
		OutputState: i.ToDocumentArrayOutputWithContext(ctx).OutputState,
	}
}

// DocumentMapInput is an input type that accepts DocumentMap and DocumentMapOutput values.
// You can construct a concrete instance of `DocumentMapInput` via:
//
//	DocumentMap{ "key": DocumentArgs{...} }
type DocumentMapInput interface {
	pulumi.Input

	ToDocumentMapOutput() DocumentMapOutput
	ToDocumentMapOutputWithContext(context.Context) DocumentMapOutput
}

type DocumentMap map[string]DocumentInput

func (DocumentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Document)(nil)).Elem()
}

func (i DocumentMap) ToDocumentMapOutput() DocumentMapOutput {
	return i.ToDocumentMapOutputWithContext(context.Background())
}

func (i DocumentMap) ToDocumentMapOutputWithContext(ctx context.Context) DocumentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentMapOutput)
}

func (i DocumentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Document] {
	return pulumix.Output[map[string]*Document]{
		OutputState: i.ToDocumentMapOutputWithContext(ctx).OutputState,
	}
}

type DocumentOutput struct{ *pulumi.OutputState }

func (DocumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Document)(nil)).Elem()
}

func (o DocumentOutput) ToDocumentOutput() DocumentOutput {
	return o
}

func (o DocumentOutput) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return o
}

func (o DocumentOutput) ToOutput(ctx context.Context) pulumix.Output[*Document] {
	return pulumix.Output[*Document]{
		OutputState: o.OutputState,
	}
}

func (o DocumentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
func (o DocumentOutput) AttachmentsSources() DocumentAttachmentsSourceArrayOutput {
	return o.ApplyT(func(v *Document) DocumentAttachmentsSourceArrayOutput { return v.AttachmentsSources }).(DocumentAttachmentsSourceArrayOutput)
}

// The JSON or YAML content of the document.
func (o DocumentOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The date the document was created.
func (o DocumentOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The default version of the document.
func (o DocumentOutput) DefaultVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.DefaultVersion }).(pulumi.StringOutput)
}

// The description of the document.
func (o DocumentOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The format of the document. Valid document types include: `JSON` and `YAML`
func (o DocumentOutput) DocumentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Document) pulumi.StringPtrOutput { return v.DocumentFormat }).(pulumi.StringPtrOutput)
}

// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
func (o DocumentOutput) DocumentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.DocumentType }).(pulumi.StringOutput)
}

// The document version.
func (o DocumentOutput) DocumentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.DocumentVersion }).(pulumi.StringOutput)
}

// The sha1 or sha256 of the document content
func (o DocumentOutput) Hash() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Hash }).(pulumi.StringOutput)
}

// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
func (o DocumentOutput) HashType() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.HashType }).(pulumi.StringOutput)
}

// The latest version of the document.
func (o DocumentOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.LatestVersion }).(pulumi.StringOutput)
}

// The name of the document.
func (o DocumentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AWS user account of the person who created the document.
func (o DocumentOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The parameters that are available to this document.
func (o DocumentOutput) Parameters() DocumentParameterArrayOutput {
	return o.ApplyT(func(v *Document) DocumentParameterArrayOutput { return v.Parameters }).(DocumentParameterArrayOutput)
}

// Additional Permissions to attach to the document. See Permissions below for details.
func (o DocumentOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Document) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
func (o DocumentOutput) PlatformTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Document) pulumi.StringArrayOutput { return v.PlatformTypes }).(pulumi.StringArrayOutput)
}

// The schema version of the document.
func (o DocumentOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.SchemaVersion }).(pulumi.StringOutput)
}

// "Creating", "Active" or "Deleting". The current status of the document.
func (o DocumentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DocumentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Document) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o DocumentOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Document) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
func (o DocumentOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Document) pulumi.StringPtrOutput { return v.TargetType }).(pulumi.StringPtrOutput)
}

// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
func (o DocumentOutput) VersionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Document) pulumi.StringPtrOutput { return v.VersionName }).(pulumi.StringPtrOutput)
}

type DocumentArrayOutput struct{ *pulumi.OutputState }

func (DocumentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Document)(nil)).Elem()
}

func (o DocumentArrayOutput) ToDocumentArrayOutput() DocumentArrayOutput {
	return o
}

func (o DocumentArrayOutput) ToDocumentArrayOutputWithContext(ctx context.Context) DocumentArrayOutput {
	return o
}

func (o DocumentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Document] {
	return pulumix.Output[[]*Document]{
		OutputState: o.OutputState,
	}
}

func (o DocumentArrayOutput) Index(i pulumi.IntInput) DocumentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Document {
		return vs[0].([]*Document)[vs[1].(int)]
	}).(DocumentOutput)
}

type DocumentMapOutput struct{ *pulumi.OutputState }

func (DocumentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Document)(nil)).Elem()
}

func (o DocumentMapOutput) ToDocumentMapOutput() DocumentMapOutput {
	return o
}

func (o DocumentMapOutput) ToDocumentMapOutputWithContext(ctx context.Context) DocumentMapOutput {
	return o
}

func (o DocumentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Document] {
	return pulumix.Output[map[string]*Document]{
		OutputState: o.OutputState,
	}
}

func (o DocumentMapOutput) MapIndex(k pulumi.StringInput) DocumentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Document {
		return vs[0].(map[string]*Document)[vs[1].(string)]
	}).(DocumentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentInput)(nil)).Elem(), &Document{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentArrayInput)(nil)).Elem(), DocumentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentMapInput)(nil)).Elem(), DocumentMap{})
	pulumi.RegisterOutputType(DocumentOutput{})
	pulumi.RegisterOutputType(DocumentArrayOutput{})
	pulumi.RegisterOutputType(DocumentMapOutput{})
}
