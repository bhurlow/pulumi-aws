# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .accelerator import *
from .endpoint_group import *
from .listener import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:globalaccelerator/accelerator:Accelerator":
                return Accelerator(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:globalaccelerator/endpointGroup:EndpointGroup":
                return EndpointGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:globalaccelerator/listener:Listener":
                return Listener(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "globalaccelerator/accelerator", _module_instance)
    pulumi.runtime.register_resource_module("aws", "globalaccelerator/endpointGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "globalaccelerator/listener", _module_instance)

_register_module()
