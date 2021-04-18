// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Outputs
{

    [OutputType]
    public sealed class InsightFilters
    {
        /// <summary>
        /// AWS account ID that a finding is generated in. See String_Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersAwsAccountId> AwsAccountIds;
        /// <summary>
        /// The name of the findings provider (company) that owns the solution (product) that generates findings. See String_Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersCompanyName> CompanyNames;
        /// <summary>
        /// Exclusive to findings that are generated as the result of a check run against a specific rule in a supported standard, such as CIS AWS Foundations. Contains security standard-related finding details. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersComplianceStatus> ComplianceStatuses;
        /// <summary>
        /// A finding's confidence. Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify. Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersConfidence> Confidences;
        /// <summary>
        /// An ISO8601-formatted timestamp that indicates when the security-findings provider captured the potential security issue that a finding captured. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersCreatedAt> CreatedAts;
        /// <summary>
        /// The level of importance assigned to the resources associated with the finding. A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersCriticality> Criticalities;
        /// <summary>
        /// A finding's description. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersDescription> Descriptions;
        /// <summary>
        /// The finding provider value for the finding confidence. Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify. Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsConfidence> FindingProviderFieldsConfidences;
        /// <summary>
        /// The finding provider value for the level of importance assigned to the resources associated with the findings. A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsCriticality> FindingProviderFieldsCriticalities;
        /// <summary>
        /// The finding identifier of a related finding that is identified by the finding provider. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsRelatedFindingsId> FindingProviderFieldsRelatedFindingsIds;
        /// <summary>
        /// The ARN of the solution that generated a related finding that is identified by the finding provider. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsRelatedFindingsProductArn> FindingProviderFieldsRelatedFindingsProductArns;
        /// <summary>
        /// The finding provider value for the severity label. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsSeverityLabel> FindingProviderFieldsSeverityLabels;
        /// <summary>
        /// The finding provider's original value for the severity. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsSeverityOriginal> FindingProviderFieldsSeverityOriginals;
        /// <summary>
        /// One or more finding types that the finding provider assigned to the finding. Uses the format of `namespace/category/classifier` that classify a finding. Valid namespace values include: `Software and Configuration Checks`, `TTPs`, `Effects`, `Unusual Behaviors`, and `Sensitive Data Identifications`. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsType> FindingProviderFieldsTypes;
        /// <summary>
        /// An ISO8601-formatted timestamp that indicates when the security-findings provider first observed the potential security issue that a finding captured. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersFirstObservedAt> FirstObservedAts;
        /// <summary>
        /// The identifier for the solution-specific component (a discrete unit of logic) that generated a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersGeneratorId> GeneratorIds;
        /// <summary>
        /// The security findings provider-specific identifier for a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersId> Ids;
        /// <summary>
        /// A keyword for a finding. See Keyword Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersKeyword> Keywords;
        /// <summary>
        /// An ISO8601-formatted timestamp that indicates when the security-findings provider most recently observed the potential security issue that a finding captured. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersLastObservedAt> LastObservedAts;
        /// <summary>
        /// The name of the malware that was observed. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersMalwareName> MalwareNames;
        /// <summary>
        /// The filesystem path of the malware that was observed. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersMalwarePath> MalwarePaths;
        /// <summary>
        /// The state of the malware that was observed. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersMalwareState> MalwareStates;
        /// <summary>
        /// The type of the malware that was observed. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersMalwareType> MalwareTypes;
        /// <summary>
        /// The destination domain of network-related information about a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkDestinationDomain> NetworkDestinationDomains;
        /// <summary>
        /// The destination IPv4 address of network-related information about a finding. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkDestinationIpv4> NetworkDestinationIpv4s;
        /// <summary>
        /// The destination IPv6 address of network-related information about a finding. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkDestinationIpv6> NetworkDestinationIpv6s;
        /// <summary>
        /// The destination port of network-related information about a finding. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkDestinationPort> NetworkDestinationPorts;
        /// <summary>
        /// Indicates the direction of network traffic associated with a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkDirection> NetworkDirections;
        /// <summary>
        /// The protocol of network-related information about a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkProtocol> NetworkProtocols;
        /// <summary>
        /// The source domain of network-related information about a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkSourceDomain> NetworkSourceDomains;
        /// <summary>
        /// The source IPv4 address of network-related information about a finding. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkSourceIpv4> NetworkSourceIpv4s;
        /// <summary>
        /// The source IPv6 address of network-related information about a finding. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkSourceIpv6> NetworkSourceIpv6s;
        /// <summary>
        /// The source media access control (MAC) address of network-related information about a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkSourceMac> NetworkSourceMacs;
        /// <summary>
        /// The source port of network-related information about a finding. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNetworkSourcePort> NetworkSourcePorts;
        /// <summary>
        /// The text of a note. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNoteText> NoteTexts;
        /// <summary>
        /// The timestamp of when the note was updated. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNoteUpdatedAt> NoteUpdatedAts;
        /// <summary>
        /// The principal that created a note. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersNoteUpdatedBy> NoteUpdatedBies;
        /// <summary>
        /// The date/time that the process was launched. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessLaunchedAt> ProcessLaunchedAts;
        /// <summary>
        /// The name of the process. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessName> ProcessNames;
        /// <summary>
        /// The parent process ID. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessParentPid> ProcessParentPids;
        /// <summary>
        /// The path to the process executable. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessPath> ProcessPaths;
        /// <summary>
        /// The process ID. See Number Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessPid> ProcessPids;
        /// <summary>
        /// The date/time that the process was terminated. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProcessTerminatedAt> ProcessTerminatedAts;
        /// <summary>
        /// The ARN generated by Security Hub that uniquely identifies a third-party company (security findings provider) after this provider's product (solution that generates findings) is registered with Security Hub. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProductArn> ProductArns;
        /// <summary>
        /// A data type where security-findings providers can include additional solution-specific details that aren't part of the defined `AwsSecurityFinding` format. See Map Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProductField> ProductFields;
        /// <summary>
        /// The name of the solution (product) that generates findings. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersProductName> ProductNames;
        /// <summary>
        /// The recommendation of what to do about the issue described in a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersRecommendationText> RecommendationTexts;
        /// <summary>
        /// The updated record state for the finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersRecordState> RecordStates;
        /// <summary>
        /// The solution-generated identifier for a related finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersRelatedFindingsId> RelatedFindingsIds;
        /// <summary>
        /// The ARN of the solution that generated a related finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersRelatedFindingsProductArn> RelatedFindingsProductArns;
        /// <summary>
        /// The IAM profile ARN of the instance. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn> ResourceAwsEc2InstanceIamInstanceProfileArns;
        /// <summary>
        /// The Amazon Machine Image (AMI) ID of the instance. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceImageId> ResourceAwsEc2InstanceImageIds;
        /// <summary>
        /// The IPv4 addresses associated with the instance. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIpv4Address> ResourceAwsEc2InstanceIpv4Addresses;
        /// <summary>
        /// The IPv6 addresses associated with the instance. See Ip Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIpv6Address> ResourceAwsEc2InstanceIpv6Addresses;
        /// <summary>
        /// The key name associated with the instance. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceKeyName> ResourceAwsEc2InstanceKeyNames;
        /// <summary>
        /// The date and time the instance was launched. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceLaunchedAt> ResourceAwsEc2InstanceLaunchedAts;
        /// <summary>
        /// The identifier of the subnet that the instance was launched in. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceSubnetId> ResourceAwsEc2InstanceSubnetIds;
        /// <summary>
        /// The instance type of the instance. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceType> ResourceAwsEc2InstanceTypes;
        /// <summary>
        /// The identifier of the VPC that the instance was launched in. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceVpcId> ResourceAwsEc2InstanceVpcIds;
        /// <summary>
        /// The creation date/time of the IAM access key related to a finding. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyCreatedAt> ResourceAwsIamAccessKeyCreatedAts;
        /// <summary>
        /// The status of the IAM access key related to a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyStatus> ResourceAwsIamAccessKeyStatuses;
        /// <summary>
        /// The user associated with the IAM access key related to a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyUserName> ResourceAwsIamAccessKeyUserNames;
        /// <summary>
        /// The canonical user ID of the owner of the S3 bucket. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsS3BucketOwnerId> ResourceAwsS3BucketOwnerIds;
        /// <summary>
        /// The display name of the owner of the S3 bucket. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceAwsS3BucketOwnerName> ResourceAwsS3BucketOwnerNames;
        /// <summary>
        /// The identifier of the image related to a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceContainerImageId> ResourceContainerImageIds;
        /// <summary>
        /// The name of the image related to a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceContainerImageName> ResourceContainerImageNames;
        /// <summary>
        /// The date/time that the container was started. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceContainerLaunchedAt> ResourceContainerLaunchedAts;
        /// <summary>
        /// The name of the container related to a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceContainerName> ResourceContainerNames;
        /// <summary>
        /// The details of a resource that doesn't have a specific subfield for the resource type defined. See Map Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceDetailsOther> ResourceDetailsOthers;
        /// <summary>
        /// The canonical identifier for the given resource type. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceId> ResourceIds;
        /// <summary>
        /// The canonical AWS partition name that the Region is assigned to. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourcePartition> ResourcePartitions;
        /// <summary>
        /// The canonical AWS external Region name where this resource is located. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceRegion> ResourceRegions;
        /// <summary>
        /// A list of AWS tags associated with a resource at the time the finding was processed. See Map Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceTag> ResourceTags;
        /// <summary>
        /// Specifies the type of the resource that details are provided for. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersResourceType> ResourceTypes;
        /// <summary>
        /// The label of a finding's severity. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersSeverityLabel> SeverityLabels;
        /// <summary>
        /// A URL that links to a page about the current finding in the security-findings provider's solution. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersSourceUrl> SourceUrls;
        /// <summary>
        /// The category of a threat intelligence indicator. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorCategory> ThreatIntelIndicatorCategories;
        /// <summary>
        /// The date/time of the last observation of a threat intelligence indicator. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorLastObservedAt> ThreatIntelIndicatorLastObservedAts;
        /// <summary>
        /// The URL for more details from the source of the threat intelligence. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorSourceUrl> ThreatIntelIndicatorSourceUrls;
        /// <summary>
        /// The source of the threat intelligence. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorSource> ThreatIntelIndicatorSources;
        /// <summary>
        /// The type of a threat intelligence indicator. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorType> ThreatIntelIndicatorTypes;
        /// <summary>
        /// The value of a threat intelligence indicator. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorValue> ThreatIntelIndicatorValues;
        /// <summary>
        /// A finding's title. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersTitle> Titles;
        /// <summary>
        /// A finding type in the format of `namespace/category/classifier` that classifies a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersType> Types;
        /// <summary>
        /// An ISO8601-formatted timestamp that indicates when the security-findings provider last updated the finding record. See Date Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersUpdatedAt> UpdatedAts;
        /// <summary>
        /// A list of name/value string pairs associated with the finding. These are custom, user-defined fields added to a finding. See Map Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersUserDefinedValue> UserDefinedValues;
        /// <summary>
        /// The veracity of a finding. See String Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersVerificationState> VerificationStates;
        /// <summary>
        /// The status of the investigation into a finding. See Workflow Status Filter below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightFiltersWorkflowStatus> WorkflowStatuses;

        [OutputConstructor]
        private InsightFilters(
            ImmutableArray<Outputs.InsightFiltersAwsAccountId> awsAccountIds,

            ImmutableArray<Outputs.InsightFiltersCompanyName> companyNames,

            ImmutableArray<Outputs.InsightFiltersComplianceStatus> complianceStatuses,

            ImmutableArray<Outputs.InsightFiltersConfidence> confidences,

            ImmutableArray<Outputs.InsightFiltersCreatedAt> createdAts,

            ImmutableArray<Outputs.InsightFiltersCriticality> criticalities,

            ImmutableArray<Outputs.InsightFiltersDescription> descriptions,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsConfidence> findingProviderFieldsConfidences,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsCriticality> findingProviderFieldsCriticalities,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsRelatedFindingsId> findingProviderFieldsRelatedFindingsIds,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsRelatedFindingsProductArn> findingProviderFieldsRelatedFindingsProductArns,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsSeverityLabel> findingProviderFieldsSeverityLabels,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsSeverityOriginal> findingProviderFieldsSeverityOriginals,

            ImmutableArray<Outputs.InsightFiltersFindingProviderFieldsType> findingProviderFieldsTypes,

            ImmutableArray<Outputs.InsightFiltersFirstObservedAt> firstObservedAts,

            ImmutableArray<Outputs.InsightFiltersGeneratorId> generatorIds,

            ImmutableArray<Outputs.InsightFiltersId> ids,

            ImmutableArray<Outputs.InsightFiltersKeyword> keywords,

            ImmutableArray<Outputs.InsightFiltersLastObservedAt> lastObservedAts,

            ImmutableArray<Outputs.InsightFiltersMalwareName> malwareNames,

            ImmutableArray<Outputs.InsightFiltersMalwarePath> malwarePaths,

            ImmutableArray<Outputs.InsightFiltersMalwareState> malwareStates,

            ImmutableArray<Outputs.InsightFiltersMalwareType> malwareTypes,

            ImmutableArray<Outputs.InsightFiltersNetworkDestinationDomain> networkDestinationDomains,

            ImmutableArray<Outputs.InsightFiltersNetworkDestinationIpv4> networkDestinationIpv4s,

            ImmutableArray<Outputs.InsightFiltersNetworkDestinationIpv6> networkDestinationIpv6s,

            ImmutableArray<Outputs.InsightFiltersNetworkDestinationPort> networkDestinationPorts,

            ImmutableArray<Outputs.InsightFiltersNetworkDirection> networkDirections,

            ImmutableArray<Outputs.InsightFiltersNetworkProtocol> networkProtocols,

            ImmutableArray<Outputs.InsightFiltersNetworkSourceDomain> networkSourceDomains,

            ImmutableArray<Outputs.InsightFiltersNetworkSourceIpv4> networkSourceIpv4s,

            ImmutableArray<Outputs.InsightFiltersNetworkSourceIpv6> networkSourceIpv6s,

            ImmutableArray<Outputs.InsightFiltersNetworkSourceMac> networkSourceMacs,

            ImmutableArray<Outputs.InsightFiltersNetworkSourcePort> networkSourcePorts,

            ImmutableArray<Outputs.InsightFiltersNoteText> noteTexts,

            ImmutableArray<Outputs.InsightFiltersNoteUpdatedAt> noteUpdatedAts,

            ImmutableArray<Outputs.InsightFiltersNoteUpdatedBy> noteUpdatedBies,

            ImmutableArray<Outputs.InsightFiltersProcessLaunchedAt> processLaunchedAts,

            ImmutableArray<Outputs.InsightFiltersProcessName> processNames,

            ImmutableArray<Outputs.InsightFiltersProcessParentPid> processParentPids,

            ImmutableArray<Outputs.InsightFiltersProcessPath> processPaths,

            ImmutableArray<Outputs.InsightFiltersProcessPid> processPids,

            ImmutableArray<Outputs.InsightFiltersProcessTerminatedAt> processTerminatedAts,

            ImmutableArray<Outputs.InsightFiltersProductArn> productArns,

            ImmutableArray<Outputs.InsightFiltersProductField> productFields,

            ImmutableArray<Outputs.InsightFiltersProductName> productNames,

            ImmutableArray<Outputs.InsightFiltersRecommendationText> recommendationTexts,

            ImmutableArray<Outputs.InsightFiltersRecordState> recordStates,

            ImmutableArray<Outputs.InsightFiltersRelatedFindingsId> relatedFindingsIds,

            ImmutableArray<Outputs.InsightFiltersRelatedFindingsProductArn> relatedFindingsProductArns,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn> resourceAwsEc2InstanceIamInstanceProfileArns,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceImageId> resourceAwsEc2InstanceImageIds,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIpv4Address> resourceAwsEc2InstanceIpv4Addresses,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceIpv6Address> resourceAwsEc2InstanceIpv6Addresses,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceKeyName> resourceAwsEc2InstanceKeyNames,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceLaunchedAt> resourceAwsEc2InstanceLaunchedAts,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceSubnetId> resourceAwsEc2InstanceSubnetIds,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceType> resourceAwsEc2InstanceTypes,

            ImmutableArray<Outputs.InsightFiltersResourceAwsEc2InstanceVpcId> resourceAwsEc2InstanceVpcIds,

            ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyCreatedAt> resourceAwsIamAccessKeyCreatedAts,

            ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyStatus> resourceAwsIamAccessKeyStatuses,

            ImmutableArray<Outputs.InsightFiltersResourceAwsIamAccessKeyUserName> resourceAwsIamAccessKeyUserNames,

            ImmutableArray<Outputs.InsightFiltersResourceAwsS3BucketOwnerId> resourceAwsS3BucketOwnerIds,

            ImmutableArray<Outputs.InsightFiltersResourceAwsS3BucketOwnerName> resourceAwsS3BucketOwnerNames,

            ImmutableArray<Outputs.InsightFiltersResourceContainerImageId> resourceContainerImageIds,

            ImmutableArray<Outputs.InsightFiltersResourceContainerImageName> resourceContainerImageNames,

            ImmutableArray<Outputs.InsightFiltersResourceContainerLaunchedAt> resourceContainerLaunchedAts,

            ImmutableArray<Outputs.InsightFiltersResourceContainerName> resourceContainerNames,

            ImmutableArray<Outputs.InsightFiltersResourceDetailsOther> resourceDetailsOthers,

            ImmutableArray<Outputs.InsightFiltersResourceId> resourceIds,

            ImmutableArray<Outputs.InsightFiltersResourcePartition> resourcePartitions,

            ImmutableArray<Outputs.InsightFiltersResourceRegion> resourceRegions,

            ImmutableArray<Outputs.InsightFiltersResourceTag> resourceTags,

            ImmutableArray<Outputs.InsightFiltersResourceType> resourceTypes,

            ImmutableArray<Outputs.InsightFiltersSeverityLabel> severityLabels,

            ImmutableArray<Outputs.InsightFiltersSourceUrl> sourceUrls,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorCategory> threatIntelIndicatorCategories,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorLastObservedAt> threatIntelIndicatorLastObservedAts,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorSourceUrl> threatIntelIndicatorSourceUrls,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorSource> threatIntelIndicatorSources,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorType> threatIntelIndicatorTypes,

            ImmutableArray<Outputs.InsightFiltersThreatIntelIndicatorValue> threatIntelIndicatorValues,

            ImmutableArray<Outputs.InsightFiltersTitle> titles,

            ImmutableArray<Outputs.InsightFiltersType> types,

            ImmutableArray<Outputs.InsightFiltersUpdatedAt> updatedAts,

            ImmutableArray<Outputs.InsightFiltersUserDefinedValue> userDefinedValues,

            ImmutableArray<Outputs.InsightFiltersVerificationState> verificationStates,

            ImmutableArray<Outputs.InsightFiltersWorkflowStatus> workflowStatuses)
        {
            AwsAccountIds = awsAccountIds;
            CompanyNames = companyNames;
            ComplianceStatuses = complianceStatuses;
            Confidences = confidences;
            CreatedAts = createdAts;
            Criticalities = criticalities;
            Descriptions = descriptions;
            FindingProviderFieldsConfidences = findingProviderFieldsConfidences;
            FindingProviderFieldsCriticalities = findingProviderFieldsCriticalities;
            FindingProviderFieldsRelatedFindingsIds = findingProviderFieldsRelatedFindingsIds;
            FindingProviderFieldsRelatedFindingsProductArns = findingProviderFieldsRelatedFindingsProductArns;
            FindingProviderFieldsSeverityLabels = findingProviderFieldsSeverityLabels;
            FindingProviderFieldsSeverityOriginals = findingProviderFieldsSeverityOriginals;
            FindingProviderFieldsTypes = findingProviderFieldsTypes;
            FirstObservedAts = firstObservedAts;
            GeneratorIds = generatorIds;
            Ids = ids;
            Keywords = keywords;
            LastObservedAts = lastObservedAts;
            MalwareNames = malwareNames;
            MalwarePaths = malwarePaths;
            MalwareStates = malwareStates;
            MalwareTypes = malwareTypes;
            NetworkDestinationDomains = networkDestinationDomains;
            NetworkDestinationIpv4s = networkDestinationIpv4s;
            NetworkDestinationIpv6s = networkDestinationIpv6s;
            NetworkDestinationPorts = networkDestinationPorts;
            NetworkDirections = networkDirections;
            NetworkProtocols = networkProtocols;
            NetworkSourceDomains = networkSourceDomains;
            NetworkSourceIpv4s = networkSourceIpv4s;
            NetworkSourceIpv6s = networkSourceIpv6s;
            NetworkSourceMacs = networkSourceMacs;
            NetworkSourcePorts = networkSourcePorts;
            NoteTexts = noteTexts;
            NoteUpdatedAts = noteUpdatedAts;
            NoteUpdatedBies = noteUpdatedBies;
            ProcessLaunchedAts = processLaunchedAts;
            ProcessNames = processNames;
            ProcessParentPids = processParentPids;
            ProcessPaths = processPaths;
            ProcessPids = processPids;
            ProcessTerminatedAts = processTerminatedAts;
            ProductArns = productArns;
            ProductFields = productFields;
            ProductNames = productNames;
            RecommendationTexts = recommendationTexts;
            RecordStates = recordStates;
            RelatedFindingsIds = relatedFindingsIds;
            RelatedFindingsProductArns = relatedFindingsProductArns;
            ResourceAwsEc2InstanceIamInstanceProfileArns = resourceAwsEc2InstanceIamInstanceProfileArns;
            ResourceAwsEc2InstanceImageIds = resourceAwsEc2InstanceImageIds;
            ResourceAwsEc2InstanceIpv4Addresses = resourceAwsEc2InstanceIpv4Addresses;
            ResourceAwsEc2InstanceIpv6Addresses = resourceAwsEc2InstanceIpv6Addresses;
            ResourceAwsEc2InstanceKeyNames = resourceAwsEc2InstanceKeyNames;
            ResourceAwsEc2InstanceLaunchedAts = resourceAwsEc2InstanceLaunchedAts;
            ResourceAwsEc2InstanceSubnetIds = resourceAwsEc2InstanceSubnetIds;
            ResourceAwsEc2InstanceTypes = resourceAwsEc2InstanceTypes;
            ResourceAwsEc2InstanceVpcIds = resourceAwsEc2InstanceVpcIds;
            ResourceAwsIamAccessKeyCreatedAts = resourceAwsIamAccessKeyCreatedAts;
            ResourceAwsIamAccessKeyStatuses = resourceAwsIamAccessKeyStatuses;
            ResourceAwsIamAccessKeyUserNames = resourceAwsIamAccessKeyUserNames;
            ResourceAwsS3BucketOwnerIds = resourceAwsS3BucketOwnerIds;
            ResourceAwsS3BucketOwnerNames = resourceAwsS3BucketOwnerNames;
            ResourceContainerImageIds = resourceContainerImageIds;
            ResourceContainerImageNames = resourceContainerImageNames;
            ResourceContainerLaunchedAts = resourceContainerLaunchedAts;
            ResourceContainerNames = resourceContainerNames;
            ResourceDetailsOthers = resourceDetailsOthers;
            ResourceIds = resourceIds;
            ResourcePartitions = resourcePartitions;
            ResourceRegions = resourceRegions;
            ResourceTags = resourceTags;
            ResourceTypes = resourceTypes;
            SeverityLabels = severityLabels;
            SourceUrls = sourceUrls;
            ThreatIntelIndicatorCategories = threatIntelIndicatorCategories;
            ThreatIntelIndicatorLastObservedAts = threatIntelIndicatorLastObservedAts;
            ThreatIntelIndicatorSourceUrls = threatIntelIndicatorSourceUrls;
            ThreatIntelIndicatorSources = threatIntelIndicatorSources;
            ThreatIntelIndicatorTypes = threatIntelIndicatorTypes;
            ThreatIntelIndicatorValues = threatIntelIndicatorValues;
            Titles = titles;
            Types = types;
            UpdatedAts = updatedAts;
            UserDefinedValues = userDefinedValues;
            VerificationStates = verificationStates;
            WorkflowStatuses = workflowStatuses;
        }
    }
}
