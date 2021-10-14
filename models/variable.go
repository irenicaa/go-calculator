package models

// VariableGroup ...
type VariableGroup map[string]float64

// Copy ...
func (variables VariableGroup) Copy() VariableGroup {
	copyOfVariables := VariableGroup{}
	for name, value := range variables {
		copyOfVariables[name] = value
	}

	return copyOfVariables
}
