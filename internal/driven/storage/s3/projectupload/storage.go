package projectupload

import (
	"github.com/isdzulqor/donation-hub/internal/common/errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage struct {
	cfg Config
}

type Config struct {
	BucketName string
	ObjectKey  string
}

func (c Config) Validate() (err error) {
	if c.BucketName == "" {
		return errors.ErrBucketNameIsRequired
	}
	if c.ObjectKey == "" {
		return errors.ErrObjectKeyIsRequired
	}
	return nil
}

func NewStorage(cfg Config) *Storage {
	err := cfg.Validate()
	if err != nil {
		panic(err)
	}
	return &Storage{
		cfg: cfg,
	}
}

func (s *Storage) GenerateUploadURL(objectKey string) (string, error) {
	// Create a new session using the default AWS configuration
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", err
	}

	// Create a new S3 service client
	svc := s3.New(sess)

	// Set the parameters for the presigned URL
	params := &s3.PutObjectInput{
		Bucket: aws.String(s.cfg.BucketName),
		Key:    aws.String(objectKey),
	}

	// Generate the presigned URL
	req, _ := svc.PutObjectRequest(params)
	url, err := req.Presign(15 * time.Minute) // Set the expiration time for the URL
	if err != nil {
		return "", err
	}

	return url, nil
}
