# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetGuestConfigurationHCRPAssignmentResult:
    """
    Guest configuration assignment is an association between a machine and guest configuration.
    """
    def __init__(__self__, location=None, name=None, properties=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Region where the VM is located.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Name of the guest configuration assignment.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the Guest configuration assignment.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource.
        """


class AwaitableGetGuestConfigurationHCRPAssignmentResult(GetGuestConfigurationHCRPAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGuestConfigurationHCRPAssignmentResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_guest_configuration_hcrp_assignment(machine_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str machine_name: The name of the ARC machine.
    :param str name: The guest configuration assignment name.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['machineName'] = machine_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:hybridcompute/v20181120:getGuestConfigurationHCRPAssignment', __args__, opts=opts).value

    return AwaitableGetGuestConfigurationHCRPAssignmentResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))