# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AliasArgs', 'Alias']

@pulumi.input_type
class AliasArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 deployment_id: pulumi.Input[str],
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Alias resource.
        :param pulumi.Input[str] alias: The Alias we want to assign to the deployment defined in the URL.
        :param pulumi.Input[str] deployment_id: The id of the Deployment the Alias should be associated with.
        :param pulumi.Input[str] team_id: The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        AliasArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            alias=alias,
            deployment_id=deployment_id,
            team_id=team_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             alias: Optional[pulumi.Input[str]] = None,
             deployment_id: Optional[pulumi.Input[str]] = None,
             team_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if alias is None:
            raise TypeError("Missing 'alias' argument")
        if deployment_id is None and 'deploymentId' in kwargs:
            deployment_id = kwargs['deploymentId']
        if deployment_id is None:
            raise TypeError("Missing 'deployment_id' argument")
        if team_id is None and 'teamId' in kwargs:
            team_id = kwargs['teamId']

        _setter("alias", alias)
        _setter("deployment_id", deployment_id)
        if team_id is not None:
            _setter("team_id", team_id)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        The Alias we want to assign to the deployment defined in the URL.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Input[str]:
        """
        The id of the Deployment the Alias should be associated with.
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _AliasState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Alias resources.
        :param pulumi.Input[str] alias: The Alias we want to assign to the deployment defined in the URL.
        :param pulumi.Input[str] deployment_id: The id of the Deployment the Alias should be associated with.
        :param pulumi.Input[str] team_id: The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        _AliasState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            alias=alias,
            deployment_id=deployment_id,
            team_id=team_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             alias: Optional[pulumi.Input[str]] = None,
             deployment_id: Optional[pulumi.Input[str]] = None,
             team_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if deployment_id is None and 'deploymentId' in kwargs:
            deployment_id = kwargs['deploymentId']
        if team_id is None and 'teamId' in kwargs:
            team_id = kwargs['teamId']

        if alias is not None:
            _setter("alias", alias)
        if deployment_id is not None:
            _setter("deployment_id", deployment_id)
        if team_id is not None:
            _setter("team_id", team_id)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The Alias we want to assign to the deployment defined in the URL.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the Deployment the Alias should be associated with.
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class Alias(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Alias resource.

        An Alias allows a `Deployment` to be accessed through a different URL.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The Alias we want to assign to the deployment defined in the URL.
        :param pulumi.Input[str] deployment_id: The id of the Deployment the Alias should be associated with.
        :param pulumi.Input[str] team_id: The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AliasArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Alias resource.

        An Alias allows a `Deployment` to be accessed through a different URL.

        :param str resource_name: The name of the resource.
        :param AliasArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AliasArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            AliasArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AliasArgs.__new__(AliasArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            if deployment_id is None and not opts.urn:
                raise TypeError("Missing required property 'deployment_id'")
            __props__.__dict__["deployment_id"] = deployment_id
            __props__.__dict__["team_id"] = team_id
        super(Alias, __self__).__init__(
            'vercel:index/alias:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            deployment_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'Alias':
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The Alias we want to assign to the deployment defined in the URL.
        :param pulumi.Input[str] deployment_id: The id of the Deployment the Alias should be associated with.
        :param pulumi.Input[str] team_id: The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AliasState.__new__(_AliasState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["deployment_id"] = deployment_id
        __props__.__dict__["team_id"] = team_id
        return Alias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The Alias we want to assign to the deployment defined in the URL.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[str]:
        """
        The id of the Deployment the Alias should be associated with.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The ID of the team the Alias and Deployment exist under. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

