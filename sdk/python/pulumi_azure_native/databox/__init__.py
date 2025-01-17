# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .get_job import *
from .job import *
from .list_job_credentials import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.databox.v20180101 as __v20180101
    v20180101 = __v20180101
    import pulumi_azure_native.databox.v20190901 as __v20190901
    v20190901 = __v20190901
    import pulumi_azure_native.databox.v20200401 as __v20200401
    v20200401 = __v20200401
    import pulumi_azure_native.databox.v20201101 as __v20201101
    v20201101 = __v20201101
    import pulumi_azure_native.databox.v20210301 as __v20210301
    v20210301 = __v20210301
    import pulumi_azure_native.databox.v20210501 as __v20210501
    v20210501 = __v20210501
    import pulumi_azure_native.databox.v20210801preview as __v20210801preview
    v20210801preview = __v20210801preview
    import pulumi_azure_native.databox.v20211201 as __v20211201
    v20211201 = __v20211201
    import pulumi_azure_native.databox.v20220201 as __v20220201
    v20220201 = __v20220201
else:
    v20180101 = _utilities.lazy_import('pulumi_azure_native.databox.v20180101')
    v20190901 = _utilities.lazy_import('pulumi_azure_native.databox.v20190901')
    v20200401 = _utilities.lazy_import('pulumi_azure_native.databox.v20200401')
    v20201101 = _utilities.lazy_import('pulumi_azure_native.databox.v20201101')
    v20210301 = _utilities.lazy_import('pulumi_azure_native.databox.v20210301')
    v20210501 = _utilities.lazy_import('pulumi_azure_native.databox.v20210501')
    v20210801preview = _utilities.lazy_import('pulumi_azure_native.databox.v20210801preview')
    v20211201 = _utilities.lazy_import('pulumi_azure_native.databox.v20211201')
    v20220201 = _utilities.lazy_import('pulumi_azure_native.databox.v20220201')

