package pkivault

import (
	"context"
	"database/sql"
	fmt "fmt"
	"strings"

	grpctoken "github.com/ZolaraProject/library/grpctoken"
	logger "github.com/ZolaraProject/library/logger"
	. "github.com/ZolaraProject/pki-vault-service/pkivaultrpc"
	_ "github.com/lib/pq"
)

func (*server) GetUsers(ctx context.Context, req *UserRequest) (*UserList, error) {
	grpcToken := grpctoken.GetToken(ctx)

	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err(grpcToken, "Open error : %v", err)
		return nil, err
	}

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
		logger.Err(grpcToken, "failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}
	defer rows.Close()

	users := []*UserInList{}
	for rows.Next() {
		user := &UserInList{}
		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)
		if err != nil {
			logger.Err(grpcToken, "failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		users = append(users, user)
	}

	var totalRows sql.NullInt64
	err = db.QueryRow(totalQuery, userTotalParams...).Scan(&totalRows)
	if err != nil {
		logger.Err(grpcToken, "failed to get total rows: %s", err)
		return nil, fmt.Errorf("failed to get total rows: %s", err)
	}

	userList := &UserList{
		Users: users,
		Total: totalRows.Int64,
	}

	return userList, nil
}

func (*server) GetUserProfile(ctx context.Context, req *GetUserRequest) (*UserInList, error) {
	grpcToken := grpctoken.GetToken(ctx)

	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err(grpcToken, "Open error : %v", err)
		return nil, err
	}

	if req.Id == 0 {
		return nil, fmt.Errorf("id is mandatory")
	}

	userQuery := `SELECT u.id AS user_id, u.username, u.email, u.role
	FROM users u
	WHERE u.id = $1`

	userRow := db.QueryRow(userQuery, req.Id)

	user := UserInList{}

	err = userRow.Scan(&user.Id, &user.Username, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Err(grpcToken, "user not found")
			return nil, fmt.Errorf("user not found")
		}

		logger.Err(grpcToken, "failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	languageActionQuery := `SELECT u.id AS user_id, l.name AS language, a.name
	FROM user_profiles up
	JOIN users u ON up.user_id = u.id
	JOIN languages l ON up.language_id = l.id
	LEFT JOIN actions a ON up.action_id = a.id
	WHERE u.id = $1
	GROUP BY u.id, u.email, u.username, l.name, a.name
	ORDER BY u.id, l.name`

	rows, err := db.Query(languageActionQuery, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Err(grpcToken, "user not found")
			return nil, fmt.Errorf("user not found")
		}

		logger.Err(grpcToken, "failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}
	defer rows.Close()

	userLanguageMap := map[string][]string{}
	for rows.Next() {
		var (
			id       int64
			language string
			// actions   []sql.NullString
			action sql.NullString
		)

		if err := rows.Scan(&id, &language, &action); err != nil {
			logger.Err(grpcToken, "failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		if _, ok := userLanguageMap[language]; !ok {
			userLanguageMap[language] = append(userLanguageMap[language], action.String)
		} else {
			userLanguageMap[language] = append(userLanguageMap[language], action.String)
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

// func (*server) GetUserInterests(ctx context.Context, req *UserInList) (*UserInterests, error) {
// 	db, err := sql.Open("postgres", DbUrl())
// 	if err != nil {
// 		logger.Err("Open error : %v", err)
// 		return nil, err
// 	}

// 	if req.Id == 0 {
// 		return nil, fmt.Errorf("id is mandatory")
// 	}

// 	query := `SELECT i.name AS interest
// 	FROM users u
// 	INNER JOIN user_interests ui ON u.id = ui.user_id
// 	INNER JOIN interests i ON ui.interest_id = i.id
// 	WHERE u.id = $1
// 	ORDER BY i.name;`

// 	rows, err := db.Query(query, req.Id)
// 	if err != nil {
// 		logger.Err("failed to execute query: %s", err)
// 		return nil, fmt.Errorf("failed to execute query: %s", err)
// 	}

// 	interests := []string{}
// 	for rows.Next() {
// 		var interest string
// 		if err := rows.Scan(&interest); err != nil {
// 			logger.Err("failed to scan row: %s", err)
// 			return nil, fmt.Errorf("failed to scan row: %s", err)
// 		}

// 		interests = append(interests, interest)
// 	}

// 	return &UserInterests{
// 		Interests: interests,
// 	}, nil
// }
