package query

const InsertUser = `
INSERT INTO "user" (id, name, email) VALUES
(@id, @name, @email);
`
const InsertAccount = `
INSERT INTO account (id, handler, created_at, last_login, user_id) VALUES
(@id, @handler, @created_at, @last_login, @user_id);
`
