package models

type PodListModel []*PodModel

type PodModel struct {
	PodName  string
	PodImage string
	PodNode  string
}

func (this *PodModel) ParseAction(action string) (*WsResponse, error) {
	return NewWsResponse("prodList", PodList()), nil
}

func (this PodListModel) ParseAction(action string) (*WsResponse, error) {
	return nil, nil
}

func PodList() []*PodModel {
	return []*PodModel{{
		PodName:  "1",
		PodImage: "1",
		PodNode:  "1",
	}, {
		PodName:  "2",
		PodImage: "2",
		PodNode:  "2",
	}}
}
