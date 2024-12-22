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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, auto_assign_custom_domains=None, automatically_expose_system_environment_variables=None, build_command=None, customer_success_code_visibility=None, dev_command=None, directory_listing=None, environments=None, framework=None, function_failover=None, git_comments=None, git_fork_protection=None, git_lfs=None, git_repository=None, id=None, ignore_command=None, install_command=None, name=None, oidc_token_config=None, options_allowlist=None, output_directory=None, password_protection=None, preview_comments=None, prioritise_production_builds=None, protection_bypass_for_automation=None, protection_bypass_for_automation_secret=None, public_source=None, resource_config=None, root_directory=None, serverless_function_region=None, skew_protection=None, team_id=None, trusted_ips=None, vercel_authentication=None):
        if auto_assign_custom_domains and not isinstance(auto_assign_custom_domains, bool):
            raise TypeError("Expected argument 'auto_assign_custom_domains' to be a bool")
        pulumi.set(__self__, "auto_assign_custom_domains", auto_assign_custom_domains)
        if automatically_expose_system_environment_variables and not isinstance(automatically_expose_system_environment_variables, bool):
            raise TypeError("Expected argument 'automatically_expose_system_environment_variables' to be a bool")
        pulumi.set(__self__, "automatically_expose_system_environment_variables", automatically_expose_system_environment_variables)
        if build_command and not isinstance(build_command, str):
            raise TypeError("Expected argument 'build_command' to be a str")
        pulumi.set(__self__, "build_command", build_command)
        if customer_success_code_visibility and not isinstance(customer_success_code_visibility, bool):
            raise TypeError("Expected argument 'customer_success_code_visibility' to be a bool")
        pulumi.set(__self__, "customer_success_code_visibility", customer_success_code_visibility)
        if dev_command and not isinstance(dev_command, str):
            raise TypeError("Expected argument 'dev_command' to be a str")
        pulumi.set(__self__, "dev_command", dev_command)
        if directory_listing and not isinstance(directory_listing, bool):
            raise TypeError("Expected argument 'directory_listing' to be a bool")
        pulumi.set(__self__, "directory_listing", directory_listing)
        if environments and not isinstance(environments, list):
            raise TypeError("Expected argument 'environments' to be a list")
        pulumi.set(__self__, "environments", environments)
        if framework and not isinstance(framework, str):
            raise TypeError("Expected argument 'framework' to be a str")
        pulumi.set(__self__, "framework", framework)
        if function_failover and not isinstance(function_failover, bool):
            raise TypeError("Expected argument 'function_failover' to be a bool")
        pulumi.set(__self__, "function_failover", function_failover)
        if git_comments and not isinstance(git_comments, dict):
            raise TypeError("Expected argument 'git_comments' to be a dict")
        pulumi.set(__self__, "git_comments", git_comments)
        if git_fork_protection and not isinstance(git_fork_protection, bool):
            raise TypeError("Expected argument 'git_fork_protection' to be a bool")
        pulumi.set(__self__, "git_fork_protection", git_fork_protection)
        if git_lfs and not isinstance(git_lfs, bool):
            raise TypeError("Expected argument 'git_lfs' to be a bool")
        pulumi.set(__self__, "git_lfs", git_lfs)
        if git_repository and not isinstance(git_repository, dict):
            raise TypeError("Expected argument 'git_repository' to be a dict")
        pulumi.set(__self__, "git_repository", git_repository)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_command and not isinstance(ignore_command, str):
            raise TypeError("Expected argument 'ignore_command' to be a str")
        pulumi.set(__self__, "ignore_command", ignore_command)
        if install_command and not isinstance(install_command, str):
            raise TypeError("Expected argument 'install_command' to be a str")
        pulumi.set(__self__, "install_command", install_command)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if oidc_token_config and not isinstance(oidc_token_config, dict):
            raise TypeError("Expected argument 'oidc_token_config' to be a dict")
        pulumi.set(__self__, "oidc_token_config", oidc_token_config)
        if options_allowlist and not isinstance(options_allowlist, dict):
            raise TypeError("Expected argument 'options_allowlist' to be a dict")
        pulumi.set(__self__, "options_allowlist", options_allowlist)
        if output_directory and not isinstance(output_directory, str):
            raise TypeError("Expected argument 'output_directory' to be a str")
        pulumi.set(__self__, "output_directory", output_directory)
        if password_protection and not isinstance(password_protection, dict):
            raise TypeError("Expected argument 'password_protection' to be a dict")
        pulumi.set(__self__, "password_protection", password_protection)
        if preview_comments and not isinstance(preview_comments, bool):
            raise TypeError("Expected argument 'preview_comments' to be a bool")
        pulumi.set(__self__, "preview_comments", preview_comments)
        if prioritise_production_builds and not isinstance(prioritise_production_builds, bool):
            raise TypeError("Expected argument 'prioritise_production_builds' to be a bool")
        pulumi.set(__self__, "prioritise_production_builds", prioritise_production_builds)
        if protection_bypass_for_automation and not isinstance(protection_bypass_for_automation, bool):
            raise TypeError("Expected argument 'protection_bypass_for_automation' to be a bool")
        pulumi.set(__self__, "protection_bypass_for_automation", protection_bypass_for_automation)
        if protection_bypass_for_automation_secret and not isinstance(protection_bypass_for_automation_secret, str):
            raise TypeError("Expected argument 'protection_bypass_for_automation_secret' to be a str")
        pulumi.set(__self__, "protection_bypass_for_automation_secret", protection_bypass_for_automation_secret)
        if public_source and not isinstance(public_source, bool):
            raise TypeError("Expected argument 'public_source' to be a bool")
        pulumi.set(__self__, "public_source", public_source)
        if resource_config and not isinstance(resource_config, dict):
            raise TypeError("Expected argument 'resource_config' to be a dict")
        pulumi.set(__self__, "resource_config", resource_config)
        if root_directory and not isinstance(root_directory, str):
            raise TypeError("Expected argument 'root_directory' to be a str")
        pulumi.set(__self__, "root_directory", root_directory)
        if serverless_function_region and not isinstance(serverless_function_region, str):
            raise TypeError("Expected argument 'serverless_function_region' to be a str")
        pulumi.set(__self__, "serverless_function_region", serverless_function_region)
        if skew_protection and not isinstance(skew_protection, str):
            raise TypeError("Expected argument 'skew_protection' to be a str")
        pulumi.set(__self__, "skew_protection", skew_protection)
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        pulumi.set(__self__, "team_id", team_id)
        if trusted_ips and not isinstance(trusted_ips, dict):
            raise TypeError("Expected argument 'trusted_ips' to be a dict")
        pulumi.set(__self__, "trusted_ips", trusted_ips)
        if vercel_authentication and not isinstance(vercel_authentication, dict):
            raise TypeError("Expected argument 'vercel_authentication' to be a dict")
        pulumi.set(__self__, "vercel_authentication", vercel_authentication)

    @property
    @pulumi.getter(name="autoAssignCustomDomains")
    def auto_assign_custom_domains(self) -> bool:
        return pulumi.get(self, "auto_assign_custom_domains")

    @property
    @pulumi.getter(name="automaticallyExposeSystemEnvironmentVariables")
    def automatically_expose_system_environment_variables(self) -> bool:
        return pulumi.get(self, "automatically_expose_system_environment_variables")

    @property
    @pulumi.getter(name="buildCommand")
    def build_command(self) -> str:
        return pulumi.get(self, "build_command")

    @property
    @pulumi.getter(name="customerSuccessCodeVisibility")
    def customer_success_code_visibility(self) -> bool:
        return pulumi.get(self, "customer_success_code_visibility")

    @property
    @pulumi.getter(name="devCommand")
    def dev_command(self) -> str:
        return pulumi.get(self, "dev_command")

    @property
    @pulumi.getter(name="directoryListing")
    def directory_listing(self) -> bool:
        return pulumi.get(self, "directory_listing")

    @property
    @pulumi.getter
    def environments(self) -> Sequence['outputs.GetProjectEnvironmentResult']:
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def framework(self) -> str:
        return pulumi.get(self, "framework")

    @property
    @pulumi.getter(name="functionFailover")
    def function_failover(self) -> bool:
        return pulumi.get(self, "function_failover")

    @property
    @pulumi.getter(name="gitComments")
    def git_comments(self) -> 'outputs.GetProjectGitCommentsResult':
        return pulumi.get(self, "git_comments")

    @property
    @pulumi.getter(name="gitForkProtection")
    def git_fork_protection(self) -> bool:
        return pulumi.get(self, "git_fork_protection")

    @property
    @pulumi.getter(name="gitLfs")
    def git_lfs(self) -> bool:
        return pulumi.get(self, "git_lfs")

    @property
    @pulumi.getter(name="gitRepository")
    def git_repository(self) -> 'outputs.GetProjectGitRepositoryResult':
        return pulumi.get(self, "git_repository")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreCommand")
    def ignore_command(self) -> str:
        return pulumi.get(self, "ignore_command")

    @property
    @pulumi.getter(name="installCommand")
    def install_command(self) -> str:
        return pulumi.get(self, "install_command")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oidcTokenConfig")
    def oidc_token_config(self) -> 'outputs.GetProjectOidcTokenConfigResult':
        return pulumi.get(self, "oidc_token_config")

    @property
    @pulumi.getter(name="optionsAllowlist")
    def options_allowlist(self) -> 'outputs.GetProjectOptionsAllowlistResult':
        return pulumi.get(self, "options_allowlist")

    @property
    @pulumi.getter(name="outputDirectory")
    def output_directory(self) -> str:
        return pulumi.get(self, "output_directory")

    @property
    @pulumi.getter(name="passwordProtection")
    def password_protection(self) -> 'outputs.GetProjectPasswordProtectionResult':
        return pulumi.get(self, "password_protection")

    @property
    @pulumi.getter(name="previewComments")
    def preview_comments(self) -> bool:
        return pulumi.get(self, "preview_comments")

    @property
    @pulumi.getter(name="prioritiseProductionBuilds")
    def prioritise_production_builds(self) -> bool:
        return pulumi.get(self, "prioritise_production_builds")

    @property
    @pulumi.getter(name="protectionBypassForAutomation")
    def protection_bypass_for_automation(self) -> bool:
        return pulumi.get(self, "protection_bypass_for_automation")

    @property
    @pulumi.getter(name="protectionBypassForAutomationSecret")
    def protection_bypass_for_automation_secret(self) -> str:
        return pulumi.get(self, "protection_bypass_for_automation_secret")

    @property
    @pulumi.getter(name="publicSource")
    def public_source(self) -> bool:
        return pulumi.get(self, "public_source")

    @property
    @pulumi.getter(name="resourceConfig")
    def resource_config(self) -> 'outputs.GetProjectResourceConfigResult':
        return pulumi.get(self, "resource_config")

    @property
    @pulumi.getter(name="rootDirectory")
    def root_directory(self) -> str:
        return pulumi.get(self, "root_directory")

    @property
    @pulumi.getter(name="serverlessFunctionRegion")
    def serverless_function_region(self) -> str:
        return pulumi.get(self, "serverless_function_region")

    @property
    @pulumi.getter(name="skewProtection")
    def skew_protection(self) -> str:
        return pulumi.get(self, "skew_protection")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="trustedIps")
    def trusted_ips(self) -> 'outputs.GetProjectTrustedIpsResult':
        return pulumi.get(self, "trusted_ips")

    @property
    @pulumi.getter(name="vercelAuthentication")
    def vercel_authentication(self) -> 'outputs.GetProjectVercelAuthenticationResult':
        return pulumi.get(self, "vercel_authentication")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            auto_assign_custom_domains=self.auto_assign_custom_domains,
            automatically_expose_system_environment_variables=self.automatically_expose_system_environment_variables,
            build_command=self.build_command,
            customer_success_code_visibility=self.customer_success_code_visibility,
            dev_command=self.dev_command,
            directory_listing=self.directory_listing,
            environments=self.environments,
            framework=self.framework,
            function_failover=self.function_failover,
            git_comments=self.git_comments,
            git_fork_protection=self.git_fork_protection,
            git_lfs=self.git_lfs,
            git_repository=self.git_repository,
            id=self.id,
            ignore_command=self.ignore_command,
            install_command=self.install_command,
            name=self.name,
            oidc_token_config=self.oidc_token_config,
            options_allowlist=self.options_allowlist,
            output_directory=self.output_directory,
            password_protection=self.password_protection,
            preview_comments=self.preview_comments,
            prioritise_production_builds=self.prioritise_production_builds,
            protection_bypass_for_automation=self.protection_bypass_for_automation,
            protection_bypass_for_automation_secret=self.protection_bypass_for_automation_secret,
            public_source=self.public_source,
            resource_config=self.resource_config,
            root_directory=self.root_directory,
            serverless_function_region=self.serverless_function_region,
            skew_protection=self.skew_protection,
            team_id=self.team_id,
            trusted_ips=self.trusted_ips,
            vercel_authentication=self.vercel_authentication)


