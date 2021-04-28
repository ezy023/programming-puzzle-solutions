package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)  // email -> person
	graph := make(map[string][]string) // email -> associated emails
	for _, account := range accounts {
		person := account[0]
		for _, email := range account[1:] {
			parent[email] = person
			var related []string
			if r, ok := graph[email]; ok {
				related = r
			} else {
				related = make([]string, 0)
			}
			for _, e := range account[1:] {
				if email != e {
					related = append(related, e)
				}
			}
			graph[email] = related
		}
	}

	output := make([][]string, 0)
	seen := make(map[string]bool)
	for email, assoc := range graph {
		if _, ok := seen[email]; ok {
			continue
		}
		buf := make([]string, 0)
		if _, ok := seen[email]; !ok {
			buf = append(buf, email)
			seen[email] = true
		}
		queue := make([]string, 0, len(assoc))
		queue = append(queue, assoc...)
		for len(queue) > 0 {
			e := queue[0]
			if _, ok := seen[e]; !ok {
				queue = append(queue, graph[e]...)
				buf = append(buf, e)
			}
			seen[e] = true
			queue = queue[1:]
		}
		if len(buf) > 0 {
			person := parent[buf[0]]
			entry := []string{person}
			sort.Strings(buf)
			entry = append(entry, buf...)
			output = append(output, entry)
		}
	}
	return output
}
