# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ReplicationNetworkMapping(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource Location
    """
    name: pulumi.Output[str]
    """
    Resource Name
    """
    properties: pulumi.Output[dict]
    """
    The Network Mapping Properties.
      * `fabric_specific_settings` (`dict`) - The fabric specific settings.
        * `instance_type` (`str`) - Gets the Instance type.

      * `primary_fabric_friendly_name` (`str`) - The primary fabric friendly name.
      * `primary_network_friendly_name` (`str`) - The primary network friendly name.
      * `primary_network_id` (`str`) - The primary network id for network mapping.
      * `recovery_fabric_arm_id` (`str`) - The recovery fabric ARM id.
      * `recovery_fabric_friendly_name` (`str`) - The recovery fabric friendly name.
      * `recovery_network_friendly_name` (`str`) - The recovery network friendly name.
      * `recovery_network_id` (`str`) - The recovery network id for network mapping.
      * `state` (`str`) - The pairing state for network mapping.
    """
    type: pulumi.Output[str]
    """
    Resource Type
    """
    def __init__(__self__, resource_name, opts=None, fabric_name=None, name=None, network_name=None, properties=None, resource_group_name=None, resource_name_=None, __props__=None, __name__=None, __opts__=None):
        """
        Network Mapping model. Ideally it should have been possible to inherit this class from prev version in InheritedModels as long as there is no difference in structure or method signature. Since there were no base Models for certain fields and methods viz NetworkMappingProperties and Load with required return type, the class has been introduced in its entirety with references to base models to facilitate extensions in subsequent versions.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fabric_name: Primary fabric name.
        :param pulumi.Input[str] name: Network mapping name.
        :param pulumi.Input[str] network_name: Primary network name.
        :param pulumi.Input[dict] properties: Input properties for creating network mapping.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[str] resource_name_: The name of the recovery services vault.

        The **properties** object supports the following:

          * `fabric_specific_details` (`pulumi.Input[dict]`) - Fabric specific input properties.
            * `instance_type` (`pulumi.Input[str]`) - The instance type.

          * `recovery_fabric_name` (`pulumi.Input[str]`) - Recovery fabric Name.
          * `recovery_network_id` (`pulumi.Input[str]`) - Recovery network Id.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if fabric_name is None:
                raise TypeError("Missing required property 'fabric_name'")
            __props__['fabric_name'] = fabric_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network_name is None:
                raise TypeError("Missing required property 'network_name'")
            __props__['network_name'] = network_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['location'] = None
            __props__['type'] = None
        super(ReplicationNetworkMapping, __self__).__init__(
            'azurerm:recoveryservices/v20180110:ReplicationNetworkMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ReplicationNetworkMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ReplicationNetworkMapping(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop