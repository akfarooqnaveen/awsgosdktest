package main

import (
	// "context"
	"log"
	aws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/service/dlm"
)



func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	// config, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create an Amazon S3 service client
	//client := ebs.NewFromConfig(config)
	session := session.New(&aws.Config{Region: aws.String("us-east-2")})
	// mySession := session.Must(session.NewSession())

	// Create a DLM client from just a session.
	client := dlm.New(session)

	var time string = "09:00"
	times := []*string{&time}
	var cr dlm.CreateRule
	cr.SetInterval(12)
	cr.SetIntervalUnit("HOURS")
	cr.SetLocation("CLOUD")
	cr.SetTimes(times)

	var rr dlm.RetainRule
	rr.SetCount(1)
	// rr.SetInterval(1)

	schedules := []*dlm.Schedule{}
	var s dlm.Schedule
	s.SetCreateRule(&cr)
	s.SetName("Schedule1")
	s.SetRetainRule(&rr)

	schedules = append(schedules, &s)

	var ResourceType string = "VOLUME"
	var ResourceLocation string ="CLOUD"
	resourceTypes := []*string{&ResourceType}
	resourceLocations := []*string{&ResourceLocation}

	var Key string = "snapshottest"
	var Value string = "true"
	var targetTag dlm.Tag = dlm.Tag{
		Key: &Key,
		Value: &Value,
	}
	targetTags := []*dlm.Tag{}
	targetTags = append(targetTags, &targetTag)
	// var targetTags map[string]*string
	// var Value string = "true"
	// targetTags["snapshottest"] = &Value
	// targetTags := []aws.Tags{}
	// var targetTag1 Tags
	// targetTag1.SetKey("snapshottest")
	// targetTag1.SetValue("true")

	// targetTags = append(targetTags, targetTag1)

	var p dlm.PolicyDetails
	// p.SetParameters()
	p.SetPolicyType("EBS_SNAPSHOT_MANAGEMENT")
	p.SetResourceLocations(resourceLocations)
	p.SetResourceTypes(resourceTypes)
	p.SetSchedules(schedules)
	p.SetTargetTags(targetTags)

	var lpi dlm.CreateLifecyclePolicyInput
	lpi.SetDescription("frq-policy")
	lpi.SetExecutionRoleArn("arn:aws:iam::731556103348:role/service-role/AWSDataLifecycleManagerDefaultRole")
	lpi.SetPolicyDetails(&p)
	lpi.SetState("ENABLED")

	req, output := client.CreateLifecyclePolicyRequest(&lpi)
	// CreateLifecyclePolicyInput{Description: "policy123",ExecutionRoleArn: "",PolicyDetails: "",State: ""}        
	// Get the first page of results for ListObjectsV2 for a bucket
	//output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	//	Bucket: aws.String("my-bucket"),
	//})

	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Println("Before send")
	//for _, object := range output.Contents {
	//	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	//}

	// req, resp := client.CreateLifecyclePolicyRequest(&ebs.CreateLifecyclePolicyInput{
	// 	Description: aws.String("frq-sdk-policy"),

	// })

	err := req.Send()
	if err == nil { // resp is now filled
		log.Println(output)
	}else{
		log.Fatal(err)
	}
}

