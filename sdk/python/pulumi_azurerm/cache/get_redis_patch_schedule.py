# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetRedisPatchScheduleResult:
    """
    Response to put/get patch schedules for Redis cache.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        List of patch schedules for a Redis cache.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableGetRedisPatchScheduleResult(GetRedisPatchScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRedisPatchScheduleResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_redis_patch_schedule(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: Default string modeled as parameter for auto generation to work correctly.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:cache:getRedisPatchSchedule', __args__, opts=opts).value

    return AwaitableGetRedisPatchScheduleResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))