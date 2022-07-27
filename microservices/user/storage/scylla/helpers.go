package scylla

import (
	"fmt"
	"strings"
	"user/models"
)

func prepareUpdateStruct(data *models.User) string {
	updateData := make([]string, 0, 4)


	//if data.ID != 0 {
	//	updateData = append(updateData, fmt.Sprintf("id = %d", data.ID))
	//}
	if data.Name != "" {
		updateData = append(updateData, fmt.Sprintf("name = '%s'", data.Name))
	}
	if data.Address != "" {
		updateData = append(updateData, fmt.Sprintf("address = '%s'", data.Address))
	}
	updateData = append(updateData, fmt.Sprintf("updated_at = '%v'", data.UpdatedAt))

	str := strings.Join(updateData[:], ", ")
	return str
}
