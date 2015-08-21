package fpp

type DetectedFaceResponse struct {
	SessionId   string `json:"session_id"`
	ImageId     string `json:"img_id"`
	ImageUrl    string `json:"url"`
	ImageWidth  int    `json:"img_width"`
	ImageHeight int    `json:"img_height"`
	Faces       []Face `json:"face"`
}

type Face struct {
	FaceId     string `json:"face_id"`
	Attributes struct {
		Age struct {
			Range int `json:"range"`
			Value int `json:"value"`
		} `json:"age"`
		Gender struct {
			Confidence float64 `json:"confidence"`
			Value      string  `json:"value"`
		} `json:"gender"`
		Race struct {
			Confidence float64 `json:"confidence"`
			Value      string  `json:"value"`
		} `json:"race"`
	} `json:"attribute"`
	// Position   struct{} `json:"position"`
	// Tag        string   `json:"tag"`
}

func DetectFace(imageUrl string) (*DetectedFaceResponse, error) {
	var detected DetectedFaceResponse
	err := GetRequest("detection/detect", map[string]string{
		"url": imageUrl,
	}, &detected)
	return &detected, err
}
