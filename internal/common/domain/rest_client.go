package domain

type RestClient interface {
	DoGet(uri string, params map[string]string) (string, error)
	DoPost(uri string, params string) (string, error)
}
