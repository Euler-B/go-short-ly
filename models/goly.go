package model

func GetAllGolies() ([]Goly, error) {
	var golies []Goly
	tx := db.Find(&golies)
	if tx.Error != nil {
		return []Goly{}, tx.Error
	}
	return golies, nil
}

func GetGoly(id uint64) (Goly, error) {
	var goly Goly
	tx := db.Where("id = ?", id).First(&goly)
	if tx.Error != nil {
		return Goly{}, tx.Error
	}

	return goly, nil
}

func CreateGoly(goly Goly) error {
	tx := db.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly Goly) error {
	tx := db.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uint64) error {
	tx := db.Unscoped().Delete(&Goly{}, id) 
	return tx.Error 
	// TODO 1: Investigar en la documentacion de GORM: 1.- Que funcion cumpple el metodo Unscoped()     2.- Cuales son los otros metodos que puedo utilizar para realiar una consulta, al menos los mas usados.
}

func FindByGolyUrl(url string) (Goly, error) {
	var goly Goly
	tx := db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}