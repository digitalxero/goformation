package resources

// AWS::EMR::Cluster.PlacementType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html
type AWSEMRClusterPlacementType struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html#aws-properties-emr-cluster-jobflowinstancesconfig-placementtype
	AvailabilityZone string `json:"AvailabilityZone"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterPlacementType) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.PlacementType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterPlacementType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}