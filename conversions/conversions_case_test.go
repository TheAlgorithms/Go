package conversions

type romanToIntegerConversionTest struct {
	input    string
	expected int
	name     string
}

var romanToIntegerTests = []romanToIntegerConversionTest{
	{input: "DCCLXXXIX", expected: 789, name: "DCCLXXXIX-789"},
	{input: "MLXVI", expected: 1066, name: "MLXVI-1066"},
	{input: "MCMXVIII", expected: 1918, name: "MCMXVIII-1918"},
	{input: "V", expected: 5, name: "V-5"},
}
