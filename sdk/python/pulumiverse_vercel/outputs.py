# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'DeploymentProjectSettings',
    'DnsRecordSrv',
    'ProjectEnvironment',
    'ProjectGitRepository',
    'ProjectPasswordProtection',
    'ProjectVercelAuthentication',
    'GetProjectEnvironmentResult',
    'GetProjectGitRepositoryResult',
    'GetProjectPasswordProtectionResult',
    'GetProjectVercelAuthenticationResult',
]

@pulumi.output_type
class DeploymentProjectSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "buildCommand":
            suggest = "build_command"
        elif key == "installCommand":
            suggest = "install_command"
        elif key == "outputDirectory":
            suggest = "output_directory"
        elif key == "rootDirectory":
            suggest = "root_directory"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DeploymentProjectSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DeploymentProjectSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DeploymentProjectSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 build_command: Optional[str] = None,
                 framework: Optional[str] = None,
                 install_command: Optional[str] = None,
                 output_directory: Optional[str] = None,
                 root_directory: Optional[str] = None):
        """
        :param str build_command: The build command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        :param str framework: The framework that is being used for this deployment. If omitted, no framework is selected.
        :param str install_command: The install command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        :param str output_directory: The output directory of the deployment. If omitted, this value will be taken from the project or automatically detected.
        :param str root_directory: The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
        """
        DeploymentProjectSettings._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            build_command=build_command,
            framework=framework,
            install_command=install_command,
            output_directory=output_directory,
            root_directory=root_directory,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             build_command: Optional[str] = None,
             framework: Optional[str] = None,
             install_command: Optional[str] = None,
             output_directory: Optional[str] = None,
             root_directory: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if build_command is not None:
            _setter("build_command", build_command)
        if framework is not None:
            _setter("framework", framework)
        if install_command is not None:
            _setter("install_command", install_command)
        if output_directory is not None:
            _setter("output_directory", output_directory)
        if root_directory is not None:
            _setter("root_directory", root_directory)

    @property
    @pulumi.getter(name="buildCommand")
    def build_command(self) -> Optional[str]:
        """
        The build command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        """
        return pulumi.get(self, "build_command")

    @property
    @pulumi.getter
    def framework(self) -> Optional[str]:
        """
        The framework that is being used for this deployment. If omitted, no framework is selected.
        """
        return pulumi.get(self, "framework")

    @property
    @pulumi.getter(name="installCommand")
    def install_command(self) -> Optional[str]:
        """
        The install command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        """
        return pulumi.get(self, "install_command")

    @property
    @pulumi.getter(name="outputDirectory")
    def output_directory(self) -> Optional[str]:
        """
        The output directory of the deployment. If omitted, this value will be taken from the project or automatically detected.
        """
        return pulumi.get(self, "output_directory")

    @property
    @pulumi.getter(name="rootDirectory")
    def root_directory(self) -> Optional[str]:
        """
        The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
        """
        return pulumi.get(self, "root_directory")


@pulumi.output_type
class DnsRecordSrv(dict):
    def __init__(__self__, *,
                 port: int,
                 priority: int,
                 target: str,
                 weight: int):
        """
        :param int port: The TCP or UDP port on which the service is to be found.
        :param int priority: The priority of the target host, lower value means more preferred.
        :param str target: The canonical hostname of the machine providing the service, ending in a dot.
        :param int weight: A relative weight for records with the same priority, higher value means higher chance of getting picked.
        """
        DnsRecordSrv._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            port=port,
            priority=priority,
            target=target,
            weight=weight,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             port: int,
             priority: int,
             target: str,
             weight: int,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("port", port)
        _setter("priority", priority)
        _setter("target", target)
        _setter("weight", weight)

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The TCP or UDP port on which the service is to be found.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of the target host, lower value means more preferred.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The canonical hostname of the machine providing the service, ending in a dot.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def weight(self) -> int:
        """
        A relative weight for records with the same priority, higher value means higher chance of getting picked.
        """
        return pulumi.get(self, "weight")


