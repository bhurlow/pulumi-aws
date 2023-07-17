# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'PolicyExcludeMap',
    'PolicyIncludeMap',
    'PolicySecurityServicePolicyData',
    'PolicySecurityServicePolicyDataPolicyOption',
    'PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy',
    'PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy',
]

@pulumi.output_type
class PolicyExcludeMap(dict):
    def __init__(__self__, *,
                 accounts: Optional[Sequence[str]] = None,
                 orgunits: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] accounts: A list of AWS Organization member Accounts that you want to include for this AWS FMS Policy.
        :param Sequence[str] orgunits: A list of IDs of the AWS Organizational Units that you want to include for this AWS FMS Policy. Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.
               
               You can specify inclusions or exclusions, but not both. If you specify an `include_map`, AWS Firewall Manager applies the policy to all accounts specified by the `include_map`, and does not evaluate any `exclude_map` specifications. If you do not specify an `include_map`, then Firewall Manager applies the policy to all accounts except for those specified by the `exclude_map`.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if orgunits is not None:
            pulumi.set(__self__, "orgunits", orgunits)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[Sequence[str]]:
        """
        A list of AWS Organization member Accounts that you want to include for this AWS FMS Policy.
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def orgunits(self) -> Optional[Sequence[str]]:
        """
        A list of IDs of the AWS Organizational Units that you want to include for this AWS FMS Policy. Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.

        You can specify inclusions or exclusions, but not both. If you specify an `include_map`, AWS Firewall Manager applies the policy to all accounts specified by the `include_map`, and does not evaluate any `exclude_map` specifications. If you do not specify an `include_map`, then Firewall Manager applies the policy to all accounts except for those specified by the `exclude_map`.
        """
        return pulumi.get(self, "orgunits")


@pulumi.output_type
class PolicyIncludeMap(dict):
    def __init__(__self__, *,
                 accounts: Optional[Sequence[str]] = None,
                 orgunits: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] accounts: A list of AWS Organization member Accounts that you want to include for this AWS FMS Policy.
        :param Sequence[str] orgunits: A list of IDs of the AWS Organizational Units that you want to include for this AWS FMS Policy. Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.
               
               You can specify inclusions or exclusions, but not both. If you specify an `include_map`, AWS Firewall Manager applies the policy to all accounts specified by the `include_map`, and does not evaluate any `exclude_map` specifications. If you do not specify an `include_map`, then Firewall Manager applies the policy to all accounts except for those specified by the `exclude_map`.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if orgunits is not None:
            pulumi.set(__self__, "orgunits", orgunits)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[Sequence[str]]:
        """
        A list of AWS Organization member Accounts that you want to include for this AWS FMS Policy.
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def orgunits(self) -> Optional[Sequence[str]]:
        """
        A list of IDs of the AWS Organizational Units that you want to include for this AWS FMS Policy. Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.

        You can specify inclusions or exclusions, but not both. If you specify an `include_map`, AWS Firewall Manager applies the policy to all accounts specified by the `include_map`, and does not evaluate any `exclude_map` specifications. If you do not specify an `include_map`, then Firewall Manager applies the policy to all accounts except for those specified by the `exclude_map`.
        """
        return pulumi.get(self, "orgunits")


