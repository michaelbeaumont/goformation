package cloudformation

// AWSECSTaskDefinition_Ulimit AWS CloudFormation Resource (AWS::ECS::TaskDefinition.Ulimit)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html
type AWSECSTaskDefinition_Ulimit struct {

	// HardLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-hardlimit
	HardLimit *Value `json:"HardLimit,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-name
	Name *Value `json:"Name,omitempty"`

	// SoftLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-softlimit
	SoftLimit *Value `json:"SoftLimit,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_Ulimit) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.Ulimit"
}
