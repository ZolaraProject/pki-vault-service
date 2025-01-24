package pkivault

import (
	. ""
	"back/logger"
	"back/utils"
	"database/sql"
	fmt "fmt"
	. "pkivaultrpc"
	"strings"
)

var (
	config utils.Config
	db     *sql.DB
)

func GetUsers(req *UserRequest) (*UserList, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	query := "SELECT id, email, username, password, role FROM users"
	totalQuery := "SELECT COUNT(*) FROM users"
	whereClause := ""
	orderClause := ""
	limitClause := ""
	userParams := []interface{}{}
	userTotalParams := []interface{}{}

	i := 1
	if req.PagingQuery != nil {
		if req.PagingQuery.Search != "" {
			if whereClause == "" {
				whereClause += " WHERE"
			} else {
				whereClause += " AND"
			}

			whereClause += fmt.Sprintf(" username ILIKE $%d", i)
			userParams = append(userParams, "%"+req.PagingQuery.Search+"%")
			userTotalParams = append(userTotalParams, "%"+req.PagingQuery.Search+"%")
			i++
		}

		if req.PagingQuery.Sort != "" {
			var order string = "ASC"
			if strings.EqualFold(req.PagingQuery.Order, "desc") {
				order = "DESC"
			}

			orderClause = fmt.Sprintf(" ORDER BY %s %s", req.PagingQuery.Sort, order)
		}

		if req.PagingQuery.Limit != 0 {
			limitClause = fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
			userParams = append(userParams, req.PagingQuery.Limit, req.PagingQuery.Offset)
			i += 2
		}
	}

	if req.Username != "" {
		if whereClause == "" {
			whereClause += " WHERE"
		} else {
			whereClause += " AND"
		}

		whereClause += fmt.Sprintf(" username = $%d", i)
		userParams = append(userParams, req.Username)
		userTotalParams = append(userTotalParams, req.Username)
		i++
	}

	if req.Email != "" {
		if whereClause == "" {
			whereClause += " WHERE"
		} else {
			whereClause += " AND"
		}

		whereClause += fmt.Sprintf(" email = $%d", i)
		userParams = append(userParams, req.Email)
		userTotalParams = append(userTotalParams, req.Email)
		i++
	}

	query += whereClause + orderClause + limitClause
	totalQuery += whereClause

	rows, err := db.Query(query, userParams...)
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}
	defer rows.Close()

	users := []*UserInList{}
	for rows.Next() {
		user := &UserInList{}
		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)
		if err != nil {
			logger.Err("failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		users = append(users, user)
	}

	var totalRows sql.NullInt64
	err = db.QueryRow(totalQuery, userTotalParams...).Scan(&totalRows)
	if err != nil {
		logger.Err("failed to get total rows: %s", err)
		return nil, fmt.Errorf("failed to get total rows: %s", err)
	}

	userList := &UserList{
		Users: users,
		Total: totalRows.Int64,
	}

	return userList, nil
}