@pulumi.output_type
class PolicySecurityServicePolicyData(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "managedServiceData":
            suggest = "managed_service_data"
        elif key == "policyOption":
            suggest = "policy_option"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicySecurityServicePolicyData. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicySecurityServicePolicyData.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicySecurityServicePolicyData.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 managed_service_data: Optional[str] = None,
                 policy_option: Optional['outputs.PolicySecurityServicePolicyDataPolicyOption'] = None):
        """
        :param str type: The service that the policy is using to protect the resources. For the current list of supported types, please refer to the [AWS Firewall Manager SecurityServicePolicyData API Type Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html#fms-Type-SecurityServicePolicyData-Type).
        :param str managed_service_data: Details about the service that are specific to the service type, in JSON format. For service type `SHIELD_ADVANCED`, this is an empty string. Examples depending on `type` can be found in the [AWS Firewall Manager SecurityServicePolicyData API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html).
        :param 'PolicySecurityServicePolicyDataPolicyOptionArgs' policy_option: Contains the Network Firewall firewall policy options to configure a centralized deployment model. Documented below.
        """
        pulumi.set(__self__, "type", type)
        if managed_service_data is not None:
            pulumi.set(__self__, "managed_service_data", managed_service_data)
        if policy_option is not None:
            pulumi.set(__self__, "policy_option", policy_option)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The service that the policy is using to protect the resources. For the current list of supported types, please refer to the [AWS Firewall Manager SecurityServicePolicyData API Type Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html#fms-Type-SecurityServicePolicyData-Type).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="managedServiceData")
    def managed_service_data(self) -> Optional[str]:
        """
        Details about the service that are specific to the service type, in JSON format. For service type `SHIELD_ADVANCED`, this is an empty string. Examples depending on `type` can be found in the [AWS Firewall Manager SecurityServicePolicyData API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html).
        """
        return pulumi.get(self, "managed_service_data")

    @property
    @pulumi.getter(name="policyOption")
    def policy_option(self) -> Optional['outputs.PolicySecurityServicePolicyDataPolicyOption']:
        """
        Contains the Network Firewall firewall policy options to configure a centralized deployment model. Documented below.
        """
        return pulumi.get(self, "policy_option")


@pulumi.output_type
class PolicySecurityServicePolicyDataPolicyOption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkFirewallPolicy":
            suggest = "network_firewall_policy"
        elif key == "thirdPartyFirewallPolicy":
            suggest = "third_party_firewall_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicySecurityServicePolicyDataPolicyOption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicySecurityServicePolicyDataPolicyOption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicySecurityServicePolicyDataPolicyOption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_firewall_policy: Optional['outputs.PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy'] = None,
                 third_party_firewall_policy: Optional['outputs.PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy'] = None):
        """
        :param 'PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicyArgs' network_firewall_policy: Defines the deployment model to use for the firewall policy. Documented below.
        """
        if network_firewall_policy is not None:
            pulumi.set(__self__, "network_firewall_policy", network_firewall_policy)
        if third_party_firewall_policy is not None:
            pulumi.set(__self__, "third_party_firewall_policy", third_party_firewall_policy)

    @property
    @pulumi.getter(name="networkFirewallPolicy")
    def network_firewall_policy(self) -> Optional['outputs.PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy']:
        """
        Defines the deployment model to use for the firewall policy. Documented below.
        """
        return pulumi.get(self, "network_firewall_policy")

    @property
    @pulumi.getter(name="thirdPartyFirewallPolicy")
    def third_party_firewall_policy(self) -> Optional['outputs.PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy']:
        return pulumi.get(self, "third_party_firewall_policy")


@pulumi.output_type
class PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "firewallDeploymentModel":
            suggest = "firewall_deployment_model"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 firewall_deployment_model: Optional[str] = None):
        """
        :param str firewall_deployment_model: Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
        """
        if firewall_deployment_model is not None:
            pulumi.set(__self__, "firewall_deployment_model", firewall_deployment_model)

    @property
    @pulumi.getter(name="firewallDeploymentModel")
    def firewall_deployment_model(self) -> Optional[str]:
        """
        Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
        """
        return pulumi.get(self, "firewall_deployment_model")


@pulumi.output_type
class PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "firewallDeploymentModel":
            suggest = "firewall_deployment_model"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 firewall_deployment_model: Optional[str] = None):
        """
        :param str firewall_deployment_model: Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
        """
        if firewall_deployment_model is not None:
            pulumi.set(__self__, "firewall_deployment_model", firewall_deployment_model)

    @property
    @pulumi.getter(name="firewallDeploymentModel")
    def firewall_deployment_model(self) -> Optional[str]:
        """
        Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
        """
        return pulumi.get(self, "firewall_deployment_model")


