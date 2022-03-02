package tflog_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-log/internal/loggertest"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const testSubsystem = "test_subsystem"

func TestSubsystemWith(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		key            string
		value          interface{}
		logMessage     string
		logArgs        []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-log-pairs": {
			key:        "test-with-key",
			value:      "test-with-value",
			logMessage: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":        hclog.Trace.String(),
					"@message":      "test message",
					"@module":       testSubsystem,
					"test-with-key": "test-with-value",
				},
			},
		},
		"mismatched-log-pair": {
			key:        "test-with-key",
			value:      "test-with-value",
			logMessage: "test message",
			logArgs:    []interface{}{"unpaired-test-log-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Trace.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					"test-with-key":  "test-with-value",
					hclog.MissingKey: "unpaired-test-log-key",
				},
			},
		},
		"mismatched-with-pair": {
			key:        "unpaired-test-with-key",
			logMessage: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":                 hclog.Trace.String(),
					"@message":               "test message",
					"@module":                testSubsystem,
					"unpaired-test-with-key": nil,
				},
			},
		},
		"with-and-log-pairs": {
			key:        "test-with-key",
			value:      "test-with-value",
			logMessage: "test message",
			logArgs: []interface{}{
				"test-log-key-1", "test-log-value-1",
				"test-log-key-2", "test-log-value-2",
				"test-log-key-3", "test-log-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Trace.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					"test-log-key-1": "test-log-value-1",
					"test-log-key-2": "test-log-value-2",
					"test-log-key-3": "test-log-value-3",
					"test-with-key":  "test-with-value",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)
			ctx = tflog.SubsystemWith(ctx, testSubsystem, testCase.key, testCase.value)

			tflog.SubsystemTrace(ctx, testSubsystem, testCase.logMessage, testCase.logArgs...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}

func TestSubsystemTrace(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		message        string
		args           []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-pairs": {
			message: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":   hclog.Trace.String(),
					"@message": "test message",
					"@module":  testSubsystem,
				},
			},
		},
		"mismatched-pair": {
			message: "test message",
			args:    []interface{}{"unpaired-test-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Trace.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					hclog.MissingKey: "unpaired-test-key",
				},
			},
		},
		"pairs": {
			message: "test message",
			args: []interface{}{
				"test-key-1", "test-value-1",
				"test-key-2", "test-value-2",
				"test-key-3", "test-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":     hclog.Trace.String(),
					"@message":   "test message",
					"@module":    testSubsystem,
					"test-key-1": "test-value-1",
					"test-key-2": "test-value-2",
					"test-key-3": "test-value-3",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)

			tflog.SubsystemTrace(ctx, testSubsystem, testCase.message, testCase.args...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}

func TestSubsystemDebug(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		message        string
		args           []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-pairs": {
			message: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":   hclog.Debug.String(),
					"@message": "test message",
					"@module":  testSubsystem,
				},
			},
		},
		"mismatched-pair": {
			message: "test message",
			args:    []interface{}{"unpaired-test-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Debug.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					hclog.MissingKey: "unpaired-test-key",
				},
			},
		},
		"pairs": {
			message: "test message",
			args: []interface{}{
				"test-key-1", "test-value-1",
				"test-key-2", "test-value-2",
				"test-key-3", "test-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":     hclog.Debug.String(),
					"@message":   "test message",
					"@module":    testSubsystem,
					"test-key-1": "test-value-1",
					"test-key-2": "test-value-2",
					"test-key-3": "test-value-3",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)

			tflog.SubsystemDebug(ctx, testSubsystem, testCase.message, testCase.args...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}

func TestSubsystemInfo(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		message        string
		args           []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-pairs": {
			message: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":   hclog.Info.String(),
					"@message": "test message",
					"@module":  testSubsystem,
				},
			},
		},
		"mismatched-pair": {
			message: "test message",
			args:    []interface{}{"unpaired-test-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Info.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					hclog.MissingKey: "unpaired-test-key",
				},
			},
		},
		"pairs": {
			message: "test message",
			args: []interface{}{
				"test-key-1", "test-value-1",
				"test-key-2", "test-value-2",
				"test-key-3", "test-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":     hclog.Info.String(),
					"@message":   "test message",
					"@module":    testSubsystem,
					"test-key-1": "test-value-1",
					"test-key-2": "test-value-2",
					"test-key-3": "test-value-3",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)

			tflog.SubsystemInfo(ctx, testSubsystem, testCase.message, testCase.args...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}

func TestSubsystemWarn(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		message        string
		args           []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-pairs": {
			message: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":   hclog.Warn.String(),
					"@message": "test message",
					"@module":  testSubsystem,
				},
			},
		},
		"mismatched-pair": {
			message: "test message",
			args:    []interface{}{"unpaired-test-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Warn.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					hclog.MissingKey: "unpaired-test-key",
				},
			},
		},
		"pairs": {
			message: "test message",
			args: []interface{}{
				"test-key-1", "test-value-1",
				"test-key-2", "test-value-2",
				"test-key-3", "test-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":     hclog.Warn.String(),
					"@message":   "test message",
					"@module":    testSubsystem,
					"test-key-1": "test-value-1",
					"test-key-2": "test-value-2",
					"test-key-3": "test-value-3",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)

			tflog.SubsystemWarn(ctx, testSubsystem, testCase.message, testCase.args...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}

func TestSubsystemError(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		message        string
		args           []interface{}
		expectedOutput []map[string]interface{}
	}{
		"no-pairs": {
			message: "test message",
			expectedOutput: []map[string]interface{}{
				{
					"@level":   hclog.Error.String(),
					"@message": "test message",
					"@module":  testSubsystem,
				},
			},
		},
		"mismatched-pair": {
			message: "test message",
			args:    []interface{}{"unpaired-test-key"},
			expectedOutput: []map[string]interface{}{
				{
					"@level":         hclog.Error.String(),
					"@message":       "test message",
					"@module":        testSubsystem,
					hclog.MissingKey: "unpaired-test-key",
				},
			},
		},
		"pairs": {
			message: "test message",
			args: []interface{}{
				"test-key-1", "test-value-1",
				"test-key-2", "test-value-2",
				"test-key-3", "test-value-3",
			},
			expectedOutput: []map[string]interface{}{
				{
					"@level":     hclog.Error.String(),
					"@message":   "test message",
					"@module":    testSubsystem,
					"test-key-1": "test-value-1",
					"test-key-2": "test-value-2",
					"test-key-3": "test-value-3",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var outputBuffer bytes.Buffer

			ctx := context.Background()
			ctx = loggertest.ProviderRoot(ctx, &outputBuffer)
			ctx = tflog.NewSubsystem(ctx, testSubsystem)

			tflog.SubsystemError(ctx, testSubsystem, testCase.message, testCase.args...)

			got, err := loggertest.MultilineJSONDecode(&outputBuffer)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedOutput, got); diff != "" {
				t.Errorf("unexpected output difference: %s", diff)
			}
		})
	}
}
