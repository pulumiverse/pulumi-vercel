# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetTeamConfigResult',
    'AwaitableGetTeamConfigResult',
    'get_team_config',
    'get_team_config_output',
]

@pulumi.output_type
class GetTeamConfigResult:
    """
    A collection of values returned by getTeamConfig.
    """
    def __init__(__self__, description=None, email_domain=None, enable_preview_feedback=None, enable_production_feedback=None, hide_ip_addresses=None, hide_ip_addresses_in_log_drains=None, id=None, invite_code=None, name=None, preview_deployment_suffix=None, remote_caching=None, saml=None, sensitive_environment_variable_policy=None, slug=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if email_domain and not isinstance(email_domain, str):
            raise TypeError("Expected argument 'email_domain' to be a str")
        pulumi.set(__self__, "email_domain", email_domain)
        if enable_preview_feedback and not isinstance(enable_preview_feedback, str):
            raise TypeError("Expected argument 'enable_preview_feedback' to be a str")
        pulumi.set(__self__, "enable_preview_feedback", enable_preview_feedback)
        if enable_production_feedback and not isinstance(enable_production_feedback, str):
            raise TypeError("Expected argument 'enable_production_feedback' to be a str")
        pulumi.set(__self__, "enable_production_feedback", enable_production_feedback)
        if hide_ip_addresses and not isinstance(hide_ip_addresses, bool):
            raise TypeError("Expected argument 'hide_ip_addresses' to be a bool")
        pulumi.set(__self__, "hide_ip_addresses", hide_ip_addresses)
        if hide_ip_addresses_in_log_drains and not isinstance(hide_ip_addresses_in_log_drains, bool):
            raise TypeError("Expected argument 'hide_ip_addresses_in_log_drains' to be a bool")
        pulumi.set(__self__, "hide_ip_addresses_in_log_drains", hide_ip_addresses_in_log_drains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invite_code and not isinstance(invite_code, str):
            raise TypeError("Expected argument 'invite_code' to be a str")
        pulumi.set(__self__, "invite_code", invite_code)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if preview_deployment_suffix and not isinstance(preview_deployment_suffix, str):
            raise TypeError("Expected argument 'preview_deployment_suffix' to be a str")
        pulumi.set(__self__, "preview_deployment_suffix", preview_deployment_suffix)
        if remote_caching and not isinstance(remote_caching, dict):
            raise TypeError("Expected argument 'remote_caching' to be a dict")
        pulumi.set(__self__, "remote_caching", remote_caching)
        if saml and not isinstance(saml, dict):
            raise TypeError("Expected argument 'saml' to be a dict")
        pulumi.set(__self__, "saml", saml)
        if sensitive_environment_variable_policy and not isinstance(sensitive_environment_variable_policy, str):
            raise TypeError("Expected argument 'sensitive_environment_variable_policy' to be a str")
        pulumi.set(__self__, "sensitive_environment_variable_policy", sensitive_environment_variable_policy)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailDomain")
    def email_domain(self) -> str:
        return pulumi.get(self, "email_domain")

    @property
    @pulumi.getter(name="enablePreviewFeedback")
    def enable_preview_feedback(self) -> str:
        return pulumi.get(self, "enable_preview_feedback")

    @property
    @pulumi.getter(name="enableProductionFeedback")
    def enable_production_feedback(self) -> str:
        return pulumi.get(self, "enable_production_feedback")

    @property
    @pulumi.getter(name="hideIpAddresses")
    def hide_ip_addresses(self) -> bool:
        return pulumi.get(self, "hide_ip_addresses")

    @property
    @pulumi.getter(name="hideIpAddressesInLogDrains")
    def hide_ip_addresses_in_log_drains(self) -> bool:
        return pulumi.get(self, "hide_ip_addresses_in_log_drains")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inviteCode")
    def invite_code(self) -> str:
        return pulumi.get(self, "invite_code")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="previewDeploymentSuffix")
    def preview_deployment_suffix(self) -> str:
        return pulumi.get(self, "preview_deployment_suffix")

    @property
    @pulumi.getter(name="remoteCaching")
    def remote_caching(self) -> 'outputs.GetTeamConfigRemoteCachingResult':
        return pulumi.get(self, "remote_caching")

    @property
    @pulumi.getter
    def saml(self) -> 'outputs.GetTeamConfigSamlResult':
        return pulumi.get(self, "saml")

    @property
    @pulumi.getter(name="sensitiveEnvironmentVariablePolicy")
    def sensitive_environment_variable_policy(self) -> str:
        return pulumi.get(self, "sensitive_environment_variable_policy")

    @property
    @pulumi.getter
    def slug(self) -> str:
        return pulumi.get(self, "slug")


class AwaitableGetTeamConfigResult(GetTeamConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTeamConfigResult(
            description=self.description,
            email_domain=self.email_domain,
            enable_preview_feedback=self.enable_preview_feedback,
            enable_production_feedback=self.enable_production_feedback,
            hide_ip_addresses=self.hide_ip_addresses,
            hide_ip_addresses_in_log_drains=self.hide_ip_addresses_in_log_drains,
            id=self.id,
            invite_code=self.invite_code,
            name=self.name,
            preview_deployment_suffix=self.preview_deployment_suffix,
            remote_caching=self.remote_caching,
            saml=self.saml,
            sensitive_environment_variable_policy=self.sensitive_environment_variable_policy,
            slug=self.slug)


def get_team_config(id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTeamConfigResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vercel:index/getTeamConfig:getTeamConfig', __args__, opts=opts, typ=GetTeamConfigResult).value

    return AwaitableGetTeamConfigResult(
        description=pulumi.get(__ret__, 'description'),
        email_domain=pulumi.get(__ret__, 'email_domain'),
        enable_preview_feedback=pulumi.get(__ret__, 'enable_preview_feedback'),
        enable_production_feedback=pulumi.get(__ret__, 'enable_production_feedback'),
        hide_ip_addresses=pulumi.get(__ret__, 'hide_ip_addresses'),
        hide_ip_addresses_in_log_drains=pulumi.get(__ret__, 'hide_ip_addresses_in_log_drains'),
        id=pulumi.get(__ret__, 'id'),
        invite_code=pulumi.get(__ret__, 'invite_code'),
        name=pulumi.get(__ret__, 'name'),
        preview_deployment_suffix=pulumi.get(__ret__, 'preview_deployment_suffix'),
        remote_caching=pulumi.get(__ret__, 'remote_caching'),
        saml=pulumi.get(__ret__, 'saml'),
        sensitive_environment_variable_policy=pulumi.get(__ret__, 'sensitive_environment_variable_policy'),
        slug=pulumi.get(__ret__, 'slug'))
def get_team_config_output(id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTeamConfigResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vercel:index/getTeamConfig:getTeamConfig', __args__, opts=opts, typ=GetTeamConfigResult)
    return __ret__.apply(lambda __response__: GetTeamConfigResult(
        description=pulumi.get(__response__, 'description'),
        email_domain=pulumi.get(__response__, 'email_domain'),
        enable_preview_feedback=pulumi.get(__response__, 'enable_preview_feedback'),
        enable_production_feedback=pulumi.get(__response__, 'enable_production_feedback'),
        hide_ip_addresses=pulumi.get(__response__, 'hide_ip_addresses'),
        hide_ip_addresses_in_log_drains=pulumi.get(__response__, 'hide_ip_addresses_in_log_drains'),
        id=pulumi.get(__response__, 'id'),
        invite_code=pulumi.get(__response__, 'invite_code'),
        name=pulumi.get(__response__, 'name'),
        preview_deployment_suffix=pulumi.get(__response__, 'preview_deployment_suffix'),
        remote_caching=pulumi.get(__response__, 'remote_caching'),
        saml=pulumi.get(__response__, 'saml'),
        sensitive_environment_variable_policy=pulumi.get(__response__, 'sensitive_environment_variable_policy'),
        slug=pulumi.get(__response__, 'slug')))
