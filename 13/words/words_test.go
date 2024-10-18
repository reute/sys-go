package main

import (
	"strings"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	filename := "states.txt"
	unsorted := readFromFile(filename)
	expectedCities := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "Colorado", "California",
		"Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana",
		"Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada",
		"New Hampshire", "New Jersey", "New Mexico", "New York",
		"North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota",
		"Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington",
		"West Virginia", "Wisconsin", "Wyoming",
	}
	node := unsorted
	for i, expectedCity := range expectedCities {
		if node == nil {
			t.Errorf("Expected more nodes in the linked list, but got nil at index %d", i)
			break
		}
		actualCity := strings.TrimSpace(node.city)
		if actualCity != expectedCity {
			t.Errorf("Expected city '%s', but got '%s' at index %d", expectedCity, actualCity, i)
		}
		node = node.next
	}
	if node != nil {
		t.Errorf("Expected fewer nodes in the linked list, but got more")
	}
}

func TestFitsFirstChar(t *testing.T) {
	tests := []struct {
		cityNew    string
		citySorted string
		expected   bool
	}{
		{"Seattle", "Eubanks", true},
		{"Los Angeles", "New York", false},
	}
	for _, test := range tests {
		actual := fitsFirstChar(test.cityNew, test.citySorted)
		if actual != test.expected {
			t.Errorf("For cityNew='%s' and citySorted='%s', expected %v but got %v", test.cityNew, test.citySorted, test.expected, actual)
		}
	}
}

func TestFitsLastChar(t *testing.T) {
	tests := []struct {
		cityNew    string
		citySorted string
		expected   bool
	}{
		{"Seattle", "Eubanks", true},
		{"Los Angeles", "New York", false},
	}
	for _, test := range tests {
		actual := fitsLastChar(test.cityNew, test.citySorted)
		if actual != test.expected {
			t.Errorf("For cityNew='%s' and citySorted='%s', expected %v but got %v", test.cityNew, test.citySorted, test.expected, actual)
		}
	}
}

func TestCheckFit(t *testing.T) {
	tests := []struct {
		cityNew    string
		citySorted *Node
		expected   int
	}{
		{"Seattle", nil, begin},
		{"Los Angeles", &Node{city: "New York"}, nofit},
		{"Seattle", &Node{city: "Eubanks"}, begin},
		{"San Francisco", &Node{city: "Dallas"}, end},
	}
	for _, test := range tests {
		actual := checkFit(test.cityNew, test.citySorted)
		if actual != test.expected {
			t.Errorf("For city='%s' and sorted='%v', expected %v but got %v", test.cityNew, test.citySorted, test.expected, actual)
		}
	}
}

func TestFindCandidates(t *testing.T) {
	equal := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	sorted := &Node{city: "San Francisco"}
	unsorted := &Node{city: "Seattle", next: &Node{city: "Los Angeles", next: &Node{city: "Oregon"}}}
	tests := []struct {
		sorted   *Node
		unsorted *Node
		expected []string
	}{
		{sorted, nil, nil},
		{sorted, unsorted, []string{"Los Angeles", "Oregon"}},
	}
	for _, test := range tests {
		actual := findCandidates(test.unsorted, test.sorted)
		if !equal(actual, test.expected) {
			t.Errorf("For sorted='%v' and unsorted='%v', expected %v but got %v", test.sorted, test.unsorted, test.expected, actual)
		}
	}
}

func TestSearchCity(t *testing.T) {
	unsorted := &Node{
		city: "Seattle",
		next: &Node{
			city: "New York",
			next: nil,
		},
	}
	actual := searchCity("Seattle", unsorted)
	if actual == nil || actual.city != "Seattle" {
		t.Errorf("Expected 'Seattle', got %v", actual)
	}
	actual = searchCity("Los Angeles", unsorted)
	if actual != nil {
		t.Errorf("Expected nil, got %v", actual)
	}
}

func TestGetLastNode(t *testing.T) {
	list := &Node{
		city: "Seattle",
		next: &Node{
			city: "New York",
			next: nil,
		},
	}
	lastNode := list.getLastNode()
	expected := list.next
	if lastNode != expected {
		t.Errorf("Expected last node value %s, but got %s", expected.city, lastNode.city)
	}
}
