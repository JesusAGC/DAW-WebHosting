package database

import (
	models "BackendOrdinario/models"
	"database/sql"
	"errors"
	"fmt"
)

func DbConnection() (db *sql.DB, e error) {
	dbDriver := "mysql"
	dbUser := "user1"
	dbPass := "1234"
	dbName := "WebHosting"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func InsertExSuggestion(s models.ExternSuggestion) error {

	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	prepareQuery, err := db.Prepare("INSERT INTO ExSuggestions(Email,Suggestion) VALUES('?', '?');")
	if err != nil {
		return err
	}

	defer prepareQuery.Close()

	// ejecutar sentencia, un valor por cada '?'
	_, err = prepareQuery.Exec(s.Email, s.Suggestion)
	if err != nil {
		return err
	}

	fmt.Println("Se insertó usuario correctamente")
	return nil
}
func InsertInSuggestion(s models.InternSuggestion) error {

	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	prepareQuery, err := db.Prepare("INSERT INTO Insuggestions(UserID,Suggestion) VALUES('?', '?');")
	if err != nil {
		return err
	}

	defer prepareQuery.Close()

	// ejecutar sentencia, un valor por cada '?'
	_, err = prepareQuery.Exec(s.UserID, s.Suggestion)
	if err != nil {
		return err
	}

	fmt.Println("Se insertó usuario correctamente")
	return nil
}

func GetAllInSuggestions() ([]models.InternSuggestion, error) {

	suggs := []models.InternSuggestion{}
	db, err := DbConnection()
	if err != nil {
		fmt.Println("No se pudo conectar")
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, UserID, Suggestion FROM InSuggestions;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var s models.InternSuggestion

	for rows.Next() {
		err = rows.Scan(&s.ID, &s.UserID, &s.Suggestion)
		if err != nil {
			return nil, err
		}

		suggs = append(suggs, s)
	}

	return suggs, nil
}

func GetAllExSuggestions() ([]models.ExternSuggestion, error) {

	suggs := []models.ExternSuggestion{}
	db, err := DbConnection()
	if err != nil {
		fmt.Println("No se pudo conectar")
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Email, Suggestion FROM WebHosting.ExSuggestions;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var s models.ExternSuggestion

	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Name, &s.Email, &s.Suggestion)
		if err != nil {
			return nil, err
		}

		suggs = append(suggs, s)
	}

	return suggs, nil
}

func DeleteExSuggestion(ID string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	prepareQuery, err := db.Prepare("DELETE FROM ExSuggestions WHERE ID = ?")
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	res, err2 := prepareQuery.Exec(ID)
	if err2 != nil {
		return err2
	}

	rows, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	}

	if rows == 0 {
		mistake := errors.New("no rows affected")
		return mistake
	}

	fmt.Println("Se ha eliminado exitosamente la sugerencia: ", ID)
	return nil
}

func DeleteInSuggestion(ID string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	prepareQuery, err := db.Prepare("DELETE FROM Insuggestions WHERE ID = ?")
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	res, err2 := prepareQuery.Exec(ID)
	if err2 != nil {
		return err2
	}

	rows, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	}

	if rows == 0 {
		mistake := errors.New("no rows affected")
		return mistake
	}

	fmt.Println("Se ha eliminado exitosamente la sugerencia: ", ID)
	return nil
}

func NewUser(u models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	prepareQuery, err := db.Prepare("INSERT INTO Users(Nombre, Apellidos, UserName, Email, Password) VALUES(?,?,?,?,?);")
	if err != nil {
		return err
	}

	defer prepareQuery.Close()

	// ejecutar sentencia, un valor por cada '?'
	_, err = prepareQuery.Exec(u.Nombre, u.Apellidos, u.UserName, u.Email, u.Password)
	if err != nil {
		return err
	}

	fmt.Println("Se insertó usuario correctamente")
	return nil
}

func GetUser(id string) (models.User, error) {
	var u models.User

	db, err := DbConnection()
	if err != nil {
		return u, err
	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT Nombre, Apellidos, UserName, Email, Password FROM Users WHERE ID = '%s'`, id)

	rows, err := db.Query(query)
	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Nombre, &u.Apellidos, &u.UserName, &u.Email, &u.Password)
		if err != nil {
			return u, err
		}
	}

	return u, nil
}

func GetUserByMail(mail string) (models.User, error) {
	var u models.User

	db, err := DbConnection()
	if err != nil {
		return u, err
	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT ID, Nombre, Apellidos, UserName, Email FROM Users WHERE Email = '%s'`, mail)

	rows, err := db.Query(query)
	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Nombre, &u.Apellidos, &u.UserName, &u.Email, &u.Password)
		if err != nil {
			return u, err
		}
	}

	return u, nil
}

