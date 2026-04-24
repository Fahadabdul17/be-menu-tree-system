package service

import (
	"be-menu-tree-system/internal/model"
	"testing"

	"github.com/google/uuid"
)

func TestBuildTree(t *testing.T) {
	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()
	id4 := uuid.New()

	menus := []model.Menu{
		{ID: id1, Name: "Root 1", ParentID: nil, Order: 1},
		{ID: id2, Name: "Child 1.1", ParentID: &id1, Order: 1},
		{ID: id3, Name: "Root 2", ParentID: nil, Order: 2},
		{ID: id4, Name: "Grandchild 1.1.1", ParentID: &id2, Order: 1},
	}

	tree := buildTree(menus)

	// Test root level
	if len(tree) != 2 {
		t.Fatalf("Expected 2 root nodes, got %d", len(tree))
	}

	// Find Root 1 in tree
	var node1Found bool
	for i := range tree {
		if tree[i].ID == id1 {
			node1Found = true
			if len(tree[i].Children) != 1 {
				t.Errorf("Root 1 should have 1 child, got %d", len(tree[i].Children))
			} else {
				child := tree[i].Children[0]
				if child.ID != id2 {
					t.Errorf("Expected child ID %s, got %s", id2, child.ID)
				}
				if len(child.Children) != 1 {
					t.Errorf("Child 1.1 should have 1 grandchild, got %d", len(child.Children))
				}
			}
		}
	}
	
	if !node1Found {
		t.Error("Root 1 not found in tree")
	}

}

func TestBuildTreeEmpty(t *testing.T) {
	tree := buildTree([]model.Menu{})
	if len(tree) != 0 {
		t.Errorf("Expected empty tree, got %d nodes", len(tree))
	}
}
