package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type link struct {
	dest  string
	count int
}

// count only matters for traversing DOWN the tree for part 2, while only ID is needed for traversing UP the tree
// that's why outgoing is a link type, and incoming is a string type
type graph_node struct {
	id       string
	incoming []string
	outgoing []link
}

func sum_bags(node graph_node, graph map[string]graph_node) int {
	if len(node.outgoing) == 0 {
		return 0
	} else {
		total := 0
		for i := 0; i < len(node.outgoing); i++ {
			new_node_link := node.outgoing[i]
			total += new_node_link.count                                                // add direct descendent bags
			total += (new_node_link.count * sum_bags(graph[new_node_link.dest], graph)) // then add the bags inside those bags
		}
		return total
	}
}

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	line_regex, err := regexp.Compile(`^(?P<root_id>[a-z]+ [a-z]+) bags contain(?P<target_ids>(?: [0-9] [a-z]+ [a-z]+ bag[s]?[\.,])+|(?: no other bags\.))$`)
	check(err)
	target_id_regex, err := regexp.Compile(`((?P<colour_count>[0-9]) (?P<colour_id>[a-z]+ [a-z]+) bag[s]?)+`)
	check(err)

	graph_node_map := map[string]graph_node{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := line_regex.FindStringSubmatch(line)

		root_id := strings.Replace(matches[line_regex.SubexpIndex("root_id")], " ", "_", 1)
		target_ids := strings.Trim(matches[line_regex.SubexpIndex("target_ids")], " .")

		if target_ids != "no other bags" {
			matches := target_id_regex.FindAllStringSubmatch(target_ids, -1)
			for i := 0; i < len(matches); i++ {
				target_colour_id := strings.Replace(matches[i][target_id_regex.SubexpIndex("colour_id")], " ", "_", 1)
				target_colour_count, err := strconv.Atoi(matches[i][target_id_regex.SubexpIndex("colour_count")])
				check(err)

				// build root -> (outgoing) -> target
				node, exists := graph_node_map[root_id]
				if exists {
					node.outgoing = append(node.outgoing, link{dest: target_colour_id, count: target_colour_count})
					graph_node_map[root_id] = node
				} else {
					new_graph_node := graph_node{id: root_id, incoming: []string{}, outgoing: []link{{dest: target_colour_id, count: target_colour_count}}}
					graph_node_map[root_id] = new_graph_node
				}

				// build target -> (incoming) -> root
				node, exists = graph_node_map[target_colour_id]
				if exists {
					node.incoming = append(node.incoming, root_id)
					graph_node_map[target_colour_id] = node
				} else {
					new_graph_node := graph_node{id: target_colour_id, incoming: []string{root_id}, outgoing: []link{}}
					graph_node_map[target_colour_id] = new_graph_node
				}
			}
		}
	}

	// built the tree. Get shiny_gold, then BFS up the incoming of every node found

	// using map keys as a set
	visited_nodes := map[string]bool{}

	// keep pulling values from this list until it is empty
	nodes_to_check := []string{"shiny_gold"}

	for len(nodes_to_check) != 0 {
		// get the node
		graph_node_to_check := graph_node_map[nodes_to_check[0]]
		// remove this element from the node to check list
		nodes_to_check = nodes_to_check[1:]
		// add all the incoming nodes to the check list
		nodes_to_check = append(nodes_to_check, graph_node_to_check.incoming...)
		// add the node we just looked at to the visited nodes
		visited_nodes[graph_node_to_check.id] = true
	}

	// part 1 answer. exclude shiny_gold from the count
	fmt.Println("Total bags that can possibly contain a shiny gold:", len(visited_nodes)-1)

	// part 2 answer. start at shiny_gold and recursively iterate over all of a nodes outgoings.
	result := sum_bags(graph_node_map["shiny_gold"], graph_node_map)
	fmt.Println("Total bags contained within a shiny gold bag:", result)
}
