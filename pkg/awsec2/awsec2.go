package awsec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AwsEc2 struct {
	svc *ec2.Client
}

type Ec2 struct {
	Id         string     `json:"id"`
	LaunchTime *time.Time `json:"launch_time"`
}

type Ec2s []Ec2

func New() *AwsEc2 {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	svc := ec2.NewFromConfig(cfg)
	return &AwsEc2{
		svc: svc,
	}
}

func (a *AwsEc2) Ls(params *ec2.DescribeInstancesInput) Ec2s {

	l := a.describeEc2(params)
	es := Ec2s{}
	for _, v := range l.Reservations {
		for _, v2 := range v.Instances {

			e := Ec2{
				Id:         aws.ToString(v2.InstanceId),
				LaunchTime: v2.LaunchTime,
			}
			es = append(es, e)
		}
	}
	return es
}

func (a *AwsEc2) describeEc2(params *ec2.DescribeInstancesInput) *ec2.DescribeInstancesOutput {
	svc := a.svc

	resp, err := svc.DescribeInstances(context.Background(), params)

	if err != nil {
		panic("failed to describe instances, " + err.Error())
	}

	//	j, _ := json.Marshal(resp)
	//	fmt.Println(string(j))

	return resp
}
