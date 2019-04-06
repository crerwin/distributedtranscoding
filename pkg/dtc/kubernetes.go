package dtc

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	kubeconfig string
	clientset  *kubernetes.Clientset
	workspace  string
}

func NewKubeClient() *KubeClient {
	client := new(KubeClient)
	client.workspace = "/data"
	client.kubeconfig = getKubeConfig()
	config, err := clientcmd.BuildConfigFromFlags("", client.kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	client.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return client
}

func getKubeConfig() string {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	return *kubeconfig
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func (c *KubeClient) GetJobs() {
	jobs, err := c.clientset.BatchV1().Jobs("").List(metav1.ListOptions{})
	for _, job := range jobs.Items {
		fmt.Println(job.Name)
	}

	if err != nil {
		panic(err.Error())
	}
}

func (c *KubeClient) Init() {
	ns, err := c.clientset.CoreV1().Namespaces().Get("dtc", metav1.GetOptions{})
	if err != nil {
		if err.Error() == "namespaces \"dtc\" not found" {
			fmt.Println("dtc namespace does not exist, creating...")
			newns := &apiv1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "dtc",
					Labels: map[string]string{
						"name": "dtc",
					},
				},
			}
			_, err2 := c.clientset.CoreV1().Namespaces().Create(newns)
			if err2 != nil {
				panic(err2.Error())
			}
		} else {
			panic(err.Error())
		}
		fmt.Println(ns.Name)
	} else {
		fmt.Println("dtc namespace already exists - ready to go.")
	}
}

func createJobName(prefix string, fileName string) string {
	maxLength := 253
	// start with a prefix
	jobName := prefix

	// get the extension
	extension := filepath.Ext(fileName)

	// remove the extension
	fileMinusExtension := strings.TrimSuffix(fileName, extension)

	// create a regex object to remove non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// remove non-alphanumeric characters
	cleanedFileName := reg.ReplaceAllString(fileMinusExtension, "")

	// make all characters lowercase
	cleanedFileName = strings.ToLower(cleanedFileName)

	// append cleanedFileName to jobName
	jobName = jobName + cleanedFileName

	// Kubernetes resource names have to be 253 characters or less
	if len(jobName) <= maxLength {
		return jobName
	} else {
		jobName = jobName[0:maxLength]
		return jobName
	}
}

func (c *KubeClient) CreateTranscodeJob(j *Job) {
	name := createJobName("dtc-", j.Item.FileName)
	kubejob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: "ntodd/video-transcoding",
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "transcode",
									MountPath: "/data",
								},
							},
							Command: []string{
								"transcode-video",
								"--crop", j.Item.Crop,
								"--no-log",
								"--output", filepath.Join(c.workspace, j.OutboxPath, j.Item.SubPath, j.Item.FileName),
								filepath.Join(c.workspace, j.InboxPath, j.Item.SubPath, j.Item.FileName),
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []apiv1.Volume{
						apiv1.Volume{
							Name: "transcode",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: "transcode",
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := c.clientset.BatchV1().Jobs("default").Create(kubejob)
	if err != nil {
		panic(err.Error())
	}
}
