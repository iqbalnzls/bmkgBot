package bmkg

type BMKGWrapperIFace interface {
	Get(key string) (body []byte, err error)
}
