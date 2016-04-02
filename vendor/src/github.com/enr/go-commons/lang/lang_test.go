package lang

import (
	"encoding/json"
	"testing"
)

type search struct {
	haystack []string
	needle   string
	expects  bool
}

var searches = []search{
	{[]string{"Age", "Baloon"}, "Age", true},
	{[]string{"Age", "Baloon"}, "Cover", false},
	{[]string{"Age", "Baloon"}, "age", false},
	{[]string{"Age", "Baloon"}, "", false},
	{[]string{"Age", "Baloon", ""}, "", true},
}

func TestSliceContainsString(t *testing.T) {
	for _, x := range searches {
		actual := SliceContainsString(x.haystack, x.needle)
		if actual != x.expects {
			t.Errorf(`Looking for %s in %v, got %t`, x.needle, x.haystack, actual)
		}
	}
}

func TestSliceContainsString_EmptySlice(t *testing.T) {
	haystack := []string{}
	needle := ""
	actual := SliceContainsString(haystack, needle)
	if actual {
		t.Errorf(`Looking for string in empty slice, expecting false got %t`, actual)
	}
}

func TestExtractJsonFieldValue(t *testing.T) {
	jsonStr := `{"labels":[],"versions":["0.1","0.1.1","0.4","0.9"]}`
	jsonB := []byte(jsonStr)
	versions, err := ExtractJSONFieldValue(jsonB, "versions")
	if err != nil {
		t.Errorf("unexpected error thrown %s", err)
	}
	_, fs := versions.([]interface{})
	if !fs {
		t.Errorf("versions expected `[]interface{}`, got a %T", versions)
	}
	_, fi := versions.(interface{})
	if !fi {
		t.Errorf("versions expected `interface{}`, got a %T", versions)
	}
}

func TestJsonArrayToStringSlice(t *testing.T) {
	jsonStr := `{"labels":[],"versions":["0.1","0.1.1","0.4","0.9"]}`
	expectedVersions := []string{"0.1", "0.1.1", "0.4", "0.9"}
	jsonB := []byte(jsonStr)
	var b map[string]interface{}
	err := json.Unmarshal(jsonB, &b)
	if err != nil {
		t.Errorf("unexpected error thrown %s", err)
	}
	if versions, keyExists := b["versions"]; keyExists {
		versionsSlice, err := JSONArrayToStringSlice(versions, "versions")
		if err != nil {
			t.Errorf("unexpected error thrown %s", err)
		}
		if len(versionsSlice) != len(expectedVersions) {
			t.Errorf("expected versions %d but got %d", len(expectedVersions), len(versionsSlice))
		}
		for _, v := range versionsSlice {
			if !SliceContainsString(expectedVersions, v) {
				t.Errorf("version %s not expected", v)
			}
		}
	}
}
