/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package awsdoc

var generatedParamsDoc = map[string]map[string]string{
	"attach.alarm":               {},
	"attach.classicloadbalancer": {},
	"attach.containertask":       {},
	"attach.elasticip": {
		"allow-reassociation": "For a VPC in an EC2-Classic account, specify true to allow an Elastic IP address that is already associated with an instance or network interface to be reassociated with the specified instance or network interface. Otherwise, the operation fails. In a VPC in an EC2-VPC-only account, reassociation is automatic, therefore you can specify false to ensure the operation fails if the Elastic IP address is already associated with another resource.",
		"id":                  "The allocation ID. This is required for EC2-VPC.",
		"instance":            "The ID of the instance. This is required for EC2-Classic. For EC2-VPC, you can specify either the instance ID or the network interface ID, but not both. The operation fails if you specify an instance ID unless exactly one network interface is attached.",
		"networkinterface":    "The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID. For EC2-VPC, you can specify either the instance ID or the network interface ID, but not both.",
		"privateip":           "The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.",
	},
	"attach.instance": {
		"targetgroup": "The Amazon Resource Name (ARN) of the target group.",
	},
	"attach.instanceprofile": {},
	"attach.internetgateway": {
		"id":  "The ID of the internet gateway.",
		"vpc": "The ID of the VPC.",
	},
	"attach.listener": {
		"id": "The Amazon Resource Name (ARN) of the listener.",
	},
	"attach.mfadevice": {
		"id":         "The serial number that uniquely identifies the MFA device. For virtual MFA devices, the serial number is the device ARN. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: =,.@:/-",
		"mfa-code-1": "An authentication code emitted by the device.  The format for this parameter is a string of six digits.  Submit your request immediately after generating the authentication codes. If you generate the codes and then wait too long to submit the request, the MFA device successfully associates with the user but the MFA device becomes out of sync. This happens because time-based one-time passwords (TOTP) expire after a short period of time. If this happens, you can <a href=\"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html\">resync the device.",
		"mfa-code-2": "A subsequent authentication code emitted by the device. The format for this parameter is a string of six digits.  Submit your request immediately after generating the authentication codes. If you generate the codes and then wait too long to submit the request, the MFA device successfully associates with the user but the MFA device becomes out of sync. This happens because time-based one-time passwords (TOTP) expire after a short period of time. If this happens, you can <a href=\"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html\">resync the device.",
		"user":       "The name of the IAM user for whom you want to enable the MFA device. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"attach.networkinterface": {
		"device-index": "The index of the device for the network interface attachment.",
		"id":           "The ID of the network interface.",
		"instance":     "The ID of the instance.",
	},
	"attach.policy": {},
	"attach.role": {
		"instanceprofile": "The name of the instance profile to update. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		"name":            "The name of the role to add. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"attach.routetable": {
		"id":     "The ID of the route table.",
		"subnet": "The ID of the subnet.",
	},
	"attach.securitygroup": {},
	"attach.user": {
		"group": "The name of the group to update. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		"name":  "The name of the user to add. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"attach.volume": {
		"device":   "The device name (for example, /dev/sdh or xvdh).",
		"id":       "The ID of the EBS volume. The volume and instance must be within the same Availability Zone.",
		"instance": "The ID of the instance.",
	},
	"authenticate.registry":  {},
	"check.certificate":      {},
	"check.database":         {},
	"check.distribution":     {},
	"check.instance":         {},
	"check.loadbalancer":     {},
	"check.natgateway":       {},
	"check.networkinterface": {},
	"check.scalinggroup":     {},
	"check.securitygroup":    {},
	"check.volume":           {},
	"copy.image": {
		"description":   "A description for the new AMI in the destination Region.",
		"encrypted":     "Specifies whether the destination snapshots of the copied image should be encrypted. You can encrypt a copy of an unencrypted snapshot, but you cannot create an unencrypted copy of an encrypted snapshot. The default CMK for EBS is used unless you specify a non-default AWS Key Management Service (AWS KMS) CMK using KmsKeyId. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html\">Amazon EBS Encryption in the Amazon Elastic Compute Cloud User Guide.",
		"name":          "The name of the new AMI in the destination Region.",
		"source-id":     "The ID of the AMI to copy.",
		"source-region": "The name of the Region that contains the AMI to copy.",
	},
	"copy.snapshot": {
		"description":   "A description for the EBS snapshot.",
		"encrypted":     "To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter. Otherwise, omit this parameter. Encrypted snapshots are encrypted, even if you omit this parameter and encryption by default is not enabled. You cannot set this parameter to false. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html\">Amazon EBS Encryption in the Amazon Elastic Compute Cloud User Guide.",
		"source-id":     "The ID of the EBS snapshot to copy.",
		"source-region": "The ID of the Region that contains the snapshot to be copied.",
	},
	"create.accesskey": {
		"user": "The name of the IAM user that the new key will belong to. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"create.alarm": {
		"alarm-actions":            "The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot | arn:aws:sns:region:account-id:sns-topic-name  | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name   Valid Values (for use with IAM roles): arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0",
		"description":              "The description for the alarm.",
		"dimensions":               "The dimensions for the metric specified in MetricName.",
		"enabled":                  "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
		"evaluation-periods":       "The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an \"M out of N\" alarm, this value is the N. An alarm's total current evaluation period can be no longer than one day, so this number multiplied by Period cannot be more than 86,400 seconds.",
		"insufficientdata-actions": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot | arn:aws:sns:region:account-id:sns-topic-name  | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name   Valid Values (for use with IAM roles): &gt;arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0",
		"metric":                   "The name for the metric associated with the alarm. For each PutMetricAlarm operation, you must specify either MetricName or a Metrics array. If you are creating an alarm based on a math expression, you cannot specify this parameter, or any of the Dimensions, Period, Namespace, Statistic, or ExtendedStatistic parameters. Instead, you specify all this information in the Metrics array.",
		"name":                     "The name for the alarm. This name must be unique within your AWS account.",
		"namespace":                "The namespace for the metric associated specified in MetricName.",
		"ok-actions":               "The actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid Values: arn:aws:automate:region:ec2:stop | arn:aws:automate:region:ec2:terminate | arn:aws:automate:region:ec2:recover | arn:aws:automate:region:ec2:reboot | arn:aws:sns:region:account-id:sns-topic-name  | arn:aws:autoscaling:region:account-id:scalingPolicy:policy-idautoScalingGroupName/group-friendly-name:policyName/policy-friendly-name   Valid Values (for use with IAM roles): arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 | arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0",
		"operator":                 "The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand. The values LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.",
		"period":                   "The length, in seconds, used each time the metric specified in MetricName is evaluated. Valid values are 10, 30, and any multiple of 60.  Period is required for alarms based on static thresholds. If you are creating an alarm based on a metric math expression, you specify the period for each metric within the objects in the Metrics array. Be sure to specify 10 or 30 only for metrics that are stored by a PutMetricData call with a StorageResolution of 1. If you specify a period of 10 or 30 for a metric that does not have sub-minute resolution, the alarm still attempts to gather data at the period rate that you specify. In this case, it does not receive data for the attempts that do not correspond to a one-minute data resolution, and the alarm may often lapse into INSUFFICENT_DATA status. Specifying 10 or 30 also sets this alarm as a high-resolution alarm, which has a higher charge than other alarms. For more information about pricing, see <a href=\"https://aws.amazon.com/cloudwatch/pricing/\">Amazon CloudWatch Pricing. An alarm's total current evaluation period can be no longer than one day, so Period multiplied by EvaluationPeriods cannot be more than 86,400 seconds.",
		"statistic-function":       "The statistic for the metric specified in MetricName, other than percentile. For percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm and specify a MetricName, you must specify either Statistic or ExtendedStatistic, but not both.",
		"threshold":                "The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.",
		"unit":                     "The unit of measure for the statistic. For example, the units for the Amazon EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes that an instance receives on all network interfaces. You can also specify a unit when you create a custom metric. Units help provide conceptual meaning to your data. Metric data points that specify a unit of measure, such as Percent, are aggregated separately. If you don't specify Unit, CloudWatch retrieves all unit types that have been published for the metric and attempts to evaluate the alarm. Usually metrics are published with only one unit, so the alarm will work as intended. However, if the metric is published with multiple types of units and you don't specify a unit, the alarm's behavior is not defined and will behave un-predictably. We recommend omitting Unit so that you don't inadvertently specify an incorrect unit that is not published for this metric. Doing so causes the alarm to be stuck in the INSUFFICIENT DATA state.",
	},
	"create.appscalingpolicy": {
		"dimension":         "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.    ecs:service:DesiredCount - The desired task count of an ECS service.    ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet request.    elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.    appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.    dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.    dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.    dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.    dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.    rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.    sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an Amazon SageMaker model endpoint variant.    custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.",
		"name":              "The name of the scaling policy.",
		"resource":          "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.   ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.   Spot Fleet request - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.   EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.   AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.   DynamoDB table - The resource type is table and the unique identifier is the resource ID. Example: table/my-table.   DynamoDB global secondary index - The resource type is index and the unique identifier is the resource ID. Example: table/my-table/index/my-table-index.   Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.   Amazon SageMaker endpoint variants - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.   Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our <a href=\"https://github.com/aws/aws-auto-scaling-custom-resource\">GitHub repository.",
		"service-namespace": "The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service. For more information, see <a href=\"http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces\">AWS Service Namespaces in the Amazon Web Services General Reference.",
		"type":              "The policy type. This parameter is required if you are creating a scaling policy. The following policy types are supported:   TargetTrackingScaling—Not supported for Amazon EMR or AppStream  StepScaling—Not supported for Amazon DynamoDB For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html\">Step Scaling Policies for Application Auto Scaling and <a href=\"https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html\">Target Tracking Scaling Policies for Application Auto Scaling in the Application Auto Scaling User Guide.",
	},
	"create.appscalingtarget": {
		"dimension":         "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.    ecs:service:DesiredCount - The desired task count of an ECS service.    ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet request.    elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.    appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.    dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.    dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.    dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.    dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.    rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.    sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an Amazon SageMaker model endpoint variant.    custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.",
		"max-capacity":      "The maximum value to scale to in response to a scale-out event. MaxCapacity is required to register a scalable target.",
		"min-capacity":      "The minimum value to scale to in response to a scale-in event. MinCapacity is required to register a scalable target.",
		"resource":          "The identifier of the resource that is associated with the scalable target. This string consists of the resource type and unique identifier.   ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.   Spot Fleet request - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.   EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.   AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.   DynamoDB table - The resource type is table and the unique identifier is the resource ID. Example: table/my-table.   DynamoDB global secondary index - The resource type is index and the unique identifier is the resource ID. Example: table/my-table/index/my-table-index.   Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.   Amazon SageMaker endpoint variants - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.   Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our <a href=\"https://github.com/aws/aws-auto-scaling-custom-resource\">GitHub repository.",
		"role":              "Application Auto Scaling creates a service-linked role that grants it permissions to modify the scalable target on your behalf. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html\">Service-Linked Roles for Application Auto Scaling. For resources that are not supported using a service-linked role, this parameter is required, and it must specify the ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.",
		"service-namespace": "The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service. For more information, see <a href=\"http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces\">AWS Service Namespaces in the Amazon Web Services General Reference.",
	},
	"create.bucket": {
		"acl":  "The canned ACL to apply to the bucket.",
		"name": "<p/>",
	},
	"create.certificate": {},
	"create.classicloadbalancer": {
		"scheme":         "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet. The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can only route requests from clients with access to the VPC for the load balancer. The default is an Internet-facing load balancer.",
		"securitygroups": "[Application Load Balancers] The IDs of the security groups for the load balancer.",
		"subnets":        "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings. [Application Load Balancers] You must specify subnets from at least two Availability Zones. [Network Load Balancers] You can specify subnets from one or more Availability Zones.",
		"tags":           "One or more tags to assign to the load balancer.",
	},
	"create.containercluster": {
		"name": "The name of your cluster. If you do not specify a name for your cluster, you create a cluster named default. Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.",
	},
	"create.database":      {},
	"create.dbsubnetgroup": {},
	"create.distribution":  {},
	"create.elasticip": {
		"domain": "Set to vpc to allocate the address for use with instances in a VPC. Default: The address is for use with instances in EC2-Classic.",
	},
	"create.function": {
		"description": "A description of the function.",
		"handler":     "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see <a href=\"https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html\">Programming Model.",
		"memory":      "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
		"name":        "The name of the Lambda function. <p class=\"title\"> Name formats     Function name - my-function.    Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.    Partial ARN - 123456789012:function:my-function.   The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
		"publish":     "Set to true to publish the first version of the function during creation.",
		"role":        "The Amazon Resource Name (ARN) of the function's execution role.",
		"runtime":     "The identifier of the function's <a href=\"https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html\">runtime.",
		"timeout":     "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
	},
	"create.group": {
		"name": "The name of the group to create. Do not include the path in this value. IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both \"MyResource\" and \"myresource\".",
	},
	"create.image": {
		"description": "A description for the new image.",
		"instance":    "The ID of the instance.",
		"name":        "A name for the new image. Constraints: 3-128 alphanumeric characters, parentheses (()), square brackets ([]), spaces ( ), periods (.), slashes (/), dashes (-), single quotes ('), at-signs (@), or underscores(_)",
		"reboot":      "By default, Amazon EC2 attempts to shut down and reboot the instance before creating the image. If the 'No Reboot' option is set, Amazon EC2 doesn't shut down the instance before creating the image. When this option is used, file system integrity on the created image can't be guaranteed.",
	},
	"create.instance": {
		"image":         "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",
		"ip":            "The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet. Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request. You cannot specify this option and the network interfaces option in the same request.",
		"keypair":       "The name of the key pair. You can create a key pair using <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html\">CreateKeyPair or <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html\">ImportKeyPair.  If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",
		"lock":          "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html\">ModifyInstanceAttribute. Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance by running the shutdown command from the instance. Default: false",
		"securitygroup": "The IDs of the security groups. You can create a security group using <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html\">CreateSecurityGroup. If you specify a network interface, you must specify any security groups as part of the network interface.",
		"subnet":        "The ID of the subnet to launch the instance into. If you specify a network interface, you must specify any subnets as part of the network interface.",
		"type":          "The instance type. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html\">Instance Types in the Amazon Elastic Compute Cloud User Guide. Default: m1.small",
		"userdata":      "The user data to make available to the instance. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html\">Running Commands on Your Linux Instance at Launch (Linux) and <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html#instancedata-add-user-data\">Adding User Data (Windows). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",
	},
	"create.instanceprofile": {
		"name": "The name of the instance profile to create. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"create.internetgateway": {},
	"create.launchconfiguration": {
		"image":          "The ID of the Amazon Machine Image (AMI) that was assigned during registration. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html\">Finding an AMI in the Amazon EC2 User Guide for Linux Instances. If you do not specify InstanceId, you must specify ImageId.",
		"keypair":        "The name of the key pair. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html\">Amazon EC2 Key Pairs in the Amazon EC2 User Guide for Linux Instances.",
		"name":           "The name of the launch configuration. This name must be unique per Region per account.",
		"public":         "For Auto Scaling groups that are running in a virtual private cloud (VPC), specifies whether to assign a public IP address to the group's instances. If you specify true, each instance in the Auto Scaling group receives a unique public IP address. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html\">Launching Auto Scaling Instances in a VPC in the Amazon EC2 Auto Scaling User Guide. If you specify this parameter, you must specify at least one subnet for VPCZoneIdentifier when you create your group.  If the instance is launched into a default subnet, the default is to assign a public IP address, unless you disabled the option to assign a public IP address on the subnet. If the instance is launched into a nondefault subnet, the default is not to assign a public IP address, unless you enabled the option to assign a public IP address on the subnet.",
		"role":           "The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role.  For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html\">IAM Role for Applications That Run on Amazon EC2 Instances in the Amazon EC2 Auto Scaling User Guide.",
		"securitygroups": "A list that contains the security groups to assign to the instances in the Auto Scaling group.  Specify the security group IDs. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html\">Security Groups for Your VPC in the Amazon Virtual Private Cloud User Guide.  Specify either the security group names or the security group IDs. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html\">Amazon EC2 Security Groups in the Amazon EC2 User Guide for Linux Instances.",
		"spotprice":      "The maximum hourly price to be paid for any Spot Instance launched to fulfill the request. Spot Instances are launched when the price you specify exceeds the current Spot market price. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html\">Launching Spot Instances in Your Auto Scaling Group in the Amazon EC2 Auto Scaling User Guide. If a Spot price is set, then the Auto Scaling group will only launch instances when the Spot price has been met, regardless of the setting in the Auto Scaling group's DesiredCapacity.   When you change your Spot price by creating a new launch configuration, running instances will continue to run as long as the Spot price for those running instances is higher than the current Spot market price.",
		"type":           "Specifies the instance type of the EC2 instance. For information about available instance types, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes\">Available Instance Types in the Amazon EC2 User Guide for Linux Instances.  If you do not specify InstanceId, you must specify InstanceType.",
		"userdata":       "The Base64-encoded user data to make available to the launched EC2 instances. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html\">Instance Metadata and User Data in the Amazon EC2 User Guide for Linux Instances.",
	},
	"create.listener": {
		"loadbalancer": "The Amazon Resource Name (ARN) of the load balancer.",
		"port":         "The port on which the load balancer is listening.",
		"protocol":     "The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP.",
		"sslpolicy":    "[HTTPS and TLS listeners] The security policy that defines which ciphers and protocols are supported. The default is the current predefined security policy.",
	},
	"create.loadbalancer": {
		"iptype":          "[Application Load Balancers] The type of IP addresses used by the subnets for your load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses). Internal load balancers must use ipv4.",
		"name":            "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".",
		"scheme":          "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet. The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can only route requests from clients with access to the VPC for the load balancer. The default is an Internet-facing load balancer.",
		"securitygroups":  "[Application Load Balancers] The IDs of the security groups for the load balancer.",
		"subnet-mappings": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings. [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets. [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet.",
		"subnets":         "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings. [Application Load Balancers] You must specify subnets from at least two Availability Zones. [Network Load Balancers] You can specify subnets from one or more Availability Zones.",
		"type":            "The type of load balancer. The default is application.",
	},
	"create.loginprofile": {
		"password":       "The new password for the user. The <a href=\"http://wikipedia.org/wiki/regex\">regex pattern that is used to validate this parameter is a string of characters. That string can include almost any printable ASCII character from the space (\u0020) through the end of the ASCII character range (\u00FF). You can also include the tab (\u0009), line feed (\u000A), and carriage return (\u000D) characters. Any of these characters are valid in a password. However, many tools, such as the AWS Management Console, might restrict the ability to type certain characters because they have special meaning within that tool.",
		"password-reset": "Specifies whether the user is required to set a new password on next sign-in.",
		"username":       "The name of the IAM user to create a password for. The user must already exist. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"create.mfadevice": {},
	"create.natgateway": {
		"elasticip-id": "The allocation ID of an Elastic IP address to associate with the NAT gateway. If the Elastic IP address is associated with another resource, you must first disassociate it.",
		"subnet":       "The subnet in which to create the NAT gateway.",
	},
	"create.networkinterface": {
		"description":    "A description for the network interface.",
		"privateip":      "The primary private IPv4 address of the network interface. If you don't specify an IPv4 address, Amazon EC2 selects one for you from the subnet's IPv4 CIDR range. If you specify an IP address, you cannot indicate any IP addresses specified in privateIpAddresses as primary (only one IP address can be designated as primary).",
		"securitygroups": "The IDs of one or more security groups.",
		"subnet":         "The ID of the subnet to associate with the network interface.",
	},
	"create.policy": {
		"description": "A friendly description of the policy. Typically used to store information about the permissions defined in the policy. For example, \"Grants access to production DynamoDB tables.\" The policy description is immutable. After a value is assigned, it cannot be changed.",
		"name":        "The friendly name of the policy. IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both \"MyResource\" and \"myresource\".",
	},
	"create.queue": {
		"name": "The name of the new queue. The following limits apply to this name:   A queue name can have up to 80 characters.   Valid values: alphanumeric characters, hyphens (-), and underscores (_).   A FIFO queue name must end with the .fifo suffix.   Queue URLs and names are case-sensitive.",
	},
	"create.record": {},
	"create.repository": {
		"name": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app).",
	},
	"create.role": {},
	"create.route": {
		"cidr":    "The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match.",
		"gateway": "The ID of an internet gateway or virtual private gateway attached to your VPC.",
		"table":   "The ID of the route table for the route.",
	},
	"create.routetable": {
		"vpc": "The ID of the VPC.",
	},
	"create.s3object": {},
	"create.scalinggroup": {
		"cooldown":                 "The amount of time, in seconds, after a scaling activity completes before another scaling activity can start. The default value is 300. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html\">Scaling Cooldowns in the Amazon EC2 Auto Scaling User Guide.",
		"desired-capacity":         "The number of Amazon EC2 instances that the Auto Scaling group attempts to maintain. This number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity, the default is the minimum size of the group.",
		"healthcheck-grace-period": "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service. During this time, any health check failures for the instance are ignored. The default value is 0. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period\">Health Check Grace Period in the Amazon EC2 Auto Scaling User Guide. Conditional: This parameter is required if you are adding an ELB health check.",
		"healthcheck-type":         "The service to use for the health checks. The valid values are EC2 and ELB. The default value is EC2. If you configure an Auto Scaling group to use ELB health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html\">Health Checks for Auto Scaling Instances in the Amazon EC2 Auto Scaling User Guide.",
		"launchconfiguration":      "The name of the launch configuration. If you do not specify LaunchConfigurationName, you must specify one of the following parameters: InstanceId, LaunchTemplate, or MixedInstancesPolicy.",
		"max-size":                 "The maximum size of the group.",
		"min-size":                 "The minimum size of the group.",
		"name":                     "The name of the Auto Scaling group. This name must be unique per Region per account.",
		"new-instances-protected":  "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in. For more information about preventing instances from terminating on scale in, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection\">Instance Protection in the Amazon EC2 Auto Scaling User Guide.",
		"subnets":                  "A comma-separated list of subnet IDs for your virtual private cloud (VPC). If you specify VPCZoneIdentifier with AvailabilityZones, the subnets that you specify for this parameter must reside in those Availability Zones. Conditional: If your account supports EC2-Classic and VPC, this parameter is required to launch instances into a VPC.",
		"targetgroups":             "The Amazon Resource Names (ARN) of the target groups to associate with the Auto Scaling group. Instances are registered as targets in a target group, and traffic is routed to the target group. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html\">Using a Load Balancer with an Auto Scaling Group in the Amazon EC2 Auto Scaling User Guide.",
	},
	"create.scalingpolicy": {
		"adjustment-magnitude": "The minimum number of instances to scale. If the value of AdjustmentType is PercentChangeInCapacity, the scaling policy changes the DesiredCapacity of the Auto Scaling group by at least this many instances. Otherwise, the error is ValidationError. This property replaces the MinAdjustmentStep property. For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances. Valid only if the policy type is SimpleScaling or StepScaling.",
		"adjustment-scaling":   "The amount by which a simple scaling policy scales the Auto Scaling group in response to an alarm breach. The adjustment is based on the value that you specified in the AdjustmentType parameter (either an absolute number or a percentage). A positive value adds to the current capacity and a negative value subtracts from the current capacity. For exact capacity, you must specify a positive value.  Conditional: If you specify SimpleScaling for the policy type, you must specify this parameter. (Not used with any other policy type.)",
		"adjustment-type":      "Specifies whether the ScalingAdjustment parameter is an absolute number or a percentage of the current capacity. The valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity. Valid only if the policy type is StepScaling or SimpleScaling. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment\">Scaling Adjustment Types in the Amazon EC2 Auto Scaling User Guide.",
		"cooldown":             "The amount of time, in seconds, after a scaling activity completes before any further dynamic scaling activities can start. If this parameter is not specified, the default cooldown period for the group applies. Valid only if the policy type is SimpleScaling. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html\">Scaling Cooldowns in the Amazon EC2 Auto Scaling User Guide.",
		"name":                 "The name of the scaling policy.",
		"scalinggroup":         "The name of the Auto Scaling group.",
	},
	"create.securitygroup": {
		"description": "A description for the security group. This is informational only. Constraints: Up to 255 characters in length Constraints for EC2-Classic: ASCII characters Constraints for EC2-VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&amp;;{}!$*",
		"name":        "The name of the security group. Constraints: Up to 255 characters in length. Cannot start with sg-. Constraints for EC2-Classic: ASCII characters Constraints for EC2-VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&amp;;{}!$*",
		"vpc":         "The ID of the VPC. Required for EC2-VPC.",
	},
	"create.snapshot": {
		"description": "A description for the snapshot.",
		"volume":      "The ID of the EBS volume.",
	},
	"create.subnet": {
		"availabilityzone": "The Availability Zone for the subnet. Default: AWS selects one for you. If you create more than one subnet in your VPC, we may not necessarily select a different zone for each subnet.",
		"cidr":             "The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24.",
		"vpc":              "The ID of the VPC.",
	},
	"create.subscription": {
		"endpoint": "The endpoint that you want to receive notifications. Endpoints vary by protocol:   For the http protocol, the endpoint is an URL beginning with \"https://\"   For the https protocol, the endpoint is a URL beginning with \"https://\"   For the email protocol, the endpoint is an email address   For the email-json protocol, the endpoint is an email address   For the sms protocol, the endpoint is a phone number of an SMS-enabled device   For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue   For the application protocol, the endpoint is the EndpointArn of a mobile app and device.   For the lambda protocol, the endpoint is the ARN of an AWS Lambda function.",
		"protocol": "The protocol you want to use. Supported protocols include:    http – delivery of JSON-encoded message via HTTP POST    https – delivery of JSON-encoded message via HTTPS POST    email – delivery of message via SMTP    email-json – delivery of JSON-encoded message via SMTP    sms – delivery of message via SMS    sqs – delivery of JSON-encoded message to an Amazon SQS queue    application – delivery of JSON-encoded message to an EndpointArn for a mobile app and device.    lambda – delivery of JSON-encoded message to an AWS Lambda function.",
		"topic":    "The ARN of the topic you want to subscribe to.",
	},
	"create.tag": {},
	"create.targetgroup": {
		"healthcheckinterval": "The approximate amount of time, in seconds, between health checks of an individual target. For HTTP and HTTPS health checks, the range is 5–300 seconds. For TCP health checks, the supported values are 10 and 30 seconds. If the target type is instance or ip, the default is 30 seconds. If the target type is lambda, the default is 35 seconds.",
		"healthcheckpath":     "[HTTP/HTTPS health checks] The ping path that is the destination on the targets for health checks. The default is /.",
		"healthcheckport":     "The port the load balancer uses when performing health checks on targets. The default is traffic-port, which is the port on which each target receives traffic from the load balancer.",
		"healthcheckprotocol": "The protocol the load balancer uses when performing health checks on targets. For Application Load Balancers, the default is HTTP. For Network Load Balancers, the default is TCP. The TCP protocol is supported for health checks only if the protocol of the target group is TCP, TLS, UDP, or TCP_UDP. The TLS, UDP, and TCP_UDP protocols are not supported for health checks.",
		"healthchecktimeout":  "The amount of time, in seconds, during which no response from a target means a failed health check. For target groups with a protocol of HTTP or HTTPS, the default is 5 seconds. For target groups with a protocol of TCP or TLS, this value must be 6 seconds for HTTP health checks and 10 seconds for TCP and HTTPS health checks. If the target type is lambda, the default is 30 seconds.",
		"healthythreshold":    "The number of consecutive health checks successes required before considering an unhealthy target healthy. For target groups with a protocol of HTTP or HTTPS, the default is 5. For target groups with a protocol of TCP or TLS, the default is 3. If the target type is lambda, the default is 5.",
		"name":                "The name of the target group. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.",
		"port":                "The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply.",
		"protocol":            "The protocol to use for routing traffic to the targets. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.",
		"unhealthythreshold":  "The number of consecutive health check failures required before considering a target unhealthy. For target groups with a protocol of HTTP or HTTPS, the default is 2. For target groups with a protocol of TCP or TLS, this value must be the same as the healthy threshold count. If the target type is lambda, the default is 2.",
		"vpc":                 "The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply.",
	},
	"create.topic": {
		"name": "The name of the topic you want to create. Constraints: Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long.",
	},
	"create.user": {
		"name": "The name of the user to create. IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both \"MyResource\" and \"myresource\".",
	},
	"create.volume": {
		"availabilityzone": "The Availability Zone in which to create the volume.",
		"size":             "The size of the volume, in GiBs. Constraints: 1-16,384 for gp2, 4-16,384 for io1, 500-16,384 for st1, 500-16,384 for sc1, and 1-1,024 for standard. If you specify a snapshot, the volume size must be equal to or larger than the snapshot size. Default: If you're creating the volume from a snapshot and don't specify a volume size, the default is the snapshot size.  At least one of Size or SnapshotId is required.",
	},
	"create.vpc": {
		"cidr": "The IPv4 network range for the VPC, in CIDR notation. For example, 10.0.0.0/16.",
	},
	"create.zone": {
		"callerreference": "A unique string that identifies the request and that allows failed CreateHostedZone requests to be retried without the risk of executing the operation twice. You must use a unique CallerReference string every time you submit a CreateHostedZone request. CallerReference can be any unique string, for example, a date/time stamp.",
		"delegationsetid": "If you want to associate a reusable delegation set with this hosted zone, the ID that Amazon Route 53 assigned to the reusable delegation set when you created it. For more information about reusable delegation sets, see <a href=\"https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html\">CreateReusableDelegationSet.",
		"name":            "The name of the domain. Specify a fully qualified domain name, for example, www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats www.example.com (without a trailing dot) and www.example.com. (with a trailing dot) as identical. If you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of NameServers that CreateHostedZone returns in DelegationSet.",
	},
	"delete.accesskey": {
		"id":   "The access key ID for the access key ID and secret access key you want to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters that can consist of any upper or lowercased letter or digit.",
		"user": "The name of the user whose access key pair you want to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"delete.alarm": {
		"name": "The alarms to be deleted.",
	},
	"delete.appscalingpolicy": {
		"dimension":         "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.    ecs:service:DesiredCount - The desired task count of an ECS service.    ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet request.    elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.    appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.    dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.    dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.    dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.    dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.    rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.    sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an Amazon SageMaker model endpoint variant.    custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.",
		"name":              "The name of the scaling policy.",
		"resource":          "The identifier of the resource associated with the scalable target. This string consists of the resource type and unique identifier.   ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.   Spot Fleet request - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.   EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.   AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.   DynamoDB table - The resource type is table and the unique identifier is the resource ID. Example: table/my-table.   DynamoDB global secondary index - The resource type is index and the unique identifier is the resource ID. Example: table/my-table/index/my-table-index.   Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.   Amazon SageMaker endpoint variants - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.   Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our <a href=\"https://github.com/aws/aws-auto-scaling-custom-resource\">GitHub repository.",
		"service-namespace": "The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service. For more information, see <a href=\"http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces\">AWS Service Namespaces in the Amazon Web Services General Reference.",
	},
	"delete.appscalingtarget": {
		"dimension":         "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.    ecs:service:DesiredCount - The desired task count of an ECS service.    ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet request.    elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.    appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.    dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.    dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.    dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.    dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.    rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.    sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an Amazon SageMaker model endpoint variant.    custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.",
		"resource":          "The identifier of the resource associated with the scalable target. This string consists of the resource type and unique identifier.   ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.   Spot Fleet request - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.   EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.   AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.   DynamoDB table - The resource type is table and the unique identifier is the resource ID. Example: table/my-table.   DynamoDB global secondary index - The resource type is index and the unique identifier is the resource ID. Example: table/my-table/index/my-table-index.   Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.   Amazon SageMaker endpoint variants - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.   Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our <a href=\"https://github.com/aws/aws-auto-scaling-custom-resource\">GitHub repository.",
		"service-namespace": "The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service. For more information, see <a href=\"http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces\">AWS Service Namespaces in the Amazon Web Services General Reference.",
	},
	"delete.bucket": {
		"name": "<p/>",
	},
	"delete.certificate": {
		"arn": "String that contains the ARN of the ACM certificate to be deleted. This must be of the form:  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012  For more information about ARNs, see <a href=\"https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html\">Amazon Resource Names (ARNs) and AWS Service Namespaces.",
	},
	"delete.classicloadbalancer": {},
	"delete.containercluster": {
		"id": "The short name or full Amazon Resource Name (ARN) of the cluster to delete.",
	},
	"delete.containertask": {},
	"delete.database": {
		"id": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
	},
	"delete.dbsubnetgroup": {},
	"delete.distribution":  {},
	"delete.elasticip": {
		"id": "The allocation ID. Required for EC2-VPC.",
		"ip": "The Elastic IP address. Required for EC2-Classic.",
	},
	"delete.function": {
		"id":      "The name of the Lambda function or version. <p class=\"title\"> Name formats     Function name - my-function (name-only), my-function:1 (with version).    Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.    Partial ARN - 123456789012:function:my-function.   You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
		"version": "Specify a version to delete. You can't delete a version that's referenced by an alias.",
	},
	"delete.group": {
		"name": "The name of the IAM group to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"delete.image": {},
	"delete.instance": {
		"ids": "The IDs of the instances. Constraints: Up to 1000 instance IDs. We recommend breaking up this request into smaller batches.",
	},
	"delete.instanceprofile": {
		"name": "The name of the instance profile to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"delete.internetgateway": {
		"id": "The ID of the internet gateway.",
	},
	"delete.launchconfiguration": {},
	"delete.listener": {
		"id": "The Amazon Resource Name (ARN) of the listener.",
	},
	"delete.loadbalancer": {
		"id": "The Amazon Resource Name (ARN) of the load balancer.",
	},
	"delete.loginprofile": {
		"username": "The name of the user whose password you want to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"delete.mfadevice": {
		"id": "The serial number that uniquely identifies the MFA device. For virtual MFA devices, the serial number is the same as the ARN. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: =,.@:/-",
	},
	"delete.natgateway": {
		"id": "The ID of the NAT gateway.",
	},
	"delete.networkinterface": {
		"id": "The ID of the network interface.",
	},
	"delete.policy": {
		"arn": "The Amazon Resource Name (ARN) of the IAM policy you want to delete. For more information about ARNs, see <a href=\"https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html\">Amazon Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.",
	},
	"delete.queue": {
		"url": "The URL of the Amazon SQS queue to delete. Queue URLs and names are case-sensitive.",
	},
	"delete.record": {},
	"delete.repository": {
		"account": "The AWS account ID associated with the registry that contains the repository to delete. If you do not specify a registry, the default registry is assumed.",
		"force":   "If a repository contains images, forces the deletion.",
		"name":    "The name of the repository to delete.",
	},
	"delete.role": {},
	"delete.route": {
		"cidr":  "The IPv4 CIDR range for the route. The value you specify must match the CIDR for the route exactly.",
		"table": "The ID of the route table.",
	},
	"delete.routetable": {
		"id": "The ID of the route table.",
	},
	"delete.s3object": {
		"bucket": "<p/>",
		"name":   "<p/>",
	},
	"delete.scalinggroup": {
		"force": "Specifies that the group is to be deleted along with all instances associated with the group, without waiting for all instances to be terminated. This parameter also deletes any lifecycle actions associated with the group.",
		"name":  "The name of the Auto Scaling group.",
	},
	"delete.scalingpolicy": {
		"id": "The name or Amazon Resource Name (ARN) of the policy.",
	},
	"delete.securitygroup": {
		"id": "The ID of the security group. Required for a nondefault VPC.",
	},
	"delete.snapshot": {
		"id": "The ID of the EBS snapshot.",
	},
	"delete.subnet": {
		"id": "The ID of the subnet.",
	},
	"delete.subscription": {
		"id": "The ARN of the subscription to be deleted.",
	},
	"delete.tag": {},
	"delete.targetgroup": {
		"id": "The Amazon Resource Name (ARN) of the target group.",
	},
	"delete.topic": {
		"id": "The ARN of the topic you want to delete.",
	},
	"delete.user": {
		"name": "The name of the user to delete. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"delete.volume": {
		"id": "The ID of the volume.",
	},
	"delete.vpc": {
		"id": "The ID of the VPC.",
	},
	"delete.zone": {
		"id": "The ID of the hosted zone you want to delete.",
	},
	"detach.alarm":               {},
	"detach.classicloadbalancer": {},
	"detach.containertask":       {},
	"detach.elasticip": {
		"association": "The association ID. Required for EC2-VPC.",
	},
	"detach.instance": {
		"targetgroup": "The Amazon Resource Name (ARN) of the target group.",
	},
	"detach.instanceprofile": {},
	"detach.internetgateway": {
		"id":  "The ID of the internet gateway.",
		"vpc": "The ID of the VPC.",
	},
	"detach.mfadevice": {
		"id":   "The serial number that uniquely identifies the MFA device. For virtual MFA devices, the serial number is the device ARN. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: =,.@:/-",
		"user": "The name of the user whose MFA device you want to deactivate. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"detach.networkinterface": {},
	"detach.policy":           {},
	"detach.role": {
		"instanceprofile": "The name of the instance profile to update. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		"name":            "The name of the role to remove. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"detach.routetable": {
		"association": "The association ID representing the current association between the route table and subnet.",
	},
	"detach.securitygroup": {},
	"detach.user": {
		"group": "The name of the group to update. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		"name":  "The name of the user to remove. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"detach.volume": {
		"device":   "The device name.",
		"force":    "Forces detachment if the previous detachment attempt did not occur cleanly (for example, logging into an instance, unmounting the volume, and detaching normally). This option can lead to data loss or a corrupted file system. Use this option only as a last resort to detach a volume from a failed instance. The instance won't have an opportunity to flush file system caches or file system metadata. If you use this option, you must perform file system check and repair procedures.",
		"id":       "The ID of the volume.",
		"instance": "The ID of the instance.",
	},
	"import.image": {
		"architecture": "The architecture of the virtual machine. Valid values: i386 | x86_64 | arm64",
		"description":  "A description string for the import image task.",
		"license":      "The license type to be used for the Amazon Machine Image (AMI) after importing. By default, we detect the source-system operating system (OS) and apply the appropriate license. Specify AWS to replace the source-system license with an AWS license, if appropriate. Specify BYOL to retain the source-system license, if appropriate. To use BYOL, you must have existing licenses with rights to use these licenses in a third party cloud, such as AWS. For more information, see <a href=\"https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#prerequisites-image\">Prerequisites in the VM Import/Export User Guide.",
		"platform":     "The operating system of the virtual machine. Valid values: Windows | Linux",
		"role":         "The name of the role to use when not using the default role, 'vmimport'.",
	},
	"restart.database": {
		"id": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
	},
	"restart.instance": {
		"ids": "The instance IDs.",
	},
	"start.alarm": {
		"names": "The names of the alarms.",
	},
	"start.containertask": {},
	"start.database": {
		"id": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
	},
	"start.instance": {
		"ids": "The IDs of the instances.",
	},
	"stop.alarm": {
		"names": "The names of the alarms.",
	},
	"stop.containertask": {},
	"stop.database": {
		"id": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
	},
	"stop.instance": {
		"ids": "The IDs of the instances.",
	},
	"update.bucket": {},
	"update.classicloadbalancer": {
		"name": "The name of the load balancer.",
	},
	"update.containertask": {
		"cluster":         "The short name or full Amazon Resource Name (ARN) of the cluster that your service is running on. If you do not specify a cluster, the default cluster is assumed.",
		"deployment-name": "The name of the service to update.",
		"desired-count":   "The number of instantiations of the task to place and keep running in your service.",
		"name":            "The family and revision (family:revision) or full ARN of the task definition to run in your service. If a revision is not specified, the latest ACTIVE revision is used. If you modify the task definition with UpdateService, Amazon ECS spawns a task with the new version of the task definition and then stops an old task after the new version is running.",
	},
	"update.distribution": {},
	"update.image":        {},
	"update.instance": {
		"id":   "The ID of the instance.",
		"lock": "If the value is true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. You cannot use this parameter for Spot Instances.",
	},
	"update.loginprofile": {
		"password":       "The new password for the specified IAM user. The <a href=\"http://wikipedia.org/wiki/regex\">regex pattern used to validate this parameter is a string of characters consisting of the following:   Any printable ASCII character ranging from the space character (\u0020) through the end of the ASCII character range   The printable characters in the Basic Latin and Latin-1 Supplement character set (through \u00FF)   The special characters tab (\u0009), line feed (\u000A), and carriage return (\u000D)   However, the format can be further restricted by the account administrator by setting a password policy on the AWS account. For more information, see UpdateAccountPasswordPolicy.",
		"password-reset": "Allows this new password to be used only once by requiring the specified IAM user to set a new password on next sign-in.",
		"username":       "The name of the user whose password you want to update. This parameter allows (through its <a href=\"http://wikipedia.org/wiki/regex\">regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
	},
	"update.policy": {
		"arn": "The Amazon Resource Name (ARN) of the IAM policy to which you want to add a new version. For more information about ARNs, see <a href=\"https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html\">Amazon Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.",
	},
	"update.record": {},
	"update.s3object": {
		"acl":     "The canned ACL to apply to the object.",
		"bucket":  "<p/>",
		"name":    "<p/>",
		"version": "VersionId used to reference a specific version of the object.",
	},
	"update.scalinggroup": {
		"cooldown":                 "The amount of time, in seconds, after a scaling activity completes before another scaling activity can start. The default value is 300. This cooldown period is not used when a scaling-specific cooldown is specified. Cooldown periods are not supported for target tracking scaling policies, step scaling policies, or scheduled scaling. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html\">Scaling Cooldowns in the Amazon EC2 Auto Scaling User Guide.",
		"desired-capacity":         "The number of EC2 instances that should be running in the Auto Scaling group. This number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group.",
		"healthcheck-grace-period": "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service. The default value is 0. For more information, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period\">Health Check Grace Period in the Amazon EC2 Auto Scaling User Guide. Conditional: This parameter is required if you are adding an ELB health check.",
		"healthcheck-type":         "The service to use for the health checks. The valid values are EC2 and ELB. If you configure an Auto Scaling group to use ELB health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks.",
		"launchconfiguration":      "The name of the launch configuration. If you specify LaunchConfigurationName in your update request, you can't specify LaunchTemplate or MixedInstancesPolicy.  To update an Auto Scaling group with a launch configuration with InstanceMonitoring set to false, you must first disable the collection of group metrics. Otherwise, you get an error. If you have previously enabled the collection of group metrics, you can disable it using DisableMetricsCollection.",
		"max-size":                 "The maximum size of the Auto Scaling group.",
		"min-size":                 "The minimum size of the Auto Scaling group.",
		"name":                     "The name of the Auto Scaling group.",
		"new-instances-protected":  "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in. For more information about preventing instances from terminating on scale in, see <a href=\"https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection\">Instance Protection in the Amazon EC2 Auto Scaling User Guide.",
		"subnets":                  "A comma-separated list of subnet IDs for virtual private cloud (VPC). If you specify VPCZoneIdentifier with AvailabilityZones, the subnets that you specify for this parameter must reside in those Availability Zones.",
	},
	"update.securitygroup": {},
	"update.subnet": {
		"id":     "The ID of the subnet.",
		"public": "Specify true to indicate that ENIs attached to instances created in the specified subnet should be assigned a public IPv4 address.",
	},
	"update.targetgroup": {},
}
