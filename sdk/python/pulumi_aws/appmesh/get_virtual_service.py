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
    'GetVirtualServiceResult',
    'AwaitableGetVirtualServiceResult',
    'get_virtual_service',
    'get_virtual_service_output',
]

@pulumi.output_type
class GetVirtualServiceResult:
    """
    A collection of values returned by getVirtualService.
    """
    def __init__(__self__, arn=None, created_date=None, id=None, last_updated_date=None, mesh_name=None, mesh_owner=None, name=None, resource_owner=None, specs=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_updated_date and not isinstance(last_updated_date, str):
            raise TypeError("Expected argument 'last_updated_date' to be a str")
        pulumi.set(__self__, "last_updated_date", last_updated_date)
        if mesh_name and not isinstance(mesh_name, str):
            raise TypeError("Expected argument 'mesh_name' to be a str")
        pulumi.set(__self__, "mesh_name", mesh_name)
        if mesh_owner and not isinstance(mesh_owner, str):
            raise TypeError("Expected argument 'mesh_owner' to be a str")
        pulumi.set(__self__, "mesh_owner", mesh_owner)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_owner and not isinstance(resource_owner, str):
            raise TypeError("Expected argument 'resource_owner' to be a str")
        pulumi.set(__self__, "resource_owner", resource_owner)
        if specs and not isinstance(specs, list):
            raise TypeError("Expected argument 'specs' to be a list")
        pulumi.set(__self__, "specs", specs)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        ARN of the virtual service.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        Creation date of the virtual service.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> str:
        """
        Last update date of the virtual service.
        """
        return pulumi.get(self, "last_updated_date")

    @property
    @pulumi.getter(name="meshName")
    def mesh_name(self) -> str:
        return pulumi.get(self, "mesh_name")

    @property
    @pulumi.getter(name="meshOwner")
    def mesh_owner(self) -> str:
        return pulumi.get(self, "mesh_owner")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceOwner")
    def resource_owner(self) -> str:
        """
        Resource owner's AWS account ID.
        """
        return pulumi.get(self, "resource_owner")

    @property
    @pulumi.getter
    def specs(self) -> Sequence['outputs.GetVirtualServiceSpecResult']:
        """
        Virtual service specification
        """
        return pulumi.get(self, "specs")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Map of tags.
        """
        return pulumi.get(self, "tags")


class AwaitableGetVirtualServiceResult(GetVirtualServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualServiceResult(
            arn=self.arn,
            created_date=self.created_date,
            id=self.id,
            last_updated_date=self.last_updated_date,
            mesh_name=self.mesh_name,
            mesh_owner=self.mesh_owner,
            name=self.name,
            resource_owner=self.resource_owner,
            specs=self.specs,
            tags=self.tags)


def get_virtual_service(mesh_name: Optional[str] = None,
                        mesh_owner: Optional[str] = None,
                        name: Optional[str] = None,
                        tags: Optional[Mapping[str, str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualServiceResult:
    """
    The App Mesh Virtual Service data source allows details of an App Mesh Virtual Service to be retrieved by its name, mesh_name, and optionally the mesh_owner.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.appmesh.get_virtual_service(mesh_name="example-mesh",
        name="example.mesh.local")
    ```

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_caller_identity()
    test = aws.appmesh.get_virtual_service(name="example.mesh.local",
        mesh_name="example-mesh",
        mesh_owner=current.account_id)
    ```


    :param str mesh_name: Name of the service mesh in which the virtual service exists.
    :param str mesh_owner: AWS account ID of the service mesh's owner.
    :param str name: Name of the virtual service.
    :param Mapping[str, str] tags: Map of tags.
    """
    __args__ = dict()
    __args__['meshName'] = mesh_name
    __args__['meshOwner'] = mesh_owner
    __args__['name'] = name
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:appmesh/getVirtualService:getVirtualService', __args__, opts=opts, typ=GetVirtualServiceResult).value

    return AwaitableGetVirtualServiceResult(
        arn=__ret__.arn,
        created_date=__ret__.created_date,
        id=__ret__.id,
        last_updated_date=__ret__.last_updated_date,
        mesh_name=__ret__.mesh_name,
        mesh_owner=__ret__.mesh_owner,
        name=__ret__.name,
        resource_owner=__ret__.resource_owner,
        specs=__ret__.specs,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_virtual_service)
def get_virtual_service_output(mesh_name: Optional[pulumi.Input[str]] = None,
                               mesh_owner: Optional[pulumi.Input[Optional[str]]] = None,
                               name: Optional[pulumi.Input[str]] = None,
                               tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualServiceResult]:
    """
    The App Mesh Virtual Service data source allows details of an App Mesh Virtual Service to be retrieved by its name, mesh_name, and optionally the mesh_owner.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.appmesh.get_virtual_service(mesh_name="example-mesh",
        name="example.mesh.local")
    ```

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_caller_identity()
    test = aws.appmesh.get_virtual_service(name="example.mesh.local",
        mesh_name="example-mesh",
        mesh_owner=current.account_id)
    ```


    :param str mesh_name: Name of the service mesh in which the virtual service exists.
    :param str mesh_owner: AWS account ID of the service mesh's owner.
    :param str name: Name of the virtual service.
    :param Mapping[str, str] tags: Map of tags.
    """
    ...
