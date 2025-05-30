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
    'GetEdgeConfigItemResult',
    'AwaitableGetEdgeConfigItemResult',
    'get_edge_config_item',
    'get_edge_config_item_output',
]

@pulumi.output_type
class GetEdgeConfigItemResult:
    """
    A collection of values returned by getEdgeConfigItem.
    """
    def __init__(__self__, id=None, key=None, team_id=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        pulumi.set(__self__, "team_id", team_id)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Edge Config that the item should exist under.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The name of the key you want to retrieve within your Edge Config.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        """
        The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value assigned to the key.
        """
        return pulumi.get(self, "value")


class AwaitableGetEdgeConfigItemResult(GetEdgeConfigItemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEdgeConfigItemResult(
            id=self.id,
            key=self.key,
            team_id=self.team_id,
            value=self.value)


def get_edge_config_item(id: Optional[str] = None,
                         key: Optional[str] = None,
                         team_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEdgeConfigItemResult:
    """
    Provides the value of an existing Edge Config Item.

    An Edge Config is a global data store that enables experimentation with feature flags, A/B testing, critical redirects, and more.

    An Edge Config Item is a value within an Edge Config.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vercel as vercel

    example = vercel.get_edge_config(id="ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
    test = vercel.get_edge_config_item(id=example.id,
        key="foobar")
    ```


    :param str id: The ID of the Edge Config that the item should exist under.
    :param str key: The name of the key you want to retrieve within your Edge Config.
    :param str team_id: The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['key'] = key
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vercel:index/getEdgeConfigItem:getEdgeConfigItem', __args__, opts=opts, typ=GetEdgeConfigItemResult).value

    return AwaitableGetEdgeConfigItemResult(
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        team_id=pulumi.get(__ret__, 'team_id'),
        value=pulumi.get(__ret__, 'value'))
def get_edge_config_item_output(id: Optional[pulumi.Input[str]] = None,
                                key: Optional[pulumi.Input[str]] = None,
                                team_id: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEdgeConfigItemResult]:
    """
    Provides the value of an existing Edge Config Item.

    An Edge Config is a global data store that enables experimentation with feature flags, A/B testing, critical redirects, and more.

    An Edge Config Item is a value within an Edge Config.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vercel as vercel

    example = vercel.get_edge_config(id="ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
    test = vercel.get_edge_config_item(id=example.id,
        key="foobar")
    ```


    :param str id: The ID of the Edge Config that the item should exist under.
    :param str key: The name of the key you want to retrieve within your Edge Config.
    :param str team_id: The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['key'] = key
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vercel:index/getEdgeConfigItem:getEdgeConfigItem', __args__, opts=opts, typ=GetEdgeConfigItemResult)
    return __ret__.apply(lambda __response__: GetEdgeConfigItemResult(
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        team_id=pulumi.get(__response__, 'team_id'),
        value=pulumi.get(__response__, 'value')))
