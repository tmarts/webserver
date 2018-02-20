package data

import (
	"testing"
)

func TestUglyXmlHasRightTestPlanName(t *testing.T) {

	names, err := getTestPlanNames("./ugly.xml")
	if err != nil {
		t.Fatalf("Error getting Test Plan Names from ./ugly.xml: " + err.Error())
	}

	// make sure xml only has one test plan
	if len(names) != 1 {
		t.Fatalf("Unexpected number of test plans in ugly.xml.  Expected: %d Actual: %d", 1, len(names))
	}

	// make sure test plan is named properly
	expectedName := "Hash (SHA-256) and Encode (Base64) variable"
	if names[0] != expectedName {
		t.Fatalf("Unexpected test plan name in ugly.xml.  Expected: \"%s\" Actual: \"%s\"", expectedName, names[0])
	}
}

func TestUglyXmlHasTestPlanEnabled(t *testing.T) {
	enabledVals, err := getTestPlanEnabledVals("./ugly.xml")
	if err != nil {
		t.Fatalf("Error getting Test Plan Names from ./ugly.xml: " + err.Error())
	}
	// make sure xml only has one test plan
	if len(enabledVals) != 1 {
		t.Fatalf("Unexpected number of test plans in ugly.xml.  Expected: %d Actual: %d", 1, len(enabledVals))
	}

	if enabledVals[0] != "true" {
		t.Fatalf("Test plan not enabled in ugly.xml.  Expected: \"%s\" Actual: \"%s\"", "true", enabledVals[0])
	}
}
