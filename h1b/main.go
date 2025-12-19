package main

type Data struct {
	CaseNumber                  string `json:"case_number"`
	JobTitle                    string `json:"job_title"`
	EmployerName                string `json:"employer_name"`
	SecondaryEntityBusinessName string `json:"secondary_entity_business_name"`
	EmployerPocEmail            string `json:"employer_poc_email"`
	EmployerPocPhone            string `json:"employer_poc_phone"`
	EmployerPocFirstName        string `json:"employer_poc_first_name"`
	EmployerPocLastName         string `json:"employer_poc_last_name"`
	WorksiteAddress1            string `json:"worksite_address1"`
	WorksiteAddress2            string `json:"worksite_address2"`
	WorksiteCity                string `json:"worksite_city"`
	WorksiteCounty              string `json:"worksite_county"`
	LawfirmNameBusinessName     string `json:"lawfirm_name_business_name"`
	EmployerAddress1            string `json:"employer_address1"`
	EmployerAddress2            string `json:"employer_address2"`
	EmployerCity                string `json:"employer_city"`
	EmployerState               string `json:"employer_state"`
	EmployerPostalCode          string `json:"employer_postal_code"`
	BeginDate                   string `json:"begin_date"`
	EndDate                     string `json:"end_date"`
	TotalWorkerPositions        string `json:"total_worker_positions"`
	H1bDependent                string `json:"h_1b_dependent"`
}

func main() {

}
