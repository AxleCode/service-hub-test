package middleware

// import (
// 	"context"
// 	"database/sql"

// 	"gitlab.com/wit-id/service-hub-test/common/constants"
// 	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
// 	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
// )

// type IAMAccessPayload struct {
// 	SiePelayananID       string
// 	DepartmentID         string
// 	DivisiID             string
// 	FamilyAltarAnggotaID string
// }

// func (v *EnsureToken) getIAMAccessToken(request IAMAccessPayload) (iamData payload.ReadIAM, err error) {
// 	q := sqlc.New(v.mainDB)

// 	var filterIAM sqlc.GetIAMAccessMddwParams
// 	var hasAccessDropdown payload.HasAccessDropdown

// 	// if request.SiePelayananID != "" {
// 	// 	filterIAM.SetPelayananID = true
// 	// 	filterIAM.SiePelayananID = request.SiePelayananID
// 	// }

// 	// // if request.PropertyID != "" {
// 	// // 	filterIAM.SetPropertyID = true
// 	// // 	filterIAM.PropertyID = sql.NullString{
// 	// // 		String: request.PropertyID,
// 	// // 		Valid:  true,
// 	// // 	}
// 	// // }

// 	// if request.DepartmentID != "" {
// 	// 	filterIAM.SetDepartmentID = true
// 	// 	filterIAM.DepartmentID = sql.NullString{
// 	// 		String: request.DepartmentID,
// 	// 		Valid:  true,
// 	// 	}
// 	// }

// 	// if request.DivisiID != "" {
// 	// 	filterIAM.SetDivisiID = true
// 	// 	filterIAM.DivisiID = sql.NullString{
// 	// 		String: request.DivisiID,
// 	// 		Valid:  true,
// 	// 	}
// 	// }

// 	filterIAM.SetDivisiID = true
// 	filterIAM.DivisiID = sql.NullString{
// 		String: request.DivisiID,
// 		Valid:  true,
// 	}
// 	filterIAM.SetDepartmentID = true
// 	filterIAM.DepartmentID = sql.NullString{
// 		String: request.DepartmentID,
// 		Valid:  true,
// 	}

// 	filterIAM.SetPelayananID = true
// 	filterIAM.SiePelayananID = request.SiePelayananID

// 	data, err := q.GetIAMAccessMddw(context.Background(), filterIAM)
// 	if err != nil {
// 		return
// 	}

// 	hasAccess, err := q.GetIAMHasAccess(context.Background(), data.Guid)
// 	if err != nil {
// 		return
// 	}

// 	// dataGroup, err := q.GetMasterdataValues(context.Background(), request.DivisiID)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	if data.SiePelayananID == constants.SuperadminConstantGuid {
// 		// TODO:: INI HARD CODE TOLONG DIPERBAIKI
// 		hasAccessDropdown = payload.HasAccessDropdown{
// 			Divisi:               true,
// 			Department:           true,
// 			SiePelayanan:         true,
// 			Scanner:              false,
// 			EventScanner:         false,
// 			IsFamliyAltarAbsence: false,
// 		}
// 		// } else if strings.ToLower(dataGroup.Value) == constants.GroupCorporate {
// 		// 	hasAccessDropdown = payload.HasAccessDropdown{
// 		// 		Divisi:       true,
// 		// 		Department:   true,
// 		// 		SiePelayanan: true,
// 		// 	}
// 	} else if data.DepartmentID.String != "" {
// 		if data.SiePelayananID == constants.SeksiPastoralFollowUpGuid {
// 			hasAccessDropdown = payload.HasAccessDropdown{
// 				Divisi:               false,
// 				Department:           false,
// 				Scanner:              true,
// 				EventScanner:         true,
// 				SiePelayanan:         false,
// 				IsFamliyAltarAbsence: false,
// 			}
// 		} else {
// 			hasAccessDropdown = payload.HasAccessDropdown{
// 				Divisi:               true,
// 				Department:           true,
// 				Scanner:              false,
// 				EventScanner:         false,
// 				SiePelayanan:         false,
// 				IsFamliyAltarAbsence: false,
// 			}
// 		}
// 	} else {
// 		hasAccessDropdown = payload.HasAccessDropdown{
// 			Divisi:               false,
// 			Department:           false,
// 			SiePelayanan:         false,
// 			Scanner:              false,
// 			EventScanner:         false,
// 			IsFamliyAltarAbsence: false,
// 		}
// 	}

// 	if request.FamilyAltarAnggotaID != "" {
// 		// if request.FamilyAltarAnggotaID == constants.PengurusFamilyAltarIDGembala ||
// 		// 	request.FamilyAltarAnggotaID == constants.PengurusFamilyAltarIDWakil ||
// 		// 	request.FamilyAltarAnggotaID == constants.PengurusFamilyAltarIDSekretaris {
// 		// 	hasAccessDropdown.IsFamliyAltarAbsence = true
// 		// }
// 		hasAccessDropdown.IsFamliyAltarAbsence = true
// 	}

// 	iamData = payload.ToPayloadIAM(sqlc.GetIAMAccessRow(data))
// 	// iamData.MenuAccess = payload.BuildMenuItems(payload.ToPayloadIAMHasAccess(hasAccess))
// 	iamData.MenuAccess = payload.ToPayloadIAMHasAccess(hasAccess, []sqlc.ListMenuRow{})
// 	iamData.HasAccessDropdown = hasAccessDropdown

// 	return
// }
