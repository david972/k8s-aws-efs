package efsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
)

// CreateMount creates EFS a mount point.
func CreateMount(svc *efs.EFS, id, subnet, securities) (*efs.MountTargetDescription, error) {
	// Check if a mount exists in this subnet.
	mnts, err := svc.DescribeMountTargets(&efs.DescribeMountTargetsInput{
		FileSystemId: aws.String(id),
	})
	if err != nil {
		return nil, err
	}

	for _, mount := range mnts.MountTargets {
		// Check if we have already setup a mount point on a specific subnet.
		if *mount.SubnetId == subnet {
			return mount, nil
		}
	}

	// Create one if it does not exist.
	return svc.CreateMountTarget(&efs.CreateMountTargetInput{
		FileSystemId:   aws.String(id),
		SubnetId:       aws.String(subnet),
		SecurityGroups: securities,
	})
}
