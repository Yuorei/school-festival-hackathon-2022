package controller


type User_register struct {
	Jwd string `json:"jwd"`
}

type User_res struct {
	IsNewUser bool `json:"isNewUser"`
}

type Rent_lists struct {
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	Deadline    string `json:"deadline"`
}
type Upload_image_url struct {
	Image_url string `json:"image_url"`
}

type Res_lists struct {
	Uuid        string `json:"id"`
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	Deadline    string `json:"deadline"`
}