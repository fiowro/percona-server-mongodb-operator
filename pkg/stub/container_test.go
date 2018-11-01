package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestGetContainer(t *testing.T) {
	pod := corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name: t.Name(),
				},
			},
		},
	}
	assert.NotNil(t, getContainer(pod, t.Name()))
	assert.Nil(t, getContainer(pod, "doesnt exist"))
}

func TestGetMongodPort(t *testing.T) {
	assert.Equal(t, "27017", getMongodPort(&corev1.Container{
		Ports: []corev1.ContainerPort{
			{
				Name:          mongodPortName,
				ContainerPort: int32(27017),
			},
		},
	}))
	assert.Equal(t, "", getMongodPort(&corev1.Container{
		Ports: []corev1.ContainerPort{},
	}))
}

func TestGetWiredTigerCacheSizeGB(t *testing.T) {
	memoryQuantity := resource.NewQuantity(gigaByte/2, resource.DecimalSI)
	assert.Equal(t, 0.25, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))

	memoryQuantity = resource.NewQuantity(1*gigaByte, resource.DecimalSI)
	assert.Equal(t, 0.25, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))

	memoryQuantity = resource.NewQuantity(4*gigaByte, resource.DecimalSI)
	assert.Equal(t, 1.5, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))

	memoryQuantity = resource.NewQuantity(64*gigaByte, resource.DecimalSI)
	assert.Equal(t, 31.5, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))

	memoryQuantity = resource.NewQuantity(128*gigaByte, resource.DecimalSI)
	assert.Equal(t, 63.5, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))

	memoryQuantity = resource.NewQuantity(256*gigaByte, resource.DecimalSI)
	assert.Equal(t, 127.5, getWiredTigerCacheSizeGB(corev1.ResourceList{corev1.ResourceMemory: *memoryQuantity}, 0.5))
}
