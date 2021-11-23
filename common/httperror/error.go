package httperror

type XmoError struct {
	Biz *BizCode
	Msg string
}

func (xe XmoError) Error() string {
	return xe.Msg
}

func Create(err error) *XmoError {
	xe := new(XmoError)
	switch v := err.(type) {
	case *XmoError:
		xe = v
	default:
		xe = &XmoError{
			Msg: err.Error(),
			Biz: BIZ_DEFAULT_ERROR,
		}
	}
	return xe
}

func (xe *XmoError) WithBiz(biz *BizCode) *XmoError {
	xe.Biz = biz
	xe.Msg = biz.Msg
	return xe
}

func (xe *XmoError) WithMsg(msg string) *XmoError {
	xe.Msg = msg
	return xe
}
