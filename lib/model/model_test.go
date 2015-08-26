package model

import (
	"testing"

	"github.com/syncthing/protocol"
)

func TestModelSingleIndex(t *testing.T) {
	// Arrange
	model := NewModel()

	deviceID := protocol.DeviceID{}
	folder := "syncthingfusetest"
	flags := uint32(0)
	options := []protocol.Option{}

	files := []protocol.FileInfo{
		protocol.FileInfo{Name: "file1"},
		protocol.FileInfo{Name: "file2"},
		protocol.FileInfo{Name: "dir1", Flags: protocol.FlagDirectory},
		protocol.FileInfo{Name: "dir1/dirfile1"},
		protocol.FileInfo{Name: "dir1/dirfile2"},
	}

	// Act
	model.Index(deviceID, folder, files, flags, options)

	// Assert
	children := model.GetChildren(folder, ".")
	assertContainsChild(t, children, "file2", 0)
	assertContainsChild(t, children, "file2", 0)
	assertContainsChild(t, children, "dir1", protocol.FlagDirectory)
	if len(children) != 3 {
		t.Error("expected 3 children, but got", len(children))
	}

	children = model.GetChildren(folder, "dir1")
	assertContainsChild(t, children, "dir1/dirfile1", 0)
	assertContainsChild(t, children, "dir1/dirfile2", 0)
	if len(children) != 2 {
		t.Error("expected 2 children, but got", len(children))
	}

	assertEntry(t, model.GetEntry(folder, "file1"), "file1", 0)
	assertEntry(t, model.GetEntry(folder, "file2"), "file2", 0)
	assertEntry(t, model.GetEntry(folder, "dir1"), "dir1", protocol.FlagDirectory)
	assertEntry(t, model.GetEntry(folder, "dir1/dirfile1"), "dir1/dirfile1", 0)
	assertEntry(t, model.GetEntry(folder, "dir1/dirfile2"), "dir1/dirfile2", 0)
}

func assertContainsChild(t *testing.T, children []protocol.FileInfo, name string, flags uint32) {
	for _, child := range children {
		if child.Name == name && child.Flags == flags {
			return
		}
	}

	t.Error("Missing file", name)
}

func assertEntry(t *testing.T, entry protocol.FileInfo, name string, flags uint32) {
	if entry.Name == name && entry.Flags == flags {
		return
	}

	t.Error("incorrect entry for file", name)
}
