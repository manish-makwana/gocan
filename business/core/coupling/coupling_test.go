package coupling

import (
	"com.fha.gocan/business/data/store/coupling"
	"com.fha.gocan/business/data/store/stat"
	"reflect"
	"testing"
)

func TestCouplingForTwoFiles(t *testing.T) {
	stats := []stat.Stat{
		{CommitId: "123", File: "file1"},
		{CommitId: "123", File: "file2"},
		{CommitId: "456", File: "file1"},
	}

	want := []coupling.Coupling{
		{
			Entity:           "file1",
			Coupled:          "file2",
			Degree:           0.6666666666666666,
			AverageRevisions: 1.5,
		},
	}
	got := CalculateCouplings(stats, 0, 0)

	assertEqual(t, want, got)
}

func TestCouplingForManyFiles(t *testing.T) {
	stats := []stat.Stat{
		{CommitId: "123", File: "file1"},
		{CommitId: "123", File: "file2"},

		{CommitId: "456", File: "file1"},
		{CommitId: "456", File: "file3"},
		{CommitId: "456", File: "file4"},

		{CommitId: "789", File: "file4"},
		{CommitId: "789", File: "file2"},
		{CommitId: "789", File: "file1"},

		{CommitId: "876", File: "file4"},
		{CommitId: "876", File: "file1"},
		{CommitId: "876", File: "file2"},

		{CommitId: "888", File: "file4"},
		{CommitId: "888", File: "file3"},
	}

	want := []coupling.Coupling{
		{
			Entity:           "file1",
			Coupled:          "file2",
			Degree:           0.8571428571428571,
			AverageRevisions: 3.5,
		},
		{
			Entity:           "file1",
			Coupled:          "file4",
			Degree:           0.75,
			AverageRevisions: 4,
		},
		{
			Entity:           "file3",
			Coupled:          "file4",
			Degree:           0.6666666666666666,
			AverageRevisions: 3,
		},
		{
			Entity:           "file4",
			Coupled:          "file2",
			Degree:           0.5714285714285714,
			AverageRevisions: 3.5,
		},
		{
			Entity:           "file3",
			Coupled:          "file1",
			Degree:           0.3333333333333333,
			AverageRevisions: 3,
		},
	}
	got := CalculateCouplings(stats, 0, 0)

	assertEqual(t, want, got)
}

func assertEqual(t *testing.T, want []coupling.Coupling, got []coupling.Coupling) {
	if len(want) != len(got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	for i, actual := range got {
		expected := want[i]
		if actual.Entity == expected.Entity {
			if actual.Coupled != expected.Coupled ||
				actual.Degree != expected.Degree ||
				actual.AverageRevisions != expected.AverageRevisions {
				t.Errorf("Wanted %v, got %v", expected, actual)
			}
		} else if actual.Entity == expected.Coupled {
			if actual.Coupled != expected.Entity ||
				actual.Degree != expected.Degree ||
				actual.AverageRevisions != expected.AverageRevisions {
				t.Errorf("Wanted %v, got %v", expected, actual)
			}
		} else {
			t.Errorf("Wanted %v, got %v", expected, actual)
		}
	}
}

func isEqual(aa, bb []coupling.Coupling) bool {
	eqCtr := 0
	for _, a := range aa {
		for _, b := range bb {
			if reflect.DeepEqual(a, b) {
				eqCtr++
			}
		}
	}
	if eqCtr != len(bb) || len(aa) != len(bb) {
		return false
	}
	return true
}
