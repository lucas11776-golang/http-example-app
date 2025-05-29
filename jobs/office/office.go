package office

// TODO: May need to rename and move to workspace so everyone can have access to each one
// e.g When a JuniorAnalyst is done the can submit to SeniorAnalyst.
type Office struct {
	OperationManager OperationManager
	SeniorAnalyst    SeniorAnalyst
	JuniorAnalyst    JuniorAnalyst
	GraphicDesigner  GraphicDesigner
}

func NewOffice() (*Office, error) {
	office := &Office{}

	return office, nil
}

// Comment
func (ctx *Office) LoadLatestNews() {

}

// Comment
func (ctx *Office) LoadLatestNewsClients() {

}
