package atom_auth

type UserDataListResponseModel struct {
	Id_User          int    `json:"id_user"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Have_Filled_Form int    `json:"have_filled_form"`
	Is_Admin         int    `json:"is_admin"`
	Is_Active        int    `json:"is_active"`
}

type UsernameInputRequestModel struct {
	Username string `json:"username"`
}

type UserCredentialsRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FatSecretTokenResponseModel struct {
	Access_Token string `json:"access_token"`
	Expires_In   int    `json:"expires_in"`
}
