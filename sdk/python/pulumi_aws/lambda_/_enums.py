# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'Runtime',
]


class Runtime(str, Enum):
    """
    See https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
    """
    DOTNET6 = "dotnet6"
    DOTNET7 = "dotnet7"
    DOTNET8 = "dotnet8"
    JAVA11 = "java11"
    JAVA17 = "java17"
    JAVA21 = "java21"
    JAVA8_AL2 = "java8.al2"
    NODE_JS16D_X = "nodejs16.x"
    NODE_JS18D_X = "nodejs18.x"
    NODE_JS20D_X = "nodejs20.x"
    CUSTOM_AL2 = "provided.al2"
    CUSTOM_AL2023 = "provided.al2023"
    PYTHON3D10 = "python3.10"
    PYTHON3D11 = "python3.11"
    PYTHON3D12 = "python3.12"
    PYTHON3D8 = "python3.8"
    PYTHON3D9 = "python3.9"
    RUBY3D2 = "ruby3.2"
    DOTNET5D0 = "dotnet5.0"
    DOTNET_CORE2D1 = "dotnetcore2.1"
    DOTNET_CORE3D1 = "dotnetcore3.1"
    GO1DX = "go1.x"
    JAVA8 = "java8"
    NODE_JS10D_X = "nodejs10.x"
    NODE_JS12D_X = "nodejs12.x"
    NODE_JS14D_X = "nodejs14.x"
    CUSTOM = "provided"
    PYTHON2D7 = "python2.7"
    PYTHON3D6 = "python3.6"
    PYTHON3D7 = "python3.7"
    RUBY2D5 = "ruby2.5"
    RUBY2D7 = "ruby2.7"
