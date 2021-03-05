package graph_playground

import "testing"

func TestExample(t *testing.T) {
	// Scenario:
	// A camera is watching a person who is wearing a hat.
	// Camera "watches" Person
	// Person is "watched by" Camera
	// Person "wears" Hat
	// Hat is "worn by" Person
	// Create an empty graph
	g := InitializeGraph()
	// Add Camera
	cameraId := g.AddVertice(InitializeVertice("camera"))
	// Add Person
	personId := g.AddVertice(InitializeVertice("person"))
	// Add Hat
	hatId := g.AddVertice(InitializeVertice("hat"))
	g.AddEdge(cameraId, personId, "watches", "watched by")
	g.AddEdge(personId, hatId, "wears", "worn by")

	// Traverse the relationships forwards

	camera := g.Vertices[cameraId]
	watched := camera.Edges["watches"].Vertice
	if watched.Name != "person" {
		t.Error("Camera is not watching person")
	}
	worn := watched.Edges["wears"].Vertice
	if worn.Name != "hat" {
		t.Error("Person is not wearing hat")
	}
}
