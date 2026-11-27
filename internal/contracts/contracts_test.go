package contracts

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Go runs tests with the package directory as the working directory.
const schemasDir = "../../schemas"

// schemaFor maps a fixture's filename prefix to the schema it must satisfy.
func schemaFor(t *testing.T, fixture string) string {
	t.Helper()
	switch {
	case strings.HasPrefix(fixture, "inventory-"):
		return "inventory.schema.json"
	case strings.HasPrefix(fixture, "findings-"):
		return "findings.schema.json"
	default:
		t.Fatalf("fixture %q has no known schema prefix", fixture)
		return ""
	}
}

func runFixtures(t *testing.T, subdir string, wantValid bool) {
	t.Helper()

	dir := filepath.Join(schemasDir, "fixtures", subdir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read fixtures dir: %v", err)
	}
	// Guard against a silent pass: zero fixtures would "succeed" trivially.
	if len(entries) == 0 {
		t.Fatalf("no fixtures found in %s", dir)
	}

	for _, entry := range entries {
		name := entry.Name()
		t.Run(name, func(t *testing.T) {
			// Compile errors are test failures, never "expected rejections".
			schema, err := jsonschema.NewCompiler().
				Compile(filepath.Join(schemasDir, schemaFor(t, name)))
			if err != nil {
				t.Fatalf("schema does not compile: %v", err)
			}

			data, err := os.ReadFile(filepath.Join(dir, name))
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}
			instance, err := jsonschema.UnmarshalJSON(bytes.NewReader(data))
			if err != nil {
				t.Fatalf("fixture is not valid JSON: %v", err)
			}

			err = schema.Validate(instance)
			if wantValid && err != nil {
				t.Errorf("expected valid, got: %v", err)
			}
			if !wantValid && err == nil {
				t.Errorf("expected rejection, but the fixture validated")
			}
		})
	}
}

func TestValidFixturesPass(t *testing.T)          { runFixtures(t, "valid", true) }
func TestInvalidFixturesAreRejected(t *testing.T) { runFixtures(t, "invalid", false) }
