# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_secret import *
from .get_secret_rotation import *
from .get_secret_version import *
from .secret import *
from .secret_policy import *
from .secret_rotation import *
from .secret_version import *
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
            if typ == "aws:secretsmanager/secret:Secret":
                return Secret(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:secretsmanager/secretPolicy:SecretPolicy":
                return SecretPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:secretsmanager/secretRotation:SecretRotation":
                return SecretRotation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:secretsmanager/secretVersion:SecretVersion":
                return SecretVersion(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "secretsmanager/secret", _module_instance)
    pulumi.runtime.register_resource_module("aws", "secretsmanager/secretPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "secretsmanager/secretRotation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "secretsmanager/secretVersion", _module_instance)

_register_module()
