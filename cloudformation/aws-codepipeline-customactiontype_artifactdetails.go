package cloudformation

import (
	"encoding/json"
)

// AWSCodePipelineCustomActionType_ArtifactDetails AWS CloudFormation Resource (AWS::CodePipeline::CustomActionType.ArtifactDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
type AWSCodePipelineCustomActionType_ArtifactDetails struct {

	// MaximumCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	MaximumCount *Value `json:"MaximumCount,omitempty"`

	// MinimumCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	MinimumCount *Value `json:"MinimumCount,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_ArtifactDetails) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.ArtifactDetails"
}

func (r *AWSCodePipelineCustomActionType_ArtifactDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
