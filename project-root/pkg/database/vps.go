package database

// AddVPS inserts a new VPS record into the database.
func AddVPS(vps VPS) (int64, error) {
	res, err := DB.Exec(`INSERT INTO vps (remark, ip, region, provider, price, total_data, reset_times, available_data, reset_date, calculation_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		vps.Remark, vps.IP, vps.Region, vps.Provider, vps.Price, vps.TotalData, vps.ResetTimes, vps.AvailableData, vps.ResetDate, vps.CalculationType)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// DeleteVPS removes a VPS record by its ID.
func DeleteVPS(vpsID int) error {
	_, err := DB.Exec(`DELETE FROM vps WHERE id = ?`, vpsID)
	return err
}

// ListVPS retrieves all VPS records.
func ListVPS() ([]VPS, error) {
	rows, err := DB.Query(`SELECT id, remark, ip, region, provider, price, total_data, reset_times, available_data, reset_date, calculation_type FROM vps`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vpsList []VPS
	for rows.Next() {
		var v VPS
		if err := rows.Scan(&v.ID, &v.Remark, &v.IP, &v.Region, &v.Provider, &v.Price, &v.TotalData, &v.ResetTimes, &v.AvailableData, &v.ResetDate, &v.CalculationType); err != nil {
			return nil, err
		}
		vpsList = append(vpsList, v)
	}
	return vpsList, nil
}
