# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetCoreNetworkPolicyDocumentResult',
    'AwaitableGetCoreNetworkPolicyDocumentResult',
    'get_core_network_policy_document',
    'get_core_network_policy_document_output',
]

@pulumi.output_type
class GetCoreNetworkPolicyDocumentResult:
    """
    A collection of values returned by getCoreNetworkPolicyDocument.
    """
    def __init__(__self__, attachment_policies=None, core_network_configurations=None, id=None, json=None, segment_actions=None, segments=None, version=None):
        if attachment_policies and not isinstance(attachment_policies, list):
            raise TypeError("Expected argument 'attachment_policies' to be a list")
        pulumi.set(__self__, "attachment_policies", attachment_policies)
        if core_network_configurations and not isinstance(core_network_configurations, list):
            raise TypeError("Expected argument 'core_network_configurations' to be a list")
        pulumi.set(__self__, "core_network_configurations", core_network_configurations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if segment_actions and not isinstance(segment_actions, list):
            raise TypeError("Expected argument 'segment_actions' to be a list")
        pulumi.set(__self__, "segment_actions", segment_actions)
        if segments and not isinstance(segments, list):
            raise TypeError("Expected argument 'segments' to be a list")
        pulumi.set(__self__, "segments", segments)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="attachmentPolicies")
    def attachment_policies(self) -> Optional[Sequence['outputs.GetCoreNetworkPolicyDocumentAttachmentPolicyResult']]:
        return pulumi.get(self, "attachment_policies")

    @property
    @pulumi.getter(name="coreNetworkConfigurations")
    def core_network_configurations(self) -> Sequence['outputs.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationResult']:
        return pulumi.get(self, "core_network_configurations")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        Standard JSON policy document rendered based on the arguments above.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="segmentActions")
    def segment_actions(self) -> Optional[Sequence['outputs.GetCoreNetworkPolicyDocumentSegmentActionResult']]:
        return pulumi.get(self, "segment_actions")

    @property
    @pulumi.getter
    def segments(self) -> Sequence['outputs.GetCoreNetworkPolicyDocumentSegmentResult']:
        return pulumi.get(self, "segments")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetCoreNetworkPolicyDocumentResult(GetCoreNetworkPolicyDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCoreNetworkPolicyDocumentResult(
            attachment_policies=self.attachment_policies,
            core_network_configurations=self.core_network_configurations,
            id=self.id,
            json=self.json,
            segment_actions=self.segment_actions,
            segments=self.segments,
            version=self.version)


def get_core_network_policy_document(attachment_policies: Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentAttachmentPolicyArgs']]] = None,
                                     core_network_configurations: Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs']]] = None,
                                     segment_actions: Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentActionArgs']]] = None,
                                     segments: Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentArgs']]] = None,
                                     version: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCoreNetworkPolicyDocumentResult:
    """
    ## Example Usage
    ### Basic Example

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.networkmanager.get_core_network_policy_document(attachment_policies=[
            aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
                action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                    association_method="constant",
                    segment="shared",
                ),
                condition_logic="or",
                conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                    key="segment",
                    operator="equals",
                    type="tag-value",
                    value="shared",
                )],
                rule_number=100,
            ),
            aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
                action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                    association_method="constant",
                    segment="prod",
                ),
                condition_logic="or",
                conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                    key="segment",
                    operator="equals",
                    type="tag-value",
                    value="prod",
                )],
                rule_number=200,
            ),
        ],
        core_network_configurations=[aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs(
            asn_ranges=["64512-64555"],
            edge_locations=[
                aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                    asn=64512,
                    location="us-east-1",
                ),
                aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                    asn=64513,
                    location="eu-central-1",
                ),
            ],
            vpn_ecmp_support=False,
        )],
        segment_actions=[aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs(
            action="share",
            mode="attachment-route",
            segment="shared",
            share_withs=["*"],
        )],
        segments=[
            aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
                description="Segment for shared services",
                name="shared",
                require_attachment_acceptance=True,
            ),
            aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
                description="Segment for prod services",
                name="prod",
                require_attachment_acceptance=True,
            ),
        ])
    ```

    `data.aws_networkmanager_core_network_policy_document.test.json` will evaluate to:

    ```python
    import pulumi
    ```


    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentAttachmentPolicyArgs']] attachment_policies: In a core network, all attachments use the block argument `attachment_policies` section to map an attachment to a segment. Instead of manually associating a segment to each attachment, attachments use tags, and then the tags are used to associate the attachment to the specified segment. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs']] core_network_configurations: The core network configuration section defines the Regions where a core network should operate. For AWS Regions that are defined in the policy, the core network creates a Core Network Edge where you can connect attachments. After it's created, each Core Network Edge is peered with every other defined Region and is configured with consistent segment and routing across all Regions. Regions cannot be removed until the associated attachments are deleted. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentActionArgs']] segment_actions: A block argument, `segment_actions` define how routing works between segments. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentArgs']] segments: A block argument that defines the different segments in the network. Here you can provide descriptions, change defaults, and provide explicit Regional operational and route filters. The names defined for each segment are used in the `segment_actions` and `attachment_policies` section. Each segment is created, and operates, as a completely separated routing domain. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
    """
    __args__ = dict()
    __args__['attachmentPolicies'] = attachment_policies
    __args__['coreNetworkConfigurations'] = core_network_configurations
    __args__['segmentActions'] = segment_actions
    __args__['segments'] = segments
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:networkmanager/getCoreNetworkPolicyDocument:getCoreNetworkPolicyDocument', __args__, opts=opts, typ=GetCoreNetworkPolicyDocumentResult).value

    return AwaitableGetCoreNetworkPolicyDocumentResult(
        attachment_policies=__ret__.attachment_policies,
        core_network_configurations=__ret__.core_network_configurations,
        id=__ret__.id,
        json=__ret__.json,
        segment_actions=__ret__.segment_actions,
        segments=__ret__.segments,
        version=__ret__.version)


@_utilities.lift_output_func(get_core_network_policy_document)
def get_core_network_policy_document_output(attachment_policies: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentAttachmentPolicyArgs']]]]] = None,
                                            core_network_configurations: Optional[pulumi.Input[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs']]]] = None,
                                            segment_actions: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentActionArgs']]]]] = None,
                                            segments: Optional[pulumi.Input[Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentArgs']]]] = None,
                                            version: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCoreNetworkPolicyDocumentResult]:
    """
    ## Example Usage
    ### Basic Example

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.networkmanager.get_core_network_policy_document(attachment_policies=[
            aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
                action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                    association_method="constant",
                    segment="shared",
                ),
                condition_logic="or",
                conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                    key="segment",
                    operator="equals",
                    type="tag-value",
                    value="shared",
                )],
                rule_number=100,
            ),
            aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
                action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                    association_method="constant",
                    segment="prod",
                ),
                condition_logic="or",
                conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                    key="segment",
                    operator="equals",
                    type="tag-value",
                    value="prod",
                )],
                rule_number=200,
            ),
        ],
        core_network_configurations=[aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs(
            asn_ranges=["64512-64555"],
            edge_locations=[
                aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                    asn=64512,
                    location="us-east-1",
                ),
                aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                    asn=64513,
                    location="eu-central-1",
                ),
            ],
            vpn_ecmp_support=False,
        )],
        segment_actions=[aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs(
            action="share",
            mode="attachment-route",
            segment="shared",
            share_withs=["*"],
        )],
        segments=[
            aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
                description="Segment for shared services",
                name="shared",
                require_attachment_acceptance=True,
            ),
            aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
                description="Segment for prod services",
                name="prod",
                require_attachment_acceptance=True,
            ),
        ])
    ```

    `data.aws_networkmanager_core_network_policy_document.test.json` will evaluate to:

    ```python
    import pulumi
    ```


    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentAttachmentPolicyArgs']] attachment_policies: In a core network, all attachments use the block argument `attachment_policies` section to map an attachment to a segment. Instead of manually associating a segment to each attachment, attachments use tags, and then the tags are used to associate the attachment to the specified segment. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs']] core_network_configurations: The core network configuration section defines the Regions where a core network should operate. For AWS Regions that are defined in the policy, the core network creates a Core Network Edge where you can connect attachments. After it's created, each Core Network Edge is peered with every other defined Region and is configured with consistent segment and routing across all Regions. Regions cannot be removed until the associated attachments are deleted. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentActionArgs']] segment_actions: A block argument, `segment_actions` define how routing works between segments. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
    :param Sequence[pulumi.InputType['GetCoreNetworkPolicyDocumentSegmentArgs']] segments: A block argument that defines the different segments in the network. Here you can provide descriptions, change defaults, and provide explicit Regional operational and route filters. The names defined for each segment are used in the `segment_actions` and `attachment_policies` section. Each segment is created, and operates, as a completely separated routing domain. By default, attachments can only communicate with other attachments in the same segment. Detailed below.
    """
    ...
