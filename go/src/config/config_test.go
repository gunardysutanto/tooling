package config

import "testing"

func TestConfig(t *testing.T) {
	expectedResult := "localhost:3306/testdb"
	actualResult := GetConfigurationFrom("/home/gunardysutanto/workspaces/tooling/go/samples/test.properties")["db.url"]
	if actualResult != expectedResult {
		t.Errorf("GetConfigurationFrom == %q, want %q", actualResult, expectedResult)
	}
}
