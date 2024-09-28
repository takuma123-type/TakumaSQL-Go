package tableoutput

type CreateTableOutput struct {
	Name string `json:"name"`
}

type TablePresenter interface {
	PresentCreateTable(output *CreateTableOutput)
}
