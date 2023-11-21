package ka

import "fmt"

func ListTopic() {
	partitions, err := MyConn.ReadPartitions()
	if err != nil {
		fmt.Println(err.Error())
	}
	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
