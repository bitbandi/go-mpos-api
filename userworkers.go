package mpos

type UserWorkers struct {
	Id              int `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Monitor         uint `json:"monitor"`
	CountAll        uint `json:"count_all"`
	CountAllArchive uint `json:"count_all_archive"`
	Hashrate        int `json:"hashrate"`
	Difficulty      int `json:"difficulty"`
}

type userWorkersResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    []UserWorkers `json:"data"`
	       } `json:"getuserworkers"`
}

func (client *MposClient) GetUserWorkers() ([]UserWorkers, error) {
	userworkers := userWorkersResponse{}
	req := &MposRequest{Page: "api", Action:"getuserworkers", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&userworkers)
	if err != nil {
		return userworkers.Result.Data, err
	}

	return userworkers.Result.Data, err
}

