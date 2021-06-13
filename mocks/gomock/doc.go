package gomock

//go:generate mockgen -source=doer/doer.go -destination=mock/mock_doer.go -package mocks

// Need to install
//go get -u github.com/golang/mock/gomock
//go get -u github.com/golang/mock/mockgen
