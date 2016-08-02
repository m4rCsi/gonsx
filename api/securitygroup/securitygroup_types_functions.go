package securitygroup

import (
	"fmt"
)

func (sgl List) String() string {
	return fmt.Sprintf("SecurityGroupList object, contains SecurityGroup objects.")
}

func (sg SecurityGroup) String() string {
	return fmt.Sprintf("objectId: %-20s name: %-20s.", sg.ObjectID, sg.Name)
}

// FilterByName returns a single securitygroup object if it matches the name in SecurityGroup
func (sgl List) FilterByName(name string) *SecurityGroup {
	var securityGroupFound SecurityGroup
	for _, sg := range sgl.SecurityGroups {
		if sg.Name == name {
			securityGroupFound = sg
			break
		}
	}
	return &securityGroupFound
}

// AddDynamicMemberDefinitionSet adds new DynamicSet to DynamicMemberDefinition of SecurityGroup object.
func (sg *SecurityGroup) AddDynamicMemberDefinitionSet(operator string, dynamicCriteriaList []DynamicCriteria) {
	newDynamicSet := DynamicSet{Operator: operator, DynamicCriteria: dynamicCriteriaList}
	if sg.DynamicMemberDefinition != nil && len(sg.DynamicMemberDefinition.DynamicSet) >= 1 {
		sg.DynamicMemberDefinition.DynamicSet = append(sg.DynamicMemberDefinition.DynamicSet, newDynamicSet)
	} else {
		dynamicSetList := []DynamicSet{newDynamicSet}
		sg.DynamicMemberDefinition = &DynamicMemberDefinition{DynamicSet: dynamicSetList}
	}

}