func UpdateUser(u models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE Users SET Nombre='?', Apellidos='?', UserName='?', Email='?', Password='?' WHERE ID='%s'", u.ID)
	prepareQuery, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	_, err = prepareQuery.Exec(&u.Nombre, &u.Apellidos, &u.UserName, &u.Email, &u.Password)
	if err != nil {
		return err
	}
	fmt.Println("Se modificó correctamente")

	return nil

}

func DeleteUser(ID string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	prepareQuery, err := db.Prepare("DELETE FROM Users WHERE ID = ?")
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	res, err2 := prepareQuery.Exec(ID)
	if err2 != nil {
		return err2
	}

	rows, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	}

	if rows == 0 {
		mistake := errors.New("no rows affected")
		return mistake
	}

	fmt.Println("Se ha eliminado exitosamente al usuario con ID: ", ID)
	return nil
}

func GetAllContractsFromDB() ([]models.Contract, error) {

	conts := []models.Contract{}
	db, err := DbConnection()
	if err != nil {
		fmt.Println("No se pudo conectar")
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, PlanID, UserID, DateOfContract, DateOfExpiration FROM Contracts;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var c models.Contract

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.PlanID, &c.UserID, &c.DateOfContract, &c.DateOfExpiration)
		if err != nil {
			return nil, err
		}

		conts = append(conts, c)
	}

	return conts, nil
}

func InsertContract(c models.Contract) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	prepareQuery, err := db.Prepare("INSERT INTO Contracts (PlanID, UserID, DateOfContract, DateOfExpiration) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}

	defer prepareQuery.Close()

	// ejecutar sentencia, un valor por cada '?'
	_, err = prepareQuery.Exec(c.PlanID, c.UserID, c.DateOfContract.Format("2006-01-02 15:04:05"), c.DateOfExpiration.Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	fmt.Println("Se insertó contrato correctamente")
	return nil
}

func ReadContract(ID string) (models.Contract, error) {
	var c models.Contract

	db, err := DbConnection()
	if err != nil {
		return c, err
	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT ID, PlanID, UserID, DateOfContract, DateOfExpiration FROM Contracts WHERE ID = %s`, ID)

	rows, err := db.Query(query)
	if err != nil {
		return c, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.PlanID, &c.UserID, &c.DateOfContract, &c.DateOfExpiration)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}

func DeleteContract(ID string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	prepareQuery, err := db.Prepare("DELETE FROM Contracts WHERE ID = ?")
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	res, err2 := prepareQuery.Exec(ID)
	if err2 != nil {
		return err2
	}

	rows, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	}

	if rows == 0 {
		mistake := errors.New("no rows affected")
		return mistake
	}

	fmt.Println("Se ha eliminado exitosamente el contrato con ID: ", ID)
	return nil
}

func UpdateContract(c models.Contract) error {

	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE Contracts SET PlanID=?, UserID=?, DateOfContract= ?, DateOfExpiration= ? WHERE ID= %s", c.ID)

	prepareQuery, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer prepareQuery.Close()

	_, err = prepareQuery.Exec(&c.PlanID, &c.UserID, c.DateOfContract.Format("2006-01-02 15:04:05"), c.DateOfExpiration.Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	fmt.Println("Se modificó correctamente")

	return nil

}

func GetLoginUserData(un string) (models.Login, error) {

	var l models.Login

	db, err := DbConnection()
	if err != nil {
		return l, err
	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT UserName, Password FROM Users WHERE UserName = '%s'`, un)

	rows, err := db.Query(query)
	if err != nil {
		return l, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&l.UserName, &l.Password)
		if err != nil {
			return l, err
		}
	}

	return l, nil
}

func GetLoginAdminData(UserName string) (models.Login, error) {

	var l models.Login

	db, err := DbConnection()
	if err != nil {
		return l, err
	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT UserName, Password FROM Admins WHERE UserName = '%s'`, UserName)

	rows, err := db.Query(query)
	if err != nil {
		return l, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&l.UserName, &l.Password)
		if err != nil {
			return l, err
		}
	}

	return l, nil
}
