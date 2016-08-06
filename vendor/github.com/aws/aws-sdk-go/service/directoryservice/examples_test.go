// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directoryservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDirectoryService_AddIpRoutes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.AddIpRoutesInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		IpRoutes: []*directoryservice.IpRoute{ // Required
			{ // Required
				CidrIp:      aws.String("CidrIp"),
				Description: aws.String("Description"),
			},
			// More values...
		},
		UpdateSecurityGroupForDirectoryControllers: aws.Bool(true),
	}
	resp, err := svc.AddIpRoutes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_AddTagsToResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.AddTagsToResourceInput{
		ResourceId: aws.String("ResourceId"), // Required
		Tags: []*directoryservice.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_ConnectDirectory() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.ConnectDirectoryInput{
		ConnectSettings: &directoryservice.DirectoryConnectSettings{ // Required
			CustomerDnsIps: []*string{ // Required
				aws.String("IpAddr"), // Required
				// More values...
			},
			CustomerUserName: aws.String("UserName"), // Required
			SubnetIds: []*string{ // Required
				aws.String("SubnetId"), // Required
				// More values...
			},
			VpcId: aws.String("VpcId"), // Required
		},
		Name:        aws.String("DirectoryName"),   // Required
		Password:    aws.String("ConnectPassword"), // Required
		Size:        aws.String("DirectorySize"),   // Required
		Description: aws.String("Description"),
		ShortName:   aws.String("DirectoryShortName"),
	}
	resp, err := svc.ConnectDirectory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateAliasInput{
		Alias:       aws.String("AliasName"),   // Required
		DirectoryId: aws.String("DirectoryId"), // Required
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateComputer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateComputerInput{
		ComputerName: aws.String("ComputerName"),     // Required
		DirectoryId:  aws.String("DirectoryId"),      // Required
		Password:     aws.String("ComputerPassword"), // Required
		ComputerAttributes: []*directoryservice.Attribute{
			{ // Required
				Name:  aws.String("AttributeName"),
				Value: aws.String("AttributeValue"),
			},
			// More values...
		},
		OrganizationalUnitDistinguishedName: aws.String("OrganizationalUnitDN"),
	}
	resp, err := svc.CreateComputer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateConditionalForwarder() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateConditionalForwarderInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		DnsIpAddrs: []*string{ // Required
			aws.String("IpAddr"), // Required
			// More values...
		},
		RemoteDomainName: aws.String("RemoteDomainName"), // Required
	}
	resp, err := svc.CreateConditionalForwarder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateDirectory() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateDirectoryInput{
		Name:        aws.String("DirectoryName"), // Required
		Password:    aws.String("Password"),      // Required
		Size:        aws.String("DirectorySize"), // Required
		Description: aws.String("Description"),
		ShortName:   aws.String("DirectoryShortName"),
		VpcSettings: &directoryservice.DirectoryVpcSettings{
			SubnetIds: []*string{ // Required
				aws.String("SubnetId"), // Required
				// More values...
			},
			VpcId: aws.String("VpcId"), // Required
		},
	}
	resp, err := svc.CreateDirectory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateMicrosoftAD() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateMicrosoftADInput{
		Name:     aws.String("DirectoryName"), // Required
		Password: aws.String("Password"),      // Required
		VpcSettings: &directoryservice.DirectoryVpcSettings{ // Required
			SubnetIds: []*string{ // Required
				aws.String("SubnetId"), // Required
				// More values...
			},
			VpcId: aws.String("VpcId"), // Required
		},
		Description: aws.String("Description"),
		ShortName:   aws.String("DirectoryShortName"),
	}
	resp, err := svc.CreateMicrosoftAD(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateSnapshotInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		Name:        aws.String("SnapshotName"),
	}
	resp, err := svc.CreateSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_CreateTrust() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.CreateTrustInput{
		DirectoryId:      aws.String("DirectoryId"),      // Required
		RemoteDomainName: aws.String("RemoteDomainName"), // Required
		TrustDirection:   aws.String("TrustDirection"),   // Required
		TrustPassword:    aws.String("TrustPassword"),    // Required
		ConditionalForwarderIpAddrs: []*string{
			aws.String("IpAddr"), // Required
			// More values...
		},
		TrustType: aws.String("TrustType"),
	}
	resp, err := svc.CreateTrust(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DeleteConditionalForwarder() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DeleteConditionalForwarderInput{
		DirectoryId:      aws.String("DirectoryId"),      // Required
		RemoteDomainName: aws.String("RemoteDomainName"), // Required
	}
	resp, err := svc.DeleteConditionalForwarder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DeleteDirectory() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DeleteDirectoryInput{
		DirectoryId: aws.String("DirectoryId"), // Required
	}
	resp, err := svc.DeleteDirectory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DeleteSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DeleteSnapshotInput{
		SnapshotId: aws.String("SnapshotId"), // Required
	}
	resp, err := svc.DeleteSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DeleteTrust() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DeleteTrustInput{
		TrustId: aws.String("TrustId"), // Required
		DeleteAssociatedConditionalForwarder: aws.Bool(true),
	}
	resp, err := svc.DeleteTrust(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DeregisterEventTopic() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DeregisterEventTopicInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		TopicName:   aws.String("TopicName"),   // Required
	}
	resp, err := svc.DeregisterEventTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DescribeConditionalForwarders() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DescribeConditionalForwardersInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		RemoteDomainNames: []*string{
			aws.String("RemoteDomainName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConditionalForwarders(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DescribeDirectories() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DescribeDirectoriesInput{
		DirectoryIds: []*string{
			aws.String("DirectoryId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.DescribeDirectories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DescribeEventTopics() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DescribeEventTopicsInput{
		DirectoryId: aws.String("DirectoryId"),
		TopicNames: []*string{
			aws.String("TopicName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeEventTopics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DescribeSnapshots() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DescribeSnapshotsInput{
		DirectoryId: aws.String("DirectoryId"),
		Limit:       aws.Int64(1),
		NextToken:   aws.String("NextToken"),
		SnapshotIds: []*string{
			aws.String("SnapshotId"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeSnapshots(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DescribeTrusts() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DescribeTrustsInput{
		DirectoryId: aws.String("DirectoryId"),
		Limit:       aws.Int64(1),
		NextToken:   aws.String("NextToken"),
		TrustIds: []*string{
			aws.String("TrustId"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrusts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DisableRadius() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DisableRadiusInput{
		DirectoryId: aws.String("DirectoryId"), // Required
	}
	resp, err := svc.DisableRadius(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_DisableSso() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.DisableSsoInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		Password:    aws.String("ConnectPassword"),
		UserName:    aws.String("UserName"),
	}
	resp, err := svc.DisableSso(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_EnableRadius() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.EnableRadiusInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		RadiusSettings: &directoryservice.RadiusSettings{ // Required
			AuthenticationProtocol: aws.String("RadiusAuthenticationProtocol"),
			DisplayLabel:           aws.String("RadiusDisplayLabel"),
			RadiusPort:             aws.Int64(1),
			RadiusRetries:          aws.Int64(1),
			RadiusServers: []*string{
				aws.String("Server"), // Required
				// More values...
			},
			RadiusTimeout:   aws.Int64(1),
			SharedSecret:    aws.String("RadiusSharedSecret"),
			UseSameUsername: aws.Bool(true),
		},
	}
	resp, err := svc.EnableRadius(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_EnableSso() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.EnableSsoInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		Password:    aws.String("ConnectPassword"),
		UserName:    aws.String("UserName"),
	}
	resp, err := svc.EnableSso(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_GetDirectoryLimits() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	var params *directoryservice.GetDirectoryLimitsInput
	resp, err := svc.GetDirectoryLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_GetSnapshotLimits() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.GetSnapshotLimitsInput{
		DirectoryId: aws.String("DirectoryId"), // Required
	}
	resp, err := svc.GetSnapshotLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_ListIpRoutes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.ListIpRoutesInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		Limit:       aws.Int64(1),
		NextToken:   aws.String("NextToken"),
	}
	resp, err := svc.ListIpRoutes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_ListTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.ListTagsForResourceInput{
		ResourceId: aws.String("ResourceId"), // Required
		Limit:      aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_RegisterEventTopic() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.RegisterEventTopicInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		TopicName:   aws.String("TopicName"),   // Required
	}
	resp, err := svc.RegisterEventTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_RemoveIpRoutes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.RemoveIpRoutesInput{
		CidrIps: []*string{ // Required
			aws.String("CidrIp"), // Required
			// More values...
		},
		DirectoryId: aws.String("DirectoryId"), // Required
	}
	resp, err := svc.RemoveIpRoutes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_RemoveTagsFromResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.RemoveTagsFromResourceInput{
		ResourceId: aws.String("ResourceId"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_RestoreFromSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.RestoreFromSnapshotInput{
		SnapshotId: aws.String("SnapshotId"), // Required
	}
	resp, err := svc.RestoreFromSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_UpdateConditionalForwarder() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.UpdateConditionalForwarderInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		DnsIpAddrs: []*string{ // Required
			aws.String("IpAddr"), // Required
			// More values...
		},
		RemoteDomainName: aws.String("RemoteDomainName"), // Required
	}
	resp, err := svc.UpdateConditionalForwarder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_UpdateRadius() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.UpdateRadiusInput{
		DirectoryId: aws.String("DirectoryId"), // Required
		RadiusSettings: &directoryservice.RadiusSettings{ // Required
			AuthenticationProtocol: aws.String("RadiusAuthenticationProtocol"),
			DisplayLabel:           aws.String("RadiusDisplayLabel"),
			RadiusPort:             aws.Int64(1),
			RadiusRetries:          aws.Int64(1),
			RadiusServers: []*string{
				aws.String("Server"), // Required
				// More values...
			},
			RadiusTimeout:   aws.Int64(1),
			SharedSecret:    aws.String("RadiusSharedSecret"),
			UseSameUsername: aws.Bool(true),
		},
	}
	resp, err := svc.UpdateRadius(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectoryService_VerifyTrust() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := directoryservice.New(sess)

	params := &directoryservice.VerifyTrustInput{
		TrustId: aws.String("TrustId"), // Required
	}
	resp, err := svc.VerifyTrust(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
