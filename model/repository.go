package model

type Error struct {
	Err        error
	StatusCode int
}

func ShowAllSafeBoxes() (interface{}, *Error) {
	type Result struct {
		id       int
		occupied bool
	}
	var res Result

	err := DB.Raw("select SafeBox.ID, (now() > Contract.FromTime and now() < Contract.ToTime) as occupied" +
		" from Contract inner join Rent on Contract.REID = Rent.ID  " +
		"inner join SafeBox on Rent.CUId = SafeBox.CUId").Scan(&res).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return nil, &internalError
	}
	return res, &Error{}
}

func ShowAllContracts() (interface{}, *Error) {
	type Result struct {
	}

	var res Result

	err := DB.Raw("select SafeBox.ID, Rent.CUId, (now()-Contract.fromTime) as remainingTime, Contract.BaseAmount" +
		" from Contract inner join Rent on Contract.REID = Rent.ID " +
		"inner join SafeBox on Rent.CUId = SafeBox.CUId").Scan(&res).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return nil, &internalError
	}
	return res, &Error{}
}

func AddSafeBox(maxVal float64, cuId int, shId int, priceClass int) (*uint, *Error) {
	var safeBox = SafeBox{
		MaximumValue: maxVal,
		CUId:         cuId,
		SHId:         shId,
		PriceClass:   priceClass,
	}

	err := DB.Create(&safeBox).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &safeBox.ID, &internalError
	}

	return &safeBox.ID, &Error{Err: nil}
}

func EditSafeBox(id int, maxVal float64, cuId int, shId int, priceClass int) (*uint, *Error) {
	var safeBox SafeBox

	err := DB.Find(&safeBox, id).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &safeBox.ID, &internalError
	}

	safeBox.MaximumValue = maxVal
	safeBox.CUId = cuId
	safeBox.SHId = shId
	safeBox.PriceClass = priceClass

	err = DB.Save(&safeBox).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &safeBox.ID, &internalError
	}
	return &safeBox.ID, &Error{}
}

func DeleteSafeBox(id int) *Error {
	err := DB.Where("id = ?", id).Delete(&SafeBox{}).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &internalError
	}
	return &Error{Err: nil}
}

func AssignSafeBox(safeBoxId int, customerId int) (*uint, *Error) {
	var safeBox SafeBox

	err := DB.Find(&safeBox, safeBoxId).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &safeBox.ID, &internalError
	}

	safeBox.CUId = customerId

	err = DB.Save(&safeBox).Error
	if err != nil {
		internalError := Error{
			Err:        err,
			StatusCode: 500,
		}
		return &safeBox.ID, &internalError
	}
	return &safeBox.ID, &Error{}
}

func EvacuateSafeBox() {

}
