// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Grafana
{
    /// <summary>
    /// Provides an Amazon Managed Grafana workspace resource.
    /// 
    /// ## Example Usage
    /// ### Basic configuration
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var assume = new Aws.Iam.Role("assume", new()
    ///     {
    ///         AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Action"] = "sts:AssumeRole",
    ///                     ["Effect"] = "Allow",
    ///                     ["Sid"] = "",
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["Service"] = "grafana.amazonaws.com",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var example = new Aws.Grafana.Workspace("example", new()
    ///     {
    ///         AccountAccessType = "CURRENT_ACCOUNT",
    ///         AuthenticationProviders = new[]
    ///         {
    ///             "SAML",
    ///         },
    ///         PermissionType = "SERVICE_MANAGED",
    ///         RoleArn = assume.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Grafana Workspace can be imported using the workspace's `id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:grafana/workspace:Workspace example g-2054c75a02
    /// ```
    /// </summary>
    [AwsResourceType("aws:grafana/workspace:Workspace")]
    public partial class Workspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizational_units` must also be present.
        /// </summary>
        [Output("accountAccessType")]
        public Output<string> AccountAccessType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Grafana workspace.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
        /// </summary>
        [Output("authenticationProviders")]
        public Output<ImmutableArray<string>> AuthenticationProviders { get; private set; } = null!;

        /// <summary>
        /// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
        /// </summary>
        [Output("dataSources")]
        public Output<ImmutableArray<string>> DataSources { get; private set; } = null!;

        /// <summary>
        /// The workspace description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the Grafana workspace.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The version of Grafana running on the workspace.
        /// </summary>
        [Output("grafanaVersion")]
        public Output<string> GrafanaVersion { get; private set; } = null!;

        /// <summary>
        /// The Grafana workspace name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
        /// </summary>
        [Output("notificationDestinations")]
        public Output<ImmutableArray<string>> NotificationDestinations { get; private set; } = null!;

        /// <summary>
        /// The role name that the workspace uses to access resources through Amazon Organizations.
        /// </summary>
        [Output("organizationRoleName")]
        public Output<string?> OrganizationRoleName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
        /// </summary>
        [Output("organizationalUnits")]
        public Output<ImmutableArray<string>> OrganizationalUnits { get; private set; } = null!;

        /// <summary>
        /// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
        /// </summary>
        [Output("permissionType")]
        public Output<string> PermissionType { get; private set; } = null!;

        /// <summary>
        /// The IAM role ARN that the workspace assumes.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        [Output("samlConfigurationStatus")]
        public Output<string> SamlConfigurationStatus { get; private set; } = null!;

        /// <summary>
        /// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
        /// </summary>
        [Output("stackSetName")]
        public Output<string?> StackSetName { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
        /// </summary>
        [Output("vpcConfiguration")]
        public Output<Outputs.WorkspaceVpcConfiguration?> VpcConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("aws:grafana/workspace:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:grafana/workspace:Workspace", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizational_units` must also be present.
        /// </summary>
        [Input("accountAccessType", required: true)]
        public Input<string> AccountAccessType { get; set; } = null!;

        [Input("authenticationProviders", required: true)]
        private InputList<string>? _authenticationProviders;

        /// <summary>
        /// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
        /// </summary>
        public InputList<string> AuthenticationProviders
        {
            get => _authenticationProviders ?? (_authenticationProviders = new InputList<string>());
            set => _authenticationProviders = value;
        }

        [Input("dataSources")]
        private InputList<string>? _dataSources;

        /// <summary>
        /// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
        /// </summary>
        public InputList<string> DataSources
        {
            get => _dataSources ?? (_dataSources = new InputList<string>());
            set => _dataSources = value;
        }

        /// <summary>
        /// The workspace description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Grafana workspace name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationDestinations")]
        private InputList<string>? _notificationDestinations;

        /// <summary>
        /// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
        /// </summary>
        public InputList<string> NotificationDestinations
        {
            get => _notificationDestinations ?? (_notificationDestinations = new InputList<string>());
            set => _notificationDestinations = value;
        }

        /// <summary>
        /// The role name that the workspace uses to access resources through Amazon Organizations.
        /// </summary>
        [Input("organizationRoleName")]
        public Input<string>? OrganizationRoleName { get; set; }

        [Input("organizationalUnits")]
        private InputList<string>? _organizationalUnits;

        /// <summary>
        /// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
        /// </summary>
        public InputList<string> OrganizationalUnits
        {
            get => _organizationalUnits ?? (_organizationalUnits = new InputList<string>());
            set => _organizationalUnits = value;
        }

        /// <summary>
        /// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
        /// </summary>
        [Input("permissionType", required: true)]
        public Input<string> PermissionType { get; set; } = null!;

        /// <summary>
        /// The IAM role ARN that the workspace assumes.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
        /// </summary>
        [Input("stackSetName")]
        public Input<string>? StackSetName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
        /// </summary>
        [Input("vpcConfiguration")]
        public Input<Inputs.WorkspaceVpcConfigurationArgs>? VpcConfiguration { get; set; }

        public WorkspaceArgs()
        {
        }
        public static new WorkspaceArgs Empty => new WorkspaceArgs();
    }

    public sealed class WorkspaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizational_units` must also be present.
        /// </summary>
        [Input("accountAccessType")]
        public Input<string>? AccountAccessType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Grafana workspace.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("authenticationProviders")]
        private InputList<string>? _authenticationProviders;

        /// <summary>
        /// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
        /// </summary>
        public InputList<string> AuthenticationProviders
        {
            get => _authenticationProviders ?? (_authenticationProviders = new InputList<string>());
            set => _authenticationProviders = value;
        }

        [Input("dataSources")]
        private InputList<string>? _dataSources;

        /// <summary>
        /// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
        /// </summary>
        public InputList<string> DataSources
        {
            get => _dataSources ?? (_dataSources = new InputList<string>());
            set => _dataSources = value;
        }

        /// <summary>
        /// The workspace description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The endpoint of the Grafana workspace.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The version of Grafana running on the workspace.
        /// </summary>
        [Input("grafanaVersion")]
        public Input<string>? GrafanaVersion { get; set; }

        /// <summary>
        /// The Grafana workspace name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationDestinations")]
        private InputList<string>? _notificationDestinations;

        /// <summary>
        /// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
        /// </summary>
        public InputList<string> NotificationDestinations
        {
            get => _notificationDestinations ?? (_notificationDestinations = new InputList<string>());
            set => _notificationDestinations = value;
        }

        /// <summary>
        /// The role name that the workspace uses to access resources through Amazon Organizations.
        /// </summary>
        [Input("organizationRoleName")]
        public Input<string>? OrganizationRoleName { get; set; }

        [Input("organizationalUnits")]
        private InputList<string>? _organizationalUnits;

        /// <summary>
        /// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
        /// </summary>
        public InputList<string> OrganizationalUnits
        {
            get => _organizationalUnits ?? (_organizationalUnits = new InputList<string>());
            set => _organizationalUnits = value;
        }

        /// <summary>
        /// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
        /// </summary>
        [Input("permissionType")]
        public Input<string>? PermissionType { get; set; }

        /// <summary>
        /// The IAM role ARN that the workspace assumes.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("samlConfigurationStatus")]
        public Input<string>? SamlConfigurationStatus { get; set; }

        /// <summary>
        /// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
        /// </summary>
        [Input("stackSetName")]
        public Input<string>? StackSetName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
        /// </summary>
        [Input("vpcConfiguration")]
        public Input<Inputs.WorkspaceVpcConfigurationGetArgs>? VpcConfiguration { get; set; }

        public WorkspaceState()
        {
        }
        public static new WorkspaceState Empty => new WorkspaceState();
    }
}