def get_project(name: Optional[str] = None,
                team_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vercel:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        auto_assign_custom_domains=pulumi.get(__ret__, 'auto_assign_custom_domains'),
        automatically_expose_system_environment_variables=pulumi.get(__ret__, 'automatically_expose_system_environment_variables'),
        build_command=pulumi.get(__ret__, 'build_command'),
        customer_success_code_visibility=pulumi.get(__ret__, 'customer_success_code_visibility'),
        dev_command=pulumi.get(__ret__, 'dev_command'),
        directory_listing=pulumi.get(__ret__, 'directory_listing'),
        environments=pulumi.get(__ret__, 'environments'),
        framework=pulumi.get(__ret__, 'framework'),
        function_failover=pulumi.get(__ret__, 'function_failover'),
        git_comments=pulumi.get(__ret__, 'git_comments'),
        git_fork_protection=pulumi.get(__ret__, 'git_fork_protection'),
        git_lfs=pulumi.get(__ret__, 'git_lfs'),
        git_repository=pulumi.get(__ret__, 'git_repository'),
        id=pulumi.get(__ret__, 'id'),
        ignore_command=pulumi.get(__ret__, 'ignore_command'),
        install_command=pulumi.get(__ret__, 'install_command'),
        name=pulumi.get(__ret__, 'name'),
        oidc_token_config=pulumi.get(__ret__, 'oidc_token_config'),
        options_allowlist=pulumi.get(__ret__, 'options_allowlist'),
        output_directory=pulumi.get(__ret__, 'output_directory'),
        password_protection=pulumi.get(__ret__, 'password_protection'),
        preview_comments=pulumi.get(__ret__, 'preview_comments'),
        prioritise_production_builds=pulumi.get(__ret__, 'prioritise_production_builds'),
        protection_bypass_for_automation=pulumi.get(__ret__, 'protection_bypass_for_automation'),
        protection_bypass_for_automation_secret=pulumi.get(__ret__, 'protection_bypass_for_automation_secret'),
        public_source=pulumi.get(__ret__, 'public_source'),
        resource_config=pulumi.get(__ret__, 'resource_config'),
        root_directory=pulumi.get(__ret__, 'root_directory'),
        serverless_function_region=pulumi.get(__ret__, 'serverless_function_region'),
        skew_protection=pulumi.get(__ret__, 'skew_protection'),
        team_id=pulumi.get(__ret__, 'team_id'),
        trusted_ips=pulumi.get(__ret__, 'trusted_ips'),
        vercel_authentication=pulumi.get(__ret__, 'vercel_authentication'))
