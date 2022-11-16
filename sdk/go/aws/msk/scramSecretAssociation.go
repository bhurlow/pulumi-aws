// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates SCRAM secrets stored in the Secrets Manager service with a Managed Streaming for Kafka (MSK) cluster.
//
// > **Note:** The following assumes the MSK cluster has SASL/SCRAM authentication enabled. See below for example usage or refer to the [Username/Password Authentication](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html) section of the MSK Developer Guide for more details.
//
// To set up username and password authentication for a cluster, create an `secretsmanager.Secret` resource and associate
// a username and password with the secret with an `secretsmanager.SecretVersion` resource. When creating a secret for the cluster,
// the `name` must have the prefix `AmazonMSK_` and you must either use an existing custom AWS KMS key or create a new
// custom AWS KMS key for your secret with the `kms.Key` resource. It is important to note that a policy is required for the `secretsmanager.Secret`
// resource in order for Kafka to be able to read it. This policy is attached automatically when the `msk.ScramSecretAssociation` is used,
// however, this policy will not be in the state and as such, will present a diff on plan/apply. For that reason, you must use the `secretsmanager.SecretPolicy`
// resource](/docs/providers/aws/r/secretsmanager_secret_policy.html) as shown below in order to ensure that the state is in a clean state after the creation of secret and the association to the cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kms"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/msk"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/secretsmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleCluster, err := msk.NewCluster(ctx, "exampleCluster", &msk.ClusterArgs{
//				ClientAuthentication: &msk.ClusterClientAuthenticationArgs{
//					Sasl: &msk.ClusterClientAuthenticationSaslArgs{
//						Scram: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
//				Description: pulumi.String("Example Key for MSK Cluster Scram Secret Association"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecret, err := secretsmanager.NewSecret(ctx, "exampleSecret", &secretsmanager.SecretArgs{
//				KmsKeyId: exampleKey.KeyId,
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"username": "user",
//				"password": "pass",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			exampleSecretVersion, err := secretsmanager.NewSecretVersion(ctx, "exampleSecretVersion", &secretsmanager.SecretVersionArgs{
//				SecretId:     exampleSecret.ID(),
//				SecretString: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = msk.NewScramSecretAssociation(ctx, "exampleScramSecretAssociation", &msk.ScramSecretAssociationArgs{
//				ClusterArn: exampleCluster.Arn,
//				SecretArnLists: pulumi.StringArray{
//					exampleSecret.Arn,
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleSecretVersion,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = secretsmanager.NewSecretPolicy(ctx, "exampleSecretPolicy", &secretsmanager.SecretPolicyArgs{
//				SecretArn: exampleSecret.Arn,
//				Policy: exampleSecret.Arn.ApplyT(func(arn string) (string, error) {
//					return fmt.Sprintf(`{
//	  "Version" : "2012-10-17",
//	  "Statement" : [ {
//	    "Sid": "AWSKafkaResourcePolicy",
//	    "Effect" : "Allow",
//	    "Principal" : {
//	      "Service" : "kafka.amazonaws.com"
//	    },
//	    "Action" : "secretsmanager:getSecretValue",
//	    "Resource" : "%v"
//	  } ]
//	}
//
// `, arn), nil
//
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// MSK SCRAM Secret Associations can be imported using the `id` e.g.,
//
// ```sh
//
//	$ pulumi import aws:msk/scramSecretAssociation:ScramSecretAssociation example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
//
// ```
type ScramSecretAssociation struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn pulumi.StringOutput `pulumi:"clusterArn"`
	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists pulumi.StringArrayOutput `pulumi:"secretArnLists"`
}

// NewScramSecretAssociation registers a new resource with the given unique name, arguments, and options.
func NewScramSecretAssociation(ctx *pulumi.Context,
	name string, args *ScramSecretAssociationArgs, opts ...pulumi.ResourceOption) (*ScramSecretAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterArn == nil {
		return nil, errors.New("invalid value for required argument 'ClusterArn'")
	}
	if args.SecretArnLists == nil {
		return nil, errors.New("invalid value for required argument 'SecretArnLists'")
	}
	var resource ScramSecretAssociation
	err := ctx.RegisterResource("aws:msk/scramSecretAssociation:ScramSecretAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScramSecretAssociation gets an existing ScramSecretAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScramSecretAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScramSecretAssociationState, opts ...pulumi.ResourceOption) (*ScramSecretAssociation, error) {
	var resource ScramSecretAssociation
	err := ctx.ReadResource("aws:msk/scramSecretAssociation:ScramSecretAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScramSecretAssociation resources.
type scramSecretAssociationState struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn *string `pulumi:"clusterArn"`
	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists []string `pulumi:"secretArnLists"`
}

type ScramSecretAssociationState struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn pulumi.StringPtrInput
	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists pulumi.StringArrayInput
}

func (ScramSecretAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*scramSecretAssociationState)(nil)).Elem()
}

type scramSecretAssociationArgs struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn string `pulumi:"clusterArn"`
	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists []string `pulumi:"secretArnLists"`
}

// The set of arguments for constructing a ScramSecretAssociation resource.
type ScramSecretAssociationArgs struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn pulumi.StringInput
	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists pulumi.StringArrayInput
}

func (ScramSecretAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scramSecretAssociationArgs)(nil)).Elem()
}

type ScramSecretAssociationInput interface {
	pulumi.Input

	ToScramSecretAssociationOutput() ScramSecretAssociationOutput
	ToScramSecretAssociationOutputWithContext(ctx context.Context) ScramSecretAssociationOutput
}

func (*ScramSecretAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ScramSecretAssociation)(nil)).Elem()
}

