package utils

const (
	Insert_Data_Into_Employees = "insert into employees (first_name,last_name,company_name,address,city,country,postal,phone,email,web) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	Select_All_From_Employee   = "select * from employees"
	Update_Employee_By_Email   = "update employees set company_name = $1 , first_name = $2 where email = $3"
)
