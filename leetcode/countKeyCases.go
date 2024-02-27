package leetcode

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 9fa4e2d Add testcases to matching-brackets (#1657)

type test_case struct {
	description string
	input       string
	expected    int
}

var testCases = []test_case{
	{
		description: "Special Case",
		input:       "RiQYo",
		expected:    4,
	},
	// {
	// 	description: "empty string",
	// 	input:       "",
	// 	expected:    true,
	// },
	// {
	// 	description: "unpaired brackets",
	// 	input:       "[[",
	// 	expected:    false,
	// },
	// {
	// 	description: "wrong ordered brackets",
	// 	input:       "}{",
	// 	expected:    false,
	// },
	// {
	// 	description: "wrong closing bracket",
	// 	input:       "{]",
	// 	expected:    false,
	// },
	// {
	// 	description: "paired with whitespace",
	// 	input:       "{ }",
	// 	expected:    true,
	// },
	// {
	// 	description: "partially paired brackets",
	// 	input:       "{[])",
	// 	expected:    false,
	// },
	// {
	// 	description: "simple nested brackets",
	// 	input:       "{[]}",
	// 	expected:    true,
	// },
	// {
	// 	description: "several paired brackets",
	// 	input:       "{}[]",
	// 	expected:    true,
	// },
	// {
	// 	description: "paired and nested brackets",
	// 	input:       "([{}({}[])])",
	// 	expected:    true,
	// },
	// {
	// 	description: "unopened closing brackets",
	// 	input:       "{[)][]}",
	// 	expected:    false,
	// },
	// {
	// 	description: "unpaired and nested brackets",
	// 	input:       "([{])",
	// 	expected:    false,
	// },
	// {
	// 	description: "paired and wrong nested brackets",
	// 	input:       "[({]})",
	// 	expected:    false,
	// },
	// {
	// 	description: "paired and incomplete brackets",
	// 	input:       "{}[",
	// 	expected:    false,
	// },
	// {
	// 	description: "too many closing brackets",
	// 	input:       "[]]",
	// 	expected:    false,
	// },
	// {
	// 	description: "early unexpected brackets",
	// 	input:       ")()",
	// 	expected:    false,
	// },
	// {
	// 	description: "early mismatched brackets",
	// 	input:       "{)()",
	// 	expected:    false,
	// },
	// {
	// 	description: "math expression",
	// 	input:       "(((185 + 223.85) * 15) - 543)/2",
	// 	expected:    true,
	// },
	// {
	// 	description: "complex latex expression",
	// 	input:       "\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)",
	// 	expected:    true,
	// },
}
