# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class DatabaseAccountCassandraTable(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource group to which the resource belongs.
    """
    name: pulumi.Output[str]
    """
    The name of the database account.
    """
    properties: pulumi.Output[dict]
    """
    The properties of an Azure Cosmos DB Cassandra table
      * `default_ttl` (`float`) - Time to live of the Cosmos DB Cassandra table
      * `id` (`str`) - Name of the Cosmos DB Cassandra table
      * `schema` (`dict`) - Schema of the Cosmos DB Cassandra table
        * `cluster_keys` (`list`) - List of cluster key.
          * `name` (`str`) - Name of the Cosmos DB Cassandra table cluster key
          * `order_by` (`str`) - Order of the Cosmos DB Cassandra table cluster key, only support "Asc" and "Desc"

        * `columns` (`list`) - List of Cassandra table columns.
          * `name` (`str`) - Name of the Cosmos DB Cassandra table column
          * `type` (`str`) - Type of the Cosmos DB Cassandra table column

        * `partition_keys` (`list`) - List of partition key.
          * `name` (`str`) - Name of the Cosmos DB Cassandra table partition key
    """
    tags: pulumi.Output[dict]
    """
    Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
    """
    type: pulumi.Output[str]
    """
    The type of Azure resource.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, keyspace_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        An Azure Cosmos DB Cassandra table.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Cosmos DB database account name.
        :param pulumi.Input[str] keyspace_name: Cosmos DB keyspace name.
        :param pulumi.Input[str] name: Cosmos DB table name.
        :param pulumi.Input[dict] properties: Properties to create and update Azure Cosmos DB Cassandra table.
        :param pulumi.Input[str] resource_group_name: Name of an Azure resource group.

        The **properties** object supports the following:

          * `options` (`pulumi.Input[dict]`) - A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
          * `resource` (`pulumi.Input[dict]`) - The standard JSON format of a Cassandra table
            * `default_ttl` (`pulumi.Input[float]`) - Time to live of the Cosmos DB Cassandra table
            * `id` (`pulumi.Input[str]`) - Name of the Cosmos DB Cassandra table
            * `schema` (`pulumi.Input[dict]`) - Schema of the Cosmos DB Cassandra table
              * `cluster_keys` (`pulumi.Input[list]`) - List of cluster key.
                * `name` (`pulumi.Input[str]`) - Name of the Cosmos DB Cassandra table cluster key
                * `order_by` (`pulumi.Input[str]`) - Order of the Cosmos DB Cassandra table cluster key, only support "Asc" and "Desc"

              * `columns` (`pulumi.Input[list]`) - List of Cassandra table columns.
                * `name` (`pulumi.Input[str]`) - Name of the Cosmos DB Cassandra table column
                * `type` (`pulumi.Input[str]`) - Type of the Cosmos DB Cassandra table column

              * `partition_keys` (`pulumi.Input[list]`) - List of partition key.
                * `name` (`pulumi.Input[str]`) - Name of the Cosmos DB Cassandra table partition key
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if keyspace_name is None:
                raise TypeError("Missing required property 'keyspace_name'")
            __props__['keyspace_name'] = keyspace_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['location'] = None
            __props__['tags'] = None
            __props__['type'] = None
        super(DatabaseAccountCassandraTable, __self__).__init__(
            'azurerm:documentdb/v20150401:DatabaseAccountCassandraTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DatabaseAccountCassandraTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DatabaseAccountCassandraTable(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop