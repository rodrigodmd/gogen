package remote

type Repository interface {
	Get(url, filepath string) error
	Update(url, filepath string) error
	Delete(filepath string) error
}