func (i *ScramSecretAssociation) ToScramSecretAssociationOutput() ScramSecretAssociationOutput {
	return i.ToScramSecretAssociationOutputWithContext(context.Background())
}

func (i *ScramSecretAssociation) ToScramSecretAssociationOutputWithContext(ctx context.Context) ScramSecretAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScramSecretAssociationOutput)
}

// ScramSecretAssociationArrayInput is an input type that accepts ScramSecretAssociationArray and ScramSecretAssociationArrayOutput values.
// You can construct a concrete instance of `ScramSecretAssociationArrayInput` via:
//
//	ScramSecretAssociationArray{ ScramSecretAssociationArgs{...} }
type ScramSecretAssociationArrayInput interface {
	pulumi.Input

	ToScramSecretAssociationArrayOutput() ScramSecretAssociationArrayOutput
	ToScramSecretAssociationArrayOutputWithContext(context.Context) ScramSecretAssociationArrayOutput
}

type ScramSecretAssociationArray []ScramSecretAssociationInput

func (ScramSecretAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScramSecretAssociation)(nil)).Elem()
}

func (i ScramSecretAssociationArray) ToScramSecretAssociationArrayOutput() ScramSecretAssociationArrayOutput {
	return i.ToScramSecretAssociationArrayOutputWithContext(context.Background())
}

func (i ScramSecretAssociationArray) ToScramSecretAssociationArrayOutputWithContext(ctx context.Context) ScramSecretAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScramSecretAssociationArrayOutput)
}

// ScramSecretAssociationMapInput is an input type that accepts ScramSecretAssociationMap and ScramSecretAssociationMapOutput values.
// You can construct a concrete instance of `ScramSecretAssociationMapInput` via:
//
//	ScramSecretAssociationMap{ "key": ScramSecretAssociationArgs{...} }
type ScramSecretAssociationMapInput interface {
	pulumi.Input

	ToScramSecretAssociationMapOutput() ScramSecretAssociationMapOutput
	ToScramSecretAssociationMapOutputWithContext(context.Context) ScramSecretAssociationMapOutput
}

type ScramSecretAssociationMap map[string]ScramSecretAssociationInput

func (ScramSecretAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScramSecretAssociation)(nil)).Elem()
}

func (i ScramSecretAssociationMap) ToScramSecretAssociationMapOutput() ScramSecretAssociationMapOutput {
	return i.ToScramSecretAssociationMapOutputWithContext(context.Background())
}

func (i ScramSecretAssociationMap) ToScramSecretAssociationMapOutputWithContext(ctx context.Context) ScramSecretAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScramSecretAssociationMapOutput)
}

type ScramSecretAssociationOutput struct{ *pulumi.OutputState }

func (ScramSecretAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScramSecretAssociation)(nil)).Elem()
}

func (o ScramSecretAssociationOutput) ToScramSecretAssociationOutput() ScramSecretAssociationOutput {
	return o
}

func (o ScramSecretAssociationOutput) ToScramSecretAssociationOutputWithContext(ctx context.Context) ScramSecretAssociationOutput {
	return o
}

// Amazon Resource Name (ARN) of the MSK cluster.
func (o ScramSecretAssociationOutput) ClusterArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ScramSecretAssociation) pulumi.StringOutput { return v.ClusterArn }).(pulumi.StringOutput)
}

// List of AWS Secrets Manager secret ARNs.
func (o ScramSecretAssociationOutput) SecretArnLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScramSecretAssociation) pulumi.StringArrayOutput { return v.SecretArnLists }).(pulumi.StringArrayOutput)
}

type ScramSecretAssociationArrayOutput struct{ *pulumi.OutputState }

func (ScramSecretAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScramSecretAssociation)(nil)).Elem()
}

func (o ScramSecretAssociationArrayOutput) ToScramSecretAssociationArrayOutput() ScramSecretAssociationArrayOutput {
	return o
}

func (o ScramSecretAssociationArrayOutput) ToScramSecretAssociationArrayOutputWithContext(ctx context.Context) ScramSecretAssociationArrayOutput {
	return o
}

func (o ScramSecretAssociationArrayOutput) Index(i pulumi.IntInput) ScramSecretAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScramSecretAssociation {
		return vs[0].([]*ScramSecretAssociation)[vs[1].(int)]
	}).(ScramSecretAssociationOutput)
}

type ScramSecretAssociationMapOutput struct{ *pulumi.OutputState }

func (ScramSecretAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScramSecretAssociation)(nil)).Elem()
}

func (o ScramSecretAssociationMapOutput) ToScramSecretAssociationMapOutput() ScramSecretAssociationMapOutput {
	return o
}

func (o ScramSecretAssociationMapOutput) ToScramSecretAssociationMapOutputWithContext(ctx context.Context) ScramSecretAssociationMapOutput {
	return o
}

func (o ScramSecretAssociationMapOutput) MapIndex(k pulumi.StringInput) ScramSecretAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScramSecretAssociation {
		return vs[0].(map[string]*ScramSecretAssociation)[vs[1].(string)]
	}).(ScramSecretAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScramSecretAssociationInput)(nil)).Elem(), &ScramSecretAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScramSecretAssociationArrayInput)(nil)).Elem(), ScramSecretAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScramSecretAssociationMapInput)(nil)).Elem(), ScramSecretAssociationMap{})
	pulumi.RegisterOutputType(ScramSecretAssociationOutput{})
	pulumi.RegisterOutputType(ScramSecretAssociationArrayOutput{})
	pulumi.RegisterOutputType(ScramSecretAssociationMapOutput{})
}
