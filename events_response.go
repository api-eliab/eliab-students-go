package main 

type EventsResponse struct {
		Events []struct {
			ID              int    `json:"id"`
			Title           string `json:"title"`
			Description     string `json:"description"`
			Icon            int    `json:"icon"`
			NeedAssisstance bool   `json:"need_assisstance"`
			Assisstance bool   `json:"assisstance"`
			NeedPayment     bool   `json:"need_payment"`
			Paid     bool   `json:"paid"`
			FormattedDate   string `json:"formatted_date"`
		} `json:"events"`
} 




