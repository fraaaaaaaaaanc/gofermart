package cookie

type Cookie struct {
	secretKeyJWTToken string
}

func NewCookie(secretKeyJWTToken string) Cookie {
	return Cookie{
		secretKeyJWTToken: secretKeyJWTToken,
	}
}
