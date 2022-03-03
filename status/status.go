package status

const (
	Input                         = "10"
	Success                       = "20"
	TemporaryFailure              = "40"
	PermanentFailure              = "50"
	NotFound                      = "51"
	BadRequest                    = "59"
	ClientCertificateRequired     = "60"
	TransientCertificateRequested = "61"
	AuthorisedCertificateRequired = "62"
	CertificateNotAccepted        = "63"
	FutureCertificateRejected     = "64"
	ExpiredCertificateRejected    = "65"
)
