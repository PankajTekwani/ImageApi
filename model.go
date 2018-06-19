package main

import(
	//"fmt"
	"time"
)
type Response struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Outputs []struct {
		ID     string `json:"id"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Model     struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			CreatedAt  time.Time `json:"created_at"`
			AppID      string    `json:"app_id"`
			OutputInfo struct {
				Message string `json:"message"`
				Type    string `json:"type"`
				TypeExt string `json:"type_ext"`
			} `json:"output_info"`
			ModelVersion struct {
				ID        string    `json:"id"`
				CreatedAt time.Time `json:"created_at"`
				Status    struct {
					Code        int    `json:"code"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"model_version"`
			DisplayName string `json:"display_name"`
		} `json:"model"`
		Input struct {
			ID   string `json:"id"`
			Data struct {
				Image struct {
					URL string `json:"url"`
				} `json:"image"`
			} `json:"data"`
		} `json:"input"`
		Data struct {
			Concepts []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Value float64 `json:"value"`
				AppID string  `json:"app_id"`
			} `json:"concepts"`
		} `json:"data"`
	} `json:"outputs"`
}

type URLEntry struct {
	URL        string  `json:"URL"`
	Probablity float64 `json:"probablity"`
}

type OutputJson struct {
	Tag    string `json:"tag"`
	Length int    `json:"length"`
	Data   []URLEntry `json:"data"`
	Status int `json:"status"`
}

var TAG_MAP = map[string][]URLEntry{}


/*
This function basically fills the TAG_MAP data structure. It first checks for the tag in the TAG_MAP map and if the tag is present that it appends the image to the appropriate location in the array slice based on its probability value. Otherwise, it simply inserts the image in the array slice.

Input:
	1: Tag in whihc we want to add an entry
	2: url: url of the image.
	3: value: probability that the tag relates to the image
*/
func addToTagMap(tag string, url string, value float64){
	entry := URLEntry{URL: url,Probablity: value}
	var flag = 0
	if arr, ok := TAG_MAP[tag]; ok {
		for i, element := range arr {
			if element.Probablity < value {
				if(len(arr) < max_arr_len){
					arr = append(arr, URLEntry{})
				}
				copy(arr[i+1:],arr[i:])
				arr[i] = entry
				TAG_MAP[tag] = arr
				flag = 1
				break
			}
		}
	}
	
	if flag == 0 {
		//TAG_MAP[tag] = []URLEntry{}
		TAG_MAP[tag] = append(TAG_MAP[tag], entry)
	}
	
	/*if tag == "portrait" {
			fmt.Println("if",tag,entry,len(TAG_MAP[tag]),TAG_MAP[tag])
		}*/
}

func search(tag string) *OutputJson {
	output := &OutputJson{}
	output.Tag = tag
	output.Status = -1
	if arr, ok := TAG_MAP[tag]; ok {
		if(len(arr)>search_result){
			output.Data = arr[:search_result]
			output.Length = search_result
		}else{
			output.Data = arr
			output.Length = len(arr)
		}
		output.Status = 200
	}else{
		output.Length = 0;
		output.Status = 200
		output.Data = nil
	}
	return output
}
