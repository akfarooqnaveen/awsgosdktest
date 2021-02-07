package main

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
)



func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	//client := ebs.NewFromConfig(config)
	session := 
	mySession := session.Must(session.NewSession())

	// Create a DLM client from just a session.
	svc := dlm.New(mySession)

	times := []string{"09:00"}
	var cr CreateRule
	cr.SetInterval(12)
	cr.SetIntervalUnit("HOURS")
	cr.SetLocation("CLOUD")
	cr.SetTimes(&times)

	var rr RetainRule
	rr.SetCount(1)
	rr.SetInterval(0)

	schedules := []Schedule
	var s Schedule
	s.SetCreateRule(&cr)
	s.SetName("Schedule1")
	s.SetRetainRule(&rr)

	schedules = append(schedules, s)

	resourceTypes := []string{"VOLUME"}
	resourceLocations := []string{"CLOUD"}

	targetTags := []Tags
	var targetTag1 Tags
	targetTag1.SetKey("snapshottest")
	targetTag1.SetValue("true")

	targetTags = append(targetTags, targetTag1)

	var p PolicyDetails
	// p.SetParameters()
	p.SetPolicyType("EBS_SNAPSHOT_MANAGEMENT")
	p.SetResourceLocations(&resourceLocations)
	p.SetResourceTypes(&resourceTypes)
	p.SetSchedules(&schedules)
	p.SetTargetTags(&targetTags)

	var lpi CreateLifecyclePolicyInput
	lpi.SetDescription("frq-policy")
	lpi.SetExecutionRoleArn("")
	lpi.SetPolicyDetails("")
	lpi.SetState("")

	CreateLifecyclePolicyInput{Description: "policy123",ExecutionRoleArn: "",PolicyDetails: "",State: ""}        
	// Get the first page of results for ListObjectsV2 for a bucket
	//output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	//	Bucket: aws.String("my-bucket"),
	//})

	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Println("first page results:")
	//for _, object := range output.Contents {
	//	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	//}

	req, resp := client.CreateLifecyclePolicyRequest(&ebs.CreateLifecyclePolicyInput{
		Description: aws.String("frq-sdk-policy"),

	})

	err := req.Send()
	if err == nil { // resp is now filled
		fmt.Println(resp)
	}
}

