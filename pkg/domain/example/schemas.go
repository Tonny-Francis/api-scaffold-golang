package example

type MyForm struct {
	Name  string `schema:"name" binding:"required"`
	Phone string `schema:"phone" binding:"required"`
}

type MyHeaders struct {
	Tenant string `header:"tenant" binding:"required"`
}

type MyParams struct {
	ID string `uri:"id" binding:"required,numeric"`
}

type MyQueryParams struct {
	Page  int    `form:"page" binding:"required,numeric"`
}