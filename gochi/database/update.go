package database

// ===== For reservations =====
func UpdateReservationStatusToFalse(uid string) error {
	stmt, err := db.Prepare("UPDATE reservation SET status=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(false, uid)

	if err != nil {
		return err
	}

	return nil
}

// ===== For hair saloons =====
func UpdateSaloonName(uid string, newName string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET name=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newName, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaloonAddress(uid string, newAddress string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET address=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newAddress, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaloonEmail(uid string, newEmail string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET email=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newEmail, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaloonPhone(uid string, newPhone string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET phone=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newPhone, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaloonOpeningTime(uid string, newOpeningTime string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET openingTime=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newOpeningTime, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaloonClosingTime(uid string, newClosingTime string) error {
	stmt, err := db.Prepare("UPDATE hairSaloon SET closingTime=? WHERE uid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newClosingTime, uid)
	if err != nil {
		return err
	}

	return nil
}
