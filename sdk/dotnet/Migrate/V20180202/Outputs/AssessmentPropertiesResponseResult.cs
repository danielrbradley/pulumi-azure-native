// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20180202.Outputs
{

    [OutputType]
    public sealed class AssessmentPropertiesResponseResult
    {
        /// <summary>
        /// AHUB discount on windows virtual machines.
        /// </summary>
        public readonly string AzureHybridUseBenefit;
        /// <summary>
        /// Target Azure location for which the machines should be assessed. These enums are the same as used by Compute API.
        /// </summary>
        public readonly string AzureLocation;
        /// <summary>
        /// Offer code according to which cost estimation is done.
        /// </summary>
        public readonly string AzureOfferCode;
        /// <summary>
        /// Pricing tier for Size evaluation.
        /// </summary>
        public readonly string AzurePricingTier;
        /// <summary>
        /// Storage Redundancy type offered by Azure.
        /// </summary>
        public readonly string AzureStorageRedundancy;
        /// <summary>
        /// Confidence rating percentage for assessment. Can be in the range [0, 100].
        /// </summary>
        public readonly double ConfidenceRatingInPercentage;
        /// <summary>
        /// Time when this project was created. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string CreatedTimestamp;
        /// <summary>
        /// Currency to report prices in.
        /// </summary>
        public readonly string Currency;
        /// <summary>
        /// Custom discount percentage to be applied on final costs. Can be in the range [0, 100].
        /// </summary>
        public readonly double DiscountPercentage;
        /// <summary>
        /// Monthly network cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
        /// </summary>
        public readonly double MonthlyBandwidthCost;
        /// <summary>
        /// Monthly compute cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
        /// </summary>
        public readonly double MonthlyComputeCost;
        /// <summary>
        /// Monthly storage cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
        /// </summary>
        public readonly double MonthlyStorageCost;
        /// <summary>
        /// Number of assessed machines part of this assessment.
        /// </summary>
        public readonly int NumberOfMachines;
        /// <summary>
        /// Percentile of performance data used to recommend Azure size.
        /// </summary>
        public readonly string Percentile;
        /// <summary>
        /// Time when the Azure Prices were queried. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string PricesTimestamp;
        /// <summary>
        /// Scaling factor used over utilization data to add a performance buffer for new machines to be created in Azure. Min Value = 1.0, Max value = 1.9, Default = 1.3.
        /// </summary>
        public readonly double ScalingFactor;
        /// <summary>
        /// Assessment sizing criterion.
        /// </summary>
        public readonly string SizingCriterion;
        /// <summary>
        /// User configurable setting that describes the status of the assessment.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// Whether the assessment has been created and is valid.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Time range of performance data used to recommend a size.
        /// </summary>
        public readonly string TimeRange;
        /// <summary>
        /// Time when this project was last updated. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string UpdatedTimestamp;

        [OutputConstructor]
        private AssessmentPropertiesResponseResult(
            string azureHybridUseBenefit,

            string azureLocation,

            string azureOfferCode,

            string azurePricingTier,

            string azureStorageRedundancy,

            double confidenceRatingInPercentage,

            string createdTimestamp,

            string currency,

            double discountPercentage,

            double monthlyBandwidthCost,

            double monthlyComputeCost,

            double monthlyStorageCost,

            int numberOfMachines,

            string percentile,

            string pricesTimestamp,

            double scalingFactor,

            string sizingCriterion,

            string stage,

            string status,

            string timeRange,

            string updatedTimestamp)
        {
            AzureHybridUseBenefit = azureHybridUseBenefit;
            AzureLocation = azureLocation;
            AzureOfferCode = azureOfferCode;
            AzurePricingTier = azurePricingTier;
            AzureStorageRedundancy = azureStorageRedundancy;
            ConfidenceRatingInPercentage = confidenceRatingInPercentage;
            CreatedTimestamp = createdTimestamp;
            Currency = currency;
            DiscountPercentage = discountPercentage;
            MonthlyBandwidthCost = monthlyBandwidthCost;
            MonthlyComputeCost = monthlyComputeCost;
            MonthlyStorageCost = monthlyStorageCost;
            NumberOfMachines = numberOfMachines;
            Percentile = percentile;
            PricesTimestamp = pricesTimestamp;
            ScalingFactor = scalingFactor;
            SizingCriterion = sizingCriterion;
            Stage = stage;
            Status = status;
            TimeRange = timeRange;
            UpdatedTimestamp = updatedTimestamp;
        }
    }
}