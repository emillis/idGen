package idGen

//Generator defines structure of a new Generator
type Generator struct {
	requirements Requirements
}

//Static will always return the same encoded string if the same key is supplied based on the Requirements
//supplied to the Generator
func (g Generator) Static(key string) string {
	return generateStaticString(key, &g.requirements)
}

//Random returns a randomly generated string based on the Requirements supplied to the Generator
func (g Generator) Random() string {
	return generateRandomString(&g.requirements)
}

//===========[FUNCTIONALITY]====================================================================================================

//NewGenerator returns newly initiated Generator
func NewGenerator(r *Requirements) Generator {
	return Generator{
		requirements: *makeRequirementsReasonable(r),
	}
}
