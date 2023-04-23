package user

type User struct {
	Id             int64  `db:"id"`
	IdentityNumber string `db:"identity_number"`
	Uid            string `db:"uid"`
	FullName       string `db:"full_name"`
	AvatarUrl      string `db:"avatar_url"`
	SessionId      string `db:"session_id"`
	Biography      string `db:"biography"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      int64  `db:"updated_at"`
}