def get_project_output(name: Optional[pulumi.Input[str]] = None,
                       team_id: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vercel:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult)
    return __ret__.apply(lambda __response__: GetProjectResult(
        auto_assign_custom_domains=pulumi.get(__response__, 'auto_assign_custom_domains'),
        automatically_expose_system_environment_variables=pulumi.get(__response__, 'automatically_expose_system_environment_variables'),
        build_command=pulumi.get(__response__, 'build_command'),
        customer_success_code_visibility=pulumi.get(__response__, 'customer_success_code_visibility'),
        dev_command=pulumi.get(__response__, 'dev_command'),
        directory_listing=pulumi.get(__response__, 'directory_listing'),
        environments=pulumi.get(__response__, 'environments'),
        framework=pulumi.get(__response__, 'framework'),
        function_failover=pulumi.get(__response__, 'function_failover'),
        git_comments=pulumi.get(__response__, 'git_comments'),
        git_fork_protection=pulumi.get(__response__, 'git_fork_protection'),
        git_lfs=pulumi.get(__response__, 'git_lfs'),
        git_repository=pulumi.get(__response__, 'git_repository'),
        id=pulumi.get(__response__, 'id'),
        ignore_command=pulumi.get(__response__, 'ignore_command'),
        install_command=pulumi.get(__response__, 'install_command'),
        name=pulumi.get(__response__, 'name'),
        oidc_token_config=pulumi.get(__response__, 'oidc_token_config'),
        options_allowlist=pulumi.get(__response__, 'options_allowlist'),
        output_directory=pulumi.get(__response__, 'output_directory'),
        password_protection=pulumi.get(__response__, 'password_protection'),
        preview_comments=pulumi.get(__response__, 'preview_comments'),
        prioritise_production_builds=pulumi.get(__response__, 'prioritise_production_builds'),
        protection_bypass_for_automation=pulumi.get(__response__, 'protection_bypass_for_automation'),
        protection_bypass_for_automation_secret=pulumi.get(__response__, 'protection_bypass_for_automation_secret'),
        public_source=pulumi.get(__response__, 'public_source'),
        resource_config=pulumi.get(__response__, 'resource_config'),
        root_directory=pulumi.get(__response__, 'root_directory'),
        serverless_function_region=pulumi.get(__response__, 'serverless_function_region'),
        skew_protection=pulumi.get(__response__, 'skew_protection'),
        team_id=pulumi.get(__response__, 'team_id'),
        trusted_ips=pulumi.get(__response__, 'trusted_ips'),
        vercel_authentication=pulumi.get(__response__, 'vercel_authentication')))
