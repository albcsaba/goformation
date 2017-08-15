package resources

// AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html
type AWSKinesisFirehoseDeliveryStreamKMSEncryptionConfig struct {

	// AWSKMSKeyARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig-awskmskeyarn
	AWSKMSKeyARN string `json:"AWSKMSKeyARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStreamKMSEncryptionConfig) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStreamKMSEncryptionConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}