package audit

import (
	"github.com/mitchellh/mapstructure"
)

//func ExtractActorFromContext(ctx context.Context) (Actor, error) {
//	actorCtx := ctx.Value("user")
//	if actorCtx == nil {
//		return nil, errors.New("user doesn't exist")
//	}
//	actor, ok := actorCtx.(Actor)
//	if !ok {
//		return nil, errors.New("user is not valid")
//	}
//	return actor, nil
//}

func ExtractColumns(data map[string]any, columns []string) (result map[string]any) {
	result = make(map[string]any)
	for _, val := range columns {
		result[val] = data[val]
	}
	return
}

func UpdatedColumns(oldData Auditor, newData Auditor) (result []string, err error) {
	var (
		oldMap map[string]any
		newMap map[string]any
	)
	if err := mapstructure.Decode(oldData, &oldMap); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(newData, &newMap); err != nil {
		return nil, err
	}
	for key, newValue := range newMap {
		if oldMap[key] != newValue {
			result = append(result, key)
		}
	}
	return result, nil
}
