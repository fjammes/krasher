package framework

// import (
// 	"k8s.io/api/core/v1"
//     clientset "k8s.io/client-go/kubernetes"
// )
// // Framework supports common operations used by e2e tests; it will keep a client & a namespace for you.
// // Eventual goal is to merge this with integration test framework.
// type Framework struct {
// 	BaseName string

// 	// Set together with creating the ClientSet and the namespace.
// 	// Guaranteed to be unique in the cluster even when running the same
// 	// test multiple times in parallel.
// 	UniqueName string

// 	ClientSet clientset.Interface
// 	Namespace *v1.Namespace // Every test has at least one namespace unless creation is skipped
// }
