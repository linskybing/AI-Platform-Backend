package mps

import "testing"

func TestConvertGPUToMPS(t *testing.T) {
	if got := ConvertGPUToMPS(2); got != 20 {
		t.Fatalf("expected 20 MPS units, got %d", got)
	}
}

func TestConvertMPSToGPU(t *testing.T) {
	if got := ConvertMPSToGPU(15); got != 2 {
		t.Fatalf("expected 2 GPUs for 15 units, got %d", got)
	}
}

func TestValidateMPSLimit(t *testing.T) {
	if !ValidateMPSLimit(100) {
		t.Fatalf("expected limit 100 to be valid")
	}
	if ValidateMPSLimit(-1) {
		t.Fatalf("expected negative limit to be invalid")
	}
}

func TestProjectMPSQuota(t *testing.T) {
	q := &ProjectMPSQuota{TotalMPSUnits: 100, UsedMPSUnits: 40}
	if !q.CanAllocate(50) {
		t.Fatalf("expected to allocate 50 units")
	}
	if q.CanAllocate(70) {
		t.Fatalf("expected 70 allocation to fail")
	}
	if q.UsagePercent() != 40 {
		t.Fatalf("expected usage percent 40, got %f", q.UsagePercent())
	}
}
