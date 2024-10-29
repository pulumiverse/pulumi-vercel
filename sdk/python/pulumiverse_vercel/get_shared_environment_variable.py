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

__all__ = [
    'GetSharedEnvironmentVariableResult',
    'AwaitableGetSharedEnvironmentVariableResult',
    'get_shared_environment_variable',
    'get_shared_environment_variable_output',
]

@pulumi.output_type
class GetSharedEnvironmentVariableResult:
    """
    A collection of values returned by getSharedEnvironmentVariable.
    """
    def __init__(__self__, comment=None, id=None, key=None, project_ids=None, sensitive=None, targets=None, team_id=None, value=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)
        if sensitive and not isinstance(sensitive, bool):
            raise TypeError("Expected argument 'sensitive' to be a bool")
        pulumi.set(__self__, "sensitive", sensitive)
        if targets and not isinstance(targets, list):
            raise TypeError("Expected argument 'targets' to be a list")
        pulumi.set(__self__, "targets", targets)
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        pulumi.set(__self__, "team_id", team_id)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        A comment explaining what the environment variable is for.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Environment Variable.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The name of the Environment Variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Sequence[str]:
        """
        The ID of the Vercel project.
        """
        return pulumi.get(self, "project_ids")

    @property
    @pulumi.getter
    def sensitive(self) -> bool:
        """
        Whether the Environment Variable is sensitive or not.
        """
        return pulumi.get(self, "sensitive")

    @property
    @pulumi.getter
    def targets(self) -> Sequence[str]:
        """
        The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        """
        The ID of the Vercel team. Shared environment variables require a team.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the Environment Variable.
        """
        return pulumi.get(self, "value")


class AwaitableGetSharedEnvironmentVariableResult(GetSharedEnvironmentVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSharedEnvironmentVariableResult(
            comment=self.comment,
            id=self.id,
            key=self.key,
            project_ids=self.project_ids,
            sensitive=self.sensitive,
            targets=self.targets,
            team_id=self.team_id,
            value=self.value)


def get_shared_environment_variable(id: Optional[str] = None,
                                    key: Optional[str] = None,
                                    targets: Optional[Sequence[str]] = None,
                                    team_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSharedEnvironmentVariableResult:
    """
    Provides information about an existing Shared Environment Variable within Vercel.

    A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects.

    For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vercel as vercel

    example = vercel.get_shared_environment_variable(id="xxxxxxxxxxxxxxx")
    example_by_key_and_target = vercel.get_shared_environment_variable(key="MY_ENV_VAR",
        targets=[
            "production",
            "preview",
        ])
    ```


    :param str id: The ID of the Environment Variable.
    :param str key: The name of the Environment Variable.
    :param Sequence[str] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
    :param str team_id: The ID of the Vercel team. Shared environment variables require a team.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['key'] = key
    __args__['targets'] = targets
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vercel:index/getSharedEnvironmentVariable:getSharedEnvironmentVariable', __args__, opts=opts, typ=GetSharedEnvironmentVariableResult).value

    return AwaitableGetSharedEnvironmentVariableResult(
        comment=pulumi.get(__ret__, 'comment'),
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        project_ids=pulumi.get(__ret__, 'project_ids'),
        sensitive=pulumi.get(__ret__, 'sensitive'),
        targets=pulumi.get(__ret__, 'targets'),
        team_id=pulumi.get(__ret__, 'team_id'),
        value=pulumi.get(__ret__, 'value'))
def get_shared_environment_variable_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                           key: Optional[pulumi.Input[Optional[str]]] = None,
                                           targets: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                           team_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSharedEnvironmentVariableResult]:
    """
    Provides information about an existing Shared Environment Variable within Vercel.

    A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects.

    For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vercel as vercel

    example = vercel.get_shared_environment_variable(id="xxxxxxxxxxxxxxx")
    example_by_key_and_target = vercel.get_shared_environment_variable(key="MY_ENV_VAR",
        targets=[
            "production",
            "preview",
        ])
    ```


    :param str id: The ID of the Environment Variable.
    :param str key: The name of the Environment Variable.
    :param Sequence[str] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
    :param str team_id: The ID of the Vercel team. Shared environment variables require a team.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['key'] = key
    __args__['targets'] = targets
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vercel:index/getSharedEnvironmentVariable:getSharedEnvironmentVariable', __args__, opts=opts, typ=GetSharedEnvironmentVariableResult)
    return __ret__.apply(lambda __response__: GetSharedEnvironmentVariableResult(
        comment=pulumi.get(__response__, 'comment'),
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        project_ids=pulumi.get(__response__, 'project_ids'),
        sensitive=pulumi.get(__response__, 'sensitive'),
        targets=pulumi.get(__response__, 'targets'),
        team_id=pulumi.get(__response__, 'team_id'),
        value=pulumi.get(__response__, 'value')))