func GetUserProfile(req *UserInList) (*UserInList, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	if req.Id == 0 {
		return nil, fmt.Errorf("id is mandatory")
	}

	userQuery := `SELECT u.id AS user_id, u.username, u.email, u.role
	FROM users u
	WHERE u.id = $1`

	userRow := db.QueryRow(userQuery, req.Id)

	user := UserInList{}

	err := userRow.Scan(&user.Id, &user.Username, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Err("user not found")
			return nil, fmt.Errorf("user not found")
		}

		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	languageActionQuery := `SELECT u.id AS user_id, l.name AS language, lv.name AS level, a.name, al.name AS action_level
	FROM user_profiles up
	JOIN users u ON up.user_id = u.id
	JOIN languages l ON up.language_id = l.id
	JOIN actions_levels al ON up.actions_levels_id = al.id
	JOIN levels lv ON al.level_id = lv.id
	LEFT JOIN actions a ON al.action_id = a.id
	WHERE u.id = $1
	GROUP BY u.id, u.email, u.username, l.name, lv.name, a.name, al.name
	ORDER BY u.id, l.name`

	rows, err := db.Query(languageActionQuery, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Err("user not found")
			return nil, fmt.Errorf("user not found")
		}

		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}
	defer rows.Close()

	userLanguageMap := map[string][]*UserAction{}
	for rows.Next() {
		var (
			id       int64
			language string
			level    string
			// actions   []sql.NullString
			action      sql.NullString
			actionLevel sql.NullString
		)

		if err := rows.Scan(&id, &language, &level, &action, &actionLevel); err != nil {
			logger.Err("failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		var actionString string = ""
		if action.Valid && actionLevel.Valid {
			actionString = fmt.Sprintf("%s %s", action.String, actionLevel.String)
		}

		if _, ok := userLanguageMap[language]; !ok {
			userLanguageMap[language] = append(userLanguageMap[language], &UserAction{
				Level:  level,
				Action: actionString,
			})
		} else {
			userLanguageMap[language] = append(userLanguageMap[language], &UserAction{
				Level:  level,
				Action: actionString,
			})
		}
	}

	for language, actions := range userLanguageMap {
		user.Languages = append(user.Languages, &UserLanguageProfile{
			Language: language,
			Actions:  actions,
		})
	}

	return &user, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func GetUserInterests(req *UserInList) (*UserInterests, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	if req.Id == 0 {
		return nil, fmt.Errorf("id is mandatory")
	}

	query := `SELECT i.name AS interest
	FROM users u
	INNER JOIN user_interests ui ON u.id = ui.user_id
	INNER JOIN interests i ON ui.interest_id = i.id
	WHERE u.id = $1
	ORDER BY i.name;`

	rows, err := db.Query(query, req.Id)
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	interests := []string{}
	for rows.Next() {
		var interest string
		if err := rows.Scan(&interest); err != nil {
			logger.Err("failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		interests = append(interests, interest)
	}

	return &UserInterests{
		Interests: interests,
	}, nil
}

func CreateUser(req *UserInList) (*Response, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	if req.Username == "" || req.Email == "" || (req.Password == "" && !req.IsOAuth) {
		logger.Err("all the fields are mandatory")
		return nil, fmt.Errorf("all the fields are mandatory")
	}

	query := "INSERT INTO users (username, email, password, oauth) VALUES ($1, $2, $3, $4) RETURNING id"
	createUserParams := []interface{}{req.Username, req.Email, req.Password}

	if req.IsOAuth {
		createUserParams = append(createUserParams, true)
	} else {
		createUserParams = append(createUserParams, false)
	}

	var userId sql.NullInt64
	err := db.QueryRow(query, createUserParams...).Scan(&userId)
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &Response{
		Message:   "User created successfully",
		CreatedId: userId.Int64,
	}, nil
}

func UpdateUser(req *UserUpdateRequest) (*Response, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	if req.Id == 0 {
		return nil, fmt.Errorf("id is mandatory")
	}

	query := "UPDATE users SET "
	updateUserParams := []interface{}{}

	i := 1
	if req.Username != "" {
		query += fmt.Sprintf("username = $%d", i)
		updateUserParams = append(updateUserParams, req.Username)
		i++
	}
	if req.Email != "" {
		if i != 1 {
			query += ","
		}
		query += fmt.Sprintf(" email = $%d", i)
		updateUserParams = append(updateUserParams, req.Email)
		i++
	}
	if req.Password != "" {
		if i != 1 {
			query += ","
		}
		query += fmt.Sprintf(" password = $%d", i)
		updateUserParams = append(updateUserParams, req.Password)
		i++
	}
	if len(req.Role.String()) > 0 {
		if i != 1 {
			query += ","
		}
		query += fmt.Sprintf(" role = $%d", i)
		updateUserParams = append(updateUserParams, strings.ToLower(req.Role.String()))
		i++
	}

	query += fmt.Sprintf(" WHERE id = $%d", i)
	updateUserParams = append(updateUserParams, req.Id)

	_, err := db.Exec(query, updateUserParams...)
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &Response{
		Message: "User updated successfully",
	}, nil
}

func DeleteUser(req *UserInList) (*Response, error) {
	config = utils.LoadEnv()
	db = utils.ConnectDatabase(config)
	defer db.Close()

	if req.Id == 0 {
		logger.Err("id is mandatory")
		return nil, fmt.Errorf("id is mandatory")
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, req.Id)
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &Response{
		Message: "User deleted successfully",
	}, nil
}
