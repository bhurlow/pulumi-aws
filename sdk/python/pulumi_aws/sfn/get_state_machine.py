# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetStateMachineResult',
    'AwaitableGetStateMachineResult',
    'get_state_machine',
    'get_state_machine_output',
]

@pulumi.output_type
class GetStateMachineResult:
    """
    A collection of values returned by getStateMachine.
    """
    def __init__(__self__, arn=None, creation_date=None, definition=None, description=None, id=None, name=None, revision_id=None, role_arn=None, status=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if definition and not isinstance(definition, str):
            raise TypeError("Expected argument 'definition' to be a str")
        pulumi.set(__self__, "definition", definition)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if revision_id and not isinstance(revision_id, str):
            raise TypeError("Expected argument 'revision_id' to be a str")
        pulumi.set(__self__, "revision_id", revision_id)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Set to the arn of the state function.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Date the state machine was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def definition(self) -> str:
        """
        Set to the state machine definition.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> str:
        """
        The revision identifier for the state machine.
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        Set to the role_arn used by the state function.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Set to the current status of the state machine.
        """
        return pulumi.get(self, "status")


class AwaitableGetStateMachineResult(GetStateMachineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStateMachineResult(
            arn=self.arn,
            creation_date=self.creation_date,
            definition=self.definition,
            description=self.description,
            id=self.id,
            name=self.name,
            revision_id=self.revision_id,
            role_arn=self.role_arn,
            status=self.status)


def get_state_machine(name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStateMachineResult:
    """
    Use this data source to get the ARN of a State Machine in AWS Step
    Function (SFN). By using this data source, you can reference a
    state machine without having to hard code the ARNs as input.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.sfn.get_state_machine(name="an_example_sfn_name")
    ```


    :param str name: Friendly name of the state machine to match.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:sfn/getStateMachine:getStateMachine', __args__, opts=opts, typ=GetStateMachineResult).value

    return AwaitableGetStateMachineResult(
        arn=pulumi.get(__ret__, 'arn'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        definition=pulumi.get(__ret__, 'definition'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        revision_id=pulumi.get(__ret__, 'revision_id'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_state_machine)
def get_state_machine_output(name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStateMachineResult]:
    """
    Use this data source to get the ARN of a State Machine in AWS Step
    Function (SFN). By using this data source, you can reference a
    state machine without having to hard code the ARNs as input.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.sfn.get_state_machine(name="an_example_sfn_name")
    ```


    :param str name: Friendly name of the state machine to match.
    """
    ...
