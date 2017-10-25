package nodepool

import (
	"github.com/jetstack-experimental/navigator/pkg/apis/navigator/v1alpha1"
	"github.com/jetstack-experimental/navigator/pkg/controllers/cassandra/util"
	"k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func StatefulSetForCluster(
	cluster *v1alpha1.CassandraCluster,
	nodepool *v1alpha1.CassandraClusterNodePool,
) *v1beta2.StatefulSet {
	set := &v1beta2.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:            util.NodePoolResourceName(cluster, nodepool),
			Namespace:       cluster.Namespace,
			Labels:          util.ClusterLabels(cluster),
			Annotations:     make(map[string]string),
			OwnerReferences: []metav1.OwnerReference{util.NewControllerRef(cluster)},
		},
	}
	return set
}