@pulumi.output_type
class ProjectEnvironment(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "gitBranch":
            suggest = "git_branch"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectEnvironment. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectEnvironment.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectEnvironment.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 key: str,
                 targets: Sequence[str],
                 value: str,
                 git_branch: Optional[str] = None,
                 id: Optional[str] = None):
        """
        :param str key: The name of the Environment Variable.
        :param Sequence[str] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param str value: The value of the Environment Variable.
        :param str git_branch: The git branch of the Environment Variable.
        :param str id: The ID of the Environment Variable.
        """
        ProjectEnvironment._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key=key,
            targets=targets,
            value=value,
            git_branch=git_branch,
            id=id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key: str,
             targets: Sequence[str],
             value: str,
             git_branch: Optional[str] = None,
             id: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("key", key)
        _setter("targets", targets)
        _setter("value", value)
        if git_branch is not None:
            _setter("git_branch", git_branch)
        if id is not None:
            _setter("id", id)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The name of the Environment Variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def targets(self) -> Sequence[str]:
        """
        The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the Environment Variable.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="gitBranch")
    def git_branch(self) -> Optional[str]:
        """
        The git branch of the Environment Variable.
        """
        return pulumi.get(self, "git_branch")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the Environment Variable.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class ProjectGitRepository(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "productionBranch":
            suggest = "production_branch"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectGitRepository. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectGitRepository.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectGitRepository.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 repo: str,
                 type: str,
                 production_branch: Optional[str] = None):
        """
        :param str repo: The name of the git repository. For example: `vercel/next.js`.
        :param str type: The git provider of the repository. Must be either `github`, `gitlab`, or `bitbucket`.
        :param str production_branch: By default, every commit pushed to the main branch will trigger a Production Deployment instead of the usual Preview Deployment. You can switch to a different branch here.
        """
        ProjectGitRepository._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            repo=repo,
            type=type,
            production_branch=production_branch,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             repo: str,
             type: str,
             production_branch: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("repo", repo)
        _setter("type", type)
        if production_branch is not None:
            _setter("production_branch", production_branch)

    @property
    @pulumi.getter
    def repo(self) -> str:
        """
        The name of the git repository. For example: `vercel/next.js`.
        """
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The git provider of the repository. Must be either `github`, `gitlab`, or `bitbucket`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="productionBranch")
    def production_branch(self) -> Optional[str]:
        """
        By default, every commit pushed to the main branch will trigger a Production Deployment instead of the usual Preview Deployment. You can switch to a different branch here.
        """
        return pulumi.get(self, "production_branch")


@pulumi.output_type
class ProjectPasswordProtection(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "protectProduction":
            suggest = "protect_production"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectPasswordProtection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectPasswordProtection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectPasswordProtection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 password: str,
                 protect_production: Optional[bool] = None):
        """
        :param str password: The password that visitors must enter to gain access to your Preview Deployments. Drift detection is not possible for this field.
        :param bool protect_production: If true, production deployments will also be protected
        """
        ProjectPasswordProtection._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            password=password,
            protect_production=protect_production,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             password: str,
             protect_production: Optional[bool] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("password", password)
        if protect_production is not None:
            _setter("protect_production", protect_production)

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password that visitors must enter to gain access to your Preview Deployments. Drift detection is not possible for this field.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="protectProduction")
    def protect_production(self) -> Optional[bool]:
        """
        If true, production deployments will also be protected
        """
        return pulumi.get(self, "protect_production")


@pulumi.output_type
class ProjectVercelAuthentication(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "protectProduction":
            suggest = "protect_production"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectVercelAuthentication. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectVercelAuthentication.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectVercelAuthentication.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 protect_production: Optional[bool] = None):
        """
        :param bool protect_production: If true, production deployments will also be protected
        """
        ProjectVercelAuthentication._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            protect_production=protect_production,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             protect_production: Optional[bool] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if protect_production is not None:
            _setter("protect_production", protect_production)

    @property
    @pulumi.getter(name="protectProduction")
    def protect_production(self) -> Optional[bool]:
        """
        If true, production deployments will also be protected
        """
        return pulumi.get(self, "protect_production")


@pulumi.output_type
class GetProjectEnvironmentResult(dict):
    def __init__(__self__, *,
                 git_branch: str,
                 id: str,
                 key: str,
                 targets: Sequence[str],
                 value: str):
        """
        :param str git_branch: The git branch of the environment variable.
        :param str id: The ID of the environment variable
        :param str key: The name of the environment variable.
        :param Sequence[str] targets: The environments that the environment variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param str value: The value of the environment variable.
        """
        GetProjectEnvironmentResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            git_branch=git_branch,
            id=id,
            key=key,
            targets=targets,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             git_branch: str,
             id: str,
             key: str,
             targets: Sequence[str],
             value: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("git_branch", git_branch)
        _setter("id", id)
        _setter("key", key)
        _setter("targets", targets)
        _setter("value", value)

    @property
    @pulumi.getter(name="gitBranch")
    def git_branch(self) -> str:
        """
        The git branch of the environment variable.
        """
        return pulumi.get(self, "git_branch")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the environment variable
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The name of the environment variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def targets(self) -> Sequence[str]:
        """
        The environments that the environment variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the environment variable.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetProjectGitRepositoryResult(dict):
    def __init__(__self__, *,
                 production_branch: str,
                 repo: str,
                 type: str):
        """
        :param str production_branch: By default, every commit pushed to the main branch will trigger a Production Deployment instead of the usual Preview Deployment. You can switch to a different branch here.
        :param str repo: The name of the git repository. For example: `vercel/next.js`.
        :param str type: The git provider of the repository. Must be either `github`, `gitlab`, or `bitbucket`.
        """
        GetProjectGitRepositoryResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            production_branch=production_branch,
            repo=repo,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             production_branch: str,
             repo: str,
             type: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("production_branch", production_branch)
        _setter("repo", repo)
        _setter("type", type)

    @property
    @pulumi.getter(name="productionBranch")
    def production_branch(self) -> str:
        """
        By default, every commit pushed to the main branch will trigger a Production Deployment instead of the usual Preview Deployment. You can switch to a different branch here.
        """
        return pulumi.get(self, "production_branch")

    @property
    @pulumi.getter
    def repo(self) -> str:
        """
        The name of the git repository. For example: `vercel/next.js`.
        """
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The git provider of the repository. Must be either `github`, `gitlab`, or `bitbucket`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetProjectPasswordProtectionResult(dict):
    def __init__(__self__, *,
                 protect_production: bool):
        """
        :param bool protect_production: If true, production deployments will also be protected
        """
        GetProjectPasswordProtectionResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            protect_production=protect_production,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             protect_production: bool,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("protect_production", protect_production)

    @property
    @pulumi.getter(name="protectProduction")
    def protect_production(self) -> bool:
        """
        If true, production deployments will also be protected
        """
        return pulumi.get(self, "protect_production")


@pulumi.output_type
class GetProjectVercelAuthenticationResult(dict):
    def __init__(__self__, *,
                 protect_production: bool):
        """
        :param bool protect_production: If true, production deployments will also be protected
        """
        GetProjectVercelAuthenticationResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            protect_production=protect_production,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             protect_production: bool,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("protect_production", protect_production)

    @property
    @pulumi.getter(name="protectProduction")
    def protect_production(self) -> bool:
        """
        If true, production deployments will also be protected
        """
        return pulumi.get(self, "protect_production")


